package libplanet

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/sircoon4/bencodex-go/bencodextype"
)

type TransferAsset struct {
	Sender    common.Address     `abi:"sender"`
	Recipient common.Address     `abi:"recipient"`
	Amount    FungibleAssetValue `abi:"amount"`
	Memo      string             `abi:"memo"`
}

func convertToTransferAssetEthAbi(actionValues *bencodextype.Dictionary) ([]byte, error) {
	var TupleTransferAsset, _ = abi.NewType("tuple", "struct TransferAsset", []abi.ArgumentMarshaling{
		{Name: "sender", Type: "address"},
		{Name: "recipient", Type: "address"},
		{Name: "amount", Type: "tuple", Components: []abi.ArgumentMarshaling{
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
		{Name: "memo", Type: "string"},
	})

	var arguments = abi.Arguments{
		abi.Argument{Name: "TransferAsset", Type: TupleTransferAsset, Indexed: false},
	}

	senderValue := common.BytesToAddress(actionValues.Get("sender").([]byte))
	recipientValue := common.BytesToAddress(actionValues.Get("recipient").([]byte))
	amountValue := extractFungibleAssetValue(actionValues.Get("amount").([]any))

	memoValue := ""
	if actionValues.Contains("memo") {
		if actionValues.Get("memo") != nil {
			memoValue = actionValues.Get("memo").(string)
		}
	}

	result, err := arguments.Pack(TransferAsset{
		Sender:    senderValue,
		Recipient: recipientValue,
		Amount:    amountValue,
		Memo:      memoValue,
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}
