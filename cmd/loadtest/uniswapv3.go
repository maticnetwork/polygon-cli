package loadtest

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/maticnetwork/polygon-cli/contracts/uniswapv3"
	"github.com/rs/zerolog/log"
)

const (
	// The fee amount to enable for one basic point.
	// https://github.com/Uniswap/deploy-v3/blob/b7aac0f1c5353b36802dc0cf95c426d2ef0c3252/src/steps/add-1bp-fee-tier.ts#L5
	ONE_BP_FEE int64 = 100

	// The spacing between ticks to be enforced for all pools with the given fee amount.
	// https://github.com/Uniswap/deploy-v3/blob/b7aac0f1c5353b36802dc0cf95c426d2ef0c3252/src/steps/add-1bp-fee-tier.ts#L6
	ONE_BP_TICK_SPACING int64 = 1

	// Time units.
	ONE_MINUTE_SECONDS = 60
	ONE_HOUR_SECONDS   = ONE_MINUTE_SECONDS * 60
	ONE_DAY_SECONDS    = ONE_HOUR_SECONDS * 24
	ONE_MONTH_SECONDS  = ONE_DAY_SECONDS * 30
	ONE_YEAR_SECONDS   = ONE_DAY_SECONDS * 365

	// The max amount of seconds into the future the incentive startTime can be set.
	// https://github.com/Uniswap/deploy-v3/blob/b7aac0f1c5353b36802dc0cf95c426d2ef0c3252/src/steps/deploy-v3-staker.ts#L11
	MAX_INCENTIVE_START_LEAD_TIME = ONE_MONTH_SECONDS

	// The max duration of an incentive in seconds.
	// https://github.com/Uniswap/deploy-v3/blob/b7aac0f1c5353b36802dc0cf95c426d2ef0c3252/src/steps/deploy-v3-staker.ts#L13
	MAX_INCENTIVE_DURATION = ONE_YEAR_SECONDS * 2

	// The minimum tick that may be passed to `getSqrtRatioAtTick` computed from log base 1.0001 of 2**-128.
	// https://github.com/Uniswap/v3-core/blob/d8b1c635c275d2a9450bd6a78f3fa2484fef73eb/contracts/libraries/TickMath.sol#L9
	MIN_TICK = -887272
	// The maximum tick that may be passed to `getSqrtRatioAtTick` computed from log base 1.0001 of 2**128.
	// https://github.com/Uniswap/v3-core/blob/d8b1c635c275d2a9450bd6a78f3fa2484fef73eb/contracts/libraries/TickMath.sol#L11
	MAX_TICK = -MIN_TICK
)

type UniswapV3Addresses struct {
	FactoryV3, Multicall, ProxyAdmin, TickLens, NFTDescriptorLib, NFTDescriptor, TransparentUpgradeableProxy, NonfungiblePositionManager, Migrator, Staker, QuoterV2, SwapRouter02 common.Address
	WETH9, TokenA, TokenB                                                                                                                                                          common.Address
}

type UniswapV3Config struct {
	FactoryV3                   contractConfig[uniswapv3.UniswapV3Factory]
	Multicall                   contractConfig[uniswapv3.UniswapInterfaceMulticall]
	ProxyAdmin                  contractConfig[uniswapv3.ProxyAdmin]
	TickLens                    contractConfig[uniswapv3.TickLens]
	NFTDescriptor               contractConfig[uniswapv3.NonfungibleTokenPositionDescriptor]
	TransparentUpgradeableProxy contractConfig[uniswapv3.TransparentUpgradeableProxy]
	NonfungiblePositionManager  contractConfig[uniswapv3.NonfungiblePositionManager]
	Migrator                    contractConfig[uniswapv3.V3Migrator]
	Staker                      contractConfig[uniswapv3.UniswapV3Staker]
	QuoterV2                    contractConfig[uniswapv3.QuoterV2]
	SwapRouter02                contractConfig[uniswapv3.SwapRouter02]

	WETH9 contractConfig[uniswapv3.WETH9]
}

type contractConfig[T uniswapV3Contract] struct {
	Address  common.Address
	contract *T
}

