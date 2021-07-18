package ginx

import (
	"github.com/dapr-ddd-action/app/pkg/conf"
	"github.com/gin-gonic/gin"
)

func NewGinEngine(cfg *conf.Config) *gin.Engine {
	if cfg.Mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) // gin设置成发布模式
	}
	r := gin.New()

	return r
}
