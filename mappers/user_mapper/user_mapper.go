package user_mapper

import (
	"github.com/suthanth/bookstore_users_api/domain/users"
	"github.com/suthanth/bookstore_users_api/dto/user_dto"
)

type UserMapper struct{}

func NewUserMapper() *UserMapper {
	mapper := &UserMapper{}
	return mapper
}

func (u UserMapper) MapUserDomainToDto(user users.User) user_dto.User_dto {
	user_dto := user_dto.User_dto{}
	user_dto.FirstName = user.FirstName
	user_dto.LastName = user.LastName
	user_dto.Email = user.Email
	return user_dto
}
