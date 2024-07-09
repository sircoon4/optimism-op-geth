package actions

import (
	"encoding/base64"
	"fmt"
	"math/big"
	"reflect"

	"github.com/sircoon4/bencodex-go"
	"github.com/sircoon4/bencodex-go/bencodextype"
)

func ExtractActionFromSerializedPayload(serializedPayload []byte) (*bencodextype.Dictionary, error) {
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
	fmt.Println("action:", action)
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
