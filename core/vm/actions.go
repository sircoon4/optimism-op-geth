package vm

import (
	"encoding/base64"
	"fmt"
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
	for _, r := range actionValues.Get("r").([]any) {
		rInfo := []int64{}
		rFirstValue, _ := strconv.Atoi(r.([]any)[0].(string))
		rSecondValue, _ := strconv.Atoi(r.([]any)[1].(string))
		rInfo = append(rInfo, int64(rFirstValue), int64(rSecondValue))
		rList = append(rList, rInfo)
	}
	worldIdValue, _ := strconv.Atoi(actionValues.Get("worldId").(string))
	stageIdValue, _ := strconv.Atoi(actionValues.Get("stageId").(string))
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
		AvatarAddress:  avatarAddressValue,
		TotalPlayCount: int64(totalPlayCountValue),
		ApStoneCount:   int64(apStoneCountValue),
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}
