package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/dapr-ddd-action/pkg/server"

	"github.com/dapr-ddd-action/internal/user/adapters/repository/data/dao"

	"github.com/dapr-ddd-action/pkg/util/file"

	"github.com/dapr-ddd-action/internal/user/adapters/repository"
	"github.com/dapr-ddd-action/internal/user/app"
	"github.com/dapr-ddd-action/internal/user/ports"
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

	if path := file.GetCurrentPath(); path != "" {
		defaultConfigFilePath = path + "/configs/config.yaml"
	}

	// init config
	appConf := conf.Init(defaultConfigFilePath)

	// init logger
	logger := zapLogger.InitZap(appConf)

	// init gorm client
	gormClient := dao.Use(database.Init(&appConf.Database))

	// init user bussiness
	userRepo := repository.NewUserRepo(client, logger, gormClient)
	userApp := app.NewApplication(userRepo)
	userController := ports.NewUserController(userApp)

	httpServer := server.NewHttpServer(userController)
	// start server
	if err := httpServer.Listen(fmt.Sprintf(":%d", appConf.Port)); err != nil {
		log.Panic(err)
	}

	c := make(chan os.Signal, 1)                    // Create channel to signify a signal being sent
	signal.Notify(c, os.Interrupt, syscall.SIGTERM) // When an interrupt or termination signal is sent, notify the channel

	<-c
	logger.Info("Gracefully shutting down...")
	_ = httpServer.Shutdown()

	client.Close()
	logger.Info("Fiber was successful shutdown.")
}
