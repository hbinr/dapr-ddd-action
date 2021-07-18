package internal

import (
	"net/http"

	"github.com/dapr-ddd-action/app/user/service/internal/controller"
	"github.com/gin-gonic/gin"
)

func RegisterUserRouter(ctl controller.UserController) *http.ServeMux {
	netMux := http.NewServeMux()

	r := gin.New()
	r.GET("/user/:id", ctl.GetUser)
	netMux.Handle("/", r)

	return netMux
}
