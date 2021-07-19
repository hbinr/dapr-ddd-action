package jsonx

import (
	"github.com/bytedance/sonic"
)

func Marshal(val interface{}) ([]byte, error) {
	return sonic.Marshal(val)
}

func Unmarshal(buf []byte, val interface{}) error {
	return sonic.Unmarshal(buf, val)
}
