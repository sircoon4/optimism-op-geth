package libplanet

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/sircoon4/bencodex-go/bencodextype"
)

type FungibleAssetValue struct {
	Currency Currency `abi:"currency"`
	RawValue *big.Int `abi:"rawValue"`
}

type Currency struct {
	DecimalPlaces        uint8            `abi:"decimalPlaces"`
	Minters              []common.Address `abi:"minters"`
	Ticker               string           `abi:"ticker"`
	TotalSupplyTrackable bool             `abi:"totalSupplyTrackable"`
	MaximumSupplyMajor   *big.Int         `abi:"maximumSupplyMajor"`
	MaximumSupplyMinor   *big.Int         `abi:"maximumSupplyMinor"`
}

func extractFungibleAssetValue(value []any) FungibleAssetValue {
	currencyDict := value[0].(*bencodextype.Dictionary)
	preDecimalPlacesValue, _ := currencyDict.Get("decimalPlaces").([]byte)
	decimalPlacesValue := preDecimalPlacesValue[0]
	mintersList := []common.Address{}
	if currencyDict.Get("minters") != nil {
		for _, minters := range currencyDict.Get("minters").([]any) {
			mintersValue := common.BytesToAddress(minters.([]byte))
			mintersList = append(mintersList, mintersValue)
		}
	}
	tickerValue := currencyDict.Get("ticker").(string)
	totalSupplyTrackable := false
	maximumSupplyMajor := big.NewInt(0)
	maximumSupplyMinor := big.NewInt(0)
	if currencyDict.Contains("totalSupplyTrackable") {
		totalSupplyTrackable = currencyDict.Get("totalSupplyTrackable").(bool)
		if currencyDict.Contains("maximumSupplyMajor") {
			if currencyDict.Get("maximumSupplyMajor") != nil {
				maximumSupplyMajor = getAsBigInt(currencyDict.Get("maximumSupplyMajor"))
			}
		}
		if currencyDict.Contains("maximumSupplyMinor") {
			if currencyDict.Get("maximumSupplyMinor") != nil {
				maximumSupplyMinor = getAsBigInt(currencyDict.Get("maximumSupplyMinor"))
			}
		}
	}
	rawValue := getAsBigInt(value[1])
	return FungibleAssetValue{
		Currency: Currency{
			DecimalPlaces:        decimalPlacesValue,
			Minters:              mintersList,
			Ticker:               tickerValue,
			TotalSupplyTrackable: totalSupplyTrackable,
			MaximumSupplyMajor:   maximumSupplyMajor,
			MaximumSupplyMinor:   maximumSupplyMinor,
		},
		RawValue: rawValue,
	}
}
