package libplanet

import (
	"strconv"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/sircoon4/bencodex-go/bencodextype"
)

type RapidCombination struct {
	Id            [16]byte       `abi:"id"`
	AvatarAddress common.Address `abi:"avatarAddress"`
	SlotIndex     int64          `abi:"slotIndex"`
}

func convertToRapidCombinationEthAbi(actionValues *bencodextype.Dictionary) ([]byte, error) {
	var TupleRapidCombination, _ = abi.NewType("tuple", "struct RapidCombination", []abi.ArgumentMarshaling{
		{Name: "id", Type: "bytes16"},
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
