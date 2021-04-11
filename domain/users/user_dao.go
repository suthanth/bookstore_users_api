package users

import (
	"fmt"

	"github.com/suthanth/bookstore_users_api/utils/rest_errors"
)

var userDb = make(map[int64]*User)

func (user *User) GetUser() *rest_errors.RestErr {
	currentUser := userDb[user.Id]
	if currentUser == nil {
		return rest_errors.NewNotFoundError(fmt.Sprintf("User Id %d not found", user.Id))
	}
	user.Id = currentUser.Id
	user.Email = currentUser.Email
	user.FirstName = currentUser.FirstName
	user.LastName = currentUser.LastName
	user.DateCreated = currentUser.DateCreated
	return nil
}

func (user *User) Save() *rest_errors.RestErr {
	currentUser := userDb[user.Id]
	if currentUser != nil || currentUser.Email == user.Email {
		return rest_errors.NewBadRequest("User already exists")
	}
	userDb[user.Id] = user
	return nil
}
