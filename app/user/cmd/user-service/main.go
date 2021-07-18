package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dapr-ddd-action/app/user/internal/controller"

	"github.com/dapr-ddd-action/pkg/ginx"

	"go.uber.org/zap"

	zapLogger "github.com/dapr-ddd-action/pkg/loggger"

	"github.com/dapr-ddd-action/app/pkg/conf"

	"github.com/dapr-ddd-action/app/user/internal/repository"
	"github.com/dapr-ddd-action/app/user/internal/service"

	daprCommon "github.com/dapr/go-sdk/service/common"

	daprd "github.com/dapr/go-sdk/service/http"
)

func main() {
	if err := initApp().Start(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("main: listening error :%v+\n", err)
	}
}

// initApp 初识化服务
func initApp() daprCommon.Service {
	appConf, err := conf.Init()
	if err != nil {
		log.Fatalf("main: init config error: %v+\n", err)
	}

	logger, err := zapLogger.InitZap(appConf)
	if err != nil {
		log.Fatalf("main: init config zap log error :%v+\n", err)
	}

	userRepo, err := repository.NewUserRepo(logger)
	if err != nil {
		logger.Fatal("main: NewUserRepo error :", zap.Error(err))
	}

	userService := service.NewUserService(userRepo, logger)
	engine := ginx.NewGinEngine(appConf)
	controller.RegisterUserRouter(engine, userService)

	mux := http.NewServeMux()
	mux.Handle("/", engine)
	appServer := daprd.NewServiceWithMux(fmt.Sprintf(":%d", appConf.Port), mux)
	return appServer
}
