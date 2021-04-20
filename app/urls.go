package app

import (
	"github.com/gin-gonic/gin"
	"github.com/suthanth/bookstore_users_api/controllers/ping_controller"
	"github.com/suthanth/bookstore_users_api/db/repositories"
	"github.com/suthanth/bookstore_users_api/services/userService"

	"github.com/suthanth/bookstore_users_api/controllers/user_controller"
)

func NewRouter() *gin.Engine {
	router := gin.New()

	pingController := new(ping_controller.PingController)
	router.GET("/api/ping", pingController.Ping())

	userRepository := repositories.NewUserRepository()
	userService := userService.NewUserService(userRepository)
	userController := user_controller.NewUserController(userService)

	users := router.Group("api")
	{
		users.GET("/users/:user_id", userController.GetUser())
		users.GET("/users/search", userController.SearchUser())
		users.POST("/users", userController.CreateUser())
	}
	return router
}
