package app

import (
	"github.com/suthanth/bookstore_users_api/controllers/ping_controller"

	"github.com/suthanth/bookstore_users_api/controllers/user_controller"
)

func mapUrls() {
	router.GET("/ping", ping_controller.Ping)

	router.GET("/users/:user_id", user_controller.GetUser)
	router.GET("/users/search", user_controller.SearchUser)
	router.POST("/users", user_controller.CreateUser)
}
