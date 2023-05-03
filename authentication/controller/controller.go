package authentication_controller

import (
	authentication_service "api-users/authentication/service"
)

type m map[string]interface{}

type AuthenticationController struct {
	AuthenticationService *authentication_service.AuthenticationService
}

func NewAuthenticationController(authenticationService *authentication_service.AuthenticationService) *AuthenticationController {
	return &AuthenticationController{AuthenticationService: authenticationService}
}
