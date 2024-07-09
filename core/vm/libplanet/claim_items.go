package libplanet

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/sircoon4/bencodex-go/bencodextype"
)

type ClaimItems struct {
	Id [16]byte    `abi:"id"`
	Cd []ClaimData `abi:"cd"`
	M  string      `abi:"m"`
}

type ClaimData struct {
	AvatarAddress       common.Address       `abi:"avatarAddress"`
	FungibleAssetValues []FungibleAssetValue `abi:"fungibleAssetValues"`
}

func convertToClaimItemsEthAbi(actionValues *bencodextype.Dictionary) ([]byte, error) {
	var TupleClaimItems, _ = abi.NewType("tuple", "struct ClaimItems", []abi.ArgumentMarshaling{
		{Name: "id", Type: "bytes16"},
		{Name: "cd", Type: "tuple[]", Components: []abi.ArgumentMarshaling{
			{Name: "avatarAddress", Type: "address"},
			{Name: "fungibleAssetValues", Type: "tuple[]", Components: []abi.ArgumentMarshaling{
				{Name: "currency", Type: "tuple", Components: []abi.ArgumentMarshaling{
					{Name: "decimalPlaces", Type: "uint8"},
					{Name: "minters", Type: "address[]"},
					{Name: "ticker", Type: "string"},
					{Name: "totalSupplyTrackable", Type: "bool"},
					{Name: "maximumSupplyMajor", Type: "uint256"},
					{Name: "maximumSupplyMinor", Type: "uint256"},
				}},
				{Name: "rawValue", Type: "uint256"},
			}},
		}},
		{Name: "m", Type: "string"},
	})

	var arguments = abi.Arguments{
		abi.Argument{Name: "ClaimItems", Type: TupleClaimItems, Indexed: false},
	}

	idValue, _ := actionValues.Get("id").([]byte)

	claimDataList := []ClaimData{}
	claimDataValues := actionValues.Get("cd").([]any)
	for _, claimData := range claimDataValues {
		claimDataValue := claimData.([]any)
		avatarAddressValue := common.BytesToAddress(claimDataValue[0].([]byte))
		fungibleAssetValuesList := []FungibleAssetValue{}
		for _, fungibleAssetValue := range claimDataValue[1].([]any) {
			fungibleAssetValuesList = append(fungibleAssetValuesList, extractFungibleAssetValue(fungibleAssetValue.([]any)))
		}
		claimDataList = append(claimDataList, ClaimData{
			AvatarAddress:       avatarAddressValue,
			FungibleAssetValues: fungibleAssetValuesList,
		})
	}

	mValue := ""
	if actionValues.Contains("m") {
		if actionValues.Get("m") != nil {
			mValue = actionValues.Get("m").(string)
		}
	}

	result, err := arguments.Pack(ClaimItems{
		Id: [16]byte(idValue),
		Cd: claimDataList,
		M:  mValue,
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}
