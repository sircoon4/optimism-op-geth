package libplanet

import (
	"encoding/base64"
	"fmt"
	"math/big"
	"reflect"

	"github.com/ethereum/go-ethereum/common"
	"github.com/sircoon4/bencodex-go"
	"github.com/sircoon4/bencodex-go/bencodextype"
)

func ExtractActionEthAbi(action *bencodextype.Dictionary) ([]byte, error) {
	actionType, ok := action.Get("type_id").(string)
	if !ok {
		return nil, fmt.Errorf("error while getting type_id")
	}
	actionValues, ok := action.Get("values").(*bencodextype.Dictionary)
	if !ok {
		return nil, fmt.Errorf("error while getting values")
	}

	var abi []byte
	var err error
	switch actionType {
	case "hack_and_slash22":
		abi, err = convertToHackAndSlashEthAbi(actionValues)
		if err != nil {
			return nil, err
		}
	case "grinding2":
		abi, err = convertToGrindingEthAbi(actionValues)
		if err != nil {
			return nil, err
		}
	case "combination_equipment17":
		abi, err = convertToCombinationEquipmentEthAbi(actionValues)
		if err != nil {
			return nil, err
		}
	case "rapid_combination10":
		abi, err = convertToRapidCombinationEthAbi(actionValues)
		if err != nil {
			return nil, err
		}
	case "hack_and_slash_sweep10":
		abi, err = convertToHackAndSlashSweepEthAbi(actionValues)
		if err != nil {
			return nil, err
		}
	case "transfer_asset5":
		abi, err = convertToTransferAssetEthAbi(actionValues)
		if err != nil {
			return nil, err
		}
	case "claim_items":
		abi, err = convertToClaimItemsEthAbi(actionValues)
		if err != nil {
			return nil, err
		}
	case "daily_reward7":
		abi, err = convertToDailyRewardEthAbi(actionValues)
		if err != nil {
			return nil, err
		}
	case "aura_summon":
		abi, err = convertToAuraSummonEthAbi(actionValues)
		if err != nil {
			return nil, err
		}
	case "explore_adventure_boss":
		abi, err = convertToExploreAdventureBossEthAbi(actionValues)
		if err != nil {
			return nil, err
		}
	default:
		abi, err = bencodex.Encode(actionValues)
		if err != nil {
			return nil, err
		}
	}

	return common.CopyBytes(abi), nil
}

func ExtractActionDictFromSerializedPayload(serializedPayload []byte) (*bencodextype.Dictionary, error) {
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
