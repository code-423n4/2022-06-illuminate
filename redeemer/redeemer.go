// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package redeemer

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
)

// RedeemerMetaData contains all meta data concerning the Redeemer contract.
var RedeemerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"l\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"s\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"p\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"Exists\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"Invalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"principal\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Redeem\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"apwineAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"f\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"authRedeem\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lender\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"marketPlace\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendleAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"}],\"name\":\"redeem\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"o\",\"type\":\"address\"}],\"name\":\"redeem\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"i\",\"type\":\"bytes32\"}],\"name\":\"redeem\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"d\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"o\",\"type\":\"address\"}],\"name\":\"redeem\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"l\",\"type\":\"address\"}],\"name\":\"setLenderAddress\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"m\",\"type\":\"address\"}],\"name\":\"setMarketPlaceAddress\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"swivelAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tempusAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162001d5938038062001d598339810160408190526200003491620000be565b60008054336001600160a01b0319918216179091556002805482166001600160a01b0397881617905560038054821695871695909517909455600480548516938616939093179092556005805484169185169190911790556006805490921692169190911790556200012e565b80516001600160a01b0381168114620000b957600080fd5b919050565b600080600080600060a08688031215620000d757600080fd5b620000e286620000a1565b9450620000f260208701620000a1565b93506200010260408701620000a1565b92506200011260608701620000a1565b91506200012260808701620000a1565b90509295509295909350565b611c1b806200013e6000396000f3fe608060405234801561001057600080fd5b50600436106100ea5760003560e01c8063cc98ef961161008c578063e36e4b5611610066578063e36e4b561461021b578063ea08c0311461022e578063ef6035691461024e578063f851a4401461026e57600080fd5b8063cc98ef96146101d5578063d2ba8788146101e8578063de1d3cb5146101fb57600080fd5b8063a1b1138c116100c8578063a1b1138c1461017c578063b4d7813d1461018f578063ba24ea6c146101a2578063bcead63e146101b557600080fd5b80632e25d2a6146100ef57806337a3eeec1461013957806370a03ced14610159575b600080fd5b60015461010f9073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b60065461010f9073ffffffffffffffffffffffffffffffffffffffff1681565b61016c6101673660046119be565b61028e565b6040519015158152602001610130565b61016c61018a366004611a2f565b6105a9565b61016c61019d366004611a6e565b610aee565b61016c6101b0366004611a92565b610c0d565b60025461010f9073ffffffffffffffffffffffffffffffffffffffff1681565b61016c6101e3366004611a6e565b611134565b61016c6101f6366004611ae3565b611251565b60055461010f9073ffffffffffffffffffffffffffffffffffffffff1681565b61016c610229366004611b27565b6114ca565b60035461010f9073ffffffffffffffffffffffffffffffffffffffff1681565b60045461010f9073ffffffffffffffffffffffffffffffffffffffff1681565b60005461010f9073ffffffffffffffffffffffffffffffffffffffff1681565b6001546040517fca1695f000000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff87811660048301526024820187905260006044830181905292169063ca1695f0906064016020604051808303816000875af115801561030f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103339190611b8d565b3373ffffffffffffffffffffffffffffffffffffffff821614610382576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001546040517fca1695f000000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff89811660048301526024820189905260006044830181905292169063ca1695f0906064016020604051808303816000875af1158015610403573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104279190611b8d565b90508073ffffffffffffffffffffffffffffffffffffffff1663204f83f96040518163ffffffff1660e01b81526004016020604051808303816000875af1158015610476573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061049a9190611baa565b421015610508576040517f53a2556c00000000000000000000000000000000000000000000000000000000815260206004820152600b60248201527f6e6f74206d61747572656400000000000000000000000000000000000000000060448201526064015b60405180910390fd5b6040517f9dc29fac00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff878116600483015260248201869052821690639dc29fac90604401600060405180830381600087803b15801561057857600080fd5b505af115801561058c573d6000803e3d6000fd5b5050505061059b888686611798565b506001979650505050505050565b6001546040517fca1695f000000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff84811660048301526024820184905260ff86166044830152600092839291169063ca1695f0906064016020604051808303816000875af115801561062f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106539190611b8d565b905060ff851660011480159061066d575060ff8516600314155b801561067d575060ff8516600214155b801561068d575060ff8516600814155b156106f4576040517f53a2556c00000000000000000000000000000000000000000000000000000000815260206004820152600960248201527f7072696e636970616c000000000000000000000000000000000000000000000060448201526064016104ff565b6002546040517f70a0823100000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff91821660048201526000918316906370a08231906024016020604051808303816000875af1158015610768573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061078c9190611baa565b6002549091506107b590839073ffffffffffffffffffffffffffffffffffffffff163084611865565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60ff871601610889576003546040517f154e0f2e00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff878116600483015260248201879052604482018490529091169063154e0f2e906064016020604051808303816000875af115801561085f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108839190611bc3565b50610a8c565b7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffd60ff871601610945576001546040517f884e17f30000000000000000000000000000000000000000000000000000000081526004810183905273ffffffffffffffffffffffffffffffffffffffff91821660248201529083169063884e17f3906044015b600060405180830381600087803b15801561092857600080fd5b505af115801561093c573d6000803e3d6000fd5b50505050610a8c565b7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe60ff8716016109cf576040517f0e6dfcd5000000000000000000000000000000000000000000000000000000008152306004820181905260248201526044810182905273ffffffffffffffffffffffffffffffffffffffff831690630e6dfcd59060640161090e565b7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff860ff871601610a8c576040517fd905777e00000000000000000000000000000000000000000000000000000000815230600482015273ffffffffffffffffffffffffffffffffffffffff83169063d905777e906024016020604051808303816000875af1158015610a65573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a899190611baa565b90505b6040805160ff8816815260208101839052859173ffffffffffffffffffffffffffffffffffffffff8816917ffd23960ad64d756515d736fc98d5f473529ef08157d573f8a13c1723cc632fe9910160405180910390a350600195945050505050565b6000805473ffffffffffffffffffffffffffffffffffffffff16338114610b41576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60025473ffffffffffffffffffffffffffffffffffffffff1615610bc1576040517f1ed713cc00000000000000000000000000000000000000000000000000000000815260206004820152600660248201527f6c656e646572000000000000000000000000000000000000000000000000000060448201526064016104ff565b6002805473ffffffffffffffffffffffffffffffffffffffff85167fffffffffffffffffffffffff00000000000000000000000000000000000000009091161790556001915050919050565b6001546040517fca1695f000000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85811660048301526024820185905260ff87166044830152600092839291169063ca1695f0906064016020604051808303816000875af1158015610c93573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610cb79190611b8d565b6002546040517f70a0823100000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9182166004820152919250600091908316906370a08231906024016020604051808303816000875af1158015610d2f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d539190611baa565b905060ff871615610d8457600254610d8490879073ffffffffffffffffffffffffffffffffffffffff163084611865565b7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff960ff881601610e40576006546040517ff3fef3a300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8681166004830152602482018490529091169063f3fef3a3906044015b600060405180830381600087803b158015610e2357600080fd5b505af1158015610e37573d6000803e3d6000fd5b505050506110d3565b7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffb60ff881601610ed6576005546040517f6c8d4fa100000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8681166004830152602482018490526000604483015230606483015290911690636c8d4fa190608401610e09565b60ff87166110715760008290508073ffffffffffffffffffffffffffffffffffffffff1663204f83f96040518163ffffffff1660e01b81526004016020604051808303816000875af1158015610f30573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f549190611baa565b421015610fbd576040517f53a2556c00000000000000000000000000000000000000000000000000000000815260206004820152600b60248201527f6e6f74206d61747572656400000000000000000000000000000000000000000060448201526064016104ff565b6040517f9dc29fac00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff868116600483015260248201849052821690639dc29fac90604401600060405180830381600087803b15801561102d57600080fd5b505af1158015611041573d6000803e3d6000fd5b505060025461106b925089915073ffffffffffffffffffffffffffffffffffffffff163085611865565b506110d3565b6040517f53a2556c00000000000000000000000000000000000000000000000000000000815260206004820152600960248201527f7072696e636970616c000000000000000000000000000000000000000000000060448201526064016104ff565b604080516000815260208101839052869173ffffffffffffffffffffffffffffffffffffffff8916917ffd23960ad64d756515d736fc98d5f473529ef08157d573f8a13c1723cc632fe9910160405180910390a35060019695505050505050565b6000805473ffffffffffffffffffffffffffffffffffffffff16338114611187576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60015473ffffffffffffffffffffffffffffffffffffffff1615611207576040517f1ed713cc00000000000000000000000000000000000000000000000000000000815260206004820152600b60248201527f6d61726b6574706c61636500000000000000000000000000000000000000000060448201526064016104ff565b6001805473ffffffffffffffffffffffffffffffffffffffff85167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116178155915050919050565b600060ff85166004146112c0576040517f53a2556c00000000000000000000000000000000000000000000000000000000815260206004820152600960248201527f7072696e636970616c000000000000000000000000000000000000000000000060448201526064016104ff565b6001546040517fca1695f000000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff86811660048301526024820186905260ff88166044830152600092169063ca1695f0906064016020604051808303816000875af1158015611343573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113679190611b8d565b6002546040517f70a0823100000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9182166004820152919250600091908316906370a08231906024016020604051808303816000875af11580156113df573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114039190611baa565b60025490915061142c90839073ffffffffffffffffffffffffffffffffffffffff163084611865565b600480546040517fd230566c00000000000000000000000000000000000000000000000000000000815291820186905273ffffffffffffffffffffffffffffffffffffffff888116602484015260448301889052169063d230566c90606401600060405180830381600087803b1580156114a557600080fd5b505af11580156114b9573d6000803e3d6000fd5b5060019a9950505050505050505050565b600060ff8616600614611539576040517f53a2556c00000000000000000000000000000000000000000000000000000000815260206004820152600960248201527f7072696e636970616c000000000000000000000000000000000000000000000060448201526064016104ff565b6001546040517fca1695f000000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff87811660048301526024820187905260ff89166044830152600092169063ca1695f0906064016020604051808303816000875af11580156115bc573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115e09190611b8d565b6002546040517f70a0823100000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9182166004820152919250600091908316906370a08231906024016020604051808303816000875af1158015611658573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061167c9190611baa565b6002549091506116a590839073ffffffffffffffffffffffffffffffffffffffff163084611865565b6040517f2b83cccd00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85811660048301526024820188905260448201839052861690632b83cccd90606401600060405180830381600087803b15801561171c57600080fd5b505af1158015611730573d6000803e3d6000fd5b50506040805160ff8c1681526020810185905289935073ffffffffffffffffffffffffffffffffffffffff8b1692507ffd23960ad64d756515d736fc98d5f473529ef08157d573f8a13c1723cc632fe9910160405180910390a3506001979650505050505050565b60006040517fa9059cbb00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff841660048201528260248201526000806044836000895af19150506117f98161194f565b61185f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600f60248201527f7472616e73666572206661696c6564000000000000000000000000000000000060448201526064016104ff565b50505050565b60006040517f23b872dd00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8516600482015273ffffffffffffffffffffffffffffffffffffffff8416602482015282604482015260008060648360008a5af19150506118e28161194f565b611948576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f7472616e736665722066726f6d206661696c656400000000000000000000000060448201526064016104ff565b5050505050565b6000803d8361196257806000803e806000fd5b806020811461197a57801561198b5760009250611990565b816000803e60005115159250611990565b600192505b50909392505050565b73ffffffffffffffffffffffffffffffffffffffff811681146119bb57600080fd5b50565b600080600080600060a086880312156119d657600080fd5b85356119e181611999565b94506020860135935060408601356119f881611999565b92506060860135611a0881611999565b949793965091946080013592915050565b803560ff81168114611a2a57600080fd5b919050565b600080600060608486031215611a4457600080fd5b611a4d84611a19565b92506020840135611a5d81611999565b929592945050506040919091013590565b600060208284031215611a8057600080fd5b8135611a8b81611999565b9392505050565b60008060008060808587031215611aa857600080fd5b611ab185611a19565b93506020850135611ac181611999565b9250604085013591506060850135611ad881611999565b939692955090935050565b60008060008060808587031215611af957600080fd5b611b0285611a19565b93506020850135611b1281611999565b93969395505050506040820135916060013590565b600080600080600060a08688031215611b3f57600080fd5b611b4886611a19565b94506020860135611b5881611999565b9350604086013592506060860135611b6f81611999565b91506080860135611b7f81611999565b809150509295509295909350565b600060208284031215611b9f57600080fd5b8151611a8b81611999565b600060208284031215611bbc57600080fd5b5051919050565b600060208284031215611bd557600080fd5b81518015158114611a8b57600080fdfea2646970667358221220c90113e7e8b4b4308464f636a14a983c77f008dbbaa914a2db8c7042dc772b5c64736f6c634300080d0033",
}

