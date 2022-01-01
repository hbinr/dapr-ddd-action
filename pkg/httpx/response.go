package httpx

import (
	"net/http"

	"github.com/dapr-ddd-action/pkg/errorx"

	"github.com/dapr-ddd-action/pkg/jsonx"
)

func RespSuccess(data interface{}, w http.ResponseWriter) {
	dataByte, err := jsonx.Marshal(data)
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("dataByte marshal error"))
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(dataByte)
}

func RespWithError(err error, msg string, w http.ResponseWriter) {
	respErr, ok := err.(errorx.SlugErr)
	if !ok {
		InternalError("internal-server-error", msg, err, w)
		return
	}

	switch respErr.ErrorType() {
	case errorx.ErrorTypeAuthorization:
		Unauthorized(respErr.Slug(), respErr.Error(), respErr, w)
	case errorx.ErrorTypeIncorrectInput:
		BadRequest(respErr.Slug(), respErr.Error(), respErr, w)
	default:
		InternalError(respErr.Slug(), respErr.Error(), respErr, w)
	}
}
