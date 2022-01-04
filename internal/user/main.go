package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dapr-ddd-action/internal/user/ports"

	"github.com/dapr-ddd-action/internal/user/service"

	"github.com/gorilla/mux"

	"github.com/dapr-ddd-action/internal/user/adapters/repository"
	"github.com/dapr-ddd-action/pkg/conf"
	zapLogger "github.com/dapr-ddd-action/pkg/loggger"

	daprCommon "github.com/dapr/go-sdk/service/common"

	dapr "github.com/dapr/go-sdk/client"
	daprd "github.com/dapr/go-sdk/service/http"
)

var defaultConfigFilePath = "./configs/config.yaml"

func main() {
	client, err := dapr.NewClient()

	if err != nil {
		log.Fatalf("main: new dapr client error :%+v\n", err)
	}
	defer client.Close()

	if err := initServer(client).Start(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("main: listening error :%+v\n", err)
	}
}

// initServer 初识化服务
func initServer(client dapr.Client) daprCommon.Service {
	appConf, err := conf.Init(defaultConfigFilePath)
	if err != nil {
		log.Fatalf("main: init config error: %+v\n", err)
	}

	logger, err := zapLogger.InitZap(appConf)
	if err != nil {
		log.Fatalf("main: init config zap log error :%+v\n", err)
	}

	userRepo := repository.NewUserRepo(client, logger)

	app := service.NewApplication(userRepo)
	router := mux.NewRouter()
	ports.RegisterUserRouter(router, app)

	server := daprd.NewServiceWithMux(fmt.Sprintf(":%d", appConf.Port), router)
	return server
}