// RedeemerABI is the input ABI used to generate the binding from.
// Deprecated: Use RedeemerMetaData.ABI instead.
var RedeemerABI = RedeemerMetaData.ABI

// RedeemerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RedeemerMetaData.Bin instead.
var RedeemerBin = RedeemerMetaData.Bin

// DeployRedeemer deploys a new Ethereum contract, binding an instance of Redeemer to it.
func DeployRedeemer(auth *bind.TransactOpts, backend bind.ContractBackend, l common.Address, s common.Address, p common.Address, t common.Address, a common.Address) (common.Address, *types.Transaction, *Redeemer, error) {
	parsed, err := RedeemerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RedeemerBin), backend, l, s, p, t, a)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Redeemer{RedeemerCaller: RedeemerCaller{contract: contract}, RedeemerTransactor: RedeemerTransactor{contract: contract}, RedeemerFilterer: RedeemerFilterer{contract: contract}}, nil
}

// Redeemer is an auto generated Go binding around an Ethereum contract.
type Redeemer struct {
	RedeemerCaller     // Read-only binding to the contract
	RedeemerTransactor // Write-only binding to the contract
	RedeemerFilterer   // Log filterer for contract events
}

// RedeemerCaller is an auto generated read-only Go binding around an Ethereum contract.
type RedeemerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RedeemerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RedeemerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RedeemerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RedeemerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RedeemerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RedeemerSession struct {
	Contract     *Redeemer         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RedeemerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RedeemerCallerSession struct {
	Contract *RedeemerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// RedeemerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RedeemerTransactorSession struct {
	Contract     *RedeemerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// RedeemerRaw is an auto generated low-level Go binding around an Ethereum contract.
type RedeemerRaw struct {
	Contract *Redeemer // Generic contract binding to access the raw methods on
}

// RedeemerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RedeemerCallerRaw struct {
	Contract *RedeemerCaller // Generic read-only contract binding to access the raw methods on
}

// RedeemerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RedeemerTransactorRaw struct {
	Contract *RedeemerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRedeemer creates a new instance of Redeemer, bound to a specific deployed contract.
func NewRedeemer(address common.Address, backend bind.ContractBackend) (*Redeemer, error) {
	contract, err := bindRedeemer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Redeemer{RedeemerCaller: RedeemerCaller{contract: contract}, RedeemerTransactor: RedeemerTransactor{contract: contract}, RedeemerFilterer: RedeemerFilterer{contract: contract}}, nil
}

// NewRedeemerCaller creates a new read-only instance of Redeemer, bound to a specific deployed contract.
func NewRedeemerCaller(address common.Address, caller bind.ContractCaller) (*RedeemerCaller, error) {
	contract, err := bindRedeemer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RedeemerCaller{contract: contract}, nil
}

// NewRedeemerTransactor creates a new write-only instance of Redeemer, bound to a specific deployed contract.
func NewRedeemerTransactor(address common.Address, transactor bind.ContractTransactor) (*RedeemerTransactor, error) {
	contract, err := bindRedeemer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RedeemerTransactor{contract: contract}, nil
}

// NewRedeemerFilterer creates a new log filterer instance of Redeemer, bound to a specific deployed contract.
func NewRedeemerFilterer(address common.Address, filterer bind.ContractFilterer) (*RedeemerFilterer, error) {
	contract, err := bindRedeemer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RedeemerFilterer{contract: contract}, nil
}

// bindRedeemer binds a generic wrapper to an already deployed contract.
func bindRedeemer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RedeemerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Redeemer *RedeemerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Redeemer.Contract.RedeemerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Redeemer *RedeemerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Redeemer.Contract.RedeemerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Redeemer *RedeemerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Redeemer.Contract.RedeemerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Redeemer *RedeemerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Redeemer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Redeemer *RedeemerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Redeemer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Redeemer *RedeemerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Redeemer.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Redeemer *RedeemerCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Redeemer.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Redeemer *RedeemerSession) Admin() (common.Address, error) {
	return _Redeemer.Contract.Admin(&_Redeemer.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Redeemer *RedeemerCallerSession) Admin() (common.Address, error) {
	return _Redeemer.Contract.Admin(&_Redeemer.CallOpts)
}

// ApwineAddr is a free data retrieval call binding the contract method 0x37a3eeec.
//
// Solidity: function apwineAddr() view returns(address)
func (_Redeemer *RedeemerCaller) ApwineAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Redeemer.contract.Call(opts, &out, "apwineAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ApwineAddr is a free data retrieval call binding the contract method 0x37a3eeec.
//
// Solidity: function apwineAddr() view returns(address)
func (_Redeemer *RedeemerSession) ApwineAddr() (common.Address, error) {
	return _Redeemer.Contract.ApwineAddr(&_Redeemer.CallOpts)
}

// ApwineAddr is a free data retrieval call binding the contract method 0x37a3eeec.
//
// Solidity: function apwineAddr() view returns(address)
func (_Redeemer *RedeemerCallerSession) ApwineAddr() (common.Address, error) {
	return _Redeemer.Contract.ApwineAddr(&_Redeemer.CallOpts)
}

// Lender is a free data retrieval call binding the contract method 0xbcead63e.
//
// Solidity: function lender() view returns(address)
func (_Redeemer *RedeemerCaller) Lender(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Redeemer.contract.Call(opts, &out, "lender")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Lender is a free data retrieval call binding the contract method 0xbcead63e.
//
// Solidity: function lender() view returns(address)
func (_Redeemer *RedeemerSession) Lender() (common.Address, error) {
	return _Redeemer.Contract.Lender(&_Redeemer.CallOpts)
}

// Lender is a free data retrieval call binding the contract method 0xbcead63e.
//
// Solidity: function lender() view returns(address)
func (_Redeemer *RedeemerCallerSession) Lender() (common.Address, error) {
	return _Redeemer.Contract.Lender(&_Redeemer.CallOpts)
}

// MarketPlace is a free data retrieval call binding the contract method 0x2e25d2a6.
//
// Solidity: function marketPlace() view returns(address)
func (_Redeemer *RedeemerCaller) MarketPlace(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Redeemer.contract.Call(opts, &out, "marketPlace")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MarketPlace is a free data retrieval call binding the contract method 0x2e25d2a6.
//
// Solidity: function marketPlace() view returns(address)
func (_Redeemer *RedeemerSession) MarketPlace() (common.Address, error) {
	return _Redeemer.Contract.MarketPlace(&_Redeemer.CallOpts)
}

// MarketPlace is a free data retrieval call binding the contract method 0x2e25d2a6.
//
// Solidity: function marketPlace() view returns(address)
func (_Redeemer *RedeemerCallerSession) MarketPlace() (common.Address, error) {
	return _Redeemer.Contract.MarketPlace(&_Redeemer.CallOpts)
}

// PendleAddr is a free data retrieval call binding the contract method 0xef603569.
//
// Solidity: function pendleAddr() view returns(address)
func (_Redeemer *RedeemerCaller) PendleAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Redeemer.contract.Call(opts, &out, "pendleAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendleAddr is a free data retrieval call binding the contract method 0xef603569.
//
// Solidity: function pendleAddr() view returns(address)
func (_Redeemer *RedeemerSession) PendleAddr() (common.Address, error) {
	return _Redeemer.Contract.PendleAddr(&_Redeemer.CallOpts)
}

// PendleAddr is a free data retrieval call binding the contract method 0xef603569.
//
// Solidity: function pendleAddr() view returns(address)
func (_Redeemer *RedeemerCallerSession) PendleAddr() (common.Address, error) {
	return _Redeemer.Contract.PendleAddr(&_Redeemer.CallOpts)
}

// SwivelAddr is a free data retrieval call binding the contract method 0xea08c031.
//
// Solidity: function swivelAddr() view returns(address)
func (_Redeemer *RedeemerCaller) SwivelAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Redeemer.contract.Call(opts, &out, "swivelAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SwivelAddr is a free data retrieval call binding the contract method 0xea08c031.
//
// Solidity: function swivelAddr() view returns(address)
func (_Redeemer *RedeemerSession) SwivelAddr() (common.Address, error) {
	return _Redeemer.Contract.SwivelAddr(&_Redeemer.CallOpts)
}

// SwivelAddr is a free data retrieval call binding the contract method 0xea08c031.
//
// Solidity: function swivelAddr() view returns(address)
func (_Redeemer *RedeemerCallerSession) SwivelAddr() (common.Address, error) {
	return _Redeemer.Contract.SwivelAddr(&_Redeemer.CallOpts)
}

// TempusAddr is a free data retrieval call binding the contract method 0xde1d3cb5.
//
// Solidity: function tempusAddr() view returns(address)
func (_Redeemer *RedeemerCaller) TempusAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Redeemer.contract.Call(opts, &out, "tempusAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TempusAddr is a free data retrieval call binding the contract method 0xde1d3cb5.
//
// Solidity: function tempusAddr() view returns(address)
func (_Redeemer *RedeemerSession) TempusAddr() (common.Address, error) {
	return _Redeemer.Contract.TempusAddr(&_Redeemer.CallOpts)
}

// TempusAddr is a free data retrieval call binding the contract method 0xde1d3cb5.
//
// Solidity: function tempusAddr() view returns(address)
func (_Redeemer *RedeemerCallerSession) TempusAddr() (common.Address, error) {
	return _Redeemer.Contract.TempusAddr(&_Redeemer.CallOpts)
}

// AuthRedeem is a paid mutator transaction binding the contract method 0x70a03ced.
//
// Solidity: function authRedeem(address u, uint256 m, address f, address t, uint256 a) returns(bool)
func (_Redeemer *RedeemerTransactor) AuthRedeem(opts *bind.TransactOpts, u common.Address, m *big.Int, f common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _Redeemer.contract.Transact(opts, "authRedeem", u, m, f, t, a)
}

// AuthRedeem is a paid mutator transaction binding the contract method 0x70a03ced.
//
// Solidity: function authRedeem(address u, uint256 m, address f, address t, uint256 a) returns(bool)
func (_Redeemer *RedeemerSession) AuthRedeem(u common.Address, m *big.Int, f common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _Redeemer.Contract.AuthRedeem(&_Redeemer.TransactOpts, u, m, f, t, a)
}

// AuthRedeem is a paid mutator transaction binding the contract method 0x70a03ced.
//
// Solidity: function authRedeem(address u, uint256 m, address f, address t, uint256 a) returns(bool)
func (_Redeemer *RedeemerTransactorSession) AuthRedeem(u common.Address, m *big.Int, f common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _Redeemer.Contract.AuthRedeem(&_Redeemer.TransactOpts, u, m, f, t, a)
}

// Redeem is a paid mutator transaction binding the contract method 0xa1b1138c.
//
// Solidity: function redeem(uint8 p, address u, uint256 m) returns(bool)
func (_Redeemer *RedeemerTransactor) Redeem(opts *bind.TransactOpts, p uint8, u common.Address, m *big.Int) (*types.Transaction, error) {
	return _Redeemer.contract.Transact(opts, "redeem", p, u, m)
}

// Redeem is a paid mutator transaction binding the contract method 0xa1b1138c.
//
// Solidity: function redeem(uint8 p, address u, uint256 m) returns(bool)
func (_Redeemer *RedeemerSession) Redeem(p uint8, u common.Address, m *big.Int) (*types.Transaction, error) {
	return _Redeemer.Contract.Redeem(&_Redeemer.TransactOpts, p, u, m)
}

// Redeem is a paid mutator transaction binding the contract method 0xa1b1138c.
//
// Solidity: function redeem(uint8 p, address u, uint256 m) returns(bool)
func (_Redeemer *RedeemerTransactorSession) Redeem(p uint8, u common.Address, m *big.Int) (*types.Transaction, error) {
	return _Redeemer.Contract.Redeem(&_Redeemer.TransactOpts, p, u, m)
}

// Redeem0 is a paid mutator transaction binding the contract method 0xba24ea6c.
//
// Solidity: function redeem(uint8 p, address u, uint256 m, address o) returns(bool)
func (_Redeemer *RedeemerTransactor) Redeem0(opts *bind.TransactOpts, p uint8, u common.Address, m *big.Int, o common.Address) (*types.Transaction, error) {
	return _Redeemer.contract.Transact(opts, "redeem0", p, u, m, o)
}

// Redeem0 is a paid mutator transaction binding the contract method 0xba24ea6c.
//
// Solidity: function redeem(uint8 p, address u, uint256 m, address o) returns(bool)
func (_Redeemer *RedeemerSession) Redeem0(p uint8, u common.Address, m *big.Int, o common.Address) (*types.Transaction, error) {
	return _Redeemer.Contract.Redeem0(&_Redeemer.TransactOpts, p, u, m, o)
}

// Redeem0 is a paid mutator transaction binding the contract method 0xba24ea6c.
//
// Solidity: function redeem(uint8 p, address u, uint256 m, address o) returns(bool)
func (_Redeemer *RedeemerTransactorSession) Redeem0(p uint8, u common.Address, m *big.Int, o common.Address) (*types.Transaction, error) {
	return _Redeemer.Contract.Redeem0(&_Redeemer.TransactOpts, p, u, m, o)
}

// Redeem1 is a paid mutator transaction binding the contract method 0xd2ba8788.
//
// Solidity: function redeem(uint8 p, address u, uint256 m, bytes32 i) returns(bool)
func (_Redeemer *RedeemerTransactor) Redeem1(opts *bind.TransactOpts, p uint8, u common.Address, m *big.Int, i [32]byte) (*types.Transaction, error) {
	return _Redeemer.contract.Transact(opts, "redeem1", p, u, m, i)
}

// Redeem1 is a paid mutator transaction binding the contract method 0xd2ba8788.
//
// Solidity: function redeem(uint8 p, address u, uint256 m, bytes32 i) returns(bool)
func (_Redeemer *RedeemerSession) Redeem1(p uint8, u common.Address, m *big.Int, i [32]byte) (*types.Transaction, error) {
	return _Redeemer.Contract.Redeem1(&_Redeemer.TransactOpts, p, u, m, i)
}

// Redeem1 is a paid mutator transaction binding the contract method 0xd2ba8788.
//
// Solidity: function redeem(uint8 p, address u, uint256 m, bytes32 i) returns(bool)
func (_Redeemer *RedeemerTransactorSession) Redeem1(p uint8, u common.Address, m *big.Int, i [32]byte) (*types.Transaction, error) {
	return _Redeemer.Contract.Redeem1(&_Redeemer.TransactOpts, p, u, m, i)
}

// Redeem2 is a paid mutator transaction binding the contract method 0xe36e4b56.
//
// Solidity: function redeem(uint8 p, address u, uint256 m, address d, address o) returns(bool)
func (_Redeemer *RedeemerTransactor) Redeem2(opts *bind.TransactOpts, p uint8, u common.Address, m *big.Int, d common.Address, o common.Address) (*types.Transaction, error) {
	return _Redeemer.contract.Transact(opts, "redeem2", p, u, m, d, o)
}

// Redeem2 is a paid mutator transaction binding the contract method 0xe36e4b56.
//
// Solidity: function redeem(uint8 p, address u, uint256 m, address d, address o) returns(bool)
func (_Redeemer *RedeemerSession) Redeem2(p uint8, u common.Address, m *big.Int, d common.Address, o common.Address) (*types.Transaction, error) {
	return _Redeemer.Contract.Redeem2(&_Redeemer.TransactOpts, p, u, m, d, o)
}

// Redeem2 is a paid mutator transaction binding the contract method 0xe36e4b56.
//
// Solidity: function redeem(uint8 p, address u, uint256 m, address d, address o) returns(bool)
func (_Redeemer *RedeemerTransactorSession) Redeem2(p uint8, u common.Address, m *big.Int, d common.Address, o common.Address) (*types.Transaction, error) {
	return _Redeemer.Contract.Redeem2(&_Redeemer.TransactOpts, p, u, m, d, o)
}

// SetLenderAddress is a paid mutator transaction binding the contract method 0xb4d7813d.
//
// Solidity: function setLenderAddress(address l) returns(bool)
func (_Redeemer *RedeemerTransactor) SetLenderAddress(opts *bind.TransactOpts, l common.Address) (*types.Transaction, error) {
	return _Redeemer.contract.Transact(opts, "setLenderAddress", l)
}

// SetLenderAddress is a paid mutator transaction binding the contract method 0xb4d7813d.
//
// Solidity: function setLenderAddress(address l) returns(bool)
func (_Redeemer *RedeemerSession) SetLenderAddress(l common.Address) (*types.Transaction, error) {
	return _Redeemer.Contract.SetLenderAddress(&_Redeemer.TransactOpts, l)
}

// SetLenderAddress is a paid mutator transaction binding the contract method 0xb4d7813d.
//
// Solidity: function setLenderAddress(address l) returns(bool)
func (_Redeemer *RedeemerTransactorSession) SetLenderAddress(l common.Address) (*types.Transaction, error) {
	return _Redeemer.Contract.SetLenderAddress(&_Redeemer.TransactOpts, l)
}

// SetMarketPlaceAddress is a paid mutator transaction binding the contract method 0xcc98ef96.
//
// Solidity: function setMarketPlaceAddress(address m) returns(bool)
func (_Redeemer *RedeemerTransactor) SetMarketPlaceAddress(opts *bind.TransactOpts, m common.Address) (*types.Transaction, error) {
	return _Redeemer.contract.Transact(opts, "setMarketPlaceAddress", m)
}

// SetMarketPlaceAddress is a paid mutator transaction binding the contract method 0xcc98ef96.
//
// Solidity: function setMarketPlaceAddress(address m) returns(bool)
func (_Redeemer *RedeemerSession) SetMarketPlaceAddress(m common.Address) (*types.Transaction, error) {
	return _Redeemer.Contract.SetMarketPlaceAddress(&_Redeemer.TransactOpts, m)
}

// SetMarketPlaceAddress is a paid mutator transaction binding the contract method 0xcc98ef96.
//
// Solidity: function setMarketPlaceAddress(address m) returns(bool)
func (_Redeemer *RedeemerTransactorSession) SetMarketPlaceAddress(m common.Address) (*types.Transaction, error) {
	return _Redeemer.Contract.SetMarketPlaceAddress(&_Redeemer.TransactOpts, m)
}

// RedeemerRedeemIterator is returned from FilterRedeem and is used to iterate over the raw logs and unpacked data for Redeem events raised by the Redeemer contract.
type RedeemerRedeemIterator struct {
	Event *RedeemerRedeem // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RedeemerRedeemIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RedeemerRedeem)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RedeemerRedeem)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RedeemerRedeemIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RedeemerRedeemIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RedeemerRedeem represents a Redeem event raised by the Redeemer contract.
type RedeemerRedeem struct {
	Principal  uint8
	Underlying common.Address
	Maturity   *big.Int
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRedeem is a free log retrieval operation binding the contract event 0xfd23960ad64d756515d736fc98d5f473529ef08157d573f8a13c1723cc632fe9.
//
// Solidity: event Redeem(uint8 principal, address indexed underlying, uint256 indexed maturity, uint256 amount)
func (_Redeemer *RedeemerFilterer) FilterRedeem(opts *bind.FilterOpts, underlying []common.Address, maturity []*big.Int) (*RedeemerRedeemIterator, error) {

	var underlyingRule []interface{}
	for _, underlyingItem := range underlying {
		underlyingRule = append(underlyingRule, underlyingItem)
	}
	var maturityRule []interface{}
	for _, maturityItem := range maturity {
		maturityRule = append(maturityRule, maturityItem)
	}

	logs, sub, err := _Redeemer.contract.FilterLogs(opts, "Redeem", underlyingRule, maturityRule)
	if err != nil {
		return nil, err
	}
	return &RedeemerRedeemIterator{contract: _Redeemer.contract, event: "Redeem", logs: logs, sub: sub}, nil
}

// WatchRedeem is a free log subscription operation binding the contract event 0xfd23960ad64d756515d736fc98d5f473529ef08157d573f8a13c1723cc632fe9.
//
// Solidity: event Redeem(uint8 principal, address indexed underlying, uint256 indexed maturity, uint256 amount)
func (_Redeemer *RedeemerFilterer) WatchRedeem(opts *bind.WatchOpts, sink chan<- *RedeemerRedeem, underlying []common.Address, maturity []*big.Int) (event.Subscription, error) {

	var underlyingRule []interface{}
	for _, underlyingItem := range underlying {
		underlyingRule = append(underlyingRule, underlyingItem)
	}
	var maturityRule []interface{}
	for _, maturityItem := range maturity {
		maturityRule = append(maturityRule, maturityItem)
	}

	logs, sub, err := _Redeemer.contract.WatchLogs(opts, "Redeem", underlyingRule, maturityRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RedeemerRedeem)
				if err := _Redeemer.contract.UnpackLog(event, "Redeem", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRedeem is a log parse operation binding the contract event 0xfd23960ad64d756515d736fc98d5f473529ef08157d573f8a13c1723cc632fe9.
//
// Solidity: event Redeem(uint8 principal, address indexed underlying, uint256 indexed maturity, uint256 amount)
func (_Redeemer *RedeemerFilterer) ParseRedeem(log types.Log) (*RedeemerRedeem, error) {
	event := new(RedeemerRedeem)
	if err := _Redeemer.contract.UnpackLog(event, "Redeem", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}