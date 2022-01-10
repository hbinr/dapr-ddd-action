package server

import (
	"github.com/dapr-ddd-action/pkg/errorx"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func NewHttpServer(servers ...Server) *fiber.App {
	config := fiber.Config{
		DisableStartupMessage: true,
		// custom error handler
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			zap.L().Error("internal err", zap.Error(err))
			// fiber 的error
			fiberNativeErr, ok := err.(*fiber.Error)
			if ok {
				return fiberNativeErr
			}

			// 自定义 error
			errx := errorx.From(err)
			return c.Status(errx.Code).JSON(errx)
		},
	}
	app := fiber.New(config)

	for _, s := range servers {
		s.RegisterHTTPRouter(app)
	}

	return app
}
