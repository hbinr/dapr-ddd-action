package httpx

import (
	"net/http"

	"github.com/dapr-ddd-action/pkg/jsonx"
)

func RespSuccess(data interface{}, w http.ResponseWriter) {
	dataByte, err := jsonx.Marshal(data)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("dataByte marshal error"))
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(dataByte))
}
