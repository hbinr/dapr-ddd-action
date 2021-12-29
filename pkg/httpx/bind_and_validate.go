package httpx

import (
	"net/http"

	"github.com/bytedance/go-tagexpr/v2/binding"
)

func BindAndValidate(recvPointer interface{}, r *http.Request) error {
	binder := binding.New(nil)

	if err := binder.BindAndValidate(recvPointer, r, nil); err != nil {
		return err
	}

	return nil
}
