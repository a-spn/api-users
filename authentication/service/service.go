package authentication_service

import (
	user_service "api-users/user/service"
)

type AuthenticationService struct {
	UserService *user_service.UserService
}

func NewAuthenticationService(userService *user_service.UserService) *AuthenticationService {
	return &AuthenticationService{UserService: userService}
}