type uniswapV3Contract interface {
	uniswapv3.UniswapV3Factory | uniswapv3.UniswapInterfaceMulticall | uniswapv3.ProxyAdmin | uniswapv3.TickLens | uniswapv3.WETH9 | uniswapv3.NonfungibleTokenPositionDescriptor | uniswapv3.TransparentUpgradeableProxy | uniswapv3.NonfungiblePositionManager | uniswapv3.V3Migrator | uniswapv3.UniswapV3Staker | uniswapv3.QuoterV2 | uniswapv3.SwapRouter02 | uniswapv3.ERC20
}

func deployUniswapV3(ctx context.Context, c *ethclient.Client, tops *bind.TransactOpts, cops *bind.CallOpts, knownAddresses UniswapV3Addresses, ownerAddress common.Address) (UniswapV3Config, error) {
	config := UniswapV3Config{}
	var err error

	// 1. Deploy UniswapV3Factory.
	config.FactoryV3.Address, config.FactoryV3.contract, err = deployOrInstantiateContract(
		ctx, c, tops, cops, "Factory", knownAddresses.FactoryV3,
		uniswapv3.DeployUniswapV3Factory,
		uniswapv3.NewUniswapV3Factory,
		func(contract *uniswapv3.UniswapV3Factory) (err error) {
			_, err = contract.Owner(cops)
			return
		},
	)
	if err != nil {
		return config, err
	}

	// 2. Enable one basic point fee tier.
	if err = enableOneBPFeeTier(config.FactoryV3.contract, tops, ONE_BP_FEE, ONE_BP_TICK_SPACING); err != nil {
		return config, err
	}

	// 3. Deploy UniswapInterfaceMulticall.
	config.Multicall.Address, config.Multicall.contract, err = deployOrInstantiateContract(
		ctx, c, tops, cops, "Multicall", knownAddresses.Multicall,
		uniswapv3.DeployUniswapInterfaceMulticall,
		uniswapv3.NewUniswapInterfaceMulticall,
		func(contract *uniswapv3.UniswapInterfaceMulticall) (err error) {
			_, err = contract.GetEthBalance(cops, common.Address{})
			return
		},
	)
	if err != nil {
		return config, err
	}

	// 4. Deploy ProxyAdmin.
	config.ProxyAdmin.Address, config.ProxyAdmin.contract, err = deployOrInstantiateContract(
		ctx, c, tops, cops, "ProxyAdmin", knownAddresses.ProxyAdmin,
		uniswapv3.DeployProxyAdmin,
		uniswapv3.NewProxyAdmin,
		func(contract *uniswapv3.ProxyAdmin) (err error) {
			_, err = contract.Owner(cops)
			return
		},
	)
	if err != nil {
		return config, err
	}

	// 5. Deploy TickLens.
	config.TickLens.Address, config.TickLens.contract, err = deployOrInstantiateContract(
		ctx, c, tops, cops, "TickLens", knownAddresses.TickLens,
		uniswapv3.DeployTickLens,
		uniswapv3.NewTickLens,
		func(contract *uniswapv3.TickLens) (err error) {
			// This call will revert because no ticks are populated yet.
			_, err = contract.GetPopulatedTicksInWord(cops, common.Address{}, int16(1))
			// TODO: Compare with error instead of string.
			if err.Error() == "execution reverted" {
				err = nil
			}
			return
		},
	)
	if err != nil {
		return config, err
	}

	// 6. Deploy WETH9.
	config.WETH9.Address, config.WETH9.contract, err = deployOrInstantiateContract(
		ctx, c, tops, cops, "WETH9", knownAddresses.WETH9,
		uniswapv3.DeployWETH9,
		uniswapv3.NewWETH9,
		func(contract *uniswapv3.WETH9) (err error) {
			_, err = contract.BalanceOf(cops, common.Address{})
			return
		},
	)
	if err != nil {
		return config, err
	}

	// 7. Deploy NonfungibleTokenPositionDescriptor.
	// Note that we previously deployed the NFTDescriptor library during the build process.
	config.NFTDescriptor.Address, config.NFTDescriptor.contract, err = deployOrInstantiateContract(
		ctx, c, tops, cops, "NFTDescriptor", knownAddresses.NFTDescriptor,
		func(*bind.TransactOpts, bind.ContractBackend) (common.Address, *types.Transaction, *uniswapv3.NonfungibleTokenPositionDescriptor, error) {
			var nativeCurrencyLabelBytes [32]byte
			copy(nativeCurrencyLabelBytes[:], "ETH")
			return uniswapv3.DeployNonfungibleTokenPositionDescriptor(tops, c, config.WETH9.Address, nativeCurrencyLabelBytes)
		},
		uniswapv3.NewNonfungibleTokenPositionDescriptor,
		func(contract *uniswapv3.NonfungibleTokenPositionDescriptor) (err error) {
			_, err = contract.WETH9(cops)
			return
		},
	)
	if err != nil {
		return config, err
	}

	// 8. Deploy TransparentUpgradeableProxy.
	config.TransparentUpgradeableProxy.Address, config.TransparentUpgradeableProxy.contract, err = deployOrInstantiateContract(
		ctx, c, tops, cops, "TransparentUpgradeableProxy", knownAddresses.TransparentUpgradeableProxy,
		func(*bind.TransactOpts, bind.ContractBackend) (common.Address, *types.Transaction, *uniswapv3.TransparentUpgradeableProxy, error) {
			var data []byte
			copy(data[:], "0x")
			return uniswapv3.DeployTransparentUpgradeableProxy(tops, c, config.NFTDescriptor.Address, config.ProxyAdmin.Address, data)
		},
		uniswapv3.NewTransparentUpgradeableProxy,
		func(contract *uniswapv3.TransparentUpgradeableProxy) (err error) {
			_, err = contract.Admin(tops)
			return
		},
	)
	if err != nil {
		return config, err
	}

	// 9. Deploy NonfungiblePositionManager.
	config.NonfungiblePositionManager.Address, config.NonfungiblePositionManager.contract, err = deployOrInstantiateContract(
		ctx, c, tops, cops, "NonfungiblePositionManager", knownAddresses.NonfungiblePositionManager,
		func(*bind.TransactOpts, bind.ContractBackend) (common.Address, *types.Transaction, *uniswapv3.NonfungiblePositionManager, error) {
			return uniswapv3.DeployNonfungiblePositionManager(tops, c, config.FactoryV3.Address, config.WETH9.Address, config.TransparentUpgradeableProxy.Address)
		},
		uniswapv3.NewNonfungiblePositionManager,
		func(contract *uniswapv3.NonfungiblePositionManager) (err error) {
			_, err = contract.BaseURI(cops)
			return
		},
	)
	if err != nil {
		return config, err
	}

	// 10. Deploy Migrator.
	config.Migrator.Address, config.Migrator.contract, err = deployOrInstantiateContract(
		ctx, c, tops, cops, "V3Migrator", knownAddresses.Migrator,
		func(*bind.TransactOpts, bind.ContractBackend) (common.Address, *types.Transaction, *uniswapv3.V3Migrator, error) {
			return uniswapv3.DeployV3Migrator(tops, c, config.FactoryV3.Address, config.WETH9.Address, config.NonfungiblePositionManager.Address)
		},
		uniswapv3.NewV3Migrator,
		func(contract *uniswapv3.V3Migrator) (err error) {
			_, err = contract.WETH9(cops)
			return
		},
	)
	if err != nil {
		return config, err
	}

	// 11. Set Factory owner.
	if err = setFactoryOwner(config.FactoryV3.contract, tops, ownerAddress); err != nil {
		return config, err
	}

	// 12. Deploy Staker.
	config.Staker.Address, config.Staker.contract, err = deployOrInstantiateContract(
		ctx, c, tops, cops, "Staker", knownAddresses.Staker,
		func(*bind.TransactOpts, bind.ContractBackend) (common.Address, *types.Transaction, *uniswapv3.UniswapV3Staker, error) {
			return uniswapv3.DeployUniswapV3Staker(tops, c, config.FactoryV3.Address, config.NonfungiblePositionManager.Address, big.NewInt(MAX_INCENTIVE_START_LEAD_TIME), big.NewInt(MAX_INCENTIVE_DURATION))
		},
		uniswapv3.NewUniswapV3Staker,
		func(contract *uniswapv3.UniswapV3Staker) (err error) {
			_, err = contract.Factory(cops)
			return
		},
	)
	if err != nil {
		return config, err
	}

	// 13. Deploy QuoterV2.
	config.QuoterV2.Address, config.QuoterV2.contract, err = deployOrInstantiateContract(
		ctx, c, tops, cops, "QuoterV2", knownAddresses.QuoterV2,
		func(*bind.TransactOpts, bind.ContractBackend) (common.Address, *types.Transaction, *uniswapv3.QuoterV2, error) {
			return uniswapv3.DeployQuoterV2(tops, c, config.FactoryV3.Address, config.WETH9.Address)
		},
		uniswapv3.NewQuoterV2,
		func(contract *uniswapv3.QuoterV2) (err error) {
			_, err = contract.Factory(cops)
			return
		},
	)
	if err != nil {
		return config, err
	}

	// 14. Deploy SwapRouter02.
	config.SwapRouter02.Address, config.SwapRouter02.contract, err = deployOrInstantiateContract(
		ctx, c, tops, cops, "SwapRouter02", knownAddresses.SwapRouter02,
		func(*bind.TransactOpts, bind.ContractBackend) (common.Address, *types.Transaction, *uniswapv3.SwapRouter02, error) {
			// Note: we specify an empty address for UniswapV2Factory.
			uniswapFactoryV2Address := common.Address{}
			return uniswapv3.DeploySwapRouter02(tops, c, uniswapFactoryV2Address, config.FactoryV3.Address, config.NonfungiblePositionManager.Address, config.WETH9.Address)
		},
		uniswapv3.NewSwapRouter02,
		func(contract *uniswapv3.SwapRouter02) (err error) {
			_, err = contract.Factory(cops)
			return
		},
	)
	if err != nil {
		return config, err
	}

	// 15. Transfer ProxyAdmin ownership.
	if err = transferProxyAdminOwnership(config.ProxyAdmin.contract, tops, ownerAddress); err != nil {
		return config, err
	}

	return config, nil
}

