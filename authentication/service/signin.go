package authentication_service

import (
	authentication_model "api-users/authentication/model"
	"api-users/config"
	user_model "api-users/user/model"
	"strings"
)

func (service *AuthenticationService) SignIn(auth authentication_model.AuthLoginPassword) (tokens authentication_model.AuthTokens, authorized bool, err error) {
	var user *user_model.User
	if config.Configuration.Security.SuperUserIsEnabled && auth.Login == config.Configuration.Security.SuperUserLogin {
		user, authorized = service.SignInAsSuperUser(auth)
	} else {
		user, authorized, err = service.SignInAsNormalUser(auth)
	}
	if !authorized || err != nil {
		return tokens, false, err
	}
	tokens.RefreshTokenString, tokens.RefreshToken, err = service.GenerateJwt(user.ID, user.Username, "refreshToken", user.Role, config.RefreshDuration, config.Configuration.JWT.RefreshToken.SignKey)
	if err != nil {
		return tokens, false, err
	}
	tokens.AccessTokenString, tokens.AccessToken, err = service.GenerateJwt(user.ID, user.Username, "accessToken", user.Role, config.AccessDuration, config.Configuration.JWT.AccessToken.SignKey)
	if err != nil {
		return tokens, false, err
	}
	return tokens, true, nil
}

func (service *AuthenticationService) SignInAsSuperUser(auth authentication_model.AuthLoginPassword) (user *user_model.User, authorized bool) {
	user = &user_model.User{Username: config.Configuration.Security.SuperUserLogin, Role: "superuser"}
	user.ID = 0
	return user, auth.Password == config.Configuration.Security.SuperUserPass
}

func (service *AuthenticationService) SignInAsNormalUser(auth authentication_model.AuthLoginPassword) (user *user_model.User, authorized bool, err error) {
	if strings.Contains(auth.Login, "@") {
		user, err = service.UserService.GetUserByEmail(auth.Login)
	} else {
		user, err = service.UserService.GetUserByUsername(auth.Login)
	}
	if err != nil {
		return user, false, err
	}
	return user, service.VerifyPassword(user.PasswordHash, auth.Password), nil
}
