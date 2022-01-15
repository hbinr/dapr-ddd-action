package jsonx

import (
	"encoding/json"
	"errors"
)

// for Windows
func Marshal(val interface{}) ([]byte, error) {
	return json.Marshal(val)

}

func Unmarshal(buf []byte, val interface{}) error {
	if len(buf) == 0 {
		return errors.New("jsonx: empty data to unmarshal")
	}

	return json.Unmarshal(buf, val)
}

//  for Linux/darwin
// func Marshal(val interface{}) ([]byte, error) {
// 	return sonic.Marshal(val)
// }

// func Unmarshal(buf []byte, val interface{}) error {
// 	if len(buf) == 0 {
// 		return errors.New("jsonx: empty data to unmarshal")
// 	}

// 	return sonic.Unmarshal(buf, val)
// }
