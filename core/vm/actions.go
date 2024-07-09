package vm

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"math/big"
	"reflect"
	"strconv"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/sircoon4/bencodex-go"
	"github.com/sircoon4/bencodex-go/bencodextype"
)

func extractActionFromSerializedPayload(serializedPayload []byte) (*bencodextype.Dictionary, error) {
	encoded, err := base64.StdEncoding.DecodeString(string(serializedPayload))
	if err != nil {
		return nil, err
	}
	decoded, err := bencodex.Decode(encoded)
	if err != nil {
		return nil, err
	}
	dict, ok := decoded.(*bencodextype.Dictionary)
	if !ok {
		return nil, fmt.Errorf("error while casting to dictionary")
	}
	action, ok := dict.Get([]byte{0x61}).([]any)[0].(*bencodextype.Dictionary)
	if !ok {
		return nil, fmt.Errorf("error while getting action")
	}
	return action, nil
}

type HackAndSlash struct {
	Id             [16]byte       `abi:"id"`
	Costumes       [][16]byte     `abi:"costumes"`
	Equipments     [][16]byte     `abi:"equipments"`
	Foods          [][16]byte     `abi:"foods"`
	R              [][]int64      `abi:"r"`
	WorldId        int64          `abi:"worldId"`
	StageId        int64          `abi:"stageId"`
	StageBuffId    int64          `abi:"stageBuffId"`
	AvatarAddress  common.Address `abi:"avatarAddress"`
	TotalPlayCount int64          `abi:"totalPlayCount"`
	ApStoneCount   int64          `abi:"apStoneCount"`
}

