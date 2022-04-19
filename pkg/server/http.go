package server

import (
	"github.com/dapr-ddd-action/pkg/jsonx"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func NewHttpServer(servers ...Server) *fiber.App {
	config := fiber.Config{
		DisableStartupMessage: true,
		// custom error handler
		// ErrorHandler: func(c *fiber.Ctx, err error) error {
		// 	zap.L().Error("internal err: ", zap.Error(err))
		// 	// fiber error
		// 	fiberNativeErr, ok := err.(*fiber.Error)
		// 	if ok {
		// 		zap.L().Sugar().Infof("template string%+v：", fiberNativeErr)

		// 		return fiberNativeErr
		// 	}
		// 	zap.L().Sugar().Infof("template string%+v：", fiberNativeErr)

		// 	// custom error
		// 	errx := errorx.From(err)
		// 	return c.Status(errx.Code).JSON(errx)
		// },
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
