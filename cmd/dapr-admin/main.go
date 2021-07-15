package main

import (
	"log"
	"net/http"

	"github.com/dapr/kit/logger"

	"github.com/dapr-ddd-action/internal/controller"
	appRepo "github.com/dapr-ddd-action/internal/repository"
	appService "github.com/dapr-ddd-action/internal/service"
	daprd "github.com/dapr/go-sdk/service/http"
)

var serviceAddress = ":8090"

func main() {
	server := daprd.NewService(serviceAddress)

	userRepo, err := appRepo.NewUserRepo(logger.NewLogger("dapr-ddd-action"))
	if err != nil {
		log.Fatalf("main: dapr.NewClient error: %v+\n", err)

	}
	userService := appService.NewUserService(userRepo)

	userCtl := controller.NewUserController(userService)
	if err := server.AddServiceInvocationHandler("/user", userCtl.GetUser); err != nil {
		log.Fatalf("main: AddServiceInvocationHandler error: %v+\n", err)
	}

	if err := server.Start(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("main:  listenning error: %v+\n", err)
	}
}
