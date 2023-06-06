// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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

// DunkinCapsNFTs is an auto generated low-level Go binding around an user-defined struct.
type DunkinCapsNFTs struct {
	Nft common.Address
	Id  *big.Int
}

// ContractMetaData contains all meta data concerning the Contract contract.
var ContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_feeAccount\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_moderatorAccount\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_feeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxPlayers\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"_gameTime\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structDunkinCaps.NFTs[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"name\":\"ClaimRewards\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gameId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"player\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structDunkinCaps.NFTs[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"name\":\"CreateBet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gameId\",\"type\":\"uint256\"}],\"name\":\"CreateGame\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"SetFeeAccount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"SetFeeAmount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"gameTime\",\"type\":\"uint64\"}],\"name\":\"SetGameTime\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxPlayers\",\"type\":\"uint256\"}],\"name\":\"SetMaxPlayers\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"SetModerator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"}],\"name\":\"SetWinner\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"betsDetailsOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gameId\",\"type\":\"uint256\"}],\"name\":\"claimReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"counter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"createGame\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeAccount\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gameTime\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"gamesDetailsOf\",\"outputs\":[{\"internalType\":\"enumDunkinCaps.Status\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"expires\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gameId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getBet\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"internalType\":\"structDunkinCaps.NFTs[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_gameId\",\"type\":\"uint256\"}],\"name\":\"getGame\",\"outputs\":[{\"internalType\":\"enumDunkinCaps.Status\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"address[]\",\"name\":\"players\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"expires\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gameId\",\"type\":\"uint256\"}],\"name\":\"getNFTsForGame\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"internalType\":\"structDunkinCaps.NFTs[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gameId\",\"type\":\"uint256\"}],\"name\":\"getWinner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gameId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"internalType\":\"structDunkinCaps.NFTs[]\",\"name\":\"nfts\",\"type\":\"tuple[]\"}],\"name\":\"makeBets\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxPlayers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"moderatorAccount\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"setFeeAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"setFeeAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"time\",\"type\":\"uint64\"}],\"name\":\"setGameTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"players\",\"type\":\"uint256\"}],\"name\":\"setMaxPlayers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"setModerator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gameId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"}],\"name\":\"setWinner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractMetaData.ABI instead.
var ContractABI = ContractMetaData.ABI

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// BetsDetailsOf is a free data retrieval call binding the contract method 0x1d4ba000.
//
// Solidity: function betsDetailsOf(uint256 , address , uint256 ) view returns(address nft, uint256 id)
func (_Contract *ContractCaller) BetsDetailsOf(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address, arg2 *big.Int) (struct {
	Nft common.Address
	Id  *big.Int
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "betsDetailsOf", arg0, arg1, arg2)

	outstruct := new(struct {
		Nft common.Address
		Id  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Nft = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Id = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// BetsDetailsOf is a free data retrieval call binding the contract method 0x1d4ba000.
//
// Solidity: function betsDetailsOf(uint256 , address , uint256 ) view returns(address nft, uint256 id)
func (_Contract *ContractSession) BetsDetailsOf(arg0 *big.Int, arg1 common.Address, arg2 *big.Int) (struct {
	Nft common.Address
	Id  *big.Int
}, error) {
	return _Contract.Contract.BetsDetailsOf(&_Contract.CallOpts, arg0, arg1, arg2)
}

// BetsDetailsOf is a free data retrieval call binding the contract method 0x1d4ba000.
//
// Solidity: function betsDetailsOf(uint256 , address , uint256 ) view returns(address nft, uint256 id)
func (_Contract *ContractCallerSession) BetsDetailsOf(arg0 *big.Int, arg1 common.Address, arg2 *big.Int) (struct {
	Nft common.Address
	Id  *big.Int
}, error) {
	return _Contract.Contract.BetsDetailsOf(&_Contract.CallOpts, arg0, arg1, arg2)
}

// Counter is a free data retrieval call binding the contract method 0x61bc221a.
//
// Solidity: function counter() view returns(uint256)
func (_Contract *ContractCaller) Counter(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "counter")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Counter is a free data retrieval call binding the contract method 0x61bc221a.
//
// Solidity: function counter() view returns(uint256)
func (_Contract *ContractSession) Counter() (*big.Int, error) {
	return _Contract.Contract.Counter(&_Contract.CallOpts)
}

// Counter is a free data retrieval call binding the contract method 0x61bc221a.
//
// Solidity: function counter() view returns(uint256)
func (_Contract *ContractCallerSession) Counter() (*big.Int, error) {
	return _Contract.Contract.Counter(&_Contract.CallOpts)
}

// FeeAccount is a free data retrieval call binding the contract method 0x65e17c9d.
//
// Solidity: function feeAccount() view returns(address)
func (_Contract *ContractCaller) FeeAccount(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "feeAccount")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeAccount is a free data retrieval call binding the contract method 0x65e17c9d.
//
// Solidity: function feeAccount() view returns(address)
func (_Contract *ContractSession) FeeAccount() (common.Address, error) {
	return _Contract.Contract.FeeAccount(&_Contract.CallOpts)
}

// FeeAccount is a free data retrieval call binding the contract method 0x65e17c9d.
//
// Solidity: function feeAccount() view returns(address)
func (_Contract *ContractCallerSession) FeeAccount() (common.Address, error) {
	return _Contract.Contract.FeeAccount(&_Contract.CallOpts)
}

// FeeAmount is a free data retrieval call binding the contract method 0x69e15404.
//
// Solidity: function feeAmount() view returns(uint256)
func (_Contract *ContractCaller) FeeAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "feeAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeAmount is a free data retrieval call binding the contract method 0x69e15404.
//
// Solidity: function feeAmount() view returns(uint256)
func (_Contract *ContractSession) FeeAmount() (*big.Int, error) {
	return _Contract.Contract.FeeAmount(&_Contract.CallOpts)
}

// FeeAmount is a free data retrieval call binding the contract method 0x69e15404.
//
// Solidity: function feeAmount() view returns(uint256)
func (_Contract *ContractCallerSession) FeeAmount() (*big.Int, error) {
	return _Contract.Contract.FeeAmount(&_Contract.CallOpts)
}

// GameTime is a free data retrieval call binding the contract method 0xa5d1c0c0.
//
// Solidity: function gameTime() view returns(uint64)
func (_Contract *ContractCaller) GameTime(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "gameTime")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GameTime is a free data retrieval call binding the contract method 0xa5d1c0c0.
//
// Solidity: function gameTime() view returns(uint64)
func (_Contract *ContractSession) GameTime() (uint64, error) {
	return _Contract.Contract.GameTime(&_Contract.CallOpts)
}

// GameTime is a free data retrieval call binding the contract method 0xa5d1c0c0.
//
// Solidity: function gameTime() view returns(uint64)
func (_Contract *ContractCallerSession) GameTime() (uint64, error) {
	return _Contract.Contract.GameTime(&_Contract.CallOpts)
}

// GamesDetailsOf is a free data retrieval call binding the contract method 0x2f6eeda2.
//
// Solidity: function gamesDetailsOf(uint256 ) view returns(uint8 status, address winner, uint64 expires)
func (_Contract *ContractCaller) GamesDetailsOf(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Status  uint8
	Winner  common.Address
	Expires uint64
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "gamesDetailsOf", arg0)

	outstruct := new(struct {
		Status  uint8
		Winner  common.Address
		Expires uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Status = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.Winner = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Expires = *abi.ConvertType(out[2], new(uint64)).(*uint64)

	return *outstruct, err

}

// GamesDetailsOf is a free data retrieval call binding the contract method 0x2f6eeda2.
//
// Solidity: function gamesDetailsOf(uint256 ) view returns(uint8 status, address winner, uint64 expires)
func (_Contract *ContractSession) GamesDetailsOf(arg0 *big.Int) (struct {
	Status  uint8
	Winner  common.Address
	Expires uint64
}, error) {
	return _Contract.Contract.GamesDetailsOf(&_Contract.CallOpts, arg0)
}

// GamesDetailsOf is a free data retrieval call binding the contract method 0x2f6eeda2.
//
// Solidity: function gamesDetailsOf(uint256 ) view returns(uint8 status, address winner, uint64 expires)
func (_Contract *ContractCallerSession) GamesDetailsOf(arg0 *big.Int) (struct {
	Status  uint8
	Winner  common.Address
	Expires uint64
}, error) {
	return _Contract.Contract.GamesDetailsOf(&_Contract.CallOpts, arg0)
}

// GetBet is a free data retrieval call binding the contract method 0xe0950ddf.
//
// Solidity: function getBet(uint256 gameId, address user) view returns((address,uint256)[])
func (_Contract *ContractCaller) GetBet(opts *bind.CallOpts, gameId *big.Int, user common.Address) ([]DunkinCapsNFTs, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getBet", gameId, user)

	if err != nil {
		return *new([]DunkinCapsNFTs), err
	}

	out0 := *abi.ConvertType(out[0], new([]DunkinCapsNFTs)).(*[]DunkinCapsNFTs)

	return out0, err

}

// GetBet is a free data retrieval call binding the contract method 0xe0950ddf.
//
// Solidity: function getBet(uint256 gameId, address user) view returns((address,uint256)[])
func (_Contract *ContractSession) GetBet(gameId *big.Int, user common.Address) ([]DunkinCapsNFTs, error) {
	return _Contract.Contract.GetBet(&_Contract.CallOpts, gameId, user)
}

// GetBet is a free data retrieval call binding the contract method 0xe0950ddf.
//
// Solidity: function getBet(uint256 gameId, address user) view returns((address,uint256)[])
func (_Contract *ContractCallerSession) GetBet(gameId *big.Int, user common.Address) ([]DunkinCapsNFTs, error) {
	return _Contract.Contract.GetBet(&_Contract.CallOpts, gameId, user)
}

// GetGame is a free data retrieval call binding the contract method 0xa2f77bcc.
//
// Solidity: function getGame(uint256 _gameId) view returns(uint8 status, address[] players, address winner, uint64 expires)
func (_Contract *ContractCaller) GetGame(opts *bind.CallOpts, _gameId *big.Int) (struct {
	Status  uint8
	Players []common.Address
	Winner  common.Address
	Expires uint64
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getGame", _gameId)

	outstruct := new(struct {
		Status  uint8
		Players []common.Address
		Winner  common.Address
		Expires uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Status = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.Players = *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)
	outstruct.Winner = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Expires = *abi.ConvertType(out[3], new(uint64)).(*uint64)

	return *outstruct, err

}

// GetGame is a free data retrieval call binding the contract method 0xa2f77bcc.
//
// Solidity: function getGame(uint256 _gameId) view returns(uint8 status, address[] players, address winner, uint64 expires)
func (_Contract *ContractSession) GetGame(_gameId *big.Int) (struct {
	Status  uint8
	Players []common.Address
	Winner  common.Address
	Expires uint64
}, error) {
	return _Contract.Contract.GetGame(&_Contract.CallOpts, _gameId)
}

// GetGame is a free data retrieval call binding the contract method 0xa2f77bcc.
//
// Solidity: function getGame(uint256 _gameId) view returns(uint8 status, address[] players, address winner, uint64 expires)
func (_Contract *ContractCallerSession) GetGame(_gameId *big.Int) (struct {
	Status  uint8
	Players []common.Address
	Winner  common.Address
	Expires uint64
}, error) {
	return _Contract.Contract.GetGame(&_Contract.CallOpts, _gameId)
}

// GetNFTsForGame is a free data retrieval call binding the contract method 0x72ec0fa0.
//
// Solidity: function getNFTsForGame(uint256 gameId) view returns((address,uint256)[])
func (_Contract *ContractCaller) GetNFTsForGame(opts *bind.CallOpts, gameId *big.Int) ([]DunkinCapsNFTs, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getNFTsForGame", gameId)

	if err != nil {
		return *new([]DunkinCapsNFTs), err
	}

	out0 := *abi.ConvertType(out[0], new([]DunkinCapsNFTs)).(*[]DunkinCapsNFTs)

	return out0, err

}

// GetNFTsForGame is a free data retrieval call binding the contract method 0x72ec0fa0.
//
// Solidity: function getNFTsForGame(uint256 gameId) view returns((address,uint256)[])
func (_Contract *ContractSession) GetNFTsForGame(gameId *big.Int) ([]DunkinCapsNFTs, error) {
	return _Contract.Contract.GetNFTsForGame(&_Contract.CallOpts, gameId)
}

// GetNFTsForGame is a free data retrieval call binding the contract method 0x72ec0fa0.
//
// Solidity: function getNFTsForGame(uint256 gameId) view returns((address,uint256)[])
func (_Contract *ContractCallerSession) GetNFTsForGame(gameId *big.Int) ([]DunkinCapsNFTs, error) {
	return _Contract.Contract.GetNFTsForGame(&_Contract.CallOpts, gameId)
}

// GetWinner is a free data retrieval call binding the contract method 0x4129b2c9.
//
// Solidity: function getWinner(uint256 gameId) view returns(address)
func (_Contract *ContractCaller) GetWinner(opts *bind.CallOpts, gameId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getWinner", gameId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetWinner is a free data retrieval call binding the contract method 0x4129b2c9.
//
// Solidity: function getWinner(uint256 gameId) view returns(address)
func (_Contract *ContractSession) GetWinner(gameId *big.Int) (common.Address, error) {
	return _Contract.Contract.GetWinner(&_Contract.CallOpts, gameId)
}

// GetWinner is a free data retrieval call binding the contract method 0x4129b2c9.
//
// Solidity: function getWinner(uint256 gameId) view returns(address)
func (_Contract *ContractCallerSession) GetWinner(gameId *big.Int) (common.Address, error) {
	return _Contract.Contract.GetWinner(&_Contract.CallOpts, gameId)
}

// MaxPlayers is a free data retrieval call binding the contract method 0x4c2412a2.
//
// Solidity: function maxPlayers() view returns(uint256)
func (_Contract *ContractCaller) MaxPlayers(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "maxPlayers")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxPlayers is a free data retrieval call binding the contract method 0x4c2412a2.
//
// Solidity: function maxPlayers() view returns(uint256)
func (_Contract *ContractSession) MaxPlayers() (*big.Int, error) {
	return _Contract.Contract.MaxPlayers(&_Contract.CallOpts)
}

// MaxPlayers is a free data retrieval call binding the contract method 0x4c2412a2.
//
// Solidity: function maxPlayers() view returns(uint256)
func (_Contract *ContractCallerSession) MaxPlayers() (*big.Int, error) {
	return _Contract.Contract.MaxPlayers(&_Contract.CallOpts)
}

// ModeratorAccount is a free data retrieval call binding the contract method 0xc3d4970d.
//
// Solidity: function moderatorAccount() view returns(address)
func (_Contract *ContractCaller) ModeratorAccount(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "moderatorAccount")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ModeratorAccount is a free data retrieval call binding the contract method 0xc3d4970d.
//
// Solidity: function moderatorAccount() view returns(address)
func (_Contract *ContractSession) ModeratorAccount() (common.Address, error) {
	return _Contract.Contract.ModeratorAccount(&_Contract.CallOpts)
}

// ModeratorAccount is a free data retrieval call binding the contract method 0xc3d4970d.
//
// Solidity: function moderatorAccount() view returns(address)
func (_Contract *ContractCallerSession) ModeratorAccount() (common.Address, error) {
	return _Contract.Contract.ModeratorAccount(&_Contract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractSession) Owner() (common.Address, error) {
	return _Contract.Contract.Owner(&_Contract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractCallerSession) Owner() (common.Address, error) {
	return _Contract.Contract.Owner(&_Contract.CallOpts)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xae169a50.
//
// Solidity: function claimReward(uint256 gameId) returns()
func (_Contract *ContractTransactor) ClaimReward(opts *bind.TransactOpts, gameId *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "claimReward", gameId)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xae169a50.
//
// Solidity: function claimReward(uint256 gameId) returns()
func (_Contract *ContractSession) ClaimReward(gameId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.ClaimReward(&_Contract.TransactOpts, gameId)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xae169a50.
//
// Solidity: function claimReward(uint256 gameId) returns()
func (_Contract *ContractTransactorSession) ClaimReward(gameId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.ClaimReward(&_Contract.TransactOpts, gameId)
}

// CreateGame is a paid mutator transaction binding the contract method 0x7255d729.
//
// Solidity: function createGame() returns(uint256)
func (_Contract *ContractTransactor) CreateGame(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "createGame")
}

// CreateGame is a paid mutator transaction binding the contract method 0x7255d729.
//
// Solidity: function createGame() returns(uint256)
func (_Contract *ContractSession) CreateGame() (*types.Transaction, error) {
	return _Contract.Contract.CreateGame(&_Contract.TransactOpts)
}

// CreateGame is a paid mutator transaction binding the contract method 0x7255d729.
//
// Solidity: function createGame() returns(uint256)
func (_Contract *ContractTransactorSession) CreateGame() (*types.Transaction, error) {
	return _Contract.Contract.CreateGame(&_Contract.TransactOpts)
}

// MakeBets is a paid mutator transaction binding the contract method 0x8af04fd6.
//
// Solidity: function makeBets(uint256 gameId, (address,uint256)[] nfts) payable returns()
func (_Contract *ContractTransactor) MakeBets(opts *bind.TransactOpts, gameId *big.Int, nfts []DunkinCapsNFTs) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "makeBets", gameId, nfts)
}

// MakeBets is a paid mutator transaction binding the contract method 0x8af04fd6.
//
// Solidity: function makeBets(uint256 gameId, (address,uint256)[] nfts) payable returns()
func (_Contract *ContractSession) MakeBets(gameId *big.Int, nfts []DunkinCapsNFTs) (*types.Transaction, error) {
	return _Contract.Contract.MakeBets(&_Contract.TransactOpts, gameId, nfts)
}

// MakeBets is a paid mutator transaction binding the contract method 0x8af04fd6.
//
// Solidity: function makeBets(uint256 gameId, (address,uint256)[] nfts) payable returns()
func (_Contract *ContractTransactorSession) MakeBets(gameId *big.Int, nfts []DunkinCapsNFTs) (*types.Transaction, error) {
	return _Contract.Contract.MakeBets(&_Contract.TransactOpts, gameId, nfts)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_Contract *ContractTransactor) OnERC721Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "onERC721Received", arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_Contract *ContractSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Contract.Contract.OnERC721Received(&_Contract.TransactOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_Contract *ContractTransactorSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Contract.Contract.OnERC721Received(&_Contract.TransactOpts, arg0, arg1, arg2, arg3)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contract *ContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contract *ContractSession) RenounceOwnership() (*types.Transaction, error) {
	return _Contract.Contract.RenounceOwnership(&_Contract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contract *ContractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Contract.Contract.RenounceOwnership(&_Contract.TransactOpts)
}

// SetFeeAccount is a paid mutator transaction binding the contract method 0x4b023cf8.
//
// Solidity: function setFeeAccount(address account) returns()
func (_Contract *ContractTransactor) SetFeeAccount(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setFeeAccount", account)
}

// SetFeeAccount is a paid mutator transaction binding the contract method 0x4b023cf8.
//
// Solidity: function setFeeAccount(address account) returns()
func (_Contract *ContractSession) SetFeeAccount(account common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetFeeAccount(&_Contract.TransactOpts, account)
}

// SetFeeAccount is a paid mutator transaction binding the contract method 0x4b023cf8.
//
// Solidity: function setFeeAccount(address account) returns()
func (_Contract *ContractTransactorSession) SetFeeAccount(account common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetFeeAccount(&_Contract.TransactOpts, account)
}

// SetFeeAmount is a paid mutator transaction binding the contract method 0x6b392680.
//
// Solidity: function setFeeAmount(uint256 amount) returns()
func (_Contract *ContractTransactor) SetFeeAmount(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setFeeAmount", amount)
}

// SetFeeAmount is a paid mutator transaction binding the contract method 0x6b392680.
//
// Solidity: function setFeeAmount(uint256 amount) returns()
func (_Contract *ContractSession) SetFeeAmount(amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetFeeAmount(&_Contract.TransactOpts, amount)
}

// SetFeeAmount is a paid mutator transaction binding the contract method 0x6b392680.
//
// Solidity: function setFeeAmount(uint256 amount) returns()
func (_Contract *ContractTransactorSession) SetFeeAmount(amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetFeeAmount(&_Contract.TransactOpts, amount)
}

// SetGameTime is a paid mutator transaction binding the contract method 0xc388a962.
//
// Solidity: function setGameTime(uint64 time) returns()
func (_Contract *ContractTransactor) SetGameTime(opts *bind.TransactOpts, time uint64) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setGameTime", time)
}

// SetGameTime is a paid mutator transaction binding the contract method 0xc388a962.
//
// Solidity: function setGameTime(uint64 time) returns()
func (_Contract *ContractSession) SetGameTime(time uint64) (*types.Transaction, error) {
	return _Contract.Contract.SetGameTime(&_Contract.TransactOpts, time)
}

// SetGameTime is a paid mutator transaction binding the contract method 0xc388a962.
//
// Solidity: function setGameTime(uint64 time) returns()
func (_Contract *ContractTransactorSession) SetGameTime(time uint64) (*types.Transaction, error) {
	return _Contract.Contract.SetGameTime(&_Contract.TransactOpts, time)
}

// SetMaxPlayers is a paid mutator transaction binding the contract method 0x288dee3b.
//
// Solidity: function setMaxPlayers(uint256 players) returns()
func (_Contract *ContractTransactor) SetMaxPlayers(opts *bind.TransactOpts, players *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setMaxPlayers", players)
}

// SetMaxPlayers is a paid mutator transaction binding the contract method 0x288dee3b.
//
// Solidity: function setMaxPlayers(uint256 players) returns()
func (_Contract *ContractSession) SetMaxPlayers(players *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetMaxPlayers(&_Contract.TransactOpts, players)
}

// SetMaxPlayers is a paid mutator transaction binding the contract method 0x288dee3b.
//
// Solidity: function setMaxPlayers(uint256 players) returns()
func (_Contract *ContractTransactorSession) SetMaxPlayers(players *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetMaxPlayers(&_Contract.TransactOpts, players)
}

// SetModerator is a paid mutator transaction binding the contract method 0x75bba189.
//
// Solidity: function setModerator(address account) returns()
func (_Contract *ContractTransactor) SetModerator(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setModerator", account)
}

// SetModerator is a paid mutator transaction binding the contract method 0x75bba189.
//
// Solidity: function setModerator(address account) returns()
func (_Contract *ContractSession) SetModerator(account common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetModerator(&_Contract.TransactOpts, account)
}

// SetModerator is a paid mutator transaction binding the contract method 0x75bba189.
//
// Solidity: function setModerator(address account) returns()
func (_Contract *ContractTransactorSession) SetModerator(account common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetModerator(&_Contract.TransactOpts, account)
}

// SetWinner is a paid mutator transaction binding the contract method 0x9c623683.
//
// Solidity: function setWinner(uint256 gameId, address winner) returns()
func (_Contract *ContractTransactor) SetWinner(opts *bind.TransactOpts, gameId *big.Int, winner common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setWinner", gameId, winner)
}

// SetWinner is a paid mutator transaction binding the contract method 0x9c623683.
//
// Solidity: function setWinner(uint256 gameId, address winner) returns()
func (_Contract *ContractSession) SetWinner(gameId *big.Int, winner common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetWinner(&_Contract.TransactOpts, gameId, winner)
}

// SetWinner is a paid mutator transaction binding the contract method 0x9c623683.
//
// Solidity: function setWinner(uint256 gameId, address winner) returns()
func (_Contract *ContractTransactorSession) SetWinner(gameId *big.Int, winner common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetWinner(&_Contract.TransactOpts, gameId, winner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contract *ContractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contract *ContractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Contract.Contract.TransferOwnership(&_Contract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contract *ContractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Contract.Contract.TransferOwnership(&_Contract.TransactOpts, newOwner)
}

// ContractClaimRewardsIterator is returned from FilterClaimRewards and is used to iterate over the raw logs and unpacked data for ClaimRewards events raised by the Contract contract.
type ContractClaimRewardsIterator struct {
	Event *ContractClaimRewards // Event containing the contract specifics and raw log

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
func (it *ContractClaimRewardsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractClaimRewards)
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
		it.Event = new(ContractClaimRewards)
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
func (it *ContractClaimRewardsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractClaimRewardsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractClaimRewards represents a ClaimRewards event raised by the Contract contract.
type ContractClaimRewards struct {
	Winner common.Address
	Arg1   []DunkinCapsNFTs
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterClaimRewards is a free log retrieval operation binding the contract event 0x926374800001c35679302e7dde2ebf1e77f15281a064176b4ebf55b0e8800776.
//
// Solidity: event ClaimRewards(address winner, (address,uint256)[] arg1)
func (_Contract *ContractFilterer) FilterClaimRewards(opts *bind.FilterOpts) (*ContractClaimRewardsIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "ClaimRewards")
	if err != nil {
		return nil, err
	}
	return &ContractClaimRewardsIterator{contract: _Contract.contract, event: "ClaimRewards", logs: logs, sub: sub}, nil
}

// WatchClaimRewards is a free log subscription operation binding the contract event 0x926374800001c35679302e7dde2ebf1e77f15281a064176b4ebf55b0e8800776.
//
// Solidity: event ClaimRewards(address winner, (address,uint256)[] arg1)
func (_Contract *ContractFilterer) WatchClaimRewards(opts *bind.WatchOpts, sink chan<- *ContractClaimRewards) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "ClaimRewards")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractClaimRewards)
				if err := _Contract.contract.UnpackLog(event, "ClaimRewards", log); err != nil {
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

// ParseClaimRewards is a log parse operation binding the contract event 0x926374800001c35679302e7dde2ebf1e77f15281a064176b4ebf55b0e8800776.
//
// Solidity: event ClaimRewards(address winner, (address,uint256)[] arg1)
func (_Contract *ContractFilterer) ParseClaimRewards(log types.Log) (*ContractClaimRewards, error) {
	event := new(ContractClaimRewards)
	if err := _Contract.contract.UnpackLog(event, "ClaimRewards", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractCreateBetIterator is returned from FilterCreateBet and is used to iterate over the raw logs and unpacked data for CreateBet events raised by the Contract contract.
type ContractCreateBetIterator struct {
	Event *ContractCreateBet // Event containing the contract specifics and raw log

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
func (it *ContractCreateBetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractCreateBet)
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
		it.Event = new(ContractCreateBet)
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
func (it *ContractCreateBetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractCreateBetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractCreateBet represents a CreateBet event raised by the Contract contract.
type ContractCreateBet struct {
	GameId *big.Int
	Player common.Address
	Arg2   []DunkinCapsNFTs
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCreateBet is a free log retrieval operation binding the contract event 0xbe1d706fe3d6b6c52dcd493e2609abaa52f07fd9b36e92d691bd9aa47e0d9d48.
//
// Solidity: event CreateBet(uint256 gameId, address player, (address,uint256)[] arg2)
func (_Contract *ContractFilterer) FilterCreateBet(opts *bind.FilterOpts) (*ContractCreateBetIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "CreateBet")
	if err != nil {
		return nil, err
	}
	return &ContractCreateBetIterator{contract: _Contract.contract, event: "CreateBet", logs: logs, sub: sub}, nil
}

// WatchCreateBet is a free log subscription operation binding the contract event 0xbe1d706fe3d6b6c52dcd493e2609abaa52f07fd9b36e92d691bd9aa47e0d9d48.
//
// Solidity: event CreateBet(uint256 gameId, address player, (address,uint256)[] arg2)
func (_Contract *ContractFilterer) WatchCreateBet(opts *bind.WatchOpts, sink chan<- *ContractCreateBet) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "CreateBet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractCreateBet)
				if err := _Contract.contract.UnpackLog(event, "CreateBet", log); err != nil {
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

// ParseCreateBet is a log parse operation binding the contract event 0xbe1d706fe3d6b6c52dcd493e2609abaa52f07fd9b36e92d691bd9aa47e0d9d48.
//
// Solidity: event CreateBet(uint256 gameId, address player, (address,uint256)[] arg2)
func (_Contract *ContractFilterer) ParseCreateBet(log types.Log) (*ContractCreateBet, error) {
	event := new(ContractCreateBet)
	if err := _Contract.contract.UnpackLog(event, "CreateBet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractCreateGameIterator is returned from FilterCreateGame and is used to iterate over the raw logs and unpacked data for CreateGame events raised by the Contract contract.
type ContractCreateGameIterator struct {
	Event *ContractCreateGame // Event containing the contract specifics and raw log

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
func (it *ContractCreateGameIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractCreateGame)
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
		it.Event = new(ContractCreateGame)
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
func (it *ContractCreateGameIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractCreateGameIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractCreateGame represents a CreateGame event raised by the Contract contract.
type ContractCreateGame struct {
	GameId *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCreateGame is a free log retrieval operation binding the contract event 0x9e5237c6a8bb45df74799cf138ebda4fed0574bed906da41387acdc4a3cbe3d8.
//
// Solidity: event CreateGame(uint256 gameId)
func (_Contract *ContractFilterer) FilterCreateGame(opts *bind.FilterOpts) (*ContractCreateGameIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "CreateGame")
	if err != nil {
		return nil, err
	}
	return &ContractCreateGameIterator{contract: _Contract.contract, event: "CreateGame", logs: logs, sub: sub}, nil
}

// WatchCreateGame is a free log subscription operation binding the contract event 0x9e5237c6a8bb45df74799cf138ebda4fed0574bed906da41387acdc4a3cbe3d8.
//
// Solidity: event CreateGame(uint256 gameId)
func (_Contract *ContractFilterer) WatchCreateGame(opts *bind.WatchOpts, sink chan<- *ContractCreateGame) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "CreateGame")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractCreateGame)
				if err := _Contract.contract.UnpackLog(event, "CreateGame", log); err != nil {
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

// ParseCreateGame is a log parse operation binding the contract event 0x9e5237c6a8bb45df74799cf138ebda4fed0574bed906da41387acdc4a3cbe3d8.
//
// Solidity: event CreateGame(uint256 gameId)
func (_Contract *ContractFilterer) ParseCreateGame(log types.Log) (*ContractCreateGame, error) {
	event := new(ContractCreateGame)
	if err := _Contract.contract.UnpackLog(event, "CreateGame", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Contract contract.
type ContractOwnershipTransferredIterator struct {
	Event *ContractOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOwnershipTransferred)
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
		it.Event = new(ContractOwnershipTransferred)
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
func (it *ContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOwnershipTransferred represents a OwnershipTransferred event raised by the Contract contract.
type ContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contract *ContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractOwnershipTransferredIterator{contract: _Contract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contract *ContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOwnershipTransferred)
				if err := _Contract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contract *ContractFilterer) ParseOwnershipTransferred(log types.Log) (*ContractOwnershipTransferred, error) {
	event := new(ContractOwnershipTransferred)
	if err := _Contract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSetFeeAccountIterator is returned from FilterSetFeeAccount and is used to iterate over the raw logs and unpacked data for SetFeeAccount events raised by the Contract contract.
type ContractSetFeeAccountIterator struct {
	Event *ContractSetFeeAccount // Event containing the contract specifics and raw log

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
func (it *ContractSetFeeAccountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSetFeeAccount)
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
		it.Event = new(ContractSetFeeAccount)
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
func (it *ContractSetFeeAccountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSetFeeAccountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSetFeeAccount represents a SetFeeAccount event raised by the Contract contract.
type ContractSetFeeAccount struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSetFeeAccount is a free log retrieval operation binding the contract event 0xe7bc311385e20a1c065058e77945d29ce9498f0673c76224d11014e049724f63.
//
// Solidity: event SetFeeAccount(address account)
func (_Contract *ContractFilterer) FilterSetFeeAccount(opts *bind.FilterOpts) (*ContractSetFeeAccountIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "SetFeeAccount")
	if err != nil {
		return nil, err
	}
	return &ContractSetFeeAccountIterator{contract: _Contract.contract, event: "SetFeeAccount", logs: logs, sub: sub}, nil
}

// WatchSetFeeAccount is a free log subscription operation binding the contract event 0xe7bc311385e20a1c065058e77945d29ce9498f0673c76224d11014e049724f63.
//
// Solidity: event SetFeeAccount(address account)
func (_Contract *ContractFilterer) WatchSetFeeAccount(opts *bind.WatchOpts, sink chan<- *ContractSetFeeAccount) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "SetFeeAccount")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSetFeeAccount)
				if err := _Contract.contract.UnpackLog(event, "SetFeeAccount", log); err != nil {
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

// ParseSetFeeAccount is a log parse operation binding the contract event 0xe7bc311385e20a1c065058e77945d29ce9498f0673c76224d11014e049724f63.
//
// Solidity: event SetFeeAccount(address account)
func (_Contract *ContractFilterer) ParseSetFeeAccount(log types.Log) (*ContractSetFeeAccount, error) {
	event := new(ContractSetFeeAccount)
	if err := _Contract.contract.UnpackLog(event, "SetFeeAccount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSetFeeAmountIterator is returned from FilterSetFeeAmount and is used to iterate over the raw logs and unpacked data for SetFeeAmount events raised by the Contract contract.
type ContractSetFeeAmountIterator struct {
	Event *ContractSetFeeAmount // Event containing the contract specifics and raw log

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
func (it *ContractSetFeeAmountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSetFeeAmount)
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
		it.Event = new(ContractSetFeeAmount)
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
func (it *ContractSetFeeAmountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSetFeeAmountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSetFeeAmount represents a SetFeeAmount event raised by the Contract contract.
type ContractSetFeeAmount struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetFeeAmount is a free log retrieval operation binding the contract event 0xd854f702b4e23ab27557a75b5b4e72066f8ab6929930f813720339c72eaa8b63.
//
// Solidity: event SetFeeAmount(uint256 amount)
func (_Contract *ContractFilterer) FilterSetFeeAmount(opts *bind.FilterOpts) (*ContractSetFeeAmountIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "SetFeeAmount")
	if err != nil {
		return nil, err
	}
	return &ContractSetFeeAmountIterator{contract: _Contract.contract, event: "SetFeeAmount", logs: logs, sub: sub}, nil
}

// WatchSetFeeAmount is a free log subscription operation binding the contract event 0xd854f702b4e23ab27557a75b5b4e72066f8ab6929930f813720339c72eaa8b63.
//
// Solidity: event SetFeeAmount(uint256 amount)
func (_Contract *ContractFilterer) WatchSetFeeAmount(opts *bind.WatchOpts, sink chan<- *ContractSetFeeAmount) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "SetFeeAmount")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSetFeeAmount)
				if err := _Contract.contract.UnpackLog(event, "SetFeeAmount", log); err != nil {
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

// ParseSetFeeAmount is a log parse operation binding the contract event 0xd854f702b4e23ab27557a75b5b4e72066f8ab6929930f813720339c72eaa8b63.
//
// Solidity: event SetFeeAmount(uint256 amount)
func (_Contract *ContractFilterer) ParseSetFeeAmount(log types.Log) (*ContractSetFeeAmount, error) {
	event := new(ContractSetFeeAmount)
	if err := _Contract.contract.UnpackLog(event, "SetFeeAmount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSetGameTimeIterator is returned from FilterSetGameTime and is used to iterate over the raw logs and unpacked data for SetGameTime events raised by the Contract contract.
type ContractSetGameTimeIterator struct {
	Event *ContractSetGameTime // Event containing the contract specifics and raw log

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
func (it *ContractSetGameTimeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSetGameTime)
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
		it.Event = new(ContractSetGameTime)
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
func (it *ContractSetGameTimeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSetGameTimeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSetGameTime represents a SetGameTime event raised by the Contract contract.
type ContractSetGameTime struct {
	GameTime uint64
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetGameTime is a free log retrieval operation binding the contract event 0x820fd422c12e33cb949e60e8654075192a5b8ef40a0dd0670d93a7eb7cacf582.
//
// Solidity: event SetGameTime(uint64 gameTime)
func (_Contract *ContractFilterer) FilterSetGameTime(opts *bind.FilterOpts) (*ContractSetGameTimeIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "SetGameTime")
	if err != nil {
		return nil, err
	}
	return &ContractSetGameTimeIterator{contract: _Contract.contract, event: "SetGameTime", logs: logs, sub: sub}, nil
}

// WatchSetGameTime is a free log subscription operation binding the contract event 0x820fd422c12e33cb949e60e8654075192a5b8ef40a0dd0670d93a7eb7cacf582.
//
// Solidity: event SetGameTime(uint64 gameTime)
func (_Contract *ContractFilterer) WatchSetGameTime(opts *bind.WatchOpts, sink chan<- *ContractSetGameTime) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "SetGameTime")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSetGameTime)
				if err := _Contract.contract.UnpackLog(event, "SetGameTime", log); err != nil {
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

// ParseSetGameTime is a log parse operation binding the contract event 0x820fd422c12e33cb949e60e8654075192a5b8ef40a0dd0670d93a7eb7cacf582.
//
// Solidity: event SetGameTime(uint64 gameTime)
func (_Contract *ContractFilterer) ParseSetGameTime(log types.Log) (*ContractSetGameTime, error) {
	event := new(ContractSetGameTime)
	if err := _Contract.contract.UnpackLog(event, "SetGameTime", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSetMaxPlayersIterator is returned from FilterSetMaxPlayers and is used to iterate over the raw logs and unpacked data for SetMaxPlayers events raised by the Contract contract.
type ContractSetMaxPlayersIterator struct {
	Event *ContractSetMaxPlayers // Event containing the contract specifics and raw log

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
func (it *ContractSetMaxPlayersIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSetMaxPlayers)
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
		it.Event = new(ContractSetMaxPlayers)
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
func (it *ContractSetMaxPlayersIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSetMaxPlayersIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSetMaxPlayers represents a SetMaxPlayers event raised by the Contract contract.
type ContractSetMaxPlayers struct {
	MaxPlayers *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSetMaxPlayers is a free log retrieval operation binding the contract event 0x8b23093d10e1d32132cd0d2c5a8cf21c09d4fe3451eb9ffe934234329ea4052b.
//
// Solidity: event SetMaxPlayers(uint256 maxPlayers)
func (_Contract *ContractFilterer) FilterSetMaxPlayers(opts *bind.FilterOpts) (*ContractSetMaxPlayersIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "SetMaxPlayers")
	if err != nil {
		return nil, err
	}
	return &ContractSetMaxPlayersIterator{contract: _Contract.contract, event: "SetMaxPlayers", logs: logs, sub: sub}, nil
}

// WatchSetMaxPlayers is a free log subscription operation binding the contract event 0x8b23093d10e1d32132cd0d2c5a8cf21c09d4fe3451eb9ffe934234329ea4052b.
//
// Solidity: event SetMaxPlayers(uint256 maxPlayers)
func (_Contract *ContractFilterer) WatchSetMaxPlayers(opts *bind.WatchOpts, sink chan<- *ContractSetMaxPlayers) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "SetMaxPlayers")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSetMaxPlayers)
				if err := _Contract.contract.UnpackLog(event, "SetMaxPlayers", log); err != nil {
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

// ParseSetMaxPlayers is a log parse operation binding the contract event 0x8b23093d10e1d32132cd0d2c5a8cf21c09d4fe3451eb9ffe934234329ea4052b.
//
// Solidity: event SetMaxPlayers(uint256 maxPlayers)
func (_Contract *ContractFilterer) ParseSetMaxPlayers(log types.Log) (*ContractSetMaxPlayers, error) {
	event := new(ContractSetMaxPlayers)
	if err := _Contract.contract.UnpackLog(event, "SetMaxPlayers", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSetModeratorIterator is returned from FilterSetModerator and is used to iterate over the raw logs and unpacked data for SetModerator events raised by the Contract contract.
type ContractSetModeratorIterator struct {
	Event *ContractSetModerator // Event containing the contract specifics and raw log

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
func (it *ContractSetModeratorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSetModerator)
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
		it.Event = new(ContractSetModerator)
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
func (it *ContractSetModeratorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSetModeratorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSetModerator represents a SetModerator event raised by the Contract contract.
type ContractSetModerator struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSetModerator is a free log retrieval operation binding the contract event 0x9fdc9d265e423894c26d09e9c37af7f37111ed08c977853425c94b61c2d23721.
//
// Solidity: event SetModerator(address account)
func (_Contract *ContractFilterer) FilterSetModerator(opts *bind.FilterOpts) (*ContractSetModeratorIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "SetModerator")
	if err != nil {
		return nil, err
	}
	return &ContractSetModeratorIterator{contract: _Contract.contract, event: "SetModerator", logs: logs, sub: sub}, nil
}

// WatchSetModerator is a free log subscription operation binding the contract event 0x9fdc9d265e423894c26d09e9c37af7f37111ed08c977853425c94b61c2d23721.
//
// Solidity: event SetModerator(address account)
func (_Contract *ContractFilterer) WatchSetModerator(opts *bind.WatchOpts, sink chan<- *ContractSetModerator) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "SetModerator")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSetModerator)
				if err := _Contract.contract.UnpackLog(event, "SetModerator", log); err != nil {
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

// ParseSetModerator is a log parse operation binding the contract event 0x9fdc9d265e423894c26d09e9c37af7f37111ed08c977853425c94b61c2d23721.
//
// Solidity: event SetModerator(address account)
func (_Contract *ContractFilterer) ParseSetModerator(log types.Log) (*ContractSetModerator, error) {
	event := new(ContractSetModerator)
	if err := _Contract.contract.UnpackLog(event, "SetModerator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSetWinnerIterator is returned from FilterSetWinner and is used to iterate over the raw logs and unpacked data for SetWinner events raised by the Contract contract.
type ContractSetWinnerIterator struct {
	Event *ContractSetWinner // Event containing the contract specifics and raw log

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
func (it *ContractSetWinnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSetWinner)
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
		it.Event = new(ContractSetWinner)
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
func (it *ContractSetWinnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSetWinnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSetWinner represents a SetWinner event raised by the Contract contract.
type ContractSetWinner struct {
	Winner common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetWinner is a free log retrieval operation binding the contract event 0x29ce87337dc1aec5607c8ceed6acbc353fa260c7bd1628d1081fd2d464a5fbc1.
//
// Solidity: event SetWinner(address winner)
func (_Contract *ContractFilterer) FilterSetWinner(opts *bind.FilterOpts) (*ContractSetWinnerIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "SetWinner")
	if err != nil {
		return nil, err
	}
	return &ContractSetWinnerIterator{contract: _Contract.contract, event: "SetWinner", logs: logs, sub: sub}, nil
}

// WatchSetWinner is a free log subscription operation binding the contract event 0x29ce87337dc1aec5607c8ceed6acbc353fa260c7bd1628d1081fd2d464a5fbc1.
//
// Solidity: event SetWinner(address winner)
func (_Contract *ContractFilterer) WatchSetWinner(opts *bind.WatchOpts, sink chan<- *ContractSetWinner) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "SetWinner")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSetWinner)
				if err := _Contract.contract.UnpackLog(event, "SetWinner", log); err != nil {
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

// ParseSetWinner is a log parse operation binding the contract event 0x29ce87337dc1aec5607c8ceed6acbc353fa260c7bd1628d1081fd2d464a5fbc1.
//
// Solidity: event SetWinner(address winner)
func (_Contract *ContractFilterer) ParseSetWinner(log types.Log) (*ContractSetWinner, error) {
	event := new(ContractSetWinner)
	if err := _Contract.contract.UnpackLog(event, "SetWinner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
