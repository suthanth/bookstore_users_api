package token_service

import (
	"time"

	"github.com/form3tech-oss/jwt-go"
	"github.com/suthanth/bookstore_users_api/dto/token_dto"
	"github.com/suthanth/bookstore_users_api/utils/rest_errors"
)

type ITokenService interface {
	CreateToken(uint) (token_dto.TokenDetailsDto, *rest_errors.RestErr)
}

type TokenService struct{}

func NewTokenService() *TokenService {
	service := &TokenService{}
	return service
}

func (t TokenService) CreateToken(userId uint) (token_dto.TokenDetailsDto, *rest_errors.RestErr) {
	var tokenDetails token_dto.TokenDetailsDto
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = userId
	claims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := at.SignedString([]byte("fff"))
	if err != nil {
		return tokenDetails, rest_errors.NewBadRequest("Invalid credentials")
	}
	tokenDetails = token_dto.TokenDetailsDto{
		AccessToken: token,
	}
	return tokenDetails, nil
}
