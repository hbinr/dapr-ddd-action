package chix

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"go.uber.org/zap"
)

// QueryInt64 Parsing http parameters or query parameters
func QueryInt64(param string, r *http.Request) (intVar int64, err error) {
	intStr := chi.URLParam(r, param)

	if intVar, err = strconv.ParseInt(intStr, 10, 64); err != nil {
		zap.L().Error("strconv.ParseInt(intStr) failed", zap.Error(err), zap.String("param", param))
		return
	}

	return
}
