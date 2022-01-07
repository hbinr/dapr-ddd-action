package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/dapr-ddd-action/internal/user/app"
	"github.com/dapr-ddd-action/internal/user/domain"
	"github.com/dapr-ddd-action/internal/user/ports"
	"github.com/gofiber/fiber/v2"

	"github.com/dapr-ddd-action/internal/user/adapters/repository"
	"github.com/dapr-ddd-action/pkg/conf"
	"github.com/dapr-ddd-action/pkg/errorx"
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

	// init config
	appConf, err := conf.Init(defaultConfigFilePath)
	if err != nil {
		log.Fatalf("main: init config error: %+v\n", err)
	}

	// init logger
	logger, err := zapLogger.InitZap(appConf)
	if err != nil {
		log.Fatalf("main: init config zap log error :%+v\n", err)
	}

	// init bussiness
	userRepo := repository.NewUserRepo(client, logger)
	userDomain := domain.NewUserDomain(userRepo)
	application := app.NewApplication(userDomain)

	// init fiber
	config := fiber.Config{
		DisableStartupMessage: true,
		// custom error handler
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			errx := errorx.From(err)
			return c.Status(errx.Code).JSON(err)
		},
	}
	app := fiber.New(config)
	ports.RegisterUserRouter(app, application)

	// start server
	if err := app.Listen(fmt.Sprintf(":%d", appConf.Port)); err != nil {
		app.Shutdown()
	}
	go func() {
		if err := app.Listen(fmt.Sprintf(":%d", appConf.Port)); err != nil {
			log.Panic(err)
		}
	}()

	c := make(chan os.Signal, 1)                    // Create channel to signify a signal being sent
	signal.Notify(c, os.Interrupt, syscall.SIGTERM) // When an interrupt or termination signal is sent, notify the channel

	<-c
	fmt.Println("Gracefully shutting down...")
	_ = app.Shutdown()

	client.Close()
	fmt.Println("Fiber was successful shutdown.")

}
