package user

import (
	user "github.com/Pedro-phd/gym-routine-golang/internal/modules/user/dto"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	UserService *UserService
	Server      *fiber.App
	Validator   *validator.Validate
}

func NewUserController(server *fiber.App, userService *UserService, Validator *validator.Validate) *UserController {
	controller := &UserController{UserService: userService, Server: server, Validator: Validator}
	app := controller.Server.Group("/user")

	app.Post("/", controller.CreateUser)

	return controller
}

func (u *UserController) CreateUser(ctx *fiber.Ctx) error {
	body := user.CreateUserRequestDTO{}

	if err := ctx.BodyParser(&body); err != nil {
		return err
	}

	err := u.Validator.Struct(&body)
	if err != nil {
		return err
	}

	err = u.UserService.CreateUser(body.Name, body.SupabaseID)
	if err != nil {
		return err
	}

	return nil
}
