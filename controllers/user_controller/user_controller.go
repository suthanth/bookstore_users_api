package user_controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/suthanth/bookstore_users_api/domain/users"

	"github.com/suthanth/bookstore_users_api/services/userService"

	"github.com/suthanth/bookstore_users_api/utils/rest_errors"
)

func CreateUser(c *gin.Context) {
	var user users.User
	//Approach 1
	// bytes, err := ioutil.ReadAll(c.Request.Body)
	// if err != nil {
	// 	//TODO Handle error
	// 	return
	// }
	// if err := json.Unmarshal(bytes, &user); err != nil {
	// 	//TODO Handle json error
	// 	return
	// }

	//Approach 2
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := rest_errors.NewBadRequest("Invalid Request")
		c.JSON(restErr.Status, restErr)
		return
	}
	result, err := userService.CreateUser(user)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		userErr := rest_errors.NewBadRequest("Invalid UserId")
		c.JSON(userErr.Status, userErr)
		return
	}
	user, getErr := userService.GetUser(userId)
	fmt.Println(getErr)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
	}
	c.JSON(http.StatusOK, user)
}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement me")
}
