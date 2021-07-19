package ginx

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// QueryInt Parsing http parameters or query parameters
func QueryInt(param string, c *gin.Context) (intVar int64, err error) {
	var intStr string

	if intStr = c.Param(param); intStr == "" {
		intStr = c.Query(param)
	}

	if intVar, err = strconv.ParseInt(intStr, 10, 64); err != nil {
		zap.L().Error("strconv.ParseInt(intStr) failed", zap.Error(err), zap.String("param", param))
		return
	}

	return
}
