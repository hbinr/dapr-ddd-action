package daprhelp

import (
	"github.com/dapr-ddd-action/pkg/jsonx"
	"github.com/dapr/go-sdk/client"
)

// BuildExpireStateItem 创建含过期时间的 State item 参数. expire默认秒
func BuildExpireStateItem(key string, data interface{}, expire int) (*client.SetStateItem, error) {
	value, err := jsonx.Marshal(data)
	if err != nil {
		return nil, err
	}

	return &client.SetStateItem{
		Key:   key,
		Value: value,
		Metadata: map[string]string{
			// ttlInSeconds (time-to-live in seconds) 过期时间
			"ttlInSeconds": string(rune(expire)),
		},
	}, nil
}
