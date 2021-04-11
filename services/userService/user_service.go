package userService

import (
	"github.com/suthanth/bookstore_users_api/utils/rest_errors"

	"github.com/suthanth/bookstore_users_api/domain/users"
)

func CreateUser(user users.User) (*users.User, *rest_errors.RestErr) {
	return &user, nil
}
