package user

import (
	"github.com/Pedro-phd/gym-routine-golang/internal/database/sqlc"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func Setup(server *fiber.App, queries *sqlc.Queries, Validator *validator.Validate) {
	repository := NewUserRepository(queries)
	service := NewUserService(repository)
	NewUserController(server, service, Validator)
}
