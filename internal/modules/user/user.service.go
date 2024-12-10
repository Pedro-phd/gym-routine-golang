package user

import "context"

type UserService struct {
	UserRepository *UserRepository
}

func NewUserService(userRepository *UserRepository) *UserService {
	return &UserService{UserRepository: userRepository}
}

func (u *UserService) CreateUser(name, supabase_id string) error {
	err := u.UserRepository.CreateUser(context.Background(), name, supabase_id)
	if err != nil {
		return err
	}
	return nil
}
