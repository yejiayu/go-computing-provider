// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package account

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

// AccountMetaData contains all meta data concerning the Account contract.
var AccountMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_nodeId\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"_multiAddresses\",\"type\":\"string[]\"},{\"internalType\":\"uint8\",\"name\":\"_ubiFlag\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"_beneficiaryAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quota\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"}],\"name\":\"BeneficiaryChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"newMultiaddrs\",\"type\":\"string[]\"}],\"name\":\"MultiaddrsChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"ubiFlag\",\"type\":\"uint8\"}],\"name\":\"UBIFlagChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"submitter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"taskId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"taskType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"zkType\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"proof\",\"type\":\"string\"}],\"name\":\"UBIProofSubmitted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"beneficiary\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"beneficiaryAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"quota\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newBeneficiary\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"newQuota\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newExpiration\",\"type\":\"uint256\"}],\"name\":\"changeBeneficiary\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"newMultiaddrs\",\"type\":\"string[]\"}],\"name\":\"changeMultiaddrs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"changeOwnerAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"newUbiFlag\",\"type\":\"uint8\"}],\"name\":\"changeUbiFlag\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAccount\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"},{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"multiAddresses\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nodeId\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_taskId\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"_taskType\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"_zkType\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_proof\",\"type\":\"string\"}],\"name\":\"submitUBIProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"tasks\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"taskId\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"taskType\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"zkType\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"proof\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isSubmitted\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ubiFlag\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b50604051620026283803806200262883398181016040528101906200003791906200057a565b846000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555083600190816200008891906200088b565b508260029080519060200190620000a19291906200015b565b5081600360006101000a81548160ff021916908360ff16021790555060405180606001604052808273ffffffffffffffffffffffffffffffffffffffff168152602001600081526020016000815250600460008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001015560408201518160020155905050505050505062000972565b828054828255906000526020600020908101928215620001a8579160200282015b82811115620001a75782518290816200019691906200088b565b50916020019190600101906200017c565b5b509050620001b79190620001bb565b5090565b5b80821115620001df5760008181620001d59190620001e3565b50600101620001bc565b5090565b508054620001f1906200067a565b6000825580601f1062000205575062000226565b601f01602090049060005260206000209081019062000225919062000229565b5b50565b5b80821115620002445760008160009055506001016200022a565b5090565b6000604051905090565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600062000289826200025c565b9050919050565b6200029b816200027c565b8114620002a757600080fd5b50565b600081519050620002bb8162000290565b92915050565b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6200031682620002cb565b810181811067ffffffffffffffff82111715620003385762000337620002dc565b5b80604052505050565b60006200034d62000248565b90506200035b82826200030b565b919050565b600067ffffffffffffffff8211156200037e576200037d620002dc565b5b6200038982620002cb565b9050602081019050919050565b60005b83811015620003b657808201518184015260208101905062000399565b60008484015250505050565b6000620003d9620003d38462000360565b62000341565b905082815260208101848484011115620003f857620003f7620002c6565b5b6200040584828562000396565b509392505050565b600082601f830112620004255762000424620002c1565b5b815162000437848260208601620003c2565b91505092915050565b600067ffffffffffffffff8211156200045e576200045d620002dc565b5b602082029050602081019050919050565b600080fd5b60006200048b620004858462000440565b62000341565b90508083825260208201905060208402830185811115620004b157620004b06200046f565b5b835b81811015620004ff57805167ffffffffffffffff811115620004da57620004d9620002c1565b5b808601620004e989826200040d565b85526020850194505050602081019050620004b3565b5050509392505050565b600082601f830112620005215762000520620002c1565b5b81516200053384826020860162000474565b91505092915050565b600060ff82169050919050565b62000554816200053c565b81146200056057600080fd5b50565b600081519050620005748162000549565b92915050565b600080600080600060a0868803121562000599576200059862000252565b5b6000620005a988828901620002aa565b955050602086015167ffffffffffffffff811115620005cd57620005cc62000257565b5b620005db888289016200040d565b945050604086015167ffffffffffffffff811115620005ff57620005fe62000257565b5b6200060d8882890162000509565b9350506060620006208882890162000563565b92505060806200063388828901620002aa565b9150509295509295909350565b600081519050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806200069357607f821691505b602082108103620006a957620006a86200064b565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302620007137fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82620006d4565b6200071f8683620006d4565b95508019841693508086168417925050509392505050565b6000819050919050565b6000819050919050565b60006200076c62000766620007608462000737565b62000741565b62000737565b9050919050565b6000819050919050565b62000788836200074b565b620007a0620007978262000773565b848454620006e1565b825550505050565b600090565b620007b7620007a8565b620007c48184846200077d565b505050565b5b81811015620007ec57620007e0600082620007ad565b600181019050620007ca565b5050565b601f8211156200083b576200080581620006af565b6200081084620006c4565b8101602085101562000820578190505b620008386200082f85620006c4565b830182620007c9565b50505b505050565b600082821c905092915050565b6000620008606000198460080262000840565b1980831691505092915050565b60006200087b83836200084d565b9150826002028217905092915050565b620008968262000640565b67ffffffffffffffff811115620008b257620008b1620002dc565b5b620008be82546200067a565b620008cb828285620007f0565b600060209050601f831160018114620009035760008415620008ee578287015190505b620008fa85826200086d565b8655506200096a565b601f1984166200091386620006af565b60005b828110156200093d5784890151825560018201915060208501945060208101905062000916565b868310156200095d578489015162000959601f8916826200084d565b8355505b6001600288020188555050505b505050505050565b611ca680620009826000396000f3fe608060405234801561001057600080fd5b50600436106100b45760003560e01c806385eac05f1161007157806385eac05f146101935780638da5cb5b146101af57806394f21938146101cd578063db613e81146101e9578063e8deed971461020d578063ee2c5e891461022b576100b4565b8063139d7fed146100b957806338af3eed146100d75780633d1333a5146100f757806358709cf2146101275780636553b6ef1461015b57806377ecf4df14610177575b600080fd5b6100c1610247565b6040516100ce9190610fb4565b60405180910390f35b6100df6102d5565b6040516100ee93929190611030565b60405180910390f35b610111600480360381019061010c91906110a7565b61030d565b60405161011e9190610fb4565b60405180910390f35b610141600480360381019061013c9190611209565b6103b9565b604051610152959493929190611289565b60405180910390f35b6101756004803603810190610170919061131d565b6105b7565b005b610191600480360381019061018c919061134a565b61069a565b005b6101ad60048036038101906101a89190611431565b6108c2565b005b6101b76109ed565b6040516101c4919061145e565b60405180910390f35b6101e760048036038101906101e2919061155f565b610a11565b005b6101f1610af0565b60405161020497969594939291906116b4565b60405180910390f35b610215610cd8565b6040516102229190611731565b60405180910390f35b6102456004803603810190610240919061174c565b610ceb565b005b60018054610254906117ce565b80601f0160208091040260200160405190810160405280929190818152602001828054610280906117ce565b80156102cd5780601f106102a2576101008083540402835291602001916102cd565b820191906000526020600020905b8154815290600101906020018083116102b057829003601f168201915b505050505081565b60048060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010154908060020154905083565b6002818154811061031d57600080fd5b906000526020600020016000915090508054610338906117ce565b80601f0160208091040260200160405190810160405280929190818152602001828054610364906117ce565b80156103b15780601f10610386576101008083540402835291602001916103b1565b820191906000526020600020905b81548152906001019060200180831161039457829003601f168201915b505050505081565b6007818051602081018201805184825260208301602085012081835280955050505050506000915090508060000180546103f2906117ce565b80601f016020809104026020016040519081016040528092919081815260200182805461041e906117ce565b801561046b5780601f106104405761010080835404028352916020019161046b565b820191906000526020600020905b81548152906001019060200180831161044e57829003601f168201915b5050505050908060010160009054906101000a900460ff1690806002018054610493906117ce565b80601f01602080910402602001604051908101604052809291908181526020018280546104bf906117ce565b801561050c5780601f106104e15761010080835404028352916020019161050c565b820191906000526020600020905b8154815290600101906020018083116104ef57829003601f168201915b505050505090806003018054610521906117ce565b80601f016020809104026020016040519081016040528092919081815260200182805461054d906117ce565b801561059a5780601f1061056f5761010080835404028352916020019161059a565b820191906000526020600020905b81548152906001019060200180831161057d57829003601f168201915b5050505050908060040160009054906101000a900460ff16905085565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610645576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161063c90611871565b60405180910390fd5b80600360006101000a81548160ff021916908360ff1602179055507f2504aabda145c6cbc7a70eee13f3abeb510de8e65904c53b824d58ff9fc5b84a8160405161068f9190611731565b60405180910390a150565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610728576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161071f90611871565b60405180910390fd5b60078460405161073891906118cd565b908152602001604051809103902060040160009054906101000a900460ff1615610797576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161078e90611956565b60405180910390fd5b6040518060a001604052808581526020018460ff168152602001838152602001828152602001600115158152506007856040516107d491906118cd565b908152602001604051809103902060008201518160000190816107f79190611b22565b5060208201518160010160006101000a81548160ff021916908360ff160217905550604082015181600201908161082e9190611b22565b5060608201518160030190816108449190611b22565b5060808201518160040160006101000a81548160ff0219169083151502179055509050503373ffffffffffffffffffffffffffffffffffffffff167f18a95158362c8c4a1cacd8dc12f81e1fd419696c958630a3c06e3466f4daad00858585856040516108b49493929190611bf4565b60405180910390a250505050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610950576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161094790611871565b60405180910390fd5b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a350565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610a9f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a9690611871565b60405180910390fd5b8060029080519060200190610ab5929190610e4a565b507f34f0cbd777d8d394caa00ae618bb83ee053819e2943bbd10c6c61019d4c5758281604051610ae59190611c4e565b60405180910390a150565b600060608060008060008060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1660016002600360009054906101000a900460ff16600460000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600460010154600460020154858054610b6e906117ce565b80601f0160208091040260200160405190810160405280929190818152602001828054610b9a906117ce565b8015610be75780601f10610bbc57610100808354040283529160200191610be7565b820191906000526020600020905b815481529060010190602001808311610bca57829003601f168201915b5050505050955084805480602002602001604051908101604052809291908181526020016000905b82821015610cbb578382906000526020600020018054610c2e906117ce565b80601f0160208091040260200160405190810160405280929190818152602001828054610c5a906117ce565b8015610ca75780601f10610c7c57610100808354040283529160200191610ca7565b820191906000526020600020905b815481529060010190602001808311610c8a57829003601f168201915b505050505081526020019060010190610c0f565b505050509450965096509650965096509650965090919293949596565b600360009054906101000a900460ff1681565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610d79576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d7090611871565b60405180910390fd5b60405180606001604052808473ffffffffffffffffffffffffffffffffffffffff16815260200183815260200182815250600460008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010155604082015181600201559050507f21c5f4c4149658f3fb981d0f3b26cb82ad95e6aab98e5a3dc73f9a2d6ab6b311838383604051610e3d93929190611030565b60405180910390a1505050565b828054828255906000526020600020908101928215610e92579160200282015b82811115610e91578251829081610e819190611b22565b5091602001919060010190610e6a565b5b509050610e9f9190610ea3565b5090565b5b80821115610ec35760008181610eba9190610ec7565b50600101610ea4565b5090565b508054610ed3906117ce565b6000825580601f10610ee55750610f04565b601f016020900490600052602060002090810190610f039190610f07565b5b50565b5b80821115610f20576000816000905550600101610f08565b5090565b600081519050919050565b600082825260208201905092915050565b60005b83811015610f5e578082015181840152602081019050610f43565b60008484015250505050565b6000601f19601f8301169050919050565b6000610f8682610f24565b610f908185610f2f565b9350610fa0818560208601610f40565b610fa981610f6a565b840191505092915050565b60006020820190508181036000830152610fce8184610f7b565b905092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061100182610fd6565b9050919050565b61101181610ff6565b82525050565b6000819050919050565b61102a81611017565b82525050565b60006060820190506110456000830186611008565b6110526020830185611021565b61105f6040830184611021565b949350505050565b6000604051905090565b600080fd5b600080fd5b61108481611017565b811461108f57600080fd5b50565b6000813590506110a18161107b565b92915050565b6000602082840312156110bd576110bc611071565b5b60006110cb84828501611092565b91505092915050565b600080fd5b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b61111682610f6a565b810181811067ffffffffffffffff82111715611135576111346110de565b5b80604052505050565b6000611148611067565b9050611154828261110d565b919050565b600067ffffffffffffffff821115611174576111736110de565b5b61117d82610f6a565b9050602081019050919050565b82818337600083830152505050565b60006111ac6111a784611159565b61113e565b9050828152602081018484840111156111c8576111c76110d9565b5b6111d384828561118a565b509392505050565b600082601f8301126111f0576111ef6110d4565b5b8135611200848260208601611199565b91505092915050565b60006020828403121561121f5761121e611071565b5b600082013567ffffffffffffffff81111561123d5761123c611076565b5b611249848285016111db565b91505092915050565b600060ff82169050919050565b61126881611252565b82525050565b60008115159050919050565b6112838161126e565b82525050565b600060a08201905081810360008301526112a38188610f7b565b90506112b2602083018761125f565b81810360408301526112c48186610f7b565b905081810360608301526112d88185610f7b565b90506112e7608083018461127a565b9695505050505050565b6112fa81611252565b811461130557600080fd5b50565b600081359050611317816112f1565b92915050565b60006020828403121561133357611332611071565b5b600061134184828501611308565b91505092915050565b6000806000806080858703121561136457611363611071565b5b600085013567ffffffffffffffff81111561138257611381611076565b5b61138e878288016111db565b945050602061139f87828801611308565b935050604085013567ffffffffffffffff8111156113c0576113bf611076565b5b6113cc878288016111db565b925050606085013567ffffffffffffffff8111156113ed576113ec611076565b5b6113f9878288016111db565b91505092959194509250565b61140e81610ff6565b811461141957600080fd5b50565b60008135905061142b81611405565b92915050565b60006020828403121561144757611446611071565b5b60006114558482850161141c565b91505092915050565b60006020820190506114736000830184611008565b92915050565b600067ffffffffffffffff821115611494576114936110de565b5b602082029050602081019050919050565b600080fd5b60006114bd6114b884611479565b61113e565b905080838252602082019050602084028301858111156114e0576114df6114a5565b5b835b8181101561152757803567ffffffffffffffff811115611505576115046110d4565b5b80860161151289826111db565b855260208501945050506020810190506114e2565b5050509392505050565b600082601f830112611546576115456110d4565b5b81356115568482602086016114aa565b91505092915050565b60006020828403121561157557611574611071565b5b600082013567ffffffffffffffff81111561159357611592611076565b5b61159f84828501611531565b91505092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b600082825260208201905092915050565b60006115f082610f24565b6115fa81856115d4565b935061160a818560208601610f40565b61161381610f6a565b840191505092915050565b600061162a83836115e5565b905092915050565b6000602082019050919050565b600061164a826115a8565b61165481856115b3565b935083602082028501611666856115c4565b8060005b858110156116a25784840389528151611683858261161e565b945061168e83611632565b925060208a0199505060018101905061166a565b50829750879550505050505092915050565b600060e0820190506116c9600083018a611008565b81810360208301526116db8189610f7b565b905081810360408301526116ef818861163f565b90506116fe606083018761125f565b61170b6080830186611008565b61171860a0830185611021565b61172560c0830184611021565b98975050505050505050565b6000602082019050611746600083018461125f565b92915050565b60008060006060848603121561176557611764611071565b5b60006117738682870161141c565b935050602061178486828701611092565b925050604061179586828701611092565b9150509250925092565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806117e657607f821691505b6020821081036117f9576117f861179f565b5b50919050565b7f4f6e6c79206f776e65722063616e2063616c6c20746869732066756e6374696f60008201527f6e2e000000000000000000000000000000000000000000000000000000000000602082015250565b600061185b602283610f2f565b9150611866826117ff565b604082019050919050565b6000602082019050818103600083015261188a8161184e565b9050919050565b600081905092915050565b60006118a782610f24565b6118b18185611891565b93506118c1818560208601610f40565b80840191505092915050565b60006118d9828461189c565b915081905092915050565b7f50726f6f6620666f722074686973207461736b20697320616c7265616479207360008201527f75626d69747465642e0000000000000000000000000000000000000000000000602082015250565b6000611940602983610f2f565b915061194b826118e4565b604082019050919050565b6000602082019050818103600083015261196f81611933565b9050919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b6000600883026119d87fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8261199b565b6119e2868361199b565b95508019841693508086168417925050509392505050565b6000819050919050565b6000611a1f611a1a611a1584611017565b6119fa565b611017565b9050919050565b6000819050919050565b611a3983611a04565b611a4d611a4582611a26565b8484546119a8565b825550505050565b600090565b611a62611a55565b611a6d818484611a30565b505050565b5b81811015611a9157611a86600082611a5a565b600181019050611a73565b5050565b601f821115611ad657611aa781611976565b611ab08461198b565b81016020851015611abf578190505b611ad3611acb8561198b565b830182611a72565b50505b505050565b600082821c905092915050565b6000611af960001984600802611adb565b1980831691505092915050565b6000611b128383611ae8565b9150826002028217905092915050565b611b2b82610f24565b67ffffffffffffffff811115611b4457611b436110de565b5b611b4e82546117ce565b611b59828285611a95565b600060209050601f831160018114611b8c5760008415611b7a578287015190505b611b848582611b06565b865550611bec565b601f198416611b9a86611976565b60005b82811015611bc257848901518255600182019150602085019450602081019050611b9d565b86831015611bdf5784890151611bdb601f891682611ae8565b8355505b6001600288020188555050505b505050505050565b60006080820190508181036000830152611c0e8187610f7b565b9050611c1d602083018661125f565b8181036040830152611c2f8185610f7b565b90508181036060830152611c438184610f7b565b905095945050505050565b60006020820190508181036000830152611c68818461163f565b90509291505056fea26469706673582212207aee7ec0665439330d0fe945486aa26d5bb9d61ac549cc689af2754563757a2c64736f6c63430008140033",
}

