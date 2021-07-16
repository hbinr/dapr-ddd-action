package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/dapr/kit/logger"
	"github.com/gorilla/mux"

	daprCommon "github.com/dapr/go-sdk/service/common"
	daprd "github.com/dapr/go-sdk/service/http"

	"github.com/dapr-ddd-action/internal/controller"
	appRepo "github.com/dapr-ddd-action/internal/repository"
	appService "github.com/dapr-ddd-action/internal/service"
)

var serviceAddress = ":8090"

func main() {
	userRepo, err := appRepo.NewUserRepo(logger.NewLogger("dapr-ddd-action"))
	if err != nil {
		log.Fatalf("main: dapr.NewClient error: %v+\n", err)

	}
	userService := appService.NewUserService(userRepo)

	userCtl := controller.NewUserController(userService)

	//server := RegisterHttpServiceByDaprHttpService(userCtl)

	//server := RegisterHttpServiceByMux(userCtl)

	server := RegisterHttpServiceByGin(userCtl)
	if err := server.Start(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("main:  listenning error: %v+\n", err)
	}
}

// RegisterHttpServiceByDaprHttpService 使用 darp 提供的 http service 接口注册http服务
func RegisterHttpServiceByDaprHttpService(ctl controller.UserController) daprCommon.Service {
	server := daprd.NewService(serviceAddress)
	if err := server.AddServiceInvocationHandler("/user", ctl.GetUser); err != nil {
		log.Fatalf("main: AddServiceInvocationHandler error: %v+\n", err)
	}
	return server
}

// RegisterHttpServiceByMux 使用 mux 注册 http 服务
func RegisterHttpServiceByMux(ctl controller.UserController) daprCommon.Service {
	netMux := http.NewServeMux()

	r := mux.NewRouter()
	r.HandleFunc("/hello/{world}", ctl.SayHi).Methods("GET")
	netMux.Handle("/", r)

	server := daprd.NewServiceWithMux(serviceAddress, netMux)
	return server
}

// RegisterHttpServiceByGin 使用 gin 注册 http 服务
func RegisterHttpServiceByGin(ctl controller.UserController) daprCommon.Service {
	netMux := http.NewServeMux()

	r := gin.New()
	r.GET("/hello", ctl.SayHello)
	netMux.Handle("/", r)

	server := daprd.NewServiceWithMux(serviceAddress, netMux)
	return server
}
