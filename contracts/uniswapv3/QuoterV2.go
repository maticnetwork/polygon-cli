// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package uniswapv3

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// IQuoterV2QuoteExactInputSingleParams is an auto generated low-level Go binding around an user-defined struct.
type IQuoterV2QuoteExactInputSingleParams struct {
	TokenIn           common.Address
	TokenOut          common.Address
	AmountIn          *big.Int
	Fee               *big.Int
	SqrtPriceLimitX96 *big.Int
}

// IQuoterV2QuoteExactOutputSingleParams is an auto generated low-level Go binding around an user-defined struct.
type IQuoterV2QuoteExactOutputSingleParams struct {
	TokenIn           common.Address
	TokenOut          common.Address
	Amount            *big.Int
	Fee               *big.Int
	SqrtPriceLimitX96 *big.Int
}

// QuoterV2MetaData contains all meta data concerning the QuoterV2 contract.
var QuoterV2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_WETH9\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"WETH9\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"name\":\"quoteExactInput\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint160[]\",\"name\":\"sqrtPriceX96AfterList\",\"type\":\"uint160[]\"},{\"internalType\":\"uint32[]\",\"name\":\"initializedTicksCrossedList\",\"type\":\"uint32[]\"},{\"internalType\":\"uint256\",\"name\":\"gasEstimate\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\"}],\"internalType\":\"structIQuoterV2.QuoteExactInputSingleParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"quoteExactInputSingle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceX96After\",\"type\":\"uint160\"},{\"internalType\":\"uint32\",\"name\":\"initializedTicksCrossed\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"gasEstimate\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"name\":\"quoteExactOutput\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint160[]\",\"name\":\"sqrtPriceX96AfterList\",\"type\":\"uint160[]\"},{\"internalType\":\"uint32[]\",\"name\":\"initializedTicksCrossedList\",\"type\":\"uint32[]\"},{\"internalType\":\"uint256\",\"name\":\"gasEstimate\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\"}],\"internalType\":\"structIQuoterV2.QuoteExactOutputSingleParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"quoteExactOutputSingle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceX96After\",\"type\":\"uint160\"},{\"internalType\":\"uint32\",\"name\":\"initializedTicksCrossed\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"gasEstimate\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"amount0Delta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"amount1Delta\",\"type\":\"int256\"},{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"}],\"name\":\"uniswapV3SwapCallback\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60c06040523480156200001157600080fd5b5060405162001bf238038062001bf2833981016040819052620000349162000070565b6001600160601b0319606092831b8116608052911b1660a052620000a7565b80516001600160a01b03811681146200006b57600080fd5b919050565b6000806040838503121562000083578182fd5b6200008e8362000053565b91506200009e6020840162000053565b90509250929050565b60805160601c60a05160601c611b17620000db600039806102e05250806104d7528061083752806109ef5250611b176000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c8063c45a01551161005b578063c45a0155146100e6578063c6a5026a146100ee578063cdca175314610101578063fa461e33146101145761007d565b80632f80bb1d146100825780634aa4a4fc146100ae578063bd21704a146100c3575b600080fd5b61009561009036600461167c565b610129565b6040516100a5949392919061198e565b60405180910390f35b6100b66102de565b6040516100a591906118f7565b6100d66100d136600461179a565b610302565b6040516100a59493929190611a29565b6100b66104d5565b6100d66100fc36600461179a565b6104f9565b61009561010f36600461167c565b61066f565b6101276101223660046116e2565b610806565b005b6000606080600061013986610981565b67ffffffffffffffff8111801561014f57600080fd5b50604051908082528060200260200182016040528015610179578160200160208202803683370190505b50925061018586610981565b67ffffffffffffffff8111801561019b57600080fd5b506040519080825280602002602001820160405280156101c5578160200160208202803683370190505b50915060005b60008060006101d98a610992565b9250925092506000806000806102356040518060a00160405280886001600160a01b03168152602001896001600160a01b031681526020018f81526020018762ffffff16815260200160006001600160a01b0316815250610302565b9350935093509350828b898151811061024a57fe5b60200260200101906001600160a01b031690816001600160a01b031681525050818a898151811061027757fe5b63ffffffff90921660209283029190910190910152929b50968201966001909601958b926102a48e6109c3565b156102b9576102b28e6109cb565b9d506102c9565b8c9b5050505050505050506102d5565b505050505050506101cb565b92959194509250565b7f000000000000000000000000000000000000000000000000000000000000000081565b6020810151815160608301516000928392839283926001600160a01b038082169084161092849261033392906109e8565b905086608001516001600160a01b0316600014156103545760408701516000555b60005a9050816001600160a01b031663128acb0830856103778c60400151610a26565b6000038c608001516001600160a01b0316600014610399578c608001516103bf565b876103b85773fffd8963efd1fc6a506488495d951d5263988d256103bf565b6401000276a45b8d602001518e606001518f600001516040516020016103e0939291906118bc565b6040516020818303038152906040526040518663ffffffff1660e01b815260040161040f95949392919061190b565b6040805180830381600087803b15801561042857600080fd5b505af1925050508015610458575060408051601f3d908101601f19168201909252610455918101906116bf565b60015b6104c8573d808015610486576040519150601f19603f3d011682016040523d82523d6000602084013e61048b565b606091505b505a8203945088608001516001600160a01b0316600014156104ac57600080555b6104b7818487610a3c565b9750975097509750505050506104ce565b50505050505b9193509193565b7f000000000000000000000000000000000000000000000000000000000000000081565b6020810151815160608301516000928392839283926001600160a01b038082169084161092849261052a92906109e8565b905060005a9050816001600160a01b031663128acb08308561054f8c60400151610a26565b60808d01516001600160a01b03161561056c578c60800151610592565b8761058b5773fffd8963efd1fc6a506488495d951d5263988d25610592565b6401000276a45b8d600001518e606001518f602001516040516020016105b3939291906118bc565b6040516020818303038152906040526040518663ffffffff1660e01b81526004016105e295949392919061190b565b6040805180830381600087803b1580156105fb57600080fd5b505af192505050801561062b575060408051601f3d908101601f19168201909252610628918101906116bf565b60015b6104c8573d808015610659576040519150601f19603f3d011682016040523d82523d6000602084013e61065e565b606091505b505a820394506104b7818487610a3c565b6000606080600061067f86610981565b67ffffffffffffffff8111801561069557600080fd5b506040519080825280602002602001820160405280156106bf578160200160208202803683370190505b5092506106cb86610981565b67ffffffffffffffff811180156106e157600080fd5b5060405190808252806020026020018201604052801561070b578160200160208202803683370190505b50915060005b600080600061071f8a610992565b92509250925060008060008061077b6040518060a00160405280896001600160a01b03168152602001886001600160a01b031681526020018f81526020018762ffffff16815260200160006001600160a01b03168152506104f9565b9350935093509350828b898151811061079057fe5b60200260200101906001600160a01b031690816001600160a01b031681525050818a89815181106107bd57fe5b63ffffffff90921660209283029190910190910152929b50968201966001909601958b926107ea8e6109c3565b156102b9576107f88e6109cb565b9d5050505050505050610711565b60008313806108155750600082135b61081e57600080fd5b600080600061082c84610992565b92509250925061085e7f0000000000000000000000000000000000000000000000000000000000000000848484610af6565b50600080600080891361088a57856001600160a01b0316856001600160a01b031610888a6000036108a5565b846001600160a01b0316866001600160a01b03161089896000035b92509250925060006108b88787876109e8565b9050600080826001600160a01b0316633850c7bd6040518163ffffffff1660e01b815260040160e06040518083038186803b1580156108f657600080fd5b505afa15801561090a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061092e91906117bc565b505050505091509150851561095457604051848152826020820152816040820152606081fd5b6000541561096a57600054841461096a57600080fd5b604051858152826020820152816040820152606081fd5b80516017601319909101045b919050565b600080806109a08482610b15565b92506109ad846014610bc5565b90506109ba846017610b15565b91509193909250565b516042111590565b80516060906109e290839060179060161901610c6c565b92915050565b6000610a1e7f0000000000000000000000000000000000000000000000000000000000000000610a19868686610dbd565b610e13565b949350505050565b6000600160ff1b8210610a3857600080fd5b5090565b600080600080600080876001600160a01b0316633850c7bd6040518163ffffffff1660e01b815260040160e06040518083038186803b158015610a7e57600080fd5b505afa158015610a92573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ab691906117bc565b50939650610acb94508d9350610ef792505050565b91975095509050610ae66001600160a01b0389168383610f84565b9350869250505093509350935093565b6000610b0c85610b07868686610dbd565b61150d565b95945050505050565b600081826014011015610b64576040805162461bcd60e51b8152602060048201526012602482015271746f416464726573735f6f766572666c6f7760701b604482015290519081900360640190fd5b8160140183511015610bb5576040805162461bcd60e51b8152602060048201526015602482015274746f416464726573735f6f75744f66426f756e647360581b604482015290519081900360640190fd5b500160200151600160601b900490565b600081826003011015610c13576040805162461bcd60e51b8152602060048201526011602482015270746f55696e7432345f6f766572666c6f7760781b604482015290519081900360640190fd5b8160030183511015610c63576040805162461bcd60e51b8152602060048201526014602482015273746f55696e7432345f6f75744f66426f756e647360601b604482015290519081900360640190fd5b50016003015190565b60608182601f011015610cb7576040805162461bcd60e51b815260206004820152600e60248201526d736c6963655f6f766572666c6f7760901b604482015290519081900360640190fd5b828284011015610cff576040805162461bcd60e51b815260206004820152600e60248201526d736c6963655f6f766572666c6f7760901b604482015290519081900360640190fd5b81830184511015610d4b576040805162461bcd60e51b8152602060048201526011602482015270736c6963655f6f75744f66426f756e647360781b604482015290519081900360640190fd5b606082158015610d6a5760405191506000825260208201604052610db4565b6040519150601f8416801560200281840101858101878315602002848b0101015b81831015610da3578051835260209283019201610d8b565b5050858452601f01601f1916604052505b50949350505050565b610dc561154b565b826001600160a01b0316846001600160a01b03161115610de3579192915b50604080516060810182526001600160a01b03948516815292909316602083015262ffffff169181019190915290565b600081602001516001600160a01b031682600001516001600160a01b031610610e3b57600080fd5b50805160208083015160409384015184516001600160a01b0394851681850152939091168385015262ffffff166060808401919091528351808403820181526080840185528051908301206001600160f81b031960a085015294901b6bffffffffffffffffffffffff191660a183015260b58201939093527fe34f199b19b2b4f47f68442619d555527d244f78a3297ea89325f843f87b8b5460d5808301919091528251808303909101815260f5909101909152805191012090565b60008060008351606014610f6357604484511015610f305760405162461bcd60e51b8152600401610f2790611964565b60405180910390fd5b60048401935083806020019051810190610f4a9190611730565b60405162461bcd60e51b8152600401610f279190611951565b83806020019051810190610f779190611853565b9250925092509193909250565b60008060008060008060008060088b6001600160a01b031663d0c93a7c6040518163ffffffff1660e01b815260040160206040518083038186803b158015610fcb57600080fd5b505afa158015610fdf573d6000803e3d6000fd5b505050506040513d6020811015610ff557600080fd5b5051600290810b908c900b8161100757fe5b0560020b901d905060006101008c6001600160a01b031663d0c93a7c6040518163ffffffff1660e01b815260040160206040518083038186803b15801561104d57600080fd5b505afa158015611061573d6000803e3d6000fd5b505050506040513d602081101561107757600080fd5b5051600290810b908d900b8161108957fe5b0560020b8161109457fe5b079050600060088d6001600160a01b031663d0c93a7c6040518163ffffffff1660e01b815260040160206040518083038186803b1580156110d457600080fd5b505afa1580156110e8573d6000803e3d6000fd5b505050506040513d60208110156110fe57600080fd5b5051600290810b908d900b8161111057fe5b0560020b901d905060006101008e6001600160a01b031663d0c93a7c6040518163ffffffff1660e01b815260040160206040518083038186803b15801561115657600080fd5b505afa15801561116a573d6000803e3d6000fd5b505050506040513d602081101561118057600080fd5b5051600290810b908e900b8161119257fe5b0560020b8161119d57fe5b07905060008160ff166001901b8f6001600160a01b0316635339c296856040518263ffffffff1660e01b8152600401808260010b815260200191505060206040518083038186803b1580156111f157600080fd5b505afa158015611205573d6000803e3d6000fd5b505050506040513d602081101561121b57600080fd5b5051161180156112a157508d6001600160a01b031663d0c93a7c6040518163ffffffff1660e01b815260040160206040518083038186803b15801561125f57600080fd5b505afa158015611273573d6000803e3d6000fd5b505050506040513d602081101561128957600080fd5b5051600290810b908d900b8161129b57fe5b0760020b155b80156112b257508b60020b8d60020b135b945060008360ff166001901b8f6001600160a01b0316635339c296876040518263ffffffff1660e01b8152600401808260010b815260200191505060206040518083038186803b15801561130557600080fd5b505afa158015611319573d6000803e3d6000fd5b505050506040513d602081101561132f57600080fd5b5051161180156113b557508d6001600160a01b031663d0c93a7c6040518163ffffffff1660e01b815260040160206040518083038186803b15801561137357600080fd5b505afa158015611387573d6000803e3d6000fd5b505050506040513d602081101561139d57600080fd5b5051600290810b908e900b816113af57fe5b0760020b155b80156113c657508b60020b8d60020b125b95508160010b8460010b12806113f257508160010b8460010b1480156113f257508060ff168360ff1611155b1561140857839950829750819850809650611415565b8199508097508398508296505b505060001960ff87161b9150505b8560010b8760010b136114e5578560010b8760010b141561144a5760001960ff858103161c165b6000818c6001600160a01b0316635339c2968a6040518263ffffffff1660e01b8152600401808260010b815260200191505060206040518083038186803b15801561149457600080fd5b505afa1580156114a8573d6000803e3d6000fd5b505050506040513d60208110156114be57600080fd5b50511690506114cc81611530565b61ffff1698909801975050600190950194600019611423565b81156114f2576001880397505b82156114ff576001880397505b505050505050509392505050565b60006115198383610e13565b9050336001600160a01b038216146109e257600080fd5b6000805b82156109e257600019830190921691600101611534565b604080516060810182526000808252602082018190529181019190915290565b600082601f83011261157b578081fd5b813561158e61158982611a77565b611a53565b8181528460208386010111156115a2578283fd5b816020850160208301379081016020019190915292915050565b8051600281900b811461098d57600080fd5b600060a082840312156115df578081fd5b60405160a0810181811067ffffffffffffffff821117156115fc57fe5b604052905080823561160d81611ac9565b8152602083013561161d81611ac9565b602082015260408381013590820152606083013562ffffff8116811461164257600080fd5b60608201526116536080840161165f565b60808201525092915050565b803561098d81611ac9565b805161ffff8116811461098d57600080fd5b6000806040838503121561168e578182fd5b823567ffffffffffffffff8111156116a4578283fd5b6116b08582860161156b565b95602094909401359450505050565b600080604083850312156116d1578182fd5b505080516020909101519092909150565b6000806000606084860312156116f6578081fd5b8335925060208401359150604084013567ffffffffffffffff81111561171a578182fd5b6117268682870161156b565b9150509250925092565b600060208284031215611741578081fd5b815167ffffffffffffffff811115611757578182fd5b8201601f81018413611767578182fd5b805161177561158982611a77565b818152856020838501011115611789578384fd5b610b0c826020830160208601611a99565b600060a082840312156117ab578081fd5b6117b583836115ce565b9392505050565b600080600080600080600060e0888a0312156117d6578283fd5b87516117e181611ac9565b96506117ef602089016115bc565b95506117fd6040890161166a565b945061180b6060890161166a565b93506118196080890161166a565b925060a088015160ff8116811461182e578283fd5b60c08901519092508015158114611843578182fd5b8091505092959891949750929550565b600080600060608486031215611867578081fd5b83519250602084015161187981611ac9565b9150611887604085016115bc565b90509250925092565b600081518084526118a8816020860160208601611a99565b601f01601f19169290920160200192915050565b606093841b6bffffffffffffffffffffffff19908116825260e89390931b6001600160e81b0319166014820152921b166017820152602b0190565b6001600160a01b0391909116815260200190565b6001600160a01b0386811682528515156020830152604082018590528316606082015260a06080820181905260009061194690830184611890565b979650505050505050565b6000602082526117b56020830184611890565b60208082526010908201526f2ab732bc3832b1ba32b21032b93937b960811b604082015260600190565b600060808201868352602060808185015281875180845260a0860191508289019350845b818110156119d75784516001600160a01b0316835293830193918301916001016119b2565b505084810360408601528651808252908201925081870190845b81811015611a1357825163ffffffff16855293830193918301916001016119f1565b5050505060609290920192909252949350505050565b9384526001600160a01b0392909216602084015263ffffffff166040830152606082015260800190565b60405181810167ffffffffffffffff81118282101715611a6f57fe5b604052919050565b600067ffffffffffffffff821115611a8b57fe5b50601f01601f191660200190565b60005b83811015611ab4578181015183820152602001611a9c565b83811115611ac3576000848401525b50505050565b6001600160a01b0381168114611ade57600080fd5b5056fea264697066735822122072d2a676397b0beba9ceb73476223f85e7f8e17d636ce0c47f29f5d3b1eacbc864736f6c63430007060033",
}

