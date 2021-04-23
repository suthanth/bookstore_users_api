package app

import (
	"github.com/gin-gonic/gin"
	"github.com/suthanth/bookstore_users_api/auth"
	"github.com/suthanth/bookstore_users_api/cache"
	"github.com/suthanth/bookstore_users_api/controllers/ping_controller"
	"github.com/suthanth/bookstore_users_api/db/repositories"
	"github.com/suthanth/bookstore_users_api/mappers/user_mapper"
	"github.com/suthanth/bookstore_users_api/services/cache_service"
	"github.com/suthanth/bookstore_users_api/services/token_service"
	"github.com/suthanth/bookstore_users_api/services/userService"

	"github.com/suthanth/bookstore_users_api/controllers/user_controller"
)

func NewRouter() *gin.Engine {
	router := gin.New()

	pingController := new(ping_controller.PingController)
	router.GET("/api/ping", pingController.Ping())
	userRepository := repositories.NewUserRepository()
	userMapper := user_mapper.NewUserMapper()
	tokenService := token_service.NewTokenService()
	redisClient := cache.GetRedisClient()
	cacheService := cache_service.NewCacheService(redisClient)
	userService := userService.NewUserService(userRepository, *userMapper, tokenService, cacheService)
	userController := user_controller.NewUserController(userService)

	noAuthUsers := router.Group("api")
	{
		noAuthUsers.POST("/users", userController.CreateUser())
		noAuthUsers.POST("/users/login", userController.Login())
	}
	router.Use(auth.AuthMiddleWare())
	authUsers := router.Group("api")
	{
		authUsers.GET("/users/:user_id", userController.GetUser())
		authUsers.GET("/users/search", userController.SearchUser())
	}
	return router
}