// Deploy or instantiate any UniswapV3 contract.
// This method will either deploy a contract if the known address is empty (equal to `common.Address{}` or `0x0“)
// or instantiate the contract if the known address is specified.
func deployOrInstantiateContract[T uniswapV3Contract](
	ctx context.Context,
	c *ethclient.Client,
	tops *bind.TransactOpts,
	cops *bind.CallOpts,
	name string,
	knownAddress common.Address,
	deployFunc func(*bind.TransactOpts, bind.ContractBackend) (common.Address, *types.Transaction, *T, error),
	instantiateFunc func(common.Address, bind.ContractBackend) (*T, error),
	callFunc func(*T) error,
) (address common.Address, contract *T, err error) {
	if knownAddress == (common.Address{}) {
		// Deploy the contract if known address is empty.
		address, _, contract, err = deployFunc(tops, c)
		if err != nil {
			log.Error().Err(err).Msg(fmt.Sprintf("Unable to deploy %s contract", name))
			return
		}
		log.Trace().Interface("address", address).Msg(fmt.Sprintf("%s contract deployed", name))
	} else {
		// Otherwise, instantiate the contract.
		address = knownAddress
		contract, err = instantiateFunc(address, c)
		if err != nil {
			log.Error().Err(err).Msg(fmt.Sprintf("Unable to instantiate %s contract", name))
			return
		}
		log.Trace().Msg(fmt.Sprintf("%s contract instantiated", name))
	}

	// Check that the contract is deployed and ready.
	if err = blockUntilSuccessful(ctx, c, func() error {
		log.Trace().Msg(fmt.Sprintf("%s contract is not deployed yet", name))
		return callFunc(contract)
	}); err != nil {
		return
	}
	return
}

