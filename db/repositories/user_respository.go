package repositories

import (
	"github.com/suthanth/bookstore_users_api/db"
	"github.com/suthanth/bookstore_users_api/domain/users"
	"github.com/suthanth/bookstore_users_api/utils/rest_errors"
)

type IUserRepository interface {
	FindUserById(uint64) (users.User, *rest_errors.RestErr)
	CreateUser(users.User) *rest_errors.RestErr
	FindByUserEmail(string) (users.User, *rest_errors.RestErr)
}

type UserRepository struct {
	DbService db.UserDbService
}

func NewUserRepository() *UserRepository {
	repo := &UserRepository{
		DbService: db.UserDbService{},
	}
	return repo
}

func (u UserRepository) CreateUser(user users.User) *rest_errors.RestErr {
	db := u.DbService.GetDb()
	if db == nil {
		return rest_errors.NewDBConnectError("Failed to get DB connection")
	}
	if err := db.Create(&user).Error; err != nil {
		return rest_errors.NewFailedToCreateUser(err.Error())
	}
	return nil
}

func (u UserRepository) FindUserById(id uint64) (users.User, *rest_errors.RestErr) {
	db := u.DbService.GetDb()
	var user users.User
	if db == nil {
		return user, rest_errors.NewDBConnectError("Failed to get DB connection")
	}
	db.Where("id = ?", id).Find(&user)
	return user, nil
}

func (u UserRepository) FindByUserEmail(email string) (users.User, *rest_errors.RestErr) {
	db := u.DbService.GetDb()
	var user users.User
	if db == nil {
		return user, rest_errors.NewDBConnectError("Failed to get DB connection")
	}
	db.Where("email = ?", email).Find(&user)
	return user, nil
}
