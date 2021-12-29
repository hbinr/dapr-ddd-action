package httpx

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"go.uber.org/zap"
)

// QueryInt64 Parsing http parameters or query parameters
func QueryInt64(param string, r *http.Request) (intVar int64, err error) {
	varsMap := mux.Vars(r)
	intStr := varsMap[param]

	if intVar, err = strconv.ParseInt(intStr, 10, 64); err != nil {
		zap.L().Error("strconv.ParseInt(intStr) failed", zap.Error(err), zap.String("param", param))
		return
	}

	return
}
