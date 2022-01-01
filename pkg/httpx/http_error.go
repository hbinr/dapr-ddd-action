package httpx

import (
	"net/http"

	"github.com/dapr-ddd-action/pkg/jsonx"

	"go.uber.org/zap"
)

type ErrorResponse struct {
	Slug       string `json:"slug"`
	Msg        string `json:"msg"`
	httpStatus int
}

func InternalError(slug string, msg string, err error, w http.ResponseWriter) {
	httpRespondWithError(err, slug, msg, w, "Internal server error", http.StatusInternalServerError)
}

func Unauthorized(slug string, msg string, err error, w http.ResponseWriter) {
	httpRespondWithError(err, slug, msg, w, "Unauthorized", http.StatusUnauthorized)
}

func BadRequest(slug string, msg string, err error, w http.ResponseWriter) {
	httpRespondWithError(err, slug, msg, w, "Bad request", http.StatusBadRequest)
}

func httpRespondWithError(err error, slug string, msg string,
	w http.ResponseWriter, logMsg string, status int) {
	resp := ErrorResponse{slug, msg, status}

	zap.L().Error(logMsg, zap.Error(err))

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)

	dataByte, err := jsonx.Marshal(resp)
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("dataByte marshal error"))
	}
	w.Write(dataByte)
}
