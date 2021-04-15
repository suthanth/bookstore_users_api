package users

import (
	"fmt"

	"github.com/suthanth/bookstore_users_api/utils/rest_errors"

	"github.com/suthanth/bookstore_users_api/utils/date_utils"

	"github.com/suthanth/bookstore_users_api/db"
)

var UserDbService = db.UserDbService{}.GetDb()

func (user *User) GetUser() *rest_errors.RestErr {
	var currentUser = &User{}
	UserDbService.Where("id = ?", user.Id).First(&currentUser)
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
	var currentUser = &User{}
	UserDbService.Where("email = ?", user.Email).First(&currentUser)
	fmt.Println(currentUser)
	if currentUser != nil {
		if currentUser.Email == user.Email {
			return rest_errors.NewBadRequest("User already exists")
		}
	}
	user.DateCreated = date_utils.GetDateNowString()
	UserDbService.Create(&user)
	return nil
}
