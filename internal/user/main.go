package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	file2 "github.com/dapr-ddd-action/pkg/util/file"

	"github.com/dapr-ddd-action/internal/user/app"
	"github.com/dapr-ddd-action/internal/user/domain"
	"github.com/dapr-ddd-action/internal/user/domain/data/dao"
	"github.com/dapr-ddd-action/internal/user/ports"
	"github.com/dapr-ddd-action/internal/user/server"

	"github.com/dapr-ddd-action/internal/user/adapters/repository"
	"github.com/dapr-ddd-action/pkg/conf"
	"github.com/dapr-ddd-action/pkg/database"
	zapLogger "github.com/dapr-ddd-action/pkg/loggger"

	dapr "github.com/dapr/go-sdk/client"
)

var defaultConfigFilePath = "./configs/config.yaml"

func main() {
	// init dapr client
	client, err := dapr.NewClient()
	if err != nil {
		log.Fatalf("main: new dapr client error :%+v\n", err)
	}

	if path := file2.GetCurrentPath(); path != "" {
		defaultConfigFilePath = path + "/configs/config.yaml"
	}

	// init config
	appConf := conf.Init(defaultConfigFilePath)

	// init logger
	logger := zapLogger.InitZap(appConf)

	// init gorm client
	gormClient := dao.Use(database.Init(&appConf.Database))

	// init bussiness
	userRepo := repository.NewUserRepo(client, logger, gormClient)
	userDomain := domain.NewUserDomain(userRepo)
	userApp := app.NewApplication(userDomain)
	userController := ports.NewUserController(userApp)
	app := server.NewHttpServer(userController)

	// start server
	if err := app.Listen(fmt.Sprintf(":%d", appConf.Port)); err != nil {
		log.Panic(err)
	}

	c := make(chan os.Signal, 1)                    // Create channel to signify a signal being sent
	signal.Notify(c, os.Interrupt, syscall.SIGTERM) // When an interrupt or termination signal is sent, notify the channel

	<-c
	logger.Info("Gracefully shutting down...")
	_ = app.Shutdown()

	client.Close()
	logger.Info("Fiber was successful shutdown.")
}