// QuoterV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use QuoterV2MetaData.ABI instead.
var QuoterV2ABI = QuoterV2MetaData.ABI

// QuoterV2Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use QuoterV2MetaData.Bin instead.
var QuoterV2Bin = QuoterV2MetaData.Bin

// DeployQuoterV2 deploys a new Ethereum contract, binding an instance of QuoterV2 to it.
func DeployQuoterV2(auth *bind.TransactOpts, backend bind.ContractBackend, _factory common.Address, _WETH9 common.Address) (common.Address, *types.Transaction, *QuoterV2, error) {
	parsed, err := QuoterV2MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(QuoterV2Bin), backend, _factory, _WETH9)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &QuoterV2{QuoterV2Caller: QuoterV2Caller{contract: contract}, QuoterV2Transactor: QuoterV2Transactor{contract: contract}, QuoterV2Filterer: QuoterV2Filterer{contract: contract}}, nil
}

// QuoterV2 is an auto generated Go binding around an Ethereum contract.
type QuoterV2 struct {
	QuoterV2Caller     // Read-only binding to the contract
	QuoterV2Transactor // Write-only binding to the contract
	QuoterV2Filterer   // Log filterer for contract events
}

// QuoterV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type QuoterV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QuoterV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type QuoterV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QuoterV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type QuoterV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QuoterV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type QuoterV2Session struct {
	Contract     *QuoterV2         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// QuoterV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type QuoterV2CallerSession struct {
	Contract *QuoterV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// QuoterV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type QuoterV2TransactorSession struct {
	Contract     *QuoterV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// QuoterV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type QuoterV2Raw struct {
	Contract *QuoterV2 // Generic contract binding to access the raw methods on
}

