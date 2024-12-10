package middleware

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type ErrorResponse struct {
	Error      bool              `json:"error"`
	Message    string            `json:"message"`
	Validation []ValidationError `json:"validation,omitempty"`
}

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func getValidationErrors(err error) []ValidationError {
	var errors []ValidationError
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, e := range validationErrors {
			errors = append(errors, ValidationError{Field: e.Field(), Message: e.Tag()})
		}
	}
	return errors
}

func ErrorMiddleware(ctx *fiber.Ctx, err error) error {
	if validationError, ok := err.(validator.ValidationErrors); ok {
		return ctx.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Error:      true,
			Message:    "Validation error",
			Validation: getValidationErrors(validationError),
		})
	}

	return ctx.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{
		Error:   true,
		Message: err.Error(),
	})
}
