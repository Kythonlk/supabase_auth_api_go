package types

type RefreshTokenRequest struct {
	RefreshToken string `json:"refreshtoken"`
	UserID       string `json:"userid"`
}

type RefreshTokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
