package ping_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PingController struct{}

func (p PingController) Ping() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		c.String(http.StatusOK, "Pong")
	}
	return fn
}