func convertToHackAndSlashEthAbi(actionValues *bencodextype.Dictionary) ([]byte, error) {
	var TupleHackAndSlash, _ = abi.NewType("tuple", "struct HackAndSlash", []abi.ArgumentMarshaling{
		{Name: "id", Type: "uint8[16]"},
		{Name: "costumes", Type: "uint8[16][]"},
		{Name: "equipments", Type: "uint8[16][]"},
		{Name: "foods", Type: "uint8[16][]"},
		{Name: "r", Type: "int64[][]"},
		{Name: "worldId", Type: "int64"},
		{Name: "stageId", Type: "int64"},
		{Name: "stageBuffId", Type: "int64"},
		{Name: "avatarAddress", Type: "address"},
		{Name: "totalPlayCount", Type: "int64"},
		{Name: "apStoneCount", Type: "int64"},
	})

	var arguments = abi.Arguments{
		abi.Argument{Name: "HackAndSlash", Type: TupleHackAndSlash, Indexed: false},
	}

	idValue, _ := actionValues.Get("id").([]byte)
	costumesList := [][16]byte{}
	for _, costume := range actionValues.Get("costumes").([]any) {
		costumeValue, _ := costume.([]byte)
		costumesList = append(costumesList, [16]byte(costumeValue))
	}
	equipmentsList := [][16]byte{}
	for _, equipment := range actionValues.Get("equipments").([]any) {
		equipmentValue, _ := equipment.([]byte)
		equipmentsList = append(equipmentsList, [16]byte(equipmentValue))
	}
	foodsList := [][16]byte{}
	for _, food := range actionValues.Get("foods").([]any) {
		foodValue, _ := food.([]byte)
		foodsList = append(foodsList, [16]byte(foodValue))
	}
	rList := [][]int64{}
	for _, rInfo := range actionValues.Get("r").([]any) {
		rInfoList := []int64{}
		for _, r := range rInfo.([]any) {
			rValue, _ := strconv.Atoi(r.(string))
			rInfoList = append(rInfoList, int64(rValue))
		}
		rList = append(rList, rInfoList)
	}
	worldIdValue, _ := strconv.Atoi(actionValues.Get("worldId").(string))
	stageIdValue, _ := strconv.Atoi(actionValues.Get("stageId").(string))
	stageBuffIdValue := -1
	if actionValues.Contains("stageBuffId") {
		if actionValues.Get("stageBuffId") != nil {
			stageBuffIdValue, _ = strconv.Atoi(actionValues.Get("stageBuffId").(string))
		}
	}
	avatarAddressValue := common.BytesToAddress(actionValues.Get("avatarAddress").([]byte))
	totalPlayCountValue, _ := strconv.Atoi(actionValues.Get("totalPlayCount").(string))
	apStoneCountValue, _ := strconv.Atoi(actionValues.Get("apStoneCount").(string))

	result, err := arguments.Pack(HackAndSlash{
		Id:             [16]byte(idValue),
		Costumes:       costumesList,
		Equipments:     equipmentsList,
		Foods:          foodsList,
		R:              rList,
		WorldId:        int64(worldIdValue),
		StageId:        int64(stageIdValue),
		StageBuffId:    int64(stageBuffIdValue),
		AvatarAddress:  avatarAddressValue,
		TotalPlayCount: int64(totalPlayCountValue),
		ApStoneCount:   int64(apStoneCountValue),
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}

type Grinding struct {
	Id [16]byte       `abi:"id"`
	A  common.Address `abi:"a"`
	E  [][16]byte     `abi:"e"`
	C  bool           `abi:"c"`
}

func convertToGrindingEthAbi(actionValues *bencodextype.Dictionary) ([]byte, error) {
	var TupleGrinding, _ = abi.NewType("tuple", "struct Grinding", []abi.ArgumentMarshaling{
		{Name: "id", Type: "uint8[16]"},
		{Name: "a", Type: "address"},
		{Name: "e", Type: "uint8[16][]"},
		{Name: "c", Type: "bool"},
	})

	var arguments = abi.Arguments{
		abi.Argument{Name: "Grinding", Type: TupleGrinding, Indexed: false},
	}

	idValue, _ := actionValues.Get("id").([]byte)
	aValue := common.BytesToAddress(actionValues.Get("a").([]byte))
	eList := [][16]byte{}
	for _, e := range actionValues.Get("e").([]any) {
		eValue, _ := e.([]byte)
		eList = append(eList, [16]byte(eValue))
	}
	cValue, _ := actionValues.Get("c").(bool)

	result, err := arguments.Pack(Grinding{
		Id: [16]byte(idValue),
		A:  aValue,
		E:  eList,
		C:  cValue,
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}

type CombinationEquipment struct {
	Id  [16]byte       `abi:"id"`
	A   common.Address `abi:"a"`
	S   int64          `abi:"s"`
	R   int64          `abi:"r"`
	I   int64          `abi:"i"`
	P   bool           `abi:"p"`
	H   bool           `abi:"h"`
	Pid int64          `abi:"pid"`
}

func convertToCombinationEquipmentEthAbi(actionValues *bencodextype.Dictionary) ([]byte, error) {
	var TupleCombinationEquipment, _ = abi.NewType("tuple", "struct CombinationEquipment", []abi.ArgumentMarshaling{
		{Name: "id", Type: "uint8[16]"},
		{Name: "a", Type: "address"},
		{Name: "s", Type: "int64"},
		{Name: "r", Type: "int64"},
		{Name: "i", Type: "int64"},
		{Name: "p", Type: "bool"},
		{Name: "h", Type: "bool"},
		{Name: "pid", Type: "int64"},
	})

	var arguments = abi.Arguments{
		abi.Argument{Name: "CombinationEquipment", Type: TupleCombinationEquipment, Indexed: false},
	}

	idValue, _ := actionValues.Get("id").([]byte)
	aValue := common.BytesToAddress(actionValues.Get("a").([]byte))
	sValue, _ := strconv.Atoi(actionValues.Get("s").(string))
	rValue, _ := strconv.Atoi(actionValues.Get("r").(string))
	iValue := -1
	if actionValues.Get("i") != nil {
		iValue, _ = strconv.Atoi(actionValues.Get("i").(string))
	}
	pValue, _ := actionValues.Get("p").(bool)
	hValue, _ := actionValues.Get("h").(bool)
	pidValue := -1
	if actionValues.Get("pid") != nil {
		pidValue, _ = strconv.Atoi(actionValues.Get("pid").(string))
	}

	result, err := arguments.Pack(CombinationEquipment{
		Id:  [16]byte(idValue),
		A:   aValue,
		S:   int64(sValue),
		R:   int64(rValue),
		I:   int64(iValue),
		P:   pValue,
		H:   hValue,
		Pid: int64(pidValue),
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RapidCombination struct {
	Id            [16]byte       `abi:"id"`
	AvatarAddress common.Address `abi:"avatarAddress"`
	SlotIndex     int64          `abi:"slotIndex"`
}

func convertToRapidCombinationEthAbi(actionValues *bencodextype.Dictionary) ([]byte, error) {
	var TupleRapidCombination, _ = abi.NewType("tuple", "struct RapidCombination", []abi.ArgumentMarshaling{
		{Name: "id", Type: "uint8[16]"},
		{Name: "avatarAddress", Type: "address"},
		{Name: "slotIndex", Type: "int64"},
	})

	var arguments = abi.Arguments{
		abi.Argument{Name: "RapidCombination", Type: TupleRapidCombination, Indexed: false},
	}

	idValue, _ := actionValues.Get("id").([]byte)
	avatarAddressValue := common.BytesToAddress(actionValues.Get("avatarAddress").([]byte))
	slotIndexValue, _ := strconv.Atoi(actionValues.Get("slotIndex").(string))

	result, err := arguments.Pack(RapidCombination{
		Id:            [16]byte(idValue),
		AvatarAddress: avatarAddressValue,
		SlotIndex:     int64(slotIndexValue),
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}

type HackAndSlashSweep struct {
	Id            [16]byte       `abi:"id"`
	Costumes      [][16]byte     `abi:"costumes"`
	Equipments    [][16]byte     `abi:"equipments"`
	RuneInfos     [][]int64      `abi:"runeInfos"`
	AvatarAddress common.Address `abi:"avatarAddress"`
	ApStoneCount  int64          `abi:"apStoneCount"`
	ActionPoint   int64          `abi:"actionPoint"`
	WorldId       int64          `abi:"worldId"`
	StageId       int64          `abi:"stageId"`
}

func convertToHackAndSlashSweepEthAbi(actionValues *bencodextype.Dictionary) ([]byte, error) {
	var TupleHackAndSlashSweep, _ = abi.NewType("tuple", "struct HackAndSlashSweep", []abi.ArgumentMarshaling{
		{Name: "id", Type: "uint8[16]"},
		{Name: "costumes", Type: "uint8[16][]"},
		{Name: "equipments", Type: "uint8[16][]"},
		{Name: "runeInfos", Type: "int64[][]"},
		{Name: "avatarAddress", Type: "address"},
		{Name: "apStoneCount", Type: "int64"},
		{Name: "actionPoint", Type: "int64"},
		{Name: "worldId", Type: "int64"},
		{Name: "stageId", Type: "int64"},
	})

	var arguments = abi.Arguments{
		abi.Argument{Name: "HackAndSlashSweep", Type: TupleHackAndSlashSweep, Indexed: false},
	}

	idValue, _ := actionValues.Get("id").([]byte)
	costumesList := [][16]byte{}
	for _, costume := range actionValues.Get("costumes").([]any) {
		costumeValue, _ := costume.([]byte)
		costumesList = append(costumesList, [16]byte(costumeValue))
	}
	equipmentsList := [][16]byte{}
	for _, equipment := range actionValues.Get("equipments").([]any) {
		equipmentValue, _ := equipment.([]byte)
		equipmentsList = append(equipmentsList, [16]byte(equipmentValue))
	}
	runeInfosList := [][]int64{}
	for _, runeInfo := range actionValues.Get("runeInfos").([]any) {
		runeInfoList := []int64{}
		for _, rune := range runeInfo.([]any) {
			runeValue, _ := strconv.Atoi(rune.(string))
			runeInfoList = append(runeInfoList, int64(runeValue))
		}
		runeInfosList = append(runeInfosList, runeInfoList)
	}
	avatarAddressValue := common.BytesToAddress(actionValues.Get("avatarAddress").([]byte))
	apStoneCountValue, _ := strconv.Atoi(actionValues.Get("apStoneCount").(string))
	actionPointValue, _ := strconv.Atoi(actionValues.Get("actionPoint").(string))
	worldIdValue, _ := strconv.Atoi(actionValues.Get("worldId").(string))
	stageIdValue, _ := strconv.Atoi(actionValues.Get("stageId").(string))

	result, err := arguments.Pack(HackAndSlashSweep{
		Id:            [16]byte(idValue),
		Costumes:      costumesList,
		Equipments:    equipmentsList,
		RuneInfos:     runeInfosList,
		AvatarAddress: avatarAddressValue,
		ApStoneCount:  int64(apStoneCountValue),
		ActionPoint:   int64(actionPointValue),
		WorldId:       int64(worldIdValue),
		StageId:       int64(stageIdValue),
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}

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
			},
			},
			{Name: "rawValue", Type: "uint256"},
		},
		},
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
		{Name: "id", Type: "uint8[16]"},
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

type DailyReward struct {
	Id [16]byte       `abi:"id"`
	A  common.Address `abi:"a"`
}

func convertToDailyRewardEthAbi(actionValues *bencodextype.Dictionary) ([]byte, error) {
	var TupleDailyReward, _ = abi.NewType("tuple", "struct DailyReward", []abi.ArgumentMarshaling{
		{Name: "id", Type: "uint8[16]"},
		{Name: "a", Type: "address"},
	})

	var arguments = abi.Arguments{
		abi.Argument{Name: "DailyReward", Type: TupleDailyReward, Indexed: false},
	}

	idValue, _ := actionValues.Get("id").([]byte)
	aValue := common.BytesToAddress(actionValues.Get("a").([]byte))

	result, err := arguments.Pack(DailyReward{
		Id: [16]byte(idValue),
		A:  aValue,
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}

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

func getAsBigInt(value any) *big.Int {
	if value == nil {
		return nil
	}
	valueBigInt, ok := value.(*big.Int)
	if ok {
		return valueBigInt
	}
	switch reflect.TypeOf(value).Kind() {
	case reflect.Int:
		return big.NewInt(int64(value.(int)))
	default:
		return nil
	}
}

// for testing purpose
type Simple struct {
	Id       [16]byte `abi:"id"`
	SimpleId int64    `abi:"simpleId"`
}

func convertToSimpleEthAbi() ([]byte, error) {
	var TupleSimple, _ = abi.NewType("tuple[]", "struct array Simple", []abi.ArgumentMarshaling{
		{Name: "id", Type: "uint8[16]"},
		{Name: "simpleId", Type: "int64"},
	})

	var arguments = abi.Arguments{
		abi.Argument{Name: "Simple", Type: TupleSimple, Indexed: false},
	}

	id1Value, _ := hex.DecodeString("b07f20d260297a42a5cadf4b835c2a01")
	simple1 := Simple{
		Id:       [16]byte(id1Value),
		SimpleId: 1,
	}

	id2Value, _ := hex.DecodeString("b07f20d260297a42a5cadf4b835c2a01")
	simple2 := Simple{
		Id:       [16]byte(id2Value),
		SimpleId: 1,
	}

	result, err := arguments.Pack([]Simple{
		simple1, simple2, simple1,
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}
