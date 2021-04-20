package userService

import (
	"github.com/suthanth/bookstore_users_api/db/repositories"
	"github.com/suthanth/bookstore_users_api/dto/user_dto"
	"github.com/suthanth/bookstore_users_api/mappers/user_mapper"
	"github.com/suthanth/bookstore_users_api/utils/crypto_utils"
	"github.com/suthanth/bookstore_users_api/utils/rest_errors"

	"github.com/suthanth/bookstore_users_api/domain/users"
)

type IUserService interface {
	CreateUser(users.User) (user_dto.User_dto, *rest_errors.RestErr)
	GetUser(int64) (user_dto.User_dto, *rest_errors.RestErr)
}

type UserService struct {
	UserRepository repositories.IUserRepository
	UserMapper     user_mapper.UserMapper
}

func NewUserService(userRepository repositories.IUserRepository,
	userMapper user_mapper.UserMapper) *UserService {
	service := &UserService{
		UserRepository: userRepository,
		UserMapper:     userMapper,
	}
	return service
}

func (u UserService) CreateUser(user users.User) (user_dto.User_dto, *rest_errors.RestErr) {
	var user_dto user_dto.User_dto
	if err := user.Validate(); err != nil {
		return user_dto, err
	}
	existingUser, err := u.UserRepository.FindByUserEmail(user.Email)
	if err != nil {
		return user_dto, err
	}
	if existingUser.Email == user.Email {
		return user_dto, rest_errors.NewFailedToCreateUser("User already exists")
	}
	hashedPwd, err := crypto_utils.EncryptPassword(user.Password)
	if err != nil {
		return user_dto, err
	}
	user.Password = hashedPwd
	// user.CreatedOn = time.Now()
	// user.UpdatedOn = time.Now()
	if err := u.UserRepository.CreateUser(user); err != nil {
		return user_dto, err
	}
	return u.UserMapper.MapUserDomainToDto(user), nil
}

func (u UserService) GetUser(userId int64) (user_dto.User_dto, *rest_errors.RestErr) {
	var user_dto user_dto.User_dto
	user, err := u.UserRepository.FindUserById(userId)
	if err != nil {
		return user_dto, err
	}
	if user.Email == "" {
		return user_dto, rest_errors.NewNotFoundError("User not found")
	}
	return u.UserMapper.MapUserDomainToDto(user), nil
}
