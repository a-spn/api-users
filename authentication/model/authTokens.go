package authentication_model

type AuthTokens struct {
	AccessTokenString  string     `json:"access_token,omitempty"`
	AccessToken        *JwtClaims `json:"-"`
	RefreshTokenString string     `json:"refresh_token,omitempty"`
	RefreshToken       *JwtClaims `json:"-"`
}
