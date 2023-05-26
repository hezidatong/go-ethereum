// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package erc4907

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

// Erc4907MetaData contains all meta data concerning the Erc4907 contract.
var Erc4907MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"expires\",\"type\":\"uint64\"}],\"name\":\"UpdateUser\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"expires\",\"type\":\"uint64\"}],\"name\":\"setUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"userExpires\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"userOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162001c2938038062001c29833981016040819052620000349162000127565b81818181600062000046838262000220565b50600162000055828262000220565b50505050505050620002ec565b634e487b7160e01b600052604160045260246000fd5b600082601f8301126200008a57600080fd5b81516001600160401b0380821115620000a757620000a762000062565b604051601f8301601f19908116603f01168101908282118183101715620000d257620000d262000062565b81604052838152602092508683858801011115620000ef57600080fd5b600091505b83821015620001135785820183015181830184015290820190620000f4565b600093810190920192909252949350505050565b600080604083850312156200013b57600080fd5b82516001600160401b03808211156200015357600080fd5b620001618683870162000078565b935060208501519150808211156200017857600080fd5b50620001878582860162000078565b9150509250929050565b600181811c90821680620001a657607f821691505b602082108103620001c757634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156200021b57600081815260208120601f850160051c81016020861015620001f65750805b601f850160051c820191505b81811015620002175782815560010162000202565b5050505b505050565b81516001600160401b038111156200023c576200023c62000062565b62000254816200024d845462000191565b84620001cd565b602080601f8311600181146200028c5760008415620002735750858301515b600019600386901b1c1916600185901b17855562000217565b600085815260208120601f198616915b82811015620002bd578886015182559484019460019091019084016200029c565b5085821015620002dc5787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b61192d80620002fc6000396000f3fe608060405234801561001057600080fd5b506004361061011b5760003560e01c80638fc88c48116100b2578063b88d4fde11610081578063c87b56dd11610066578063c87b56dd1461028d578063e030565e146102a0578063e985e9c5146102b357600080fd5b8063b88d4fde14610267578063c2f1f14a1461027a57600080fd5b80638fc88c48146101f757806394bf804d1461023957806395d89b411461024c578063a22cb4651461025457600080fd5b806323b872dd116100ee57806323b872dd1461019d57806342842e0e146101b05780636352211e146101c357806370a08231146101d657600080fd5b806301ffc9a71461012057806306fdde0314610148578063081812fc1461015d578063095ea7b314610188575b600080fd5b61013361012e36600461146b565b6102ef565b60405190151581526020015b60405180910390f35b61015061034b565b60405161013f91906114d8565b61017061016b3660046114eb565b6103dd565b6040516001600160a01b03909116815260200161013f565b61019b61019636600461151b565b610404565b005b61019b6101ab366004611545565b610558565b61019b6101be366004611545565b6105df565b6101706101d13660046114eb565b6105fa565b6101e96101e4366004611581565b61065f565b60405190815260200161013f565b6101e96102053660046114eb565b60009081526006602052604090205474010000000000000000000000000000000000000000900467ffffffffffffffff1690565b61019b61024736600461159c565b6106f9565b610150610707565b61019b6102623660046115c8565b610716565b61019b61027536600461161a565b610721565b6101706102883660046114eb565b6107af565b61015061029b3660046114eb565b610810565b61019b6102ae3660046116f6565b610884565b6101336102c1366004611743565b6001600160a01b03918216600090815260056020908152604080832093909416825291909152205460ff1690565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167fad092b5c0000000000000000000000000000000000000000000000000000000014806103455750610345826109aa565b92915050565b60606000805461035a9061176d565b80601f01602080910402602001604051908101604052809291908181526020018280546103869061176d565b80156103d35780601f106103a8576101008083540402835291602001916103d3565b820191906000526020600020905b8154815290600101906020018083116103b657829003601f168201915b5050505050905090565b60006103e882610a8d565b506000908152600460205260409020546001600160a01b031690565b600061040f826105fa565b9050806001600160a01b0316836001600160a01b03160361049d5760405162461bcd60e51b815260206004820152602160248201527f4552433732313a20617070726f76616c20746f2063757272656e74206f776e6560448201527f720000000000000000000000000000000000000000000000000000000000000060648201526084015b60405180910390fd5b336001600160a01b03821614806104d757506001600160a01b038116600090815260056020908152604080832033845290915290205460ff165b6105495760405162461bcd60e51b815260206004820152603e60248201527f4552433732313a20617070726f76652063616c6c6572206973206e6f7420746f60448201527f6b656e206f776e6572206e6f7220617070726f76656420666f7220616c6c00006064820152608401610494565b6105538383610af4565b505050565b6105623382610b7a565b6105d45760405162461bcd60e51b815260206004820152602e60248201527f4552433732313a2063616c6c6572206973206e6f7420746f6b656e206f776e6560448201527f72206e6f7220617070726f7665640000000000000000000000000000000000006064820152608401610494565b610553838383610bf9565b61055383838360405180602001604052806000815250610721565b6000818152600260205260408120546001600160a01b0316806103455760405162461bcd60e51b815260206004820152601860248201527f4552433732313a20696e76616c696420746f6b656e20494400000000000000006044820152606401610494565b60006001600160a01b0382166106dd5760405162461bcd60e51b815260206004820152602960248201527f4552433732313a2061646472657373207a65726f206973206e6f74206120766160448201527f6c6964206f776e657200000000000000000000000000000000000000000000006064820152608401610494565b506001600160a01b031660009081526003602052604090205490565b6107038183610de9565b5050565b60606001805461035a9061176d565b610703338383610f4f565b61072b3383610b7a565b61079d5760405162461bcd60e51b815260206004820152602e60248201527f4552433732313a2063616c6c6572206973206e6f7420746f6b656e206f776e6560448201527f72206e6f7220617070726f7665640000000000000000000000000000000000006064820152608401610494565b6107a98484848461103c565b50505050565b600081815260066020526040812054427401000000000000000000000000000000000000000090910467ffffffffffffffff161061080357506000908152600660205260409020546001600160a01b031690565b506000919050565b919050565b606061081b82610a8d565b600061083260408051602081019091526000815290565b90506000815111610852576040518060200160405280600081525061087d565b8061085c846110c5565b60405160200161086d9291906117a7565b6040516020818303038152906040525b9392505050565b61088e3384610b7a565b6109005760405162461bcd60e51b815260206004820152603260248201527f455243343930373a207472616e736665722063616c6c6572206973206e6f742060448201527f6f776e6572206e6f7220617070726f76656400000000000000000000000000006064820152608401610494565b60008381526006602090815260409182902080546001600160a01b0386167fffffffff0000000000000000000000000000000000000000000000000000000090911681177401000000000000000000000000000000000000000067ffffffffffffffff871690810291909117835593519384529092909186917f4e06b4e7000e659094299b3533b47b6aa8ad048e95e872d23d1f4ee55af89cfe910160405180910390a350505050565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f80ac58cd000000000000000000000000000000000000000000000000000000001480610a3d57507fffffffff0000000000000000000000000000000000000000000000000000000082167f5b5e139f00000000000000000000000000000000000000000000000000000000145b8061034557507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff00000000000000000000000000000000000000000000000000000000831614610345565b6000818152600260205260409020546001600160a01b0316610af15760405162461bcd60e51b815260206004820152601860248201527f4552433732313a20696e76616c696420746f6b656e20494400000000000000006044820152606401610494565b50565b600081815260046020526040902080547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0384169081179091558190610b41826105fa565b6001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560405160405180910390a45050565b600080610b86836105fa565b9050806001600160a01b0316846001600160a01b03161480610bcd57506001600160a01b0380821660009081526005602090815260408083209388168352929052205460ff165b80610bf15750836001600160a01b0316610be6846103dd565b6001600160a01b0316145b949350505050565b826001600160a01b0316610c0c826105fa565b6001600160a01b031614610c885760405162461bcd60e51b815260206004820152602560248201527f4552433732313a207472616e736665722066726f6d20696e636f72726563742060448201527f6f776e65720000000000000000000000000000000000000000000000000000006064820152608401610494565b6001600160a01b038216610d035760405162461bcd60e51b8152602060048201526024808201527f4552433732313a207472616e7366657220746f20746865207a65726f2061646460448201527f72657373000000000000000000000000000000000000000000000000000000006064820152608401610494565b610d0e8383836111fa565b610d19600082610af4565b6001600160a01b0383166000908152600360205260408120805460019290610d429084906117ec565b90915550506001600160a01b0382166000908152600360205260408120805460019290610d709084906117ff565b909155505060008181526002602052604080822080547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0386811691821790925591518493918716917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef91a4505050565b6001600160a01b038216610e3f5760405162461bcd60e51b815260206004820181905260248201527f4552433732313a206d696e7420746f20746865207a65726f20616464726573736044820152606401610494565b6000818152600260205260409020546001600160a01b031615610ea45760405162461bcd60e51b815260206004820152601c60248201527f4552433732313a20746f6b656e20616c7265616479206d696e746564000000006044820152606401610494565b610eb0600083836111fa565b6001600160a01b0382166000908152600360205260408120805460019290610ed99084906117ff565b909155505060008181526002602052604080822080547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b03861690811790915590518392907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef908290a45050565b816001600160a01b0316836001600160a01b031603610fb05760405162461bcd60e51b815260206004820152601960248201527f4552433732313a20617070726f766520746f2063616c6c6572000000000000006044820152606401610494565b6001600160a01b0383811660008181526005602090815260408083209487168084529482529182902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001686151590811790915591519182527f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c3191015b60405180910390a3505050565b611047848484610bf9565b6110538484848461129c565b6107a95760405162461bcd60e51b815260206004820152603260248201527f4552433732313a207472616e7366657220746f206e6f6e20455243373231526560448201527f63656976657220696d706c656d656e74657200000000000000000000000000006064820152608401610494565b60608160000361110857505060408051808201909152600181527f3000000000000000000000000000000000000000000000000000000000000000602082015290565b8160005b8115611132578061111c81611812565b915061112b9050600a83611860565b915061110c565b60008167ffffffffffffffff81111561114d5761114d611604565b6040519080825280601f01601f191660200182016040528015611177576020820181803683370190505b5090505b8415610bf15761118c6001836117ec565b9150611199600a86611874565b6111a49060306117ff565b60f81b8183815181106111b9576111b9611888565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053506111f3600a86611860565b945061117b565b816001600160a01b0316836001600160a01b03161415801561123257506000818152600660205260409020546001600160a01b031615155b1561055357600081815260066020908152604080832080547fffffffff000000000000000000000000000000000000000000000000000000001690555182815283917f4e06b4e7000e659094299b3533b47b6aa8ad048e95e872d23d1f4ee55af89cfe910161102f565b60006001600160a01b0384163b15611432576040517f150b7a020000000000000000000000000000000000000000000000000000000081526001600160a01b0385169063150b7a02906112f990339089908890889060040161189e565b6020604051808303816000875af1925050508015611334575060408051601f3d908101601f19168201909252611331918101906118da565b60015b6113e7573d808015611362576040519150601f19603f3d011682016040523d82523d6000602084013e611367565b606091505b5080516000036113df5760405162461bcd60e51b815260206004820152603260248201527f4552433732313a207472616e7366657220746f206e6f6e20455243373231526560448201527f63656976657220696d706c656d656e74657200000000000000000000000000006064820152608401610494565b805181602001fd5b7fffffffff00000000000000000000000000000000000000000000000000000000167f150b7a0200000000000000000000000000000000000000000000000000000000149050610bf1565b506001949350505050565b7fffffffff0000000000000000000000000000000000000000000000000000000081168114610af157600080fd5b60006020828403121561147d57600080fd5b813561087d8161143d565b60005b838110156114a357818101518382015260200161148b565b50506000910152565b600081518084526114c4816020860160208601611488565b601f01601f19169290920160200192915050565b60208152600061087d60208301846114ac565b6000602082840312156114fd57600080fd5b5035919050565b80356001600160a01b038116811461080b57600080fd5b6000806040838503121561152e57600080fd5b61153783611504565b946020939093013593505050565b60008060006060848603121561155a57600080fd5b61156384611504565b925061157160208501611504565b9150604084013590509250925092565b60006020828403121561159357600080fd5b61087d82611504565b600080604083850312156115af57600080fd5b823591506115bf60208401611504565b90509250929050565b600080604083850312156115db57600080fd5b6115e483611504565b9150602083013580151581146115f957600080fd5b809150509250929050565b634e487b7160e01b600052604160045260246000fd5b6000806000806080858703121561163057600080fd5b61163985611504565b935061164760208601611504565b925060408501359150606085013567ffffffffffffffff8082111561166b57600080fd5b818701915087601f83011261167f57600080fd5b81358181111561169157611691611604565b604051601f8201601f19908116603f011681019083821181831017156116b9576116b9611604565b816040528281528a60208487010111156116d257600080fd5b82602086016020830137600060208483010152809550505050505092959194509250565b60008060006060848603121561170b57600080fd5b8335925061171b60208501611504565b9150604084013567ffffffffffffffff8116811461173857600080fd5b809150509250925092565b6000806040838503121561175657600080fd5b61175f83611504565b91506115bf60208401611504565b600181811c9082168061178157607f821691505b6020821081036117a157634e487b7160e01b600052602260045260246000fd5b50919050565b600083516117b9818460208801611488565b8351908301906117cd818360208801611488565b01949350505050565b634e487b7160e01b600052601160045260246000fd5b81810381811115610345576103456117d6565b80820180821115610345576103456117d6565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611843576118436117d6565b5060010190565b634e487b7160e01b600052601260045260246000fd5b60008261186f5761186f61184a565b500490565b6000826118835761188361184a565b500690565b634e487b7160e01b600052603260045260246000fd5b60006001600160a01b038087168352808616602084015250836040830152608060608301526118d060808301846114ac565b9695505050505050565b6000602082840312156118ec57600080fd5b815161087d8161143d56fea26469706673582212203edd9b1cf8269b371393c3db815d4a69136e4dcb6f33ba21f9d1035a826aa84864736f6c63430008120033",
}

