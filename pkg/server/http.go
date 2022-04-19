package server

import (
	"github.com/dapr-ddd-action/pkg/errorx"
	"github.com/dapr-ddd-action/pkg/jsonx"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"go.uber.org/zap"
)

func NewHttpServer(servers ...Server) *fiber.App {
	config := fiber.Config{
		DisableStartupMessage: true,
		// custom error handler
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			// fiber error
			fiberNativeErr, ok := err.(*fiber.Error)
			code := fiber.StatusInternalServerError
			if ok {
				code = fiberNativeErr.Code
				err = fiberNativeErr
			} else {
				customErr := errorx.From(err)
				code = customErr.Code
				err = customErr
			}
			zap.L().Error("internal err: ", zap.Error(err),
				zap.Int("Code", code),
				zap.String("BaseURL", c.BaseURL()),
				zap.String("URL", string(c.Context().URI().Path())))
			return c.Status(code).JSON(err)
		},
		// custom  json encode or decode
		JSONEncoder: jsonx.Marshal,
		JSONDecoder: jsonx.Unmarshal,
	}

	app := fiber.New(config)

	// add middleware
	app.Use(recover.New())
	for _, s := range servers {
		s.RegisterHTTPRouter(app)
	}

	return app
}
