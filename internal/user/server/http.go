package server

import (
	"github.com/dapr-ddd-action/pkg/errorx"
	"github.com/gofiber/fiber/v2"
)

func NewHttpServer(servers ...Server) *fiber.App {
	config := fiber.Config{
		DisableStartupMessage: true,
		// custom error handler
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			errx := errorx.From(err)
			return c.Status(errx.Code).JSON(err)
		},
	}
	app := fiber.New(config)

	for _, s := range servers {
		s.RegisterHTTPRouter(app)
	}

	return app
}
