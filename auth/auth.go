package auth

import (
	"net/http"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/form3tech-oss/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/suthanth/bookstore_users_api/logger"
)

var jwtMiddleware *jwtmiddleware.JWTMiddleware

func init() {
	jwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(t *jwt.Token) (interface{}, error) {
			return []byte("fff"), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})
}

func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := jwtMiddleware.CheckJWT(c.Writer, c.Request)
		if err != nil {
			logger.SugarLogger.Errorw("Toke not found in the header")
			c.Abort()
			c.Writer.WriteHeader(http.StatusUnauthorized)
			c.Writer.Write([]byte("Unauthorized"))
			return
		}
	}
}
