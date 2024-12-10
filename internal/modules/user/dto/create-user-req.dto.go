package user

import _ "github.com/go-playground/validator/v10"

type CreateUserRequestDTO struct {
	Name       string `json:"name" validate:"required,min=3,max=255"`
	SupabaseID string `json:"supabase_id" validate:"required"`
}
