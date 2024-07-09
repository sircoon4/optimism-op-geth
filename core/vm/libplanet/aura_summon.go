package libplanet

import (
	"strconv"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/sircoon4/bencodex-go/bencodextype"
)

type AuraSummon struct {
	Id  [16]byte       `abi:"id"`
	Aa  common.Address `abi:"aa"`
	Gid int64          `abi:"gid"`
	Sc  int64          `abi:"sc"`
}

func convertToAuraSummonEthAbi(actionValues *bencodextype.Dictionary) ([]byte, error) {
	var TupleAuraSummon, _ = abi.NewType("tuple", "struct AuraSummon", []abi.ArgumentMarshaling{
		{Name: "id", Type: "bytes16"},
		{Name: "aa", Type: "address"},
		{Name: "gid", Type: "int64"},
		{Name: "sc", Type: "int64"},
	})

	var arguments = abi.Arguments{
		abi.Argument{Name: "AuraSummon", Type: TupleAuraSummon, Indexed: false},
	}

	idValue, _ := actionValues.Get("id").([]byte)
	aaValue := common.BytesToAddress(actionValues.Get("aa").([]byte))
	gidValue, _ := strconv.Atoi(actionValues.Get("gid").(string))
	scValue, _ := strconv.Atoi(actionValues.Get("sc").(string))

	result, err := arguments.Pack(AuraSummon{
		Id:  [16]byte(idValue),
		Aa:  aaValue,
		Gid: int64(gidValue),
		Sc:  int64(scValue),
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}
