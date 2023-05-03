package authentication_service

import (
	authentication_model "api-users/authentication/model"
	"api-users/config"
)

func (service *AuthenticationService) Refresh(auth authentication_model.AuthTokens) (tokens authentication_model.AuthTokens, authorized bool, err error) {
	claims, err := service.DecodeJwt(auth.RefreshTokenString, config.Configuration.JWT.RefreshToken.VerifyKey)
	if err != nil {
		return tokens, false, nil
	}
	user, err := service.UserService.GetUserById(claims.UserID)
	if err != nil {
		return tokens, false, nil
	}
	tokens.AccessTokenString, tokens.AccessToken, err = service.GenerateJwt(user.ID, user.Username, "accessToken", user.Role, config.AccessDuration, config.Configuration.JWT.AccessToken.SignKey)
	if err != nil {
		return tokens, false, err
	}
	return tokens, true, err
}
