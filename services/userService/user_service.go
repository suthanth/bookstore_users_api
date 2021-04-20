package userService

import (
	"fmt"

	"github.com/suthanth/bookstore_users_api/db/repositories"
	"github.com/suthanth/bookstore_users_api/utils/rest_errors"

	"github.com/suthanth/bookstore_users_api/domain/users"
)

type IUserService interface {
	CreateUser(users.User) (users.User, *rest_errors.RestErr)
	GetUser(int64) (users.User, *rest_errors.RestErr)
}

type UserService struct {
	UserRepository repositories.IUserRepository
}

func NewUserService(userRepository repositories.IUserRepository) *UserService {
	service := &UserService{
		UserRepository: userRepository,
	}
	return service
}

func (u UserService) CreateUser(user users.User) (users.User, *rest_errors.RestErr) {
	if err := user.Validate(); err != nil {
		return user, err
	}
	existingUser, err := u.UserRepository.FindByUserEmail(user.Email)
	if err != nil {
		return user, err
	}
	fmt.Println(existingUser)
	if existingUser.Email == user.Email {
		return user, rest_errors.NewFailedToCreateUser("User already exists")
	}
	if err := u.UserRepository.CreateUser(user); err != nil {
		return user, err
	}
	return user, nil
}

func (u UserService) GetUser(userId int64) (users.User, *rest_errors.RestErr) {
	user, err := u.UserRepository.FindUserById(userId)
	if err != nil {
		return user, err
	}
	if user.Email == "" {
		return user, rest_errors.NewNotFoundError("User not found")
	}
	return user, nil
}
