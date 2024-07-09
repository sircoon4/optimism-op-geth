package actions

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/sircoon4/bencodex-go/bencodextype"
)

type DailyReward struct {
	Id [16]byte       `abi:"id"`
	A  common.Address `abi:"a"`
}

func ConvertToDailyRewardEthAbi(actionValues *bencodextype.Dictionary) ([]byte, error) {
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
