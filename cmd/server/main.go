package main

import (
	"github.com/Pedro-phd/gym-routine-golang/internal/database/sqlc"
	database "github.com/Pedro-phd/gym-routine-golang/internal/infra"
	"github.com/Pedro-phd/gym-routine-golang/internal/middleware"
	"github.com/Pedro-phd/gym-routine-golang/internal/modules/user"
	"github.com/Pedro-phd/gym-routine-golang/pkg/log"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// init project
func main() {

	log := log.InitLogger()
	defer log.Sync()

	db := database.InitDatabase(log)
	db.RunMigrations()

	queries := sqlc.New(db.DB)
	server := fiber.New(fiber.Config{
		ErrorHandler: middleware.ErrorMiddleware,
	})

	validate := validator.New()

	// handlers
	user.Setup(server, queries, validate)

	server.Listen(":3000")
}
