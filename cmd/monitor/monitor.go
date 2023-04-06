/*
Copyright © 2022 Polygon <engineering@polygon.technology>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Lesser General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Lesser General Public License for more details.

You should have received a copy of the GNU Lesser General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package monitor

import (
	"context"
	"fmt"
	"math/big"
	"net/url"
	"os"
	"sort"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	ethrpc "github.com/ethereum/go-ethereum/rpc"

	"github.com/cenkalti/backoff"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/maticnetwork/polygon-cli/metrics"
	"github.com/maticnetwork/polygon-cli/rpctypes"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var (
	batchSize   uint64
	windowSize  int
	verbosity   int64
	recentOnly  bool
	intervalStr string
	interval    time.Duration
	logFile     string
)

type (
	monitorStatus struct {
		ChainID   *big.Int
		HeadBlock *big.Int
		PeerCount uint64
		GasPrice  *big.Int

		Blocks            map[string]rpctypes.PolyBlock `json:"-"`
		BlocksLock        sync.RWMutex                  `json:"-"`
		MaxBlockRetrieved *big.Int

		WindowOffset int
	}
	chainState struct {
		HeadBlock uint64
		ChainID   *big.Int
		PeerCount uint64
		GasPrice  *big.Int
	}
	monitorMode int
)

const (
	monitorModeHelp monitorMode = iota
	monitorModeExplorer
	monitorModeBlock
)

func getChainState(ctx context.Context, ec *ethclient.Client) (*chainState, error) {
	var err error
	cs := new(chainState)
	cs.HeadBlock, err = ec.BlockNumber(ctx)
	if err != nil {
		return nil, fmt.Errorf("couldn't fetch block number: %s", err.Error())
	}

	cs.ChainID, err = ec.ChainID(ctx)
	if err != nil {
		return nil, fmt.Errorf("couldn't fetch chain id: %s", err.Error())
	}

	cs.PeerCount, err = ec.PeerCount(ctx)
	if err != nil {
		log.Info().Err(err).Msg("Using fake peer count")
		cs.PeerCount = 0
	}

	cs.GasPrice, err = ec.SuggestGasPrice(ctx)
	if err != nil {
		return nil, fmt.Errorf("couldn't estimate gas: %s", err.Error())
	}

	return cs, nil

}

// monitorCmd represents the monitor command
var MonitorCmd = &cobra.Command{
	Use:   "monitor [rpc-url]",
	Short: "A simple terminal monitor for a blockchain",
	Args:  cobra.MinimumNArgs(1),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		// validate url argument
		_, err := url.Parse(args[0])
		if err != nil {
			return err
		}

		// validate batch-size flag
		if batchSize == 0 {
			return fmt.Errorf("batch-size can't be equal to zero")
		}

		// validate interval duration
		if interval, err = time.ParseDuration(intervalStr); err != nil {
			return err
		}

		return setMonitorLogLevel(verbosity)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()

		rpc, err := ethrpc.DialContext(ctx, args[0])
		if err != nil {
			log.Error().Err(err).Msg("Unable to dial rpc")
			return err
		}
		ec := ethclient.NewClient(rpc)

		ms := new(monitorStatus)

		ms.MaxBlockRetrieved = big.NewInt(0)
		ms.BlocksLock.Lock()
		ms.Blocks = make(map[string]rpctypes.PolyBlock, 0)
		ms.BlocksLock.Unlock()
		ms.ChainID = big.NewInt(0)
		zero := big.NewInt(0)

		isUiRendered := false
		errChan := make(chan error)
		go func() {
			for {
				var cs *chainState
				cs, err = getChainState(ctx, ec)
				if err != nil {
					log.Error().Err(err).Msg("Encountered issue fetching network information")
					time.Sleep(interval)
					continue
				}

				ms.HeadBlock = new(big.Int).SetUint64(cs.HeadBlock)
				ms.ChainID = cs.ChainID
				ms.PeerCount = cs.PeerCount
				ms.GasPrice = cs.GasPrice

				from := new(big.Int).Sub(ms.HeadBlock, new(big.Int).SetUint64(batchSize))
				// Prevent getBlockRange from fetching duplicate blocks.
				if ms.MaxBlockRetrieved.Cmp(from) == 1 {
					from.Add(ms.MaxBlockRetrieved, big.NewInt(1))
				}
				// Skip next poll if the latest block is already at the head.
				if from.Cmp(ms.HeadBlock) >= 0 {
					continue
				}

				if from.Cmp(zero) < 0 {
					from.SetInt64(0)
				}

				log.Debug().
					Int64("from", from.Int64()).
					Int64("to", ms.HeadBlock.Int64()).
					Int64("max", ms.MaxBlockRetrieved.Int64()).
					Msg("Getting block range")

				err = ms.getBlockRange(ctx, from, ms.HeadBlock, rpc)
				if err != nil {
					log.Error().Err(err).Msg("There was an issue fetching the block range")
				}

				if !isUiRendered {
					go func() {
						errChan <- renderMonitorUI(ms)
					}()
					isUiRendered = true
				}

				time.Sleep(interval)
			}
		}()

		err = <-errChan
		return err
	},
}

func (ms *monitorStatus) getBlockRange(ctx context.Context, from, to *big.Int, c *ethrpc.Client) error {
	one := big.NewInt(1)
	blms := make([]ethrpc.BatchElem, 0)
	for i := from; i.Cmp(to) != 1; i.Add(i, one) {
		r := new(rpctypes.RawBlockResponse)
		var err error
		blms = append(blms, ethrpc.BatchElem{
			Method: "eth_getBlockByNumber",
			Args:   []interface{}{"0x" + i.Text(16), true},
			Result: r,
			Error:  err,
		})
	}
	b := backoff.NewExponentialBackOff()
	b.MaxElapsedTime = 3 * time.Minute
	retryable := func() error {
		err := c.BatchCallContext(ctx, blms)
		return err
	}
	err := backoff.Retry(retryable, b)
	if err != nil {
		return err
	}
	for _, b := range blms {
		if b.Error != nil {
			return b.Error
		}
		pb := rpctypes.NewPolyBlock(b.Result.(*rpctypes.RawBlockResponse))

		ms.BlocksLock.Lock()
		ms.Blocks[pb.Number().String()] = pb
		ms.BlocksLock.Unlock()

		if ms.MaxBlockRetrieved.Cmp(pb.Number()) == -1 {
			ms.MaxBlockRetrieved = pb.Number()
		}
	}

	return nil
}

func init() {
	MonitorCmd.PersistentFlags().Uint64VarP(&batchSize, "batch-size", "b", 25, "Number of requests per batch")
	MonitorCmd.PersistentFlags().IntVarP(&windowSize, "window-size", "w", 25, "Number of blocks visible in the window")
	MonitorCmd.PersistentFlags().Int64VarP(&verbosity, "verbosity", "v", 200, "0 - Silent\n100 Fatal\n200 Error\n300 Warning\n400 Info\n500 Debug\n600 Trace")
	MonitorCmd.PersistentFlags().BoolVarP(&recentOnly, "recent-only", "r", false, "Only show blocks from latest batch")
	MonitorCmd.PersistentFlags().StringVarP(&intervalStr, "interval", "i", "5s", "Amount of time between batch block rpc calls")
	MonitorCmd.PersistentFlags().StringVarP(&logFile, "log-file", "l", "", "Write logs to a file (default stdout)")
}

func renderMonitorUI(ms *monitorStatus) error {
	if err := ui.Init(); err != nil {
		return err
	}
	defer ui.Close()

	currentMode := monitorModeExplorer

	blockTable := widgets.NewList()
	blockTable.TextStyle = ui.NewStyle(ui.ColorWhite)

	h0 := widgets.NewParagraph()
	h0.Title = "Current"

	h1 := widgets.NewParagraph()
	h1.Title = "Gas Price"

	h2 := widgets.NewParagraph()
	h2.Title = "Current Peers"

	h3 := widgets.NewParagraph()
	h3.Title = "Chain ID"

	h4 := widgets.NewParagraph()
	h4.Title = "Avg Block Time"

	sl0 := widgets.NewSparkline()
	sl0.LineColor = ui.ColorRed
	slg0 := widgets.NewSparklineGroup(sl0)
	slg0.Title = "TXs / Block"

	sl1 := widgets.NewSparkline()
	sl1.LineColor = ui.ColorGreen
	slg1 := widgets.NewSparklineGroup(sl1)
	slg1.Title = "Gas Price"

	sl2 := widgets.NewSparkline()
	sl2.LineColor = ui.ColorYellow
	slg2 := widgets.NewSparklineGroup(sl2)
	slg2.Title = "Block Size"

	sl3 := widgets.NewSparkline()
	sl3.LineColor = ui.ColorBlue
	slg3 := widgets.NewSparklineGroup(sl3)
	slg3.Title = "Uncles"

	sl4 := widgets.NewSparkline()
	sl4.LineColor = ui.ColorMagenta
	slg4 := widgets.NewSparklineGroup(sl4)
	slg4.Title = "Gas Used"

	grid := ui.NewGrid()
	blockGrid := ui.NewGrid()

	b0 := widgets.NewParagraph()
	b0.Title = "Block Headers"
	b0.Text = "Use the arrow keys to scroll through the transactions. Press <Esc> to go back to the explorer view"

	b1 := widgets.NewList()
	b1.Title = "Block Info"
	b1.TextStyle = ui.NewStyle(ui.ColorYellow)
	b1.WrapText = false

	b2 := widgets.NewList()
	b2.Title = "Transactions"
	b2.TextStyle = ui.NewStyle(ui.ColorGreen)
	b2.WrapText = true

	blockGrid.Set(
		ui.NewRow(1.0/10, b0),

		ui.NewRow(9.0/10,
			ui.NewCol(1.0/2, b1),
			ui.NewCol(1.0/2, b2),
		),
	)

	grid.Set(
		ui.NewRow(1.0/10,
			ui.NewCol(1.0/5, h0),
			ui.NewCol(1.0/5, h1),
			ui.NewCol(1.0/5, h2),
			ui.NewCol(1.0/5, h3),
			ui.NewCol(1.0/5, h4),
		),

		ui.NewRow(4.0/10,
			ui.NewCol(1.0/5, slg0),
			ui.NewCol(1.0/5, slg1),
			ui.NewCol(1.0/5, slg2),
			ui.NewCol(1.0/5, slg3),
			ui.NewCol(1.0/5, slg4),
		),
		ui.NewRow(5.0/10, blockTable),
	)

	termWidth, termHeight := ui.TerminalDimensions()
	grid.SetRect(0, 0, termWidth, termHeight)
	blockGrid.SetRect(0, 0, termWidth, termHeight)

	var selectedBlock rpctypes.PolyBlock
	var setBlock = false
	var allBlocks metrics.SortableBlocks
	var renderedBlocks metrics.SortableBlocks

	redraw := func(ms *monitorStatus) {
		log.Debug().Interface("ms", ms).Msg("Redrawing")

		if currentMode == monitorModeHelp {
			// TODO add some help context?
		} else if currentMode == monitorModeBlock {
			// render a block
			b1.Rows = metrics.GetSimpleBlockFields(selectedBlock)
			b2.Rows = metrics.GetSimpleBlockTxFields(selectedBlock, ms.ChainID)

			ui.Clear()
			ui.Render(blockGrid)
			return
		}

		if blockTable.SelectedRow == 0 {
			// default
			blocks := make([]rpctypes.PolyBlock, 0)

			ms.BlocksLock.RLock()
			for _, b := range ms.Blocks {
				blocks = append(blocks, b)
			}
			ms.BlocksLock.RUnlock()

			allBlocks = metrics.SortableBlocks(blocks)
			sort.Sort(allBlocks)
		}

		start := len(allBlocks) - windowSize - ms.WindowOffset
		if start < 0 {
			start = 0
		}
		end := len(allBlocks) - ms.WindowOffset
		renderedBlocks = allBlocks[start:end]

		h0.Text = fmt.Sprintf("Height: %s\nTime: %s", ms.HeadBlock.String(), time.Now().Format("02 Jan 06 15:04:05 MST"))
		gasGwei := new(big.Int).Div(ms.GasPrice, metrics.UnitShannon)
		h1.Text = fmt.Sprintf("%s gwei", gasGwei.String())
		h2.Text = fmt.Sprintf("%d", ms.PeerCount)
		h3.Text = ms.ChainID.String()
		h4.Text = fmt.Sprintf("%0.2f", metrics.GetMeanBlockTime(renderedBlocks))

		sl0.Data = metrics.GetTxsPerBlock(renderedBlocks)
		sl1.Data = metrics.GetMeanGasPricePerBlock(renderedBlocks)
		sl2.Data = metrics.GetSizePerBlock(renderedBlocks)
		sl3.Data = metrics.GetUnclesPerBlock(renderedBlocks)
		sl4.Data = metrics.GetGasPerBlock(renderedBlocks)

		// If a row has not been selected, continue to update the list with new blocks.
		rows, title := metrics.GetSimpleBlockRecords(renderedBlocks)
		blockTable.Rows = rows
		blockTable.Title = title

		blockTable.TextStyle = ui.NewStyle(ui.ColorWhite)
		blockTable.SelectedRowStyle = ui.NewStyle(ui.ColorWhite, ui.ColorRed, ui.ModifierBold)
		if blockTable.SelectedRow > 0 && blockTable.SelectedRow <= len(blockTable.Rows) {
			// Only changed the selected block when the user presses the up down keys.
			// Otherwise this will adjust when the table is updated automatically.
			if setBlock {
				selectedBlock = renderedBlocks[len(renderedBlocks)-blockTable.SelectedRow]
				setBlock = false
			}
		}

		ui.Render(grid)
	}

	currentBn := ms.HeadBlock
	uiEvents := ui.PollEvents()
	ticker := time.NewTicker(time.Second).C

	redraw(ms)

	currIdx := 0
	for {
		select {
		case e := <-uiEvents:
			switch e.ID {
			case "q", "<C-c>":
				return nil
			case "<Escape>":
				blockTable.SelectedRow = 0
				currentMode = monitorModeExplorer
				ms.WindowOffset = 0
				redraw(ms)
			case "<Enter>":
				if blockTable.SelectedRow > 0 {
					currentMode = monitorModeBlock
				}
				redraw(ms)
			case "<Resize>":
				payload := e.Payload.(ui.Resize)
				grid.SetRect(0, 0, payload.Width, payload.Height)
				blockGrid.SetRect(0, 0, payload.Width, payload.Height)
				ui.Clear()
				redraw(ms)
			case "<Up>", "<Down>":
				if currentMode == monitorModeBlock {
					if len(b2.Rows) != 0 && e.ID == "<Down>" {
						b2.ScrollDown()
					} else if len(b2.Rows) != 0 && e.ID == "<Up>" {
						b2.ScrollUp()
					}
					redraw(ms)
					break
				}

				if blockTable.SelectedRow == 0 {
					currIdx = 1
					blockTable.SelectedRow = currIdx
					setBlock = true
					redraw(ms)
					break
				}
				currIdx = blockTable.SelectedRow

				if e.ID == "<Down>" {
					log.Debug().Int("currIdx", currIdx).Int("windowSize", windowSize).Int("renderedBlocks", len(renderedBlocks)).Msg("Down")
					if currIdx > windowSize-1 && ms.WindowOffset < len(allBlocks)-windowSize {
						ms.WindowOffset += 1
						currIdx -= 1
						redraw(ms)
						break
					}
					currIdx += 1
					setBlock = true
				} else if e.ID == "<Up>" {
					log.Debug().Int("currIdx", currIdx).Int("windowSize", windowSize).Msg("Up")
					if currIdx <= 1 && ms.WindowOffset > 0 {
						ms.WindowOffset -= 1
						currIdx += 1
						redraw(ms)
						break
					}
					currIdx -= 1
					setBlock = true
				}
				// need a better way to understand how many rows are visible
				if currIdx > 0 && currIdx <= windowSize && currIdx <= len(renderedBlocks) {
					blockTable.SelectedRow = currIdx
				}

				redraw(ms)
			case "<MouseLeft>", "<MouseRight>", "<MouseRelease>", "<MouseWheelUp>", "<MouseWheelDown>":
				break
			default:
				log.Trace().Str("id", e.ID).Msg("Unknown ui event")
			}
		case <-ticker:
			if currentBn != ms.HeadBlock {
				currentBn = ms.HeadBlock
				redraw(ms)
			}
		}
	}
}

// setMonitorLogLevel sets the log level based on the flags. If the log file flag
// is set, then output will be written to the file instead of stdout. Use
// `tail -f <log file>` to see the log output in real time.
func setMonitorLogLevel(verbosity int64) error {
	if len(logFile) > 0 {
		file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0664)
		if err != nil {
			return err
		}
		log.Logger = zerolog.New(file).With().Logger()
	}

	if verbosity < 100 {
		zerolog.SetGlobalLevel(zerolog.PanicLevel)
	} else if verbosity < 200 {
		zerolog.SetGlobalLevel(zerolog.FatalLevel)
	} else if verbosity < 300 {
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	} else if verbosity < 400 {
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	} else if verbosity < 500 {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	} else if verbosity < 600 {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.TraceLevel)
	}

	return nil
}
