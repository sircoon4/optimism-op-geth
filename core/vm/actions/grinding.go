package actions

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/sircoon4/bencodex-go/bencodextype"
)

type Grinding struct {
	Id [16]byte       `abi:"id"`
	A  common.Address `abi:"a"`
	E  [][16]byte     `abi:"e"`
	C  bool           `abi:"c"`
}

func ConvertToGrindingEthAbi(actionValues *bencodextype.Dictionary) ([]byte, error) {
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
