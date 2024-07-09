package libplanet

import (
	"strconv"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/sircoon4/bencodex-go/bencodextype"
)

type ExploreAdventureBoss struct {
	Id            [16]byte       `abi:"id"`
	Season        int64          `abi:"season"`
	AvatarAddress common.Address `abi:"avatarAddress"`
	Costumes      [][16]byte     `abi:"costumes"`
	Equipments    [][16]byte     `abi:"equipments"`
	Foods         [][16]byte     `abi:"foods"`
	R             []RuneSlotInfo `abi:"r"`
	StageBuffId   int64          `abi:"stageBuffId"`
}

func convertToExploreAdventureBossEthAbi(actionValues *bencodextype.Dictionary) ([]byte, error) {
	var TupleExploreAdventureBoss, _ = abi.NewType("tuple", "struct ExploreAdventureBoss", []abi.ArgumentMarshaling{
		{Name: "id", Type: "bytes16"},
		{Name: "season", Type: "int64"},
		{Name: "avatarAddress", Type: "address"},
		{Name: "costumes", Type: "bytes16[]"},
		{Name: "equipments", Type: "bytes16[]"},
		{Name: "foods", Type: "bytes16[]"},
		{Name: "r", Type: "tuple[]", Components: []abi.ArgumentMarshaling{
			{Name: "slotIndex", Type: "int64"},
			{Name: "runeId", Type: "int64"},
		}},
		{Name: "stageBuffId", Type: "int64"},
	})

	var arguments = abi.Arguments{
		abi.Argument{Name: "ExploreAdventureBoss", Type: TupleExploreAdventureBoss, Indexed: false},
	}

	idValue, _ := actionValues.Get("id").([]byte)
	seasonValue, _ := actionValues.Get("season").(int)
	avatarAddressValue := common.BytesToAddress(actionValues.Get("avatarAddress").([]byte))

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
	for _, r := range actionValues.Get("r").([]any) {
		rList = append(rList, extractRuneSlotInfo(r.([]any)))
	}

	stageBuffIdValue := -1
	if actionValues.Contains("stageBuffId") {
		if actionValues.Get("stageBuffId") != nil {
			stageBuffIdValue, _ = strconv.Atoi(actionValues.Get("stageBuffId").(string))
		}
	}

	result, err := arguments.Pack(ExploreAdventureBoss{
		Id:            [16]byte(idValue),
		Season:        int64(seasonValue),
		AvatarAddress: avatarAddressValue,
		Costumes:      costumesList,
		Equipments:    equipmentsList,
		Foods:         foodsList,
		R:             rList,
		StageBuffId:   int64(stageBuffIdValue),
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}
