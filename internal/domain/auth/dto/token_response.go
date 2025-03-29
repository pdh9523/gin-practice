package dto

type TokenResponseDto struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func NewTokenResponseDto(accessToken, refreshToken string) *TokenResponseDto {
	return &TokenResponseDto{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
}
