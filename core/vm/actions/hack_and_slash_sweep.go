package actions

import (
	"strconv"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/sircoon4/bencodex-go/bencodextype"
)

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

func ConvertToHackAndSlashSweepEthAbi(actionValues *bencodextype.Dictionary) ([]byte, error) {
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