// QuoterV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type QuoterV2CallerRaw struct {
	Contract *QuoterV2Caller // Generic read-only contract binding to access the raw methods on
}

// QuoterV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type QuoterV2TransactorRaw struct {
	Contract *QuoterV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewQuoterV2 creates a new instance of QuoterV2, bound to a specific deployed contract.
func NewQuoterV2(address common.Address, backend bind.ContractBackend) (*QuoterV2, error) {
	contract, err := bindQuoterV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &QuoterV2{QuoterV2Caller: QuoterV2Caller{contract: contract}, QuoterV2Transactor: QuoterV2Transactor{contract: contract}, QuoterV2Filterer: QuoterV2Filterer{contract: contract}}, nil
}

// NewQuoterV2Caller creates a new read-only instance of QuoterV2, bound to a specific deployed contract.
func NewQuoterV2Caller(address common.Address, caller bind.ContractCaller) (*QuoterV2Caller, error) {
	contract, err := bindQuoterV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &QuoterV2Caller{contract: contract}, nil
}

// NewQuoterV2Transactor creates a new write-only instance of QuoterV2, bound to a specific deployed contract.
func NewQuoterV2Transactor(address common.Address, transactor bind.ContractTransactor) (*QuoterV2Transactor, error) {
	contract, err := bindQuoterV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &QuoterV2Transactor{contract: contract}, nil
}

