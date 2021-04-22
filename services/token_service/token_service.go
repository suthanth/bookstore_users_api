package token_service

import (
	"time"

	"github.com/form3tech-oss/jwt-go"
	"github.com/suthanth/bookstore_users_api/dto/token_dto"
	"github.com/suthanth/bookstore_users_api/logger"
	"github.com/suthanth/bookstore_users_api/utils/rest_errors"
	"github.com/twinj/uuid"
)

type ITokenService interface {
	CreateToken(uint64) (token_dto.TokenDetailsDto, *rest_errors.RestErr)
}

type TokenService struct{}

func NewTokenService() *TokenService {
	service := &TokenService{}
	return service
}

func (t TokenService) CreateToken(userId uint64) (token_dto.TokenDetailsDto, *rest_errors.RestErr) {
	var err error
	tokenDetails := token_dto.TokenDetailsDto{}
	tokenDetails.AtExpires = time.Now().Add(time.Minute * 15).Unix()
	tokenDetails.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	tokenDetails.AccessUUID = uuid.NewV4().String()
	tokenDetails.RefreshUUID = uuid.NewV4().String()
	tokenDetails.AccessToken, err = createAccessToken(tokenDetails, userId)
	if err != nil {
		logger.SugarLogger.Errorf("Failed to create access token.", err.Error())
		return tokenDetails, rest_errors.NewBadRequest(err.Error())
	}
	tokenDetails.RefreshToken, err = createRefreshToken(tokenDetails, userId)
	if err != nil {
		logger.SugarLogger.Errorf("Failed to create refresh token.", err.Error())
		return tokenDetails, rest_errors.NewBadRequest(err.Error())
	}
	return tokenDetails, nil
}

func createAccessToken(tokenDetails token_dto.TokenDetailsDto, userId uint64) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userId
	atClaims["exp"] = tokenDetails.AtExpires
	atClaims["access_uuid"] = tokenDetails.AccessUUID
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	return at.SignedString([]byte("fff"))
}

func createRefreshToken(tokenDetails token_dto.TokenDetailsDto, userId uint64) (string, error) {
	rtClaims := jwt.MapClaims{}
	rtClaims["authorized"] = true
	rtClaims["user_id"] = userId
	rtClaims["exp"] = tokenDetails.RtExpires
	rtClaims["refresh_uuid"] = tokenDetails.RefreshUUID
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	return rt.SignedString([]byte("refresh"))
}
