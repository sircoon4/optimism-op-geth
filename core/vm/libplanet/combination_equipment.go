package libplanet

import (
	"strconv"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/sircoon4/bencodex-go/bencodextype"
)

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
		{Name: "id", Type: "bytes16"},
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
