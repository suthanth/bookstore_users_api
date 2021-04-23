package userService

import (
	"github.com/suthanth/bookstore_users_api/db/repositories"
	"github.com/suthanth/bookstore_users_api/dto/token_dto"
	"github.com/suthanth/bookstore_users_api/dto/user_dto"
	"github.com/suthanth/bookstore_users_api/mappers/user_mapper"
	"github.com/suthanth/bookstore_users_api/services/cache_service"
	"github.com/suthanth/bookstore_users_api/services/token_service"
	"github.com/suthanth/bookstore_users_api/utils/crypto_utils"
	"github.com/suthanth/bookstore_users_api/utils/rest_errors"

	"github.com/suthanth/bookstore_users_api/domain/users"
)

type IUserService interface {
	CreateUser(users.User) (user_dto.UserDto, *rest_errors.RestErr)
	GetUser(uint64) (user_dto.UserDto, *rest_errors.RestErr)
	Login(user_dto.UserLoginDto) (token_dto.TokenDetailsDto, *rest_errors.RestErr)
	ValidateTokenUUIDWithCache(string, uint64) *rest_errors.RestErr
}

type UserService struct {
	UserRepository repositories.IUserRepository
	UserMapper     user_mapper.UserMapper
	TokenService   token_service.ITokenService
	CacheService   cache_service.ICacheService
}

func NewUserService(userRepository repositories.IUserRepository,
	userMapper user_mapper.UserMapper, tokenService token_service.ITokenService,
	cacheService cache_service.ICacheService) *UserService {
	service := &UserService{
		UserRepository: userRepository,
		UserMapper:     userMapper,
		TokenService:   tokenService,
		CacheService:   cacheService,
	}
	return service
}

func (u UserService) CreateUser(user users.User) (user_dto.UserDto, *rest_errors.RestErr) {
	var user_dto user_dto.UserDto
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
	if err := u.UserRepository.CreateUser(user); err != nil {
		return user_dto, err
	}
	return u.UserMapper.MapUserDomainToDto(user), nil
}

func (u UserService) GetUser(userId uint64) (user_dto.UserDto, *rest_errors.RestErr) {
	var user_dto user_dto.UserDto
	user, err := u.UserRepository.FindUserById(userId)
	if err != nil {
		return user_dto, err
	}
	if user.Email == "" {
		return user_dto, rest_errors.NewNotFoundError("User not found")
	}
	return u.UserMapper.MapUserDomainToDto(user), nil
}

func (u UserService) Login(loginDto user_dto.UserLoginDto) (token_dto.TokenDetailsDto, *rest_errors.RestErr) {
	var tokenDetails token_dto.TokenDetailsDto
	if loginDto.Email == "" || loginDto.Password == "" {
		return tokenDetails, rest_errors.NewBadRequest("Email/Password cannot be empty")
	}
	existingUser, err := u.UserRepository.FindByUserEmail(loginDto.Email)
	if err != nil {
		return tokenDetails, err
	}

	if !crypto_utils.ComparePassword(existingUser.Password, loginDto.Password) {
		return tokenDetails, rest_errors.NewUnAuthorizedError("Invalid credentials")
	}
	tokenDetails, err = u.TokenService.CreateToken(existingUser.ID)
	if err != nil {
		return tokenDetails, err
	}
	cacheErr := u.CacheService.SaveTokenDetails(tokenDetails, existingUser.ID)
	if err != nil {
		return tokenDetails, rest_errors.NewBadRequest(cacheErr.Error())
	}
	return tokenDetails, nil
}

func (u UserService) ValidateTokenUUIDWithCache(tokenUUID string, userId uint64) *rest_errors.RestErr {
	val, err := u.CacheService.FetchTokenDetails(tokenUUID)
	if err != nil {
		return rest_errors.NewBadRequest("Failed to fetch userID from cache")
	}
	if val != userId {
		return rest_errors.NewBadRequest("UserID not matched with cached Id")
	}
	return nil
}
