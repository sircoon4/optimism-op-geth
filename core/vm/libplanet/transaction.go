package libplanet

import (
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/sircoon4/bencodex-go/bencodextype"
)

type Transaction struct {
	Signature        []byte             `abi:"signature"`
	Actions          []Action           `abi:"actions"`
	GenesisHash      [32]byte           `abi:"genesisHash"`
	GasLimit         int64              `abi:"gasLimit"`
	MaxGasPrice      FungibleAssetValue `abi:"maxGasPrice"`
	Nonce            int64              `abi:"nonce"`
	PublicKey        []byte             `abi:"publicKey"`
	Signer           common.Address     `abi:"signer"`
	Timestamp        *big.Int           `abi:"timestamp"`
	UpdatedAddresses []common.Address   `abi:"updatedAddresses"`
}

type Action struct {
	TypeId string `abi:"typeId"`
	Values []byte `abi:"values"`
}

func ConvertToTransactionEthAbi(actionValues *bencodextype.Dictionary) ([]byte, error) {
	var TupleTransaction, _ = abi.NewType("tuple", "struct Transaction", []abi.ArgumentMarshaling{
		{Name: "signature", Type: "bytes"},
		{Name: "actions", Type: "tuple[]", Components: []abi.ArgumentMarshaling{
			{Name: "typeId", Type: "string"},
			{Name: "values", Type: "bytes"},
		}},
		{Name: "genesisHash", Type: "bytes32"},
		{Name: "gasLimit", Type: "int64"},
		{Name: "maxGasPrice", Type: "tuple", Components: []abi.ArgumentMarshaling{
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
		{Name: "nonce", Type: "int64"},
		{Name: "publicKey", Type: "bytes"},
		{Name: "signer", Type: "address"},
		{Name: "timestamp", Type: "uint256"},
		{Name: "updatedAddresses", Type: "address[]"},
	})

	var arguments = abi.Arguments{
		abi.Argument{Name: "Transaction", Type: TupleTransaction, Indexed: false},
	}

	signatureValue, _ := actionValues.Get([]byte{0x53}).([]byte)
	actionsList := []Action{}
	for _, actionValue := range actionValues.Get([]byte{0x61}).([]any) {
		action := actionValue.(*bencodextype.Dictionary)
		typeIdValue := action.Get("type_id").(string)
		valuesValue, err := ExtractActionEthAbi(action)
		if err != nil {
			return nil, err
		}
		actionsList = append(actionsList, Action{
			TypeId: typeIdValue,
			Values: valuesValue,
		})
	}
	genesisHashValue, _ := actionValues.Get([]byte{0x67}).([]byte)
	gasLimitValue, _ := actionValues.Get([]byte{0x6c}).(int)
	maxGasPriceValue := extractFungibleAssetValue(actionValues.Get([]byte{0x6d}).([]any))
	nonceValue, _ := actionValues.Get([]byte{0x6e}).(int)
	publicKeyValue, _ := actionValues.Get([]byte{0x70}).([]byte)
	signerValue := common.BytesToAddress(actionValues.Get([]byte{0x73}).([]byte))
	timestamp, _ := time.Parse(time.RFC3339Nano, actionValues.Get([]byte{0x74}).(string))
	timestampValue := big.NewInt(timestamp.UnixNano())
	updatedAddressesList := []common.Address{}
	for _, updatedAddress := range actionValues.Get([]byte{0x75}).([]any) {
		updatedAddressesList = append(updatedAddressesList, common.BytesToAddress(updatedAddress.([]byte)))
	}

	result, err := arguments.Pack(Transaction{
		Signature:        signatureValue,
		Actions:          actionsList,
		GenesisHash:      [32]byte(genesisHashValue),
		GasLimit:         int64(gasLimitValue),
		MaxGasPrice:      maxGasPriceValue,
		Nonce:            int64(nonceValue),
		PublicKey:        publicKeyValue,
		Signer:           signerValue,
		Timestamp:        timestampValue,
		UpdatedAddresses: updatedAddressesList,
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}
