package user

import (
	"context"

	"github.com/Pedro-phd/gym-routine-golang/internal/database/sqlc"
)

type UserRepository struct {
	Queries *sqlc.Queries
}

func NewUserRepository(queries *sqlc.Queries) *UserRepository {
	return &UserRepository{Queries: queries}
}

func (u *UserRepository) CreateUser(ctx context.Context, name, supabase_id string) error {
	user := sqlc.CreateUserParams{
		Name:       name,
		SupabaseID: supabase_id,
	}

	err := u.Queries.CreateUser(ctx, user)
	if err != nil {
		return err
	}
	return nil
}
