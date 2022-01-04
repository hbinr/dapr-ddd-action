package jsonx

import (
	"encoding/json"
	"errors"
	"runtime"

	"github.com/bytedance/sonic"
)

func Marshal(val interface{}) ([]byte, error) {
	switch runtime.GOOS {
	case "windows":
		return json.Marshal(val)
	case "darwin":
		return sonic.Marshal(val)
	case "linux":
		return sonic.Marshal(val)
	default:
		return nil, errors.New("jsonx: unknown os")
	}
}

func Unmarshal(buf []byte, val interface{}) error {
	if len(buf) == 0 {
		return errors.New("jsonx: empty data to unmarshal")
	}

	switch runtime.GOOS {
	case "windows":
		return json.Unmarshal(buf, val)
	case "darwin":
		return sonic.Unmarshal(buf, val)
	case "linux":
		return sonic.Unmarshal(buf, val)
	default:
		return errors.New("jsonx: unknown os")
	}

}
