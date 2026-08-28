package users

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)


type UserService struct {
	userRepository *UserRepository
}

func NewUserService(ur *UserRepository) *UserService{
	return &UserService{
		userRepository: ur,
	}
}


func (us *UserService) CreateNewUser(ctx context.Context, name string) (*User, error) {
	// Do all the business logic here

	if len(name) < 3 {
		return  nil, fmt.Errorf("Name must contain 3 characters")
	}

	id := uuid.New()

	user := User{
		ID: id,
		Name: name,
	}

	return us.userRepository.CreateNewUser(ctx, user)
}


func (us *UserService) GetUsers(ctx context.Context) (*[]User, error) {
	return us.userRepository.GetUsers(ctx)
}