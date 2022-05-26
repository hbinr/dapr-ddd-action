package jsonx

import (
	"errors"

	"github.com/bytedance/sonic"
)

func Marshal(val interface{}) ([]byte, error) {
	return sonic.Marshal(val)
}

func Unmarshal(buf []byte, val interface{}) error {
	if len(buf) == 0 {
		return errors.New("jsonx: empty data to unmarshal")
	}

	return sonic.Unmarshal(buf, val)
}
