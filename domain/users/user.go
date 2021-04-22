package users

import (
	"strings"
	"time"

	"github.com/suthanth/bookstore_users_api/utils/rest_errors"
)

type User struct {
	ID        uint64    `gorm:"primary_key,AUTO_INCREMENT,column:id";`
	FirstName string    `gorm:"column:first_name"`
	LastName  string    `gorm:"column:last_name"`
	Email     string    `gorm:"column:email"`
	Password  string    `gorm:"column:password"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (user *User) Validate() *rest_errors.RestErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return rest_errors.NewBadRequest("Invalid email")
	}
	return nil
}