func enableOneBPFeeTier(contract *uniswapv3.UniswapV3Factory, tops *bind.TransactOpts, fee, tickSpacing int64) error {
	if _, err := contract.EnableFeeAmount(tops, big.NewInt(fee), big.NewInt(tickSpacing)); err != nil {
		log.Error().Err(err).Msg("Unable to enable one basic point fee tier")
		return err
	}
	log.Trace().Msg("Enable one basic point fee tier")
	return nil
}

func setFactoryOwner(contract *uniswapv3.UniswapV3Factory, tops *bind.TransactOpts, newOwner common.Address) error {
	if _, err := contract.SetOwner(tops, newOwner); err != nil {
		log.Error().Err(err).Msg("Unable to set new owner for Factory contract")
		return err
	}
	log.Trace().Msg("Set new owner for Factory contract")
	return nil
}

func transferProxyAdminOwnership(contract *uniswapv3.ProxyAdmin, tops *bind.TransactOpts, newOwner common.Address) error {
	if _, err := contract.TransferOwnership(tops, newOwner); err != nil {
		log.Error().Err(err).Msg("Unable to transfer ProxyAdmin ownership")
		return err
	}
	log.Trace().Msg("Transfer ProxyAdmin ownership")
	return nil
}

func deployERC20Pair(ctx context.Context, c *ethclient.Client, tops *bind.TransactOpts, cops *bind.CallOpts, config UniswapV3Config, tokenAKnownAddress, tokenBKnownAddress common.Address) (contractConfig[uniswapv3.ERC20], contractConfig[uniswapv3.ERC20], error) {
	tokensToMint := big.NewInt(1_000_000_000_000_000_000)

	var token0, token1 contractConfig[uniswapv3.ERC20]
	var err error
	token0.Address, token0.contract, err = deployOrInstantiateContract(
		ctx, c, tops, cops, "TokenA", tokenAKnownAddress,
		func(*bind.TransactOpts, bind.ContractBackend) (common.Address, *types.Transaction, *uniswapv3.ERC20, error) {
			return uniswapv3.DeployERC20(tops, c, "TokenA", "A")
		},
		uniswapv3.NewERC20,
		func(contract *uniswapv3.ERC20) error {
			return approveERC20SpendingsByUniswap(contract, tops, config.NonfungiblePositionManager.Address, config.SwapRouter02.Address, tokensToMint)
		},
	)
	if err != nil {
		return token0, token1, err
	}
	token1.Address, token1.contract, err = deployOrInstantiateContract(
		ctx, c, tops, cops, "TokenB", tokenBKnownAddress,
		func(*bind.TransactOpts, bind.ContractBackend) (common.Address, *types.Transaction, *uniswapv3.ERC20, error) {
			return uniswapv3.DeployERC20(tops, c, "TokenB", "B")
		},
		uniswapv3.NewERC20,
		func(contract *uniswapv3.ERC20) (err error) {
			return approveERC20SpendingsByUniswap(contract, tops, config.NonfungiblePositionManager.Address, config.SwapRouter02.Address, tokensToMint)
		},
	)
	if err != nil {
		return token0, token1, err
	}
	return token0, token1, nil
}

