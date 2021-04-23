package auth

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/form3tech-oss/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/suthanth/bookstore_users_api/dto/token_dto"
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

func extractToken(bearerToken string) string {
	strArr := strings.Split(bearerToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func ValidateToken(bearerToken string, userId uint64) (token_dto.TokenDetailsDto, error) {
	if strings.HasPrefix(bearerToken, "Bearer") {
		bearerToken = extractToken(bearerToken)
	}
	var tokenDetials token_dto.TokenDetailsDto
	token, err := jwt.Parse(bearerToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Claims.(jwt.MapClaims); !ok {
			return nil, errors.New("invlaid token")
		}
		return []byte("fff"), nil
	})
	if err != nil || !token.Valid {
		return tokenDetials, err
	}
	claimes, ok := token.Claims.(jwt.MapClaims)
	authUserId, err := strconv.ParseUint(fmt.Sprintf("%.f", claimes["user_id"]), 10, 64)
	if err != nil || !ok || userId != authUserId {
		return tokenDetials, errors.New("invalid token")
	}
	accessUUID, _ := claimes["access_uuid"].(string)
	refreshUUID, _ := claimes["refresh_uuid"].(string)
	tokenDetials = token_dto.TokenDetailsDto{
		AccessUUID:  accessUUID,
		RefreshUUID: refreshUUID,
	}
	return tokenDetials, nil
}
