package users

import (
	"strings"

	"github.com/suthanth/bookstore_users_api/utils/rest_errors"
)

type User struct {
	Id          int64 `gorm:"primary_key";"AUTO_INCREMENT";`
	FirstName   string
	LastName    string
	Email       string
	DateCreated string
}

func (user *User) Validate() *rest_errors.RestErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return rest_errors.NewBadRequest("Invalid email")
	}
	return nil
}