func approveERC20SpendingsByUniswap(contract *uniswapv3.ERC20, tops *bind.TransactOpts, NonfungiblePositionManagerAddress, SwapRouter02Address common.Address, amount *big.Int) error {
	_, err := contract.Approve(tops, NonfungiblePositionManagerAddress, amount)
	if err != nil {
		log.Trace().Msg("Unable to approve NonfungiblePositionManagerAddress spendings")
		return err
	}

	_, err = contract.Approve(tops, SwapRouter02Address, amount)
	if err != nil {
		log.Trace().Msg("Unable to approve SwapRouter02Address spendings")
		return err
	}

	log.Trace().Msg("Spendings approved for both NonfungiblePositionManagerAddress and SwapRouter02Address")
	return nil
}

func createPool(ctx context.Context, c *ethclient.Client, tops *bind.TransactOpts, cops *bind.CallOpts, config UniswapV3Config, tokenA, tokenB contractConfig[uniswapv3.ERC20], fees *big.Int, recipient common.Address) error {
	// Create a pool between the ERC20 contracts.
	_, err := config.FactoryV3.contract.CreatePool(tops, tokenA.Address, tokenB.Address, fees)
	if err != nil {
		log.Error().Err(err).Msg("Unable to create a TokenA-TokenB pool")
		return err
	}

	// Retrieve the pool address.
	var poolAddress common.Address
	if err = blockUntilSuccessful(ctx, c, func() (err error) {
		poolAddress, err = config.FactoryV3.contract.GetPool(cops, tokenA.Address, tokenB.Address, fees)
		if poolAddress == (common.Address{}) {
			return fmt.Errorf("the TokenA-TokenB pool address is not deployed yet")
		}

		return
	}); err != nil {
		log.Error().Err(err).Msg("Unable to get the address of the TokenA-TokenB pool")
		return err
	}
	log.Trace().Interface("address", poolAddress).Msg("New TokenA-TokenB pool created")

	// Initialize the pool.
	var poolContract *uniswapv3.UniswapV3Pool
	poolContract, err = uniswapv3.NewUniswapV3Pool(poolAddress, c)
	if err != nil {
		log.Error().Err(err).Msg("Unable to initialize the TokenA-TokenB pool contract")
		return err
	}

	// To compute this value, we set that 1 TokenB is worth 500 TokenA.
	// Then we use the handy script under `contracts/uniswapv3/helper.py`.
	// $ python3 helper.py 1 500
	// Current price: 500
	// Current price (Q64.96): 1771595571142957189036318392320
	// Tick index: 62149
	// Source: https://uniswapv3book.com/docs/milestone_1/calculating-liquidity/
	sqrtPriceX96 := new(big.Int)
	sqrtPriceX96.SetString("1771595571142957189036318392320", 10)
	if err = blockUntilSuccessful(ctx, c, func() (err error) {
		_, err = poolContract.Initialize(tops, sqrtPriceX96)
		return
	}); err != nil {
		log.Error().Err(err).Msg("Unable to initialize the TokenA-TokenB pool")
		return err
	}
	log.Trace().Msg("TokenA-TokenB pool initialized")

	// Provide liquidity.
	var blockNumber uint64
	blockNumber, err = c.BlockNumber(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Unable to get the latest block number")
		return err
	}

	var block *types.Block
	block, err = c.BlockByNumber(ctx, big.NewInt(int64(blockNumber)))
	if err != nil {
		log.Error().Err(err).Msg("Unable to get the latest block")
		return err
	}
	timestamp := int64(block.Time())

	// TODO: Understand why this call reverts.
	if _, err = config.NonfungiblePositionManager.contract.Mint(tops, uniswapv3.INonfungiblePositionManagerMintParams{
		Token0:         tokenA.Address,
		Token1:         tokenB.Address,
		Fee:            fees,
		TickLower:      big.NewInt(MIN_TICK), // We provide liquidity across the whole possible range.
		TickUpper:      big.NewInt(MAX_TICK),
		Amount0Desired: big.NewInt(1_000),
		Amount1Desired: big.NewInt(500_000),
		Amount0Min:     big.NewInt(0), // We mint without any slippage protection. Don't do this in production!
		Amount1Min:     big.NewInt(0), // Same thing here.
		Recipient:      recipient,
		Deadline:       big.NewInt(timestamp + 10), // 10 seconds
	}); err != nil {
		log.Error().Err(err).Msg("Unable to create liquidity for the TokenA-TokenB pool")
		return err
	}
	log.Trace().Msg("Liquidity provided to the TokenA-TokenB pool")
	return nil
}

func swapTokenBForTokenA(tops *bind.TransactOpts, config UniswapV3Config, tokenA, tokenB contractConfig[uniswapv3.ERC20], fees *big.Int, recipient common.Address) error {
	if _, err := config.SwapRouter02.contract.ExactInputSingle(tops, uniswapv3.IV3SwapRouterExactInputSingleParams{
		TokenIn:           tokenB.Address,
		TokenOut:          tokenA.Address,
		Fee:               fees,
		Recipient:         recipient,
		AmountIn:          big.NewInt(1000),
		AmountOutMinimum:  big.NewInt(0),
		SqrtPriceLimitX96: big.NewInt(0),
	}); err != nil {
		log.Error().Err(err).Msg("Unable to swap TokenB for TokenA")
		return err
	}
	log.Trace().Msg("Swapped TokenB for TokenA")
	return nil
}
