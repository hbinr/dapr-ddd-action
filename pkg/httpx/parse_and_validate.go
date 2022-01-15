package httpx

import (
	vd "github.com/bytedance/go-tagexpr/v2/validator"
	"github.com/dapr-ddd-action/pkg/errorx"
	"github.com/gofiber/fiber/v2"
)

func BodyParseAndValidate(c *fiber.Ctx, recvPointer interface{}) error {
	if err := c.BodyParser(recvPointer); err != nil {
		return errorx.Internal(err, "request body parse failed")
	}

	return vd.Validate(recvPointer)
}

func QueryParseAndValidate(c *fiber.Ctx, recvPointer interface{}) error {
	if err := c.QueryParser(recvPointer); err != nil {
		return errorx.Internal(err, "request query parse failed")
	}

	return vd.Validate(recvPointer)
}
