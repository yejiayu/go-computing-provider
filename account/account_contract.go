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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_nodeId\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"_multiAddresses\",\"type\":\"string[]\"},{\"internalType\":\"uint8\",\"name\":\"_ubiFlag\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"_beneficiaryAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"beneficiary\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"beneficiaryAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"quota\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newBeneficiary\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"newQuota\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newExpiration\",\"type\":\"uint256\"}],\"name\":\"changeBeneficiary\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"newMultiaddrs\",\"type\":\"string[]\"}],\"name\":\"changeMultiaddrs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"changeOwnerAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"newUbiFlag\",\"type\":\"uint8\"}],\"name\":\"changeUbiFlag\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"multiAddresses\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nodeId\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_taskId\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"_taskType\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"_zkType\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_proof\",\"type\":\"string\"}],\"name\":\"submitUBIProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"tasks\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"taskId\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"taskType\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"zkType\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"proof\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isSubmitted\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ubiFlag\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b50604051620021073803806200210783398181016040528101906200003791906200057a565b846000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555083600190816200008891906200088b565b508260029080519060200190620000a19291906200015b565b5081600360006101000a81548160ff021916908360ff16021790555060405180606001604052808273ffffffffffffffffffffffffffffffffffffffff168152602001600081526020016000815250600460008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001015560408201518160020155905050505050505062000972565b828054828255906000526020600020908101928215620001a8579160200282015b82811115620001a75782518290816200019691906200088b565b50916020019190600101906200017c565b5b509050620001b79190620001bb565b5090565b5b80821115620001df5760008181620001d59190620001e3565b50600101620001bc565b5090565b508054620001f1906200067a565b6000825580601f1062000205575062000226565b601f01602090049060005260206000209081019062000225919062000229565b5b50565b5b80821115620002445760008160009055506001016200022a565b5090565b6000604051905090565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600062000289826200025c565b9050919050565b6200029b816200027c565b8114620002a757600080fd5b50565b600081519050620002bb8162000290565b92915050565b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6200031682620002cb565b810181811067ffffffffffffffff82111715620003385762000337620002dc565b5b80604052505050565b60006200034d62000248565b90506200035b82826200030b565b919050565b600067ffffffffffffffff8211156200037e576200037d620002dc565b5b6200038982620002cb565b9050602081019050919050565b60005b83811015620003b657808201518184015260208101905062000399565b60008484015250505050565b6000620003d9620003d38462000360565b62000341565b905082815260208101848484011115620003f857620003f7620002c6565b5b6200040584828562000396565b509392505050565b600082601f830112620004255762000424620002c1565b5b815162000437848260208601620003c2565b91505092915050565b600067ffffffffffffffff8211156200045e576200045d620002dc565b5b602082029050602081019050919050565b600080fd5b60006200048b620004858462000440565b62000341565b90508083825260208201905060208402830185811115620004b157620004b06200046f565b5b835b81811015620004ff57805167ffffffffffffffff811115620004da57620004d9620002c1565b5b808601620004e989826200040d565b85526020850194505050602081019050620004b3565b5050509392505050565b600082601f830112620005215762000520620002c1565b5b81516200053384826020860162000474565b91505092915050565b600060ff82169050919050565b62000554816200053c565b81146200056057600080fd5b50565b600081519050620005748162000549565b92915050565b600080600080600060a0868803121562000599576200059862000252565b5b6000620005a988828901620002aa565b955050602086015167ffffffffffffffff811115620005cd57620005cc62000257565b5b620005db888289016200040d565b945050604086015167ffffffffffffffff811115620005ff57620005fe62000257565b5b6200060d8882890162000509565b9350506060620006208882890162000563565b92505060806200063388828901620002aa565b9150509295509295909350565b600081519050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806200069357607f821691505b602082108103620006a957620006a86200064b565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302620007137fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82620006d4565b6200071f8683620006d4565b95508019841693508086168417925050509392505050565b6000819050919050565b6000819050919050565b60006200076c62000766620007608462000737565b62000741565b62000737565b9050919050565b6000819050919050565b62000788836200074b565b620007a0620007978262000773565b848454620006e1565b825550505050565b600090565b620007b7620007a8565b620007c48184846200077d565b505050565b5b81811015620007ec57620007e0600082620007ad565b600181019050620007ca565b5050565b601f8211156200083b576200080581620006af565b6200081084620006c4565b8101602085101562000820578190505b620008386200082f85620006c4565b830182620007c9565b50505b505050565b600082821c905092915050565b6000620008606000198460080262000840565b1980831691505092915050565b60006200087b83836200084d565b9150826002028217905092915050565b620008968262000640565b67ffffffffffffffff811115620008b257620008b1620002dc565b5b620008be82546200067a565b620008cb828285620007f0565b600060209050601f831160018114620009035760008415620008ee578287015190505b620008fa85826200086d565b8655506200096a565b601f1984166200091386620006af565b60005b828110156200093d5784890151825560018201915060208501945060208101905062000916565b868310156200095d578489015162000959601f8916826200084d565b8355505b6001600288020188555050505b505050505050565b61178580620009826000396000f3fe608060405234801561001057600080fd5b50600436106100b45760003560e01c806385eac05f1161007157806385eac05f14610193578063893d20e8146101af5780638da5cb5b146101cd57806394f21938146101eb578063e8deed9714610207578063ee2c5e8914610225576100b4565b8063139d7fed146100b957806338af3eed146100d75780633d1333a5146100f757806358709cf2146101275780636553b6ef1461015b57806377ecf4df14610177575b600080fd5b6100c1610241565b6040516100ce9190610c98565b60405180910390f35b6100df6102cf565b6040516100ee93929190610d14565b60405180910390f35b610111600480360381019061010c9190610d8b565b610307565b60405161011e9190610c98565b60405180910390f35b610141600480360381019061013c9190610eed565b6103b3565b604051610152959493929190610f6d565b60405180910390f35b61017560048036038101906101709190611001565b6105b1565b005b610191600480360381019061018c919061102e565b61065d565b005b6101ad60048036038101906101a89190611115565b610831565b005b6101b7610902565b6040516101c49190611142565b60405180910390f35b6101d561092b565b6040516101e29190611142565b60405180910390f35b61020560048036038101906102009190611243565b61094f565b005b61020f6109f7565b60405161021c919061128c565b60405180910390f35b61023f600480360381019061023a91906112a7565b610a0a565b005b6001805461024e90611329565b80601f016020809104026020016040519081016040528092919081815260200182805461027a90611329565b80156102c75780601f1061029c576101008083540402835291602001916102c7565b820191906000526020600020905b8154815290600101906020018083116102aa57829003601f168201915b505050505081565b60048060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010154908060020154905083565b6002818154811061031757600080fd5b90600052602060002001600091509050805461033290611329565b80601f016020809104026020016040519081016040528092919081815260200182805461035e90611329565b80156103ab5780601f10610380576101008083540402835291602001916103ab565b820191906000526020600020905b81548152906001019060200180831161038e57829003601f168201915b505050505081565b6007818051602081018201805184825260208301602085012081835280955050505050506000915090508060000180546103ec90611329565b80601f016020809104026020016040519081016040528092919081815260200182805461041890611329565b80156104655780601f1061043a57610100808354040283529160200191610465565b820191906000526020600020905b81548152906001019060200180831161044857829003601f168201915b5050505050908060010160009054906101000a900460ff169080600201805461048d90611329565b80601f01602080910402602001604051908101604052809291908181526020018280546104b990611329565b80156105065780601f106104db57610100808354040283529160200191610506565b820191906000526020600020905b8154815290600101906020018083116104e957829003601f168201915b50505050509080600301805461051b90611329565b80601f016020809104026020016040519081016040528092919081815260200182805461054790611329565b80156105945780601f1061056957610100808354040283529160200191610594565b820191906000526020600020905b81548152906001019060200180831161057757829003601f168201915b5050505050908060040160009054906101000a900460ff16905085565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461063f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610636906113cc565b60405180910390fd5b80600360006101000a81548160ff021916908360ff16021790555050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146106eb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106e2906113cc565b60405180910390fd5b6007846040516106fb9190611428565b908152602001604051809103902060040160009054906101000a900460ff161561075a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610751906114b1565b60405180910390fd5b6040518060a001604052808581526020018460ff168152602001838152602001828152602001600115158152506007856040516107979190611428565b908152602001604051809103902060008201518160000190816107ba919061167d565b5060208201518160010160006101000a81548160ff021916908360ff16021790555060408201518160020190816107f1919061167d565b506060820151816003019081610807919061167d565b5060808201518160040160006101000a81548160ff02191690831515021790555090505050505050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146108bf576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108b6906113cc565b60405180910390fd5b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146109dd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109d4906113cc565b60405180910390fd5b80600290805190602001906109f3929190610b2e565b5050565b600360009054906101000a900460ff1681565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610a98576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a8f906113cc565b60405180910390fd5b60405180606001604052808473ffffffffffffffffffffffffffffffffffffffff16815260200183815260200182815250600460008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001015560408201518160020155905050505050565b828054828255906000526020600020908101928215610b76579160200282015b82811115610b75578251829081610b65919061167d565b5091602001919060010190610b4e565b5b509050610b839190610b87565b5090565b5b80821115610ba75760008181610b9e9190610bab565b50600101610b88565b5090565b508054610bb790611329565b6000825580601f10610bc95750610be8565b601f016020900490600052602060002090810190610be79190610beb565b5b50565b5b80821115610c04576000816000905550600101610bec565b5090565b600081519050919050565b600082825260208201905092915050565b60005b83811015610c42578082015181840152602081019050610c27565b60008484015250505050565b6000601f19601f8301169050919050565b6000610c6a82610c08565b610c748185610c13565b9350610c84818560208601610c24565b610c8d81610c4e565b840191505092915050565b60006020820190508181036000830152610cb28184610c5f565b905092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610ce582610cba565b9050919050565b610cf581610cda565b82525050565b6000819050919050565b610d0e81610cfb565b82525050565b6000606082019050610d296000830186610cec565b610d366020830185610d05565b610d436040830184610d05565b949350505050565b6000604051905090565b600080fd5b600080fd5b610d6881610cfb565b8114610d7357600080fd5b50565b600081359050610d8581610d5f565b92915050565b600060208284031215610da157610da0610d55565b5b6000610daf84828501610d76565b91505092915050565b600080fd5b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610dfa82610c4e565b810181811067ffffffffffffffff82111715610e1957610e18610dc2565b5b80604052505050565b6000610e2c610d4b565b9050610e388282610df1565b919050565b600067ffffffffffffffff821115610e5857610e57610dc2565b5b610e6182610c4e565b9050602081019050919050565b82818337600083830152505050565b6000610e90610e8b84610e3d565b610e22565b905082815260208101848484011115610eac57610eab610dbd565b5b610eb7848285610e6e565b509392505050565b600082601f830112610ed457610ed3610db8565b5b8135610ee4848260208601610e7d565b91505092915050565b600060208284031215610f0357610f02610d55565b5b600082013567ffffffffffffffff811115610f2157610f20610d5a565b5b610f2d84828501610ebf565b91505092915050565b600060ff82169050919050565b610f4c81610f36565b82525050565b60008115159050919050565b610f6781610f52565b82525050565b600060a0820190508181036000830152610f878188610c5f565b9050610f966020830187610f43565b8181036040830152610fa88186610c5f565b90508181036060830152610fbc8185610c5f565b9050610fcb6080830184610f5e565b9695505050505050565b610fde81610f36565b8114610fe957600080fd5b50565b600081359050610ffb81610fd5565b92915050565b60006020828403121561101757611016610d55565b5b600061102584828501610fec565b91505092915050565b6000806000806080858703121561104857611047610d55565b5b600085013567ffffffffffffffff81111561106657611065610d5a565b5b61107287828801610ebf565b945050602061108387828801610fec565b935050604085013567ffffffffffffffff8111156110a4576110a3610d5a565b5b6110b087828801610ebf565b925050606085013567ffffffffffffffff8111156110d1576110d0610d5a565b5b6110dd87828801610ebf565b91505092959194509250565b6110f281610cda565b81146110fd57600080fd5b50565b60008135905061110f816110e9565b92915050565b60006020828403121561112b5761112a610d55565b5b600061113984828501611100565b91505092915050565b60006020820190506111576000830184610cec565b92915050565b600067ffffffffffffffff82111561117857611177610dc2565b5b602082029050602081019050919050565b600080fd5b60006111a161119c8461115d565b610e22565b905080838252602082019050602084028301858111156111c4576111c3611189565b5b835b8181101561120b57803567ffffffffffffffff8111156111e9576111e8610db8565b5b8086016111f68982610ebf565b855260208501945050506020810190506111c6565b5050509392505050565b600082601f83011261122a57611229610db8565b5b813561123a84826020860161118e565b91505092915050565b60006020828403121561125957611258610d55565b5b600082013567ffffffffffffffff81111561127757611276610d5a565b5b61128384828501611215565b91505092915050565b60006020820190506112a16000830184610f43565b92915050565b6000806000606084860312156112c0576112bf610d55565b5b60006112ce86828701611100565b93505060206112df86828701610d76565b92505060406112f086828701610d76565b9150509250925092565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061134157607f821691505b602082108103611354576113536112fa565b5b50919050565b7f4f6e6c79206f776e65722063616e2063616c6c20746869732066756e6374696f60008201527f6e2e000000000000000000000000000000000000000000000000000000000000602082015250565b60006113b6602283610c13565b91506113c18261135a565b604082019050919050565b600060208201905081810360008301526113e5816113a9565b9050919050565b600081905092915050565b600061140282610c08565b61140c81856113ec565b935061141c818560208601610c24565b80840191505092915050565b600061143482846113f7565b915081905092915050565b7f50726f6f6620666f722074686973207461736b20697320616c7265616479207360008201527f75626d69747465642e0000000000000000000000000000000000000000000000602082015250565b600061149b602983610c13565b91506114a68261143f565b604082019050919050565b600060208201905081810360008301526114ca8161148e565b9050919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b6000600883026115337fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826114f6565b61153d86836114f6565b95508019841693508086168417925050509392505050565b6000819050919050565b600061157a61157561157084610cfb565b611555565b610cfb565b9050919050565b6000819050919050565b6115948361155f565b6115a86115a082611581565b848454611503565b825550505050565b600090565b6115bd6115b0565b6115c881848461158b565b505050565b5b818110156115ec576115e16000826115b5565b6001810190506115ce565b5050565b601f82111561163157611602816114d1565b61160b846114e6565b8101602085101561161a578190505b61162e611626856114e6565b8301826115cd565b50505b505050565b600082821c905092915050565b600061165460001984600802611636565b1980831691505092915050565b600061166d8383611643565b9150826002028217905092915050565b61168682610c08565b67ffffffffffffffff81111561169f5761169e610dc2565b5b6116a98254611329565b6116b48282856115f0565b600060209050601f8311600181146116e757600084156116d5578287015190505b6116df8582611661565b865550611747565b601f1984166116f5866114d1565b60005b8281101561171d578489015182556001820191506020850194506020810190506116f8565b8683101561173a5784890151611736601f891682611643565b8355505b6001600288020188555050505b50505050505056fea2646970667358221220b3df31cb9ce650ee01e2d096df7aadaee3450e599a8039ed428d86dff25b811864736f6c63430008140033",
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

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_Account *AccountCaller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Account.contract.Call(opts, &out, "getOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_Account *AccountSession) GetOwner() (common.Address, error) {
	return _Account.Contract.GetOwner(&_Account.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_Account *AccountCallerSession) GetOwner() (common.Address, error) {
	return _Account.Contract.GetOwner(&_Account.CallOpts)
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
