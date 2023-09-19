package loadtest

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
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
)

type uniswapV3Contract[T uniswapv3.UniswapV3Factory | uniswapv3.UniswapInterfaceMulticall | uniswapv3.ProxyAdmin] struct {
	Address  ethcommon.Address
	Contract *T
}

type UniswapV3Config struct {
	Factory    uniswapV3Contract[uniswapv3.UniswapV3Factory]
	Multicall  uniswapV3Contract[uniswapv3.UniswapInterfaceMulticall]
	ProxyAdmin uniswapV3Contract[uniswapv3.ProxyAdmin]
}

func deployUniswapV3(ctx context.Context, c *ethclient.Client, tops *bind.TransactOpts) (UniswapV3Config, error) {
	var config UniswapV3Config
	var err error

	config.Factory.Address, config.Factory.Contract, err = deployFactory(c, tops)
	if err != nil {
		return UniswapV3Config{}, err
	}

	err = enableOneBPFeeTier(config.Factory.Contract, tops, ONE_BP_FEE, ONE_BP_TICK_SPACING)
	if err != nil {
		return UniswapV3Config{}, err
	}

	config.Multicall.Address, config.Multicall.Contract, err = deployMulticall(c, tops)
	if err != nil {
		return UniswapV3Config{}, err
	}

	config.ProxyAdmin.Address, config.ProxyAdmin.Contract, err = deployProxyAdmin(c, tops)
	if err != nil {
		return UniswapV3Config{}, err
	}

	return config, nil
}

// https://github.com/Uniswap/deploy-v3/blob/b7aac0f1c5353b36802dc0cf95c426d2ef0c3252/src/steps/deploy-v3-core-factory.ts
// https://github.com/Uniswap/v3-core/blob/d8b1c635c275d2a9450bd6a78f3fa2484fef73eb/contracts/UniswapV3Factory.sol
func deployFactory(c *ethclient.Client, tops *bind.TransactOpts) (ethcommon.Address, *uniswapv3.UniswapV3Factory, error) {
	address, _, _, err := uniswapv3.DeployUniswapV3Factory(tops, c)
	if err != nil {
		log.Error().Err(err).Msg("Unable to deploy UniswapV3Factory contract")
		return ethcommon.Address{}, nil, err
	}
	log.Trace().Interface("address", address).Msg("UniswapV3Factory contract address")

	contract, err := uniswapv3.NewUniswapV3Factory(address, c)
	if err != nil {
		log.Error().Err(err).Msg("Unable to instantiate UniswapV3Factory contract")
		return ethcommon.Address{}, nil, err
	}
	return address, contract, nil
}

// https://github.com/Uniswap/deploy-v3/blob/b7aac0f1c5353b36802dc0cf95c426d2ef0c3252/src/steps/add-1bp-fee-tier.ts
func enableOneBPFeeTier(contract *uniswapv3.UniswapV3Factory, tops *bind.TransactOpts, fee, tickSpacing int64) error {
	if _, err := contract.EnableFeeAmount(tops, big.NewInt(fee), big.NewInt(tickSpacing)); err != nil {
		return err
	}
	log.Trace().Msg("Enable a one basic point fee tier")
	return nil
}

// https://github.com/Uniswap/deploy-v3/blob/b7aac0f1c5353b36802dc0cf95c426d2ef0c3252/src/steps/deploy-multicall2.ts
// https://github.com/Uniswap/v3-periphery/blob/b13f9d9d9868b98b765c4f1d8d7f486b00b22488/contracts/lens/UniswapInterfaceMulticall.sol
func deployMulticall(c *ethclient.Client, tops *bind.TransactOpts) (ethcommon.Address, *uniswapv3.UniswapInterfaceMulticall, error) {
	address, _, _, err := uniswapv3.DeployUniswapInterfaceMulticall(tops, c)
	if err != nil {
		log.Error().Err(err).Msg("Unable to deploy UniswapInterfaceMulticall contract")
		return ethcommon.Address{}, nil, err
	}
	log.Trace().Interface("address", address).Msg("UniswapInterfaceMulticall contract address")

	contract, err := uniswapv3.NewUniswapInterfaceMulticall(address, c)
	if err != nil {
		log.Error().Err(err).Msg("Unable to instantiate UniswapInterfaceMulticall contract")
		return ethcommon.Address{}, nil, err
	}
	return address, contract, nil
}

func deployProxyAdmin(c *ethclient.Client, tops *bind.TransactOpts) (ethcommon.Address, *uniswapv3.ProxyAdmin, error) {
	address, _, _, err := uniswapv3.DeployProxyAdmin(tops, c)
	if err != nil {
		log.Error().Err(err).Msg("Unable to deploy ProxyAdmin contract")
		return ethcommon.Address{}, nil, err
	}
	log.Trace().Interface("address", address).Msg("ProxyAdmin contract address")

	contract, err := uniswapv3.NewProxyAdmin(address, c)
	if err != nil {
		log.Error().Err(err).Msg("Unable to instantiate ProxyAdmin contract")
		return ethcommon.Address{}, nil, err
	}
	return address, contract, nil
}

// Create and initialise an ERC20 pool between two ERC20 contracts.
// Note that this will also deploy both ERC20 contracts.
func createPool() {
	// TODO
}

func swapTokenAForTokenB() {
	// TODO
}

func swapTokenBForTokenA() {
	// TODO
}
