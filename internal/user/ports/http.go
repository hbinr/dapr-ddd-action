package ports

import (
	"github.com/gofiber/fiber/v2"

	"github.com/dapr-ddd-action/internal/user/app"
	"github.com/dapr-ddd-action/internal/user/app/command"

	"github.com/dapr-ddd-action/pkg/errorx"
	"github.com/dapr-ddd-action/pkg/httpx"
)

type UserController struct {
	app app.Application
}

func NewUserController(app app.Application) *UserController {
	return &UserController{app: app}
}

func (u *UserController) RegisterHTTPRouter(r *fiber.App) {
	group := r.Group("/user")
	group.Get("/:id", u.GetUser)
	group.Put("/", u.UpdateUser)
}

// GerUser 获取用户信息
func (u UserController) GetUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return errorx.BadRequest(err.Error())
	}

	if id == 0 {
		return errorx.BadRequest("id is zero")
	}

	userDto, err := u.app.Queries.UserInfo.Handler(c.Context(), int64(id))
	if err != nil {
		return err
	}

	// 看前端需要怎样的数据结构，可能需要固定的数据结构，前端取值可能是: res.data.code / res.data.data
	// 如果是我开发前端，那么这里可以直接返回json，即 c.JSON(userDto) . 前端取值是: res.code / res.data
	return c.JSON(httpx.RespSuccess(userDto))
}

func (u UserController) UpdateUser(c *fiber.Ctx) error {
	req := new(UpdateUserReq)

	if err := httpx.ParseAndValidate(c, req); err != nil {
		return errorx.BadRequest(err.Error())
	}

	if err := u.app.Commands.EditUserInfo.Handler(c.Context(), command.EditUserInfoCmd{
		ID:       req.ID,
		UserName: req.UserName,
	}); err != nil {
		return err
	}

	return c.JSON(httpx.RespSuccess(nil))
	// return nil
}