// NewQuoterV2Filterer creates a new log filterer instance of QuoterV2, bound to a specific deployed contract.
func NewQuoterV2Filterer(address common.Address, filterer bind.ContractFilterer) (*QuoterV2Filterer, error) {
	contract, err := bindQuoterV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &QuoterV2Filterer{contract: contract}, nil
}

// bindQuoterV2 binds a generic wrapper to an already deployed contract.
func bindQuoterV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := QuoterV2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_QuoterV2 *QuoterV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _QuoterV2.Contract.QuoterV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_QuoterV2 *QuoterV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _QuoterV2.Contract.QuoterV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_QuoterV2 *QuoterV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _QuoterV2.Contract.QuoterV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_QuoterV2 *QuoterV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _QuoterV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_QuoterV2 *QuoterV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _QuoterV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_QuoterV2 *QuoterV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _QuoterV2.Contract.contract.Transact(opts, method, params...)
}

// WETH9 is a free data retrieval call binding the contract method 0x4aa4a4fc.
//
// Solidity: function WETH9() view returns(address)
func (_QuoterV2 *QuoterV2Caller) WETH9(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _QuoterV2.contract.Call(opts, &out, "WETH9")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH9 is a free data retrieval call binding the contract method 0x4aa4a4fc.
//
// Solidity: function WETH9() view returns(address)
func (_QuoterV2 *QuoterV2Session) WETH9() (common.Address, error) {
	return _QuoterV2.Contract.WETH9(&_QuoterV2.CallOpts)
}

// WETH9 is a free data retrieval call binding the contract method 0x4aa4a4fc.
//
// Solidity: function WETH9() view returns(address)
func (_QuoterV2 *QuoterV2CallerSession) WETH9() (common.Address, error) {
	return _QuoterV2.Contract.WETH9(&_QuoterV2.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_QuoterV2 *QuoterV2Caller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _QuoterV2.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_QuoterV2 *QuoterV2Session) Factory() (common.Address, error) {
	return _QuoterV2.Contract.Factory(&_QuoterV2.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_QuoterV2 *QuoterV2CallerSession) Factory() (common.Address, error) {
	return _QuoterV2.Contract.Factory(&_QuoterV2.CallOpts)
}

// UniswapV3SwapCallback is a free data retrieval call binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes path) view returns()
func (_QuoterV2 *QuoterV2Caller) UniswapV3SwapCallback(opts *bind.CallOpts, amount0Delta *big.Int, amount1Delta *big.Int, path []byte) error {
	var out []interface{}
	err := _QuoterV2.contract.Call(opts, &out, "uniswapV3SwapCallback", amount0Delta, amount1Delta, path)

	if err != nil {
		return err
	}

	return err

}

// UniswapV3SwapCallback is a free data retrieval call binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes path) view returns()
func (_QuoterV2 *QuoterV2Session) UniswapV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, path []byte) error {
	return _QuoterV2.Contract.UniswapV3SwapCallback(&_QuoterV2.CallOpts, amount0Delta, amount1Delta, path)
}

// UniswapV3SwapCallback is a free data retrieval call binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes path) view returns()
func (_QuoterV2 *QuoterV2CallerSession) UniswapV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, path []byte) error {
	return _QuoterV2.Contract.UniswapV3SwapCallback(&_QuoterV2.CallOpts, amount0Delta, amount1Delta, path)
}

// QuoteExactInput is a paid mutator transaction binding the contract method 0xcdca1753.
//
// Solidity: function quoteExactInput(bytes path, uint256 amountIn) returns(uint256 amountOut, uint160[] sqrtPriceX96AfterList, uint32[] initializedTicksCrossedList, uint256 gasEstimate)
func (_QuoterV2 *QuoterV2Transactor) QuoteExactInput(opts *bind.TransactOpts, path []byte, amountIn *big.Int) (*types.Transaction, error) {
	return _QuoterV2.contract.Transact(opts, "quoteExactInput", path, amountIn)
}

// QuoteExactInput is a paid mutator transaction binding the contract method 0xcdca1753.
//
// Solidity: function quoteExactInput(bytes path, uint256 amountIn) returns(uint256 amountOut, uint160[] sqrtPriceX96AfterList, uint32[] initializedTicksCrossedList, uint256 gasEstimate)
func (_QuoterV2 *QuoterV2Session) QuoteExactInput(path []byte, amountIn *big.Int) (*types.Transaction, error) {
	return _QuoterV2.Contract.QuoteExactInput(&_QuoterV2.TransactOpts, path, amountIn)
}

// QuoteExactInput is a paid mutator transaction binding the contract method 0xcdca1753.
//
// Solidity: function quoteExactInput(bytes path, uint256 amountIn) returns(uint256 amountOut, uint160[] sqrtPriceX96AfterList, uint32[] initializedTicksCrossedList, uint256 gasEstimate)
func (_QuoterV2 *QuoterV2TransactorSession) QuoteExactInput(path []byte, amountIn *big.Int) (*types.Transaction, error) {
	return _QuoterV2.Contract.QuoteExactInput(&_QuoterV2.TransactOpts, path, amountIn)
}

// QuoteExactInputSingle is a paid mutator transaction binding the contract method 0xc6a5026a.
//
// Solidity: function quoteExactInputSingle((address,address,uint256,uint24,uint160) params) returns(uint256 amountOut, uint160 sqrtPriceX96After, uint32 initializedTicksCrossed, uint256 gasEstimate)
func (_QuoterV2 *QuoterV2Transactor) QuoteExactInputSingle(opts *bind.TransactOpts, params IQuoterV2QuoteExactInputSingleParams) (*types.Transaction, error) {
	return _QuoterV2.contract.Transact(opts, "quoteExactInputSingle", params)
}

// QuoteExactInputSingle is a paid mutator transaction binding the contract method 0xc6a5026a.
//
// Solidity: function quoteExactInputSingle((address,address,uint256,uint24,uint160) params) returns(uint256 amountOut, uint160 sqrtPriceX96After, uint32 initializedTicksCrossed, uint256 gasEstimate)
func (_QuoterV2 *QuoterV2Session) QuoteExactInputSingle(params IQuoterV2QuoteExactInputSingleParams) (*types.Transaction, error) {
	return _QuoterV2.Contract.QuoteExactInputSingle(&_QuoterV2.TransactOpts, params)
}

// QuoteExactInputSingle is a paid mutator transaction binding the contract method 0xc6a5026a.
//
// Solidity: function quoteExactInputSingle((address,address,uint256,uint24,uint160) params) returns(uint256 amountOut, uint160 sqrtPriceX96After, uint32 initializedTicksCrossed, uint256 gasEstimate)
func (_QuoterV2 *QuoterV2TransactorSession) QuoteExactInputSingle(params IQuoterV2QuoteExactInputSingleParams) (*types.Transaction, error) {
	return _QuoterV2.Contract.QuoteExactInputSingle(&_QuoterV2.TransactOpts, params)
}

// QuoteExactOutput is a paid mutator transaction binding the contract method 0x2f80bb1d.
//
// Solidity: function quoteExactOutput(bytes path, uint256 amountOut) returns(uint256 amountIn, uint160[] sqrtPriceX96AfterList, uint32[] initializedTicksCrossedList, uint256 gasEstimate)
func (_QuoterV2 *QuoterV2Transactor) QuoteExactOutput(opts *bind.TransactOpts, path []byte, amountOut *big.Int) (*types.Transaction, error) {
	return _QuoterV2.contract.Transact(opts, "quoteExactOutput", path, amountOut)
}

// QuoteExactOutput is a paid mutator transaction binding the contract method 0x2f80bb1d.
//
// Solidity: function quoteExactOutput(bytes path, uint256 amountOut) returns(uint256 amountIn, uint160[] sqrtPriceX96AfterList, uint32[] initializedTicksCrossedList, uint256 gasEstimate)
func (_QuoterV2 *QuoterV2Session) QuoteExactOutput(path []byte, amountOut *big.Int) (*types.Transaction, error) {
	return _QuoterV2.Contract.QuoteExactOutput(&_QuoterV2.TransactOpts, path, amountOut)
}

// QuoteExactOutput is a paid mutator transaction binding the contract method 0x2f80bb1d.
//
// Solidity: function quoteExactOutput(bytes path, uint256 amountOut) returns(uint256 amountIn, uint160[] sqrtPriceX96AfterList, uint32[] initializedTicksCrossedList, uint256 gasEstimate)
func (_QuoterV2 *QuoterV2TransactorSession) QuoteExactOutput(path []byte, amountOut *big.Int) (*types.Transaction, error) {
	return _QuoterV2.Contract.QuoteExactOutput(&_QuoterV2.TransactOpts, path, amountOut)
}

// QuoteExactOutputSingle is a paid mutator transaction binding the contract method 0xbd21704a.
//
// Solidity: function quoteExactOutputSingle((address,address,uint256,uint24,uint160) params) returns(uint256 amountIn, uint160 sqrtPriceX96After, uint32 initializedTicksCrossed, uint256 gasEstimate)
func (_QuoterV2 *QuoterV2Transactor) QuoteExactOutputSingle(opts *bind.TransactOpts, params IQuoterV2QuoteExactOutputSingleParams) (*types.Transaction, error) {
	return _QuoterV2.contract.Transact(opts, "quoteExactOutputSingle", params)
}

// QuoteExactOutputSingle is a paid mutator transaction binding the contract method 0xbd21704a.
//
// Solidity: function quoteExactOutputSingle((address,address,uint256,uint24,uint160) params) returns(uint256 amountIn, uint160 sqrtPriceX96After, uint32 initializedTicksCrossed, uint256 gasEstimate)
func (_QuoterV2 *QuoterV2Session) QuoteExactOutputSingle(params IQuoterV2QuoteExactOutputSingleParams) (*types.Transaction, error) {
	return _QuoterV2.Contract.QuoteExactOutputSingle(&_QuoterV2.TransactOpts, params)
}

// QuoteExactOutputSingle is a paid mutator transaction binding the contract method 0xbd21704a.
//
// Solidity: function quoteExactOutputSingle((address,address,uint256,uint24,uint160) params) returns(uint256 amountIn, uint160 sqrtPriceX96After, uint32 initializedTicksCrossed, uint256 gasEstimate)
func (_QuoterV2 *QuoterV2TransactorSession) QuoteExactOutputSingle(params IQuoterV2QuoteExactOutputSingleParams) (*types.Transaction, error) {
	return _QuoterV2.Contract.QuoteExactOutputSingle(&_QuoterV2.TransactOpts, params)
}
