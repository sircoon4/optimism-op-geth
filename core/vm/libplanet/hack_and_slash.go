package libplanet

import (
	"strconv"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/sircoon4/bencodex-go/bencodextype"
)

type HackAndSlash struct {
	Id             [16]byte       `abi:"id"`
	Costumes       [][16]byte     `abi:"costumes"`
	Equipments     [][16]byte     `abi:"equipments"`
	Foods          [][16]byte     `abi:"foods"`
	R              []RuneSlotInfo `abi:"r"`
	WorldId        int64          `abi:"worldId"`
	StageId        int64          `abi:"stageId"`
	StageBuffId    int64          `abi:"stageBuffId"`
	AvatarAddress  common.Address `abi:"avatarAddress"`
	TotalPlayCount int64          `abi:"totalPlayCount"`
	ApStoneCount   int64          `abi:"apStoneCount"`
}

func convertToHackAndSlashEthAbi(actionValues *bencodextype.Dictionary) ([]byte, error) {
	var TupleHackAndSlash, _ = abi.NewType("tuple", "struct HackAndSlash", []abi.ArgumentMarshaling{
		{Name: "id", Type: "bytes16"},
		{Name: "costumes", Type: "bytes16[]"},
		{Name: "equipments", Type: "bytes16[]"},
		{Name: "foods", Type: "bytes16[]"},
		{Name: "r", Type: "tuple[]", Components: []abi.ArgumentMarshaling{
			{Name: "slotIndex", Type: "int64"},
			{Name: "runeId", Type: "int64"},
		}},
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
	rList := []RuneSlotInfo{}
	for _, rInfo := range actionValues.Get("r").([]any) {
		rList = append(rList, extractRuneSlotInfo(rInfo.([]any)))
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