// Erc4907ABI is the input ABI used to generate the binding from.
// Deprecated: Use Erc4907MetaData.ABI instead.
var Erc4907ABI = Erc4907MetaData.ABI

// Erc4907Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use Erc4907MetaData.Bin instead.
var Erc4907Bin = Erc4907MetaData.Bin

// DeployErc4907 deploys a new Ethereum contract, binding an instance of Erc4907 to it.
func DeployErc4907(auth *bind.TransactOpts, backend bind.ContractBackend, name_ string, symbol_ string) (common.Address, *types.Transaction, *Erc4907, error) {
	parsed, err := Erc4907MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(Erc4907Bin), backend, name_, symbol_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Erc4907{Erc4907Caller: Erc4907Caller{contract: contract}, Erc4907Transactor: Erc4907Transactor{contract: contract}, Erc4907Filterer: Erc4907Filterer{contract: contract}}, nil
}

// Erc4907 is an auto generated Go binding around an Ethereum contract.
type Erc4907 struct {
	Erc4907Caller     // Read-only binding to the contract
	Erc4907Transactor // Write-only binding to the contract
	Erc4907Filterer   // Log filterer for contract events
}

// Erc4907Caller is an auto generated read-only Go binding around an Ethereum contract.
type Erc4907Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc4907Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Erc4907Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc4907Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Erc4907Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc4907Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Erc4907Session struct {
	Contract     *Erc4907          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Erc4907CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Erc4907CallerSession struct {
	Contract *Erc4907Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// Erc4907TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Erc4907TransactorSession struct {
	Contract     *Erc4907Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// Erc4907Raw is an auto generated low-level Go binding around an Ethereum contract.
type Erc4907Raw struct {
	Contract *Erc4907 // Generic contract binding to access the raw methods on
}

// Erc4907CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Erc4907CallerRaw struct {
	Contract *Erc4907Caller // Generic read-only contract binding to access the raw methods on
}

// Erc4907TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Erc4907TransactorRaw struct {
	Contract *Erc4907Transactor // Generic write-only contract binding to access the raw methods on
}

