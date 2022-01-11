package ports

import (
	"github.com/gofiber/fiber/v2"

	"github.com/dapr-ddd-action/internal/user/app"
	"github.com/dapr-ddd-action/internal/user/app/command"
	"github.com/dapr-ddd-action/internal/user/app/query"

	"github.com/dapr-ddd-action/pkg/errorx"
	"github.com/dapr-ddd-action/pkg/httpx"
)

// 入参 需要转化为 Command, Query，然后给到App层
// 出参 dto
type UserController struct {
	app app.Application
}

func NewUserController(app app.Application) *UserController {
	return &UserController{app: app}
}

func (u *UserController) RegisterHTTPRouter(r *fiber.App) {
	group := r.Group("/user")
	group.Get("/:id/info", u.GetUser)
	group.Get("/list", u.GetUser)
	group.Put("/", u.UpdateUser)
}

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

	return c.JSON(httpx.RespSuccess(userDto))
}

func (u UserController) ListUser(c *fiber.Ctx) error {
	req := new(query.UsersPageQuery)
	if err := httpx.ParseAndValidate(c, req); err != nil {
		return errorx.BadRequest(err.Error())
	}

	userDto, err := u.app.Queries.UsersPage.Handler(c.Context(), req)
	if err != nil {
		return err
	}

	return c.JSON(httpx.RespSuccess(userDto))
}
func (u UserController) UpdateUser(c *fiber.Ctx) error {
	req := new(command.EditUserInfoCmd)
	if err := httpx.ParseAndValidate(c, req); err != nil {
		return errorx.BadRequest(err.Error())
	}

	if err := u.app.Commands.EditUserInfo.Handler(c.Context(), req); err != nil {
		return err
	}

	return c.JSON(httpx.RespSuccess(nil))
}
