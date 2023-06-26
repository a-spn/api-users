package authentication_service

import (
	"api-users/config"
	user_model "api-users/user/model"
)

func (service *AuthenticationService) Register(user user_model.User) error {
	user.Role = config.Configuration.Security.AttributedRoleOnRegister
	return service.UserService.CreateUser(user)
}
