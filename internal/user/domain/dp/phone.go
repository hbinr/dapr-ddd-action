package dp

import (
	"strings"

	"github.com/pkg/errors"
)

// poneNumber dp (domain primitive) 实践
// 参考: https://mp.weixin.qq.com/s/tTnj4XHy-Q0S_25VO9F7gQ
type poneNumber struct {
	number string
}

func (p poneNumber) NewPhoneNumber(number string) (*poneNumber, error) {
	if p.number == "" {
		return nil, errors.New("empty phone number")
	} else if !p.isValid() {
		return nil, errors.New("invalid phone number")

	}
	return &poneNumber{number}, nil
}

func (p poneNumber) GetAreaCode() string {
	for i, length := 0, len(p.number); i < length; i++ {
		prefix := p.number[0:i]
		if isAreaCode(prefix) {
			return prefix
		}
	}
	return ""
}

func (p poneNumber) isValid() bool {
	// 正则匹配是否符合电话格式
	return true
}

func isAreaCode(prefix string) bool {
	areas := []string{"0571", "021", "010"}
	for _, item := range areas {
		if strings.Contains(item, prefix) {
			return true
		}
	}
	return false
}