// NewErc4907 creates a new instance of Erc4907, bound to a specific deployed contract.
func NewErc4907(address common.Address, backend bind.ContractBackend) (*Erc4907, error) {
	contract, err := bindErc4907(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Erc4907{Erc4907Caller: Erc4907Caller{contract: contract}, Erc4907Transactor: Erc4907Transactor{contract: contract}, Erc4907Filterer: Erc4907Filterer{contract: contract}}, nil
}

// NewErc4907Caller creates a new read-only instance of Erc4907, bound to a specific deployed contract.
func NewErc4907Caller(address common.Address, caller bind.ContractCaller) (*Erc4907Caller, error) {
	contract, err := bindErc4907(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Erc4907Caller{contract: contract}, nil
}

// NewErc4907Transactor creates a new write-only instance of Erc4907, bound to a specific deployed contract.
func NewErc4907Transactor(address common.Address, transactor bind.ContractTransactor) (*Erc4907Transactor, error) {
	contract, err := bindErc4907(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Erc4907Transactor{contract: contract}, nil
}

// NewErc4907Filterer creates a new log filterer instance of Erc4907, bound to a specific deployed contract.
func NewErc4907Filterer(address common.Address, filterer bind.ContractFilterer) (*Erc4907Filterer, error) {
	contract, err := bindErc4907(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Erc4907Filterer{contract: contract}, nil
}

// bindErc4907 binds a generic wrapper to an already deployed contract.
func bindErc4907(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Erc4907MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc4907 *Erc4907Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Erc4907.Contract.Erc4907Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc4907 *Erc4907Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc4907.Contract.Erc4907Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc4907 *Erc4907Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc4907.Contract.Erc4907Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc4907 *Erc4907CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Erc4907.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc4907 *Erc4907TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc4907.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc4907 *Erc4907TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc4907.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Erc4907 *Erc4907Caller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Erc4907.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Erc4907 *Erc4907Session) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Erc4907.Contract.BalanceOf(&_Erc4907.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Erc4907 *Erc4907CallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Erc4907.Contract.BalanceOf(&_Erc4907.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Erc4907 *Erc4907Caller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Erc4907.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Erc4907 *Erc4907Session) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Erc4907.Contract.GetApproved(&_Erc4907.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Erc4907 *Erc4907CallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Erc4907.Contract.GetApproved(&_Erc4907.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Erc4907 *Erc4907Caller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Erc4907.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Erc4907 *Erc4907Session) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Erc4907.Contract.IsApprovedForAll(&_Erc4907.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Erc4907 *Erc4907CallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Erc4907.Contract.IsApprovedForAll(&_Erc4907.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Erc4907 *Erc4907Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Erc4907.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Erc4907 *Erc4907Session) Name() (string, error) {
	return _Erc4907.Contract.Name(&_Erc4907.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Erc4907 *Erc4907CallerSession) Name() (string, error) {
	return _Erc4907.Contract.Name(&_Erc4907.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Erc4907 *Erc4907Caller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Erc4907.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Erc4907 *Erc4907Session) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Erc4907.Contract.OwnerOf(&_Erc4907.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Erc4907 *Erc4907CallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Erc4907.Contract.OwnerOf(&_Erc4907.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Erc4907 *Erc4907Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Erc4907.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Erc4907 *Erc4907Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Erc4907.Contract.SupportsInterface(&_Erc4907.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Erc4907 *Erc4907CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Erc4907.Contract.SupportsInterface(&_Erc4907.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Erc4907 *Erc4907Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Erc4907.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Erc4907 *Erc4907Session) Symbol() (string, error) {
	return _Erc4907.Contract.Symbol(&_Erc4907.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Erc4907 *Erc4907CallerSession) Symbol() (string, error) {
	return _Erc4907.Contract.Symbol(&_Erc4907.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Erc4907 *Erc4907Caller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Erc4907.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Erc4907 *Erc4907Session) TokenURI(tokenId *big.Int) (string, error) {
	return _Erc4907.Contract.TokenURI(&_Erc4907.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Erc4907 *Erc4907CallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Erc4907.Contract.TokenURI(&_Erc4907.CallOpts, tokenId)
}

// UserExpires is a free data retrieval call binding the contract method 0x8fc88c48.
//
// Solidity: function userExpires(uint256 tokenId) view returns(uint256)
func (_Erc4907 *Erc4907Caller) UserExpires(opts *bind.CallOpts, tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Erc4907.contract.Call(opts, &out, "userExpires", tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserExpires is a free data retrieval call binding the contract method 0x8fc88c48.
//
// Solidity: function userExpires(uint256 tokenId) view returns(uint256)
func (_Erc4907 *Erc4907Session) UserExpires(tokenId *big.Int) (*big.Int, error) {
	return _Erc4907.Contract.UserExpires(&_Erc4907.CallOpts, tokenId)
}

// UserExpires is a free data retrieval call binding the contract method 0x8fc88c48.
//
// Solidity: function userExpires(uint256 tokenId) view returns(uint256)
func (_Erc4907 *Erc4907CallerSession) UserExpires(tokenId *big.Int) (*big.Int, error) {
	return _Erc4907.Contract.UserExpires(&_Erc4907.CallOpts, tokenId)
}

// UserOf is a free data retrieval call binding the contract method 0xc2f1f14a.
//
// Solidity: function userOf(uint256 tokenId) view returns(address)
func (_Erc4907 *Erc4907Caller) UserOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Erc4907.contract.Call(opts, &out, "userOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UserOf is a free data retrieval call binding the contract method 0xc2f1f14a.
//
// Solidity: function userOf(uint256 tokenId) view returns(address)
func (_Erc4907 *Erc4907Session) UserOf(tokenId *big.Int) (common.Address, error) {
	return _Erc4907.Contract.UserOf(&_Erc4907.CallOpts, tokenId)
}

// UserOf is a free data retrieval call binding the contract method 0xc2f1f14a.
//
// Solidity: function userOf(uint256 tokenId) view returns(address)
func (_Erc4907 *Erc4907CallerSession) UserOf(tokenId *big.Int) (common.Address, error) {
	return _Erc4907.Contract.UserOf(&_Erc4907.CallOpts, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Erc4907 *Erc4907Transactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Erc4907.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Erc4907 *Erc4907Session) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Erc4907.Contract.Approve(&_Erc4907.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Erc4907 *Erc4907TransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Erc4907.Contract.Approve(&_Erc4907.TransactOpts, to, tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 tokenId, address to) returns()
func (_Erc4907 *Erc4907Transactor) Mint(opts *bind.TransactOpts, tokenId *big.Int, to common.Address) (*types.Transaction, error) {
	return _Erc4907.contract.Transact(opts, "mint", tokenId, to)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 tokenId, address to) returns()
func (_Erc4907 *Erc4907Session) Mint(tokenId *big.Int, to common.Address) (*types.Transaction, error) {
	return _Erc4907.Contract.Mint(&_Erc4907.TransactOpts, tokenId, to)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 tokenId, address to) returns()
func (_Erc4907 *Erc4907TransactorSession) Mint(tokenId *big.Int, to common.Address) (*types.Transaction, error) {
	return _Erc4907.Contract.Mint(&_Erc4907.TransactOpts, tokenId, to)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Erc4907 *Erc4907Transactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Erc4907.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Erc4907 *Erc4907Session) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Erc4907.Contract.SafeTransferFrom(&_Erc4907.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Erc4907 *Erc4907TransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Erc4907.Contract.SafeTransferFrom(&_Erc4907.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Erc4907 *Erc4907Transactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Erc4907.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Erc4907 *Erc4907Session) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Erc4907.Contract.SafeTransferFrom0(&_Erc4907.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Erc4907 *Erc4907TransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Erc4907.Contract.SafeTransferFrom0(&_Erc4907.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Erc4907 *Erc4907Transactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Erc4907.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Erc4907 *Erc4907Session) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Erc4907.Contract.SetApprovalForAll(&_Erc4907.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Erc4907 *Erc4907TransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Erc4907.Contract.SetApprovalForAll(&_Erc4907.TransactOpts, operator, approved)
}

// SetUser is a paid mutator transaction binding the contract method 0xe030565e.
//
// Solidity: function setUser(uint256 tokenId, address user, uint64 expires) returns()
func (_Erc4907 *Erc4907Transactor) SetUser(opts *bind.TransactOpts, tokenId *big.Int, user common.Address, expires uint64) (*types.Transaction, error) {
	return _Erc4907.contract.Transact(opts, "setUser", tokenId, user, expires)
}

// SetUser is a paid mutator transaction binding the contract method 0xe030565e.
//
// Solidity: function setUser(uint256 tokenId, address user, uint64 expires) returns()
func (_Erc4907 *Erc4907Session) SetUser(tokenId *big.Int, user common.Address, expires uint64) (*types.Transaction, error) {
	return _Erc4907.Contract.SetUser(&_Erc4907.TransactOpts, tokenId, user, expires)
}

// SetUser is a paid mutator transaction binding the contract method 0xe030565e.
//
// Solidity: function setUser(uint256 tokenId, address user, uint64 expires) returns()
func (_Erc4907 *Erc4907TransactorSession) SetUser(tokenId *big.Int, user common.Address, expires uint64) (*types.Transaction, error) {
	return _Erc4907.Contract.SetUser(&_Erc4907.TransactOpts, tokenId, user, expires)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Erc4907 *Erc4907Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Erc4907.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Erc4907 *Erc4907Session) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Erc4907.Contract.TransferFrom(&_Erc4907.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Erc4907 *Erc4907TransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Erc4907.Contract.TransferFrom(&_Erc4907.TransactOpts, from, to, tokenId)
}

// Erc4907ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Erc4907 contract.
type Erc4907ApprovalIterator struct {
	Event *Erc4907Approval // Event containing the contract specifics and raw log

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
func (it *Erc4907ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc4907Approval)
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
		it.Event = new(Erc4907Approval)
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
func (it *Erc4907ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc4907ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc4907Approval represents a Approval event raised by the Erc4907 contract.
type Erc4907Approval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Erc4907 *Erc4907Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*Erc4907ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Erc4907.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &Erc4907ApprovalIterator{contract: _Erc4907.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Erc4907 *Erc4907Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *Erc4907Approval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Erc4907.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc4907Approval)
				if err := _Erc4907.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Erc4907 *Erc4907Filterer) ParseApproval(log types.Log) (*Erc4907Approval, error) {
	event := new(Erc4907Approval)
	if err := _Erc4907.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc4907ApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Erc4907 contract.
type Erc4907ApprovalForAllIterator struct {
	Event *Erc4907ApprovalForAll // Event containing the contract specifics and raw log

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
func (it *Erc4907ApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc4907ApprovalForAll)
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
		it.Event = new(Erc4907ApprovalForAll)
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
func (it *Erc4907ApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc4907ApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc4907ApprovalForAll represents a ApprovalForAll event raised by the Erc4907 contract.
type Erc4907ApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Erc4907 *Erc4907Filterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*Erc4907ApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Erc4907.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &Erc4907ApprovalForAllIterator{contract: _Erc4907.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Erc4907 *Erc4907Filterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *Erc4907ApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Erc4907.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc4907ApprovalForAll)
				if err := _Erc4907.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Erc4907 *Erc4907Filterer) ParseApprovalForAll(log types.Log) (*Erc4907ApprovalForAll, error) {
	event := new(Erc4907ApprovalForAll)
	if err := _Erc4907.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc4907TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Erc4907 contract.
type Erc4907TransferIterator struct {
	Event *Erc4907Transfer // Event containing the contract specifics and raw log

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
func (it *Erc4907TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc4907Transfer)
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
		it.Event = new(Erc4907Transfer)
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
func (it *Erc4907TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc4907TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc4907Transfer represents a Transfer event raised by the Erc4907 contract.
type Erc4907Transfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Erc4907 *Erc4907Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*Erc4907TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Erc4907.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &Erc4907TransferIterator{contract: _Erc4907.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Erc4907 *Erc4907Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *Erc4907Transfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Erc4907.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc4907Transfer)
				if err := _Erc4907.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Erc4907 *Erc4907Filterer) ParseTransfer(log types.Log) (*Erc4907Transfer, error) {
	event := new(Erc4907Transfer)
	if err := _Erc4907.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc4907UpdateUserIterator is returned from FilterUpdateUser and is used to iterate over the raw logs and unpacked data for UpdateUser events raised by the Erc4907 contract.
type Erc4907UpdateUserIterator struct {
	Event *Erc4907UpdateUser // Event containing the contract specifics and raw log

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
func (it *Erc4907UpdateUserIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc4907UpdateUser)
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
		it.Event = new(Erc4907UpdateUser)
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
func (it *Erc4907UpdateUserIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc4907UpdateUserIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc4907UpdateUser represents a UpdateUser event raised by the Erc4907 contract.
type Erc4907UpdateUser struct {
	TokenId *big.Int
	User    common.Address
	Expires uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUpdateUser is a free log retrieval operation binding the contract event 0x4e06b4e7000e659094299b3533b47b6aa8ad048e95e872d23d1f4ee55af89cfe.
//
// Solidity: event UpdateUser(uint256 indexed tokenId, address indexed user, uint64 expires)
func (_Erc4907 *Erc4907Filterer) FilterUpdateUser(opts *bind.FilterOpts, tokenId []*big.Int, user []common.Address) (*Erc4907UpdateUserIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Erc4907.contract.FilterLogs(opts, "UpdateUser", tokenIdRule, userRule)
	if err != nil {
		return nil, err
	}
	return &Erc4907UpdateUserIterator{contract: _Erc4907.contract, event: "UpdateUser", logs: logs, sub: sub}, nil
}

// WatchUpdateUser is a free log subscription operation binding the contract event 0x4e06b4e7000e659094299b3533b47b6aa8ad048e95e872d23d1f4ee55af89cfe.
//
// Solidity: event UpdateUser(uint256 indexed tokenId, address indexed user, uint64 expires)
func (_Erc4907 *Erc4907Filterer) WatchUpdateUser(opts *bind.WatchOpts, sink chan<- *Erc4907UpdateUser, tokenId []*big.Int, user []common.Address) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Erc4907.contract.WatchLogs(opts, "UpdateUser", tokenIdRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc4907UpdateUser)
				if err := _Erc4907.contract.UnpackLog(event, "UpdateUser", log); err != nil {
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

// ParseUpdateUser is a log parse operation binding the contract event 0x4e06b4e7000e659094299b3533b47b6aa8ad048e95e872d23d1f4ee55af89cfe.
//
// Solidity: event UpdateUser(uint256 indexed tokenId, address indexed user, uint64 expires)
func (_Erc4907 *Erc4907Filterer) ParseUpdateUser(log types.Log) (*Erc4907UpdateUser, error) {
	event := new(Erc4907UpdateUser)
	if err := _Erc4907.contract.UnpackLog(event, "UpdateUser", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
