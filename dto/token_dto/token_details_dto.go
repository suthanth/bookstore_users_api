package token_dto

type TokenDetailsDto struct {
	AccessToken  string
	RefreshToken string
	AtExpires    int64
	RtExpires    int64
	AccessUUID   string
	RefreshUUID  string
}