// AccountABI is the input ABI used to generate the binding from.
// Deprecated: Use AccountMetaData.ABI instead.
var AccountABI = AccountMetaData.ABI

// AccountBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AccountMetaData.Bin instead.
var AccountBin = AccountMetaData.Bin

// DeployAccount deploys a new Ethereum contract, binding an instance of Account to it.
func DeployAccount(auth *bind.TransactOpts, backend bind.ContractBackend, _owner common.Address, _nodeId string, _multiAddresses []string, _ubiFlag uint8, _beneficiaryAddress common.Address) (common.Address, *types.Transaction, *Account, error) {
	parsed, err := AccountMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AccountBin), backend, _owner, _nodeId, _multiAddresses, _ubiFlag, _beneficiaryAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Account{AccountCaller: AccountCaller{contract: contract}, AccountTransactor: AccountTransactor{contract: contract}, AccountFilterer: AccountFilterer{contract: contract}}, nil
}

// Account is an auto generated Go binding around an Ethereum contract.
type Account struct {
	AccountCaller     // Read-only binding to the contract
	AccountTransactor // Write-only binding to the contract
	AccountFilterer   // Log filterer for contract events
}

// AccountCaller is an auto generated read-only Go binding around an Ethereum contract.
type AccountCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AccountTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AccountFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AccountSession struct {
	Contract     *Account          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AccountCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AccountCallerSession struct {
	Contract *AccountCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// AccountTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AccountTransactorSession struct {
	Contract     *AccountTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// AccountRaw is an auto generated low-level Go binding around an Ethereum contract.
type AccountRaw struct {
	Contract *Account // Generic contract binding to access the raw methods on
}

// AccountCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AccountCallerRaw struct {
	Contract *AccountCaller // Generic read-only contract binding to access the raw methods on
}

// AccountTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AccountTransactorRaw struct {
	Contract *AccountTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAccount creates a new instance of Account, bound to a specific deployed contract.
func NewAccount(address common.Address, backend bind.ContractBackend) (*Account, error) {
	contract, err := bindAccount(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Account{AccountCaller: AccountCaller{contract: contract}, AccountTransactor: AccountTransactor{contract: contract}, AccountFilterer: AccountFilterer{contract: contract}}, nil
}

// NewAccountCaller creates a new read-only instance of Account, bound to a specific deployed contract.
func NewAccountCaller(address common.Address, caller bind.ContractCaller) (*AccountCaller, error) {
	contract, err := bindAccount(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AccountCaller{contract: contract}, nil
}

// NewAccountTransactor creates a new write-only instance of Account, bound to a specific deployed contract.
func NewAccountTransactor(address common.Address, transactor bind.ContractTransactor) (*AccountTransactor, error) {
	contract, err := bindAccount(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AccountTransactor{contract: contract}, nil
}

// NewAccountFilterer creates a new log filterer instance of Account, bound to a specific deployed contract.
func NewAccountFilterer(address common.Address, filterer bind.ContractFilterer) (*AccountFilterer, error) {
	contract, err := bindAccount(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AccountFilterer{contract: contract}, nil
}

// bindAccount binds a generic wrapper to an already deployed contract.
func bindAccount(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AccountMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Account *AccountRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Account.Contract.AccountCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Account *AccountRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Account.Contract.AccountTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Account *AccountRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Account.Contract.AccountTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Account *AccountCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Account.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Account *AccountTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Account.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Account *AccountTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Account.Contract.contract.Transact(opts, method, params...)
}

// Beneficiary is a free data retrieval call binding the contract method 0x38af3eed.
//
// Solidity: function beneficiary() view returns(address beneficiaryAddress, uint256 quota, uint256 expiration)
func (_Account *AccountCaller) Beneficiary(opts *bind.CallOpts) (struct {
	BeneficiaryAddress common.Address
	Quota              *big.Int
	Expiration         *big.Int
}, error) {
	var out []interface{}
	err := _Account.contract.Call(opts, &out, "beneficiary")

	outstruct := new(struct {
		BeneficiaryAddress common.Address
		Quota              *big.Int
		Expiration         *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BeneficiaryAddress = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Quota = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Expiration = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Beneficiary is a free data retrieval call binding the contract method 0x38af3eed.
//
// Solidity: function beneficiary() view returns(address beneficiaryAddress, uint256 quota, uint256 expiration)
func (_Account *AccountSession) Beneficiary() (struct {
	BeneficiaryAddress common.Address
	Quota              *big.Int
	Expiration         *big.Int
}, error) {
	return _Account.Contract.Beneficiary(&_Account.CallOpts)
}

// Beneficiary is a free data retrieval call binding the contract method 0x38af3eed.
//
// Solidity: function beneficiary() view returns(address beneficiaryAddress, uint256 quota, uint256 expiration)
func (_Account *AccountCallerSession) Beneficiary() (struct {
	BeneficiaryAddress common.Address
	Quota              *big.Int
	Expiration         *big.Int
}, error) {
	return _Account.Contract.Beneficiary(&_Account.CallOpts)
}

// GetAccount is a free data retrieval call binding the contract method 0xdb613e81.
//
// Solidity: function getAccount() view returns(address, string, string[], uint8, address, uint256, uint256)
func (_Account *AccountCaller) GetAccount(opts *bind.CallOpts) (common.Address, string, []string, uint8, common.Address, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _Account.contract.Call(opts, &out, "getAccount")

	if err != nil {
		return *new(common.Address), *new(string), *new([]string), *new(uint8), *new(common.Address), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(string)).(*string)
	out2 := *abi.ConvertType(out[2], new([]string)).(*[]string)
	out3 := *abi.ConvertType(out[3], new(uint8)).(*uint8)
	out4 := *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	out5 := *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	out6 := *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, out4, out5, out6, err

}

// GetAccount is a free data retrieval call binding the contract method 0xdb613e81.
//
// Solidity: function getAccount() view returns(address, string, string[], uint8, address, uint256, uint256)
func (_Account *AccountSession) GetAccount() (common.Address, string, []string, uint8, common.Address, *big.Int, *big.Int, error) {
	return _Account.Contract.GetAccount(&_Account.CallOpts)
}

// GetAccount is a free data retrieval call binding the contract method 0xdb613e81.
//
// Solidity: function getAccount() view returns(address, string, string[], uint8, address, uint256, uint256)
func (_Account *AccountCallerSession) GetAccount() (common.Address, string, []string, uint8, common.Address, *big.Int, *big.Int, error) {
	return _Account.Contract.GetAccount(&_Account.CallOpts)
}

// MultiAddresses is a free data retrieval call binding the contract method 0x3d1333a5.
//
// Solidity: function multiAddresses(uint256 ) view returns(string)
func (_Account *AccountCaller) MultiAddresses(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _Account.contract.Call(opts, &out, "multiAddresses", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// MultiAddresses is a free data retrieval call binding the contract method 0x3d1333a5.
//
// Solidity: function multiAddresses(uint256 ) view returns(string)
func (_Account *AccountSession) MultiAddresses(arg0 *big.Int) (string, error) {
	return _Account.Contract.MultiAddresses(&_Account.CallOpts, arg0)
}

// MultiAddresses is a free data retrieval call binding the contract method 0x3d1333a5.
//
// Solidity: function multiAddresses(uint256 ) view returns(string)
func (_Account *AccountCallerSession) MultiAddresses(arg0 *big.Int) (string, error) {
	return _Account.Contract.MultiAddresses(&_Account.CallOpts, arg0)
}

// NodeId is a free data retrieval call binding the contract method 0x139d7fed.
//
// Solidity: function nodeId() view returns(string)
func (_Account *AccountCaller) NodeId(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Account.contract.Call(opts, &out, "nodeId")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// NodeId is a free data retrieval call binding the contract method 0x139d7fed.
//
// Solidity: function nodeId() view returns(string)
func (_Account *AccountSession) NodeId() (string, error) {
	return _Account.Contract.NodeId(&_Account.CallOpts)
}

// NodeId is a free data retrieval call binding the contract method 0x139d7fed.
//
// Solidity: function nodeId() view returns(string)
func (_Account *AccountCallerSession) NodeId() (string, error) {
	return _Account.Contract.NodeId(&_Account.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Account *AccountCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Account.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Account *AccountSession) Owner() (common.Address, error) {
	return _Account.Contract.Owner(&_Account.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Account *AccountCallerSession) Owner() (common.Address, error) {
	return _Account.Contract.Owner(&_Account.CallOpts)
}

// Tasks is a free data retrieval call binding the contract method 0x58709cf2.
//
// Solidity: function tasks(string ) view returns(string taskId, uint8 taskType, string zkType, string proof, bool isSubmitted)
func (_Account *AccountCaller) Tasks(opts *bind.CallOpts, arg0 string) (struct {
	TaskId      string
	TaskType    uint8
	ZkType      string
	Proof       string
	IsSubmitted bool
}, error) {
	var out []interface{}
	err := _Account.contract.Call(opts, &out, "tasks", arg0)

	outstruct := new(struct {
		TaskId      string
		TaskType    uint8
		ZkType      string
		Proof       string
		IsSubmitted bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TaskId = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.TaskType = *abi.ConvertType(out[1], new(uint8)).(*uint8)
	outstruct.ZkType = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Proof = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.IsSubmitted = *abi.ConvertType(out[4], new(bool)).(*bool)

	return *outstruct, err

}

// Tasks is a free data retrieval call binding the contract method 0x58709cf2.
//
// Solidity: function tasks(string ) view returns(string taskId, uint8 taskType, string zkType, string proof, bool isSubmitted)
func (_Account *AccountSession) Tasks(arg0 string) (struct {
	TaskId      string
	TaskType    uint8
	ZkType      string
	Proof       string
	IsSubmitted bool
}, error) {
	return _Account.Contract.Tasks(&_Account.CallOpts, arg0)
}

// Tasks is a free data retrieval call binding the contract method 0x58709cf2.
//
// Solidity: function tasks(string ) view returns(string taskId, uint8 taskType, string zkType, string proof, bool isSubmitted)
func (_Account *AccountCallerSession) Tasks(arg0 string) (struct {
	TaskId      string
	TaskType    uint8
	ZkType      string
	Proof       string
	IsSubmitted bool
}, error) {
	return _Account.Contract.Tasks(&_Account.CallOpts, arg0)
}

// UbiFlag is a free data retrieval call binding the contract method 0xe8deed97.
//
// Solidity: function ubiFlag() view returns(uint8)
func (_Account *AccountCaller) UbiFlag(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Account.contract.Call(opts, &out, "ubiFlag")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// UbiFlag is a free data retrieval call binding the contract method 0xe8deed97.
//
// Solidity: function ubiFlag() view returns(uint8)
func (_Account *AccountSession) UbiFlag() (uint8, error) {
	return _Account.Contract.UbiFlag(&_Account.CallOpts)
}

// UbiFlag is a free data retrieval call binding the contract method 0xe8deed97.
//
// Solidity: function ubiFlag() view returns(uint8)
func (_Account *AccountCallerSession) UbiFlag() (uint8, error) {
	return _Account.Contract.UbiFlag(&_Account.CallOpts)
}

// ChangeBeneficiary is a paid mutator transaction binding the contract method 0xee2c5e89.
//
// Solidity: function changeBeneficiary(address newBeneficiary, uint256 newQuota, uint256 newExpiration) returns()
func (_Account *AccountTransactor) ChangeBeneficiary(opts *bind.TransactOpts, newBeneficiary common.Address, newQuota *big.Int, newExpiration *big.Int) (*types.Transaction, error) {
	return _Account.contract.Transact(opts, "changeBeneficiary", newBeneficiary, newQuota, newExpiration)
}

// ChangeBeneficiary is a paid mutator transaction binding the contract method 0xee2c5e89.
//
// Solidity: function changeBeneficiary(address newBeneficiary, uint256 newQuota, uint256 newExpiration) returns()
func (_Account *AccountSession) ChangeBeneficiary(newBeneficiary common.Address, newQuota *big.Int, newExpiration *big.Int) (*types.Transaction, error) {
	return _Account.Contract.ChangeBeneficiary(&_Account.TransactOpts, newBeneficiary, newQuota, newExpiration)
}

// ChangeBeneficiary is a paid mutator transaction binding the contract method 0xee2c5e89.
//
// Solidity: function changeBeneficiary(address newBeneficiary, uint256 newQuota, uint256 newExpiration) returns()
func (_Account *AccountTransactorSession) ChangeBeneficiary(newBeneficiary common.Address, newQuota *big.Int, newExpiration *big.Int) (*types.Transaction, error) {
	return _Account.Contract.ChangeBeneficiary(&_Account.TransactOpts, newBeneficiary, newQuota, newExpiration)
}

// ChangeMultiaddrs is a paid mutator transaction binding the contract method 0x94f21938.
//
// Solidity: function changeMultiaddrs(string[] newMultiaddrs) returns()
func (_Account *AccountTransactor) ChangeMultiaddrs(opts *bind.TransactOpts, newMultiaddrs []string) (*types.Transaction, error) {
	return _Account.contract.Transact(opts, "changeMultiaddrs", newMultiaddrs)
}

// ChangeMultiaddrs is a paid mutator transaction binding the contract method 0x94f21938.
//
// Solidity: function changeMultiaddrs(string[] newMultiaddrs) returns()
func (_Account *AccountSession) ChangeMultiaddrs(newMultiaddrs []string) (*types.Transaction, error) {
	return _Account.Contract.ChangeMultiaddrs(&_Account.TransactOpts, newMultiaddrs)
}

// ChangeMultiaddrs is a paid mutator transaction binding the contract method 0x94f21938.
//
// Solidity: function changeMultiaddrs(string[] newMultiaddrs) returns()
func (_Account *AccountTransactorSession) ChangeMultiaddrs(newMultiaddrs []string) (*types.Transaction, error) {
	return _Account.Contract.ChangeMultiaddrs(&_Account.TransactOpts, newMultiaddrs)
}

// ChangeOwnerAddress is a paid mutator transaction binding the contract method 0x85eac05f.
//
// Solidity: function changeOwnerAddress(address newOwner) returns()
func (_Account *AccountTransactor) ChangeOwnerAddress(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Account.contract.Transact(opts, "changeOwnerAddress", newOwner)
}

// ChangeOwnerAddress is a paid mutator transaction binding the contract method 0x85eac05f.
//
// Solidity: function changeOwnerAddress(address newOwner) returns()
func (_Account *AccountSession) ChangeOwnerAddress(newOwner common.Address) (*types.Transaction, error) {
	return _Account.Contract.ChangeOwnerAddress(&_Account.TransactOpts, newOwner)
}

// ChangeOwnerAddress is a paid mutator transaction binding the contract method 0x85eac05f.
//
// Solidity: function changeOwnerAddress(address newOwner) returns()
func (_Account *AccountTransactorSession) ChangeOwnerAddress(newOwner common.Address) (*types.Transaction, error) {
	return _Account.Contract.ChangeOwnerAddress(&_Account.TransactOpts, newOwner)
}

// ChangeUbiFlag is a paid mutator transaction binding the contract method 0x6553b6ef.
//
// Solidity: function changeUbiFlag(uint8 newUbiFlag) returns()
func (_Account *AccountTransactor) ChangeUbiFlag(opts *bind.TransactOpts, newUbiFlag uint8) (*types.Transaction, error) {
	return _Account.contract.Transact(opts, "changeUbiFlag", newUbiFlag)
}

// ChangeUbiFlag is a paid mutator transaction binding the contract method 0x6553b6ef.
//
// Solidity: function changeUbiFlag(uint8 newUbiFlag) returns()
func (_Account *AccountSession) ChangeUbiFlag(newUbiFlag uint8) (*types.Transaction, error) {
	return _Account.Contract.ChangeUbiFlag(&_Account.TransactOpts, newUbiFlag)
}

// ChangeUbiFlag is a paid mutator transaction binding the contract method 0x6553b6ef.
//
// Solidity: function changeUbiFlag(uint8 newUbiFlag) returns()
func (_Account *AccountTransactorSession) ChangeUbiFlag(newUbiFlag uint8) (*types.Transaction, error) {
	return _Account.Contract.ChangeUbiFlag(&_Account.TransactOpts, newUbiFlag)
}

// SubmitUBIProof is a paid mutator transaction binding the contract method 0x77ecf4df.
//
// Solidity: function submitUBIProof(string _taskId, uint8 _taskType, string _zkType, string _proof) returns()
func (_Account *AccountTransactor) SubmitUBIProof(opts *bind.TransactOpts, _taskId string, _taskType uint8, _zkType string, _proof string) (*types.Transaction, error) {
	return _Account.contract.Transact(opts, "submitUBIProof", _taskId, _taskType, _zkType, _proof)
}

// SubmitUBIProof is a paid mutator transaction binding the contract method 0x77ecf4df.
//
// Solidity: function submitUBIProof(string _taskId, uint8 _taskType, string _zkType, string _proof) returns()
func (_Account *AccountSession) SubmitUBIProof(_taskId string, _taskType uint8, _zkType string, _proof string) (*types.Transaction, error) {
	return _Account.Contract.SubmitUBIProof(&_Account.TransactOpts, _taskId, _taskType, _zkType, _proof)
}

// SubmitUBIProof is a paid mutator transaction binding the contract method 0x77ecf4df.
//
// Solidity: function submitUBIProof(string _taskId, uint8 _taskType, string _zkType, string _proof) returns()
func (_Account *AccountTransactorSession) SubmitUBIProof(_taskId string, _taskType uint8, _zkType string, _proof string) (*types.Transaction, error) {
	return _Account.Contract.SubmitUBIProof(&_Account.TransactOpts, _taskId, _taskType, _zkType, _proof)
}

// AccountBeneficiaryChangedIterator is returned from FilterBeneficiaryChanged and is used to iterate over the raw logs and unpacked data for BeneficiaryChanged events raised by the Account contract.
type AccountBeneficiaryChangedIterator struct {
	Event *AccountBeneficiaryChanged // Event containing the contract specifics and raw log

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
func (it *AccountBeneficiaryChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountBeneficiaryChanged)
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
		it.Event = new(AccountBeneficiaryChanged)
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
func (it *AccountBeneficiaryChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountBeneficiaryChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountBeneficiaryChanged represents a BeneficiaryChanged event raised by the Account contract.
type AccountBeneficiaryChanged struct {
	Beneficiary common.Address
	Quota       *big.Int
	Expiration  *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBeneficiaryChanged is a free log retrieval operation binding the contract event 0x21c5f4c4149658f3fb981d0f3b26cb82ad95e6aab98e5a3dc73f9a2d6ab6b311.
//
// Solidity: event BeneficiaryChanged(address beneficiary, uint256 quota, uint256 expiration)
func (_Account *AccountFilterer) FilterBeneficiaryChanged(opts *bind.FilterOpts) (*AccountBeneficiaryChangedIterator, error) {

	logs, sub, err := _Account.contract.FilterLogs(opts, "BeneficiaryChanged")
	if err != nil {
		return nil, err
	}
	return &AccountBeneficiaryChangedIterator{contract: _Account.contract, event: "BeneficiaryChanged", logs: logs, sub: sub}, nil
}

// WatchBeneficiaryChanged is a free log subscription operation binding the contract event 0x21c5f4c4149658f3fb981d0f3b26cb82ad95e6aab98e5a3dc73f9a2d6ab6b311.
//
// Solidity: event BeneficiaryChanged(address beneficiary, uint256 quota, uint256 expiration)
func (_Account *AccountFilterer) WatchBeneficiaryChanged(opts *bind.WatchOpts, sink chan<- *AccountBeneficiaryChanged) (event.Subscription, error) {

	logs, sub, err := _Account.contract.WatchLogs(opts, "BeneficiaryChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountBeneficiaryChanged)
				if err := _Account.contract.UnpackLog(event, "BeneficiaryChanged", log); err != nil {
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

// ParseBeneficiaryChanged is a log parse operation binding the contract event 0x21c5f4c4149658f3fb981d0f3b26cb82ad95e6aab98e5a3dc73f9a2d6ab6b311.
//
// Solidity: event BeneficiaryChanged(address beneficiary, uint256 quota, uint256 expiration)
func (_Account *AccountFilterer) ParseBeneficiaryChanged(log types.Log) (*AccountBeneficiaryChanged, error) {
	event := new(AccountBeneficiaryChanged)
	if err := _Account.contract.UnpackLog(event, "BeneficiaryChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountMultiaddrsChangedIterator is returned from FilterMultiaddrsChanged and is used to iterate over the raw logs and unpacked data for MultiaddrsChanged events raised by the Account contract.
type AccountMultiaddrsChangedIterator struct {
	Event *AccountMultiaddrsChanged // Event containing the contract specifics and raw log

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
func (it *AccountMultiaddrsChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountMultiaddrsChanged)
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
		it.Event = new(AccountMultiaddrsChanged)
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
func (it *AccountMultiaddrsChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountMultiaddrsChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountMultiaddrsChanged represents a MultiaddrsChanged event raised by the Account contract.
type AccountMultiaddrsChanged struct {
	NewMultiaddrs []string
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterMultiaddrsChanged is a free log retrieval operation binding the contract event 0x34f0cbd777d8d394caa00ae618bb83ee053819e2943bbd10c6c61019d4c57582.
//
// Solidity: event MultiaddrsChanged(string[] newMultiaddrs)
func (_Account *AccountFilterer) FilterMultiaddrsChanged(opts *bind.FilterOpts) (*AccountMultiaddrsChangedIterator, error) {

	logs, sub, err := _Account.contract.FilterLogs(opts, "MultiaddrsChanged")
	if err != nil {
		return nil, err
	}
	return &AccountMultiaddrsChangedIterator{contract: _Account.contract, event: "MultiaddrsChanged", logs: logs, sub: sub}, nil
}

// WatchMultiaddrsChanged is a free log subscription operation binding the contract event 0x34f0cbd777d8d394caa00ae618bb83ee053819e2943bbd10c6c61019d4c57582.
//
// Solidity: event MultiaddrsChanged(string[] newMultiaddrs)
func (_Account *AccountFilterer) WatchMultiaddrsChanged(opts *bind.WatchOpts, sink chan<- *AccountMultiaddrsChanged) (event.Subscription, error) {

	logs, sub, err := _Account.contract.WatchLogs(opts, "MultiaddrsChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountMultiaddrsChanged)
				if err := _Account.contract.UnpackLog(event, "MultiaddrsChanged", log); err != nil {
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

// ParseMultiaddrsChanged is a log parse operation binding the contract event 0x34f0cbd777d8d394caa00ae618bb83ee053819e2943bbd10c6c61019d4c57582.
//
// Solidity: event MultiaddrsChanged(string[] newMultiaddrs)
func (_Account *AccountFilterer) ParseMultiaddrsChanged(log types.Log) (*AccountMultiaddrsChanged, error) {
	event := new(AccountMultiaddrsChanged)
	if err := _Account.contract.UnpackLog(event, "MultiaddrsChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Account contract.
type AccountOwnershipTransferredIterator struct {
	Event *AccountOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AccountOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountOwnershipTransferred)
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
		it.Event = new(AccountOwnershipTransferred)
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
func (it *AccountOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountOwnershipTransferred represents a OwnershipTransferred event raised by the Account contract.
type AccountOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Account *AccountFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AccountOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Account.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AccountOwnershipTransferredIterator{contract: _Account.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Account *AccountFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AccountOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Account.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountOwnershipTransferred)
				if err := _Account.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Account *AccountFilterer) ParseOwnershipTransferred(log types.Log) (*AccountOwnershipTransferred, error) {
	event := new(AccountOwnershipTransferred)
	if err := _Account.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountUBIFlagChangedIterator is returned from FilterUBIFlagChanged and is used to iterate over the raw logs and unpacked data for UBIFlagChanged events raised by the Account contract.
type AccountUBIFlagChangedIterator struct {
	Event *AccountUBIFlagChanged // Event containing the contract specifics and raw log

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
func (it *AccountUBIFlagChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountUBIFlagChanged)
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
		it.Event = new(AccountUBIFlagChanged)
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
func (it *AccountUBIFlagChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountUBIFlagChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountUBIFlagChanged represents a UBIFlagChanged event raised by the Account contract.
type AccountUBIFlagChanged struct {
	UbiFlag uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUBIFlagChanged is a free log retrieval operation binding the contract event 0x2504aabda145c6cbc7a70eee13f3abeb510de8e65904c53b824d58ff9fc5b84a.
//
// Solidity: event UBIFlagChanged(uint8 ubiFlag)
func (_Account *AccountFilterer) FilterUBIFlagChanged(opts *bind.FilterOpts) (*AccountUBIFlagChangedIterator, error) {

	logs, sub, err := _Account.contract.FilterLogs(opts, "UBIFlagChanged")
	if err != nil {
		return nil, err
	}
	return &AccountUBIFlagChangedIterator{contract: _Account.contract, event: "UBIFlagChanged", logs: logs, sub: sub}, nil
}

// WatchUBIFlagChanged is a free log subscription operation binding the contract event 0x2504aabda145c6cbc7a70eee13f3abeb510de8e65904c53b824d58ff9fc5b84a.
//
// Solidity: event UBIFlagChanged(uint8 ubiFlag)
func (_Account *AccountFilterer) WatchUBIFlagChanged(opts *bind.WatchOpts, sink chan<- *AccountUBIFlagChanged) (event.Subscription, error) {

	logs, sub, err := _Account.contract.WatchLogs(opts, "UBIFlagChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountUBIFlagChanged)
				if err := _Account.contract.UnpackLog(event, "UBIFlagChanged", log); err != nil {
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

// ParseUBIFlagChanged is a log parse operation binding the contract event 0x2504aabda145c6cbc7a70eee13f3abeb510de8e65904c53b824d58ff9fc5b84a.
//
// Solidity: event UBIFlagChanged(uint8 ubiFlag)
func (_Account *AccountFilterer) ParseUBIFlagChanged(log types.Log) (*AccountUBIFlagChanged, error) {
	event := new(AccountUBIFlagChanged)
	if err := _Account.contract.UnpackLog(event, "UBIFlagChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountUBIProofSubmittedIterator is returned from FilterUBIProofSubmitted and is used to iterate over the raw logs and unpacked data for UBIProofSubmitted events raised by the Account contract.
type AccountUBIProofSubmittedIterator struct {
	Event *AccountUBIProofSubmitted // Event containing the contract specifics and raw log

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
func (it *AccountUBIProofSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountUBIProofSubmitted)
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
		it.Event = new(AccountUBIProofSubmitted)
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
func (it *AccountUBIProofSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountUBIProofSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountUBIProofSubmitted represents a UBIProofSubmitted event raised by the Account contract.
type AccountUBIProofSubmitted struct {
	Submitter common.Address
	TaskId    string
	TaskType  uint8
	ZkType    string
	Proof     string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUBIProofSubmitted is a free log retrieval operation binding the contract event 0x18a95158362c8c4a1cacd8dc12f81e1fd419696c958630a3c06e3466f4daad00.
//
// Solidity: event UBIProofSubmitted(address indexed submitter, string taskId, uint8 taskType, string zkType, string proof)
func (_Account *AccountFilterer) FilterUBIProofSubmitted(opts *bind.FilterOpts, submitter []common.Address) (*AccountUBIProofSubmittedIterator, error) {

	var submitterRule []interface{}
	for _, submitterItem := range submitter {
		submitterRule = append(submitterRule, submitterItem)
	}

	logs, sub, err := _Account.contract.FilterLogs(opts, "UBIProofSubmitted", submitterRule)
	if err != nil {
		return nil, err
	}
	return &AccountUBIProofSubmittedIterator{contract: _Account.contract, event: "UBIProofSubmitted", logs: logs, sub: sub}, nil
}

// WatchUBIProofSubmitted is a free log subscription operation binding the contract event 0x18a95158362c8c4a1cacd8dc12f81e1fd419696c958630a3c06e3466f4daad00.
//
// Solidity: event UBIProofSubmitted(address indexed submitter, string taskId, uint8 taskType, string zkType, string proof)
func (_Account *AccountFilterer) WatchUBIProofSubmitted(opts *bind.WatchOpts, sink chan<- *AccountUBIProofSubmitted, submitter []common.Address) (event.Subscription, error) {

	var submitterRule []interface{}
	for _, submitterItem := range submitter {
		submitterRule = append(submitterRule, submitterItem)
	}

	logs, sub, err := _Account.contract.WatchLogs(opts, "UBIProofSubmitted", submitterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountUBIProofSubmitted)
				if err := _Account.contract.UnpackLog(event, "UBIProofSubmitted", log); err != nil {
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

// ParseUBIProofSubmitted is a log parse operation binding the contract event 0x18a95158362c8c4a1cacd8dc12f81e1fd419696c958630a3c06e3466f4daad00.
//
// Solidity: event UBIProofSubmitted(address indexed submitter, string taskId, uint8 taskType, string zkType, string proof)
func (_Account *AccountFilterer) ParseUBIProofSubmitted(log types.Log) (*AccountUBIProofSubmitted, error) {
	event := new(AccountUBIProofSubmitted)
	if err := _Account.contract.UnpackLog(event, "UBIProofSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
