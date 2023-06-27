package authentication_routes

import (
	authentication_controller "api-users/authentication/controller"
	"api-users/config"

	"github.com/labstack/echo/v4"
)

type AuthenticationRoutes struct {
	Controller *authentication_controller.AuthenticationController
}

func NewAuthenticationRoutes(controller *authentication_controller.AuthenticationController) *AuthenticationRoutes {
	return &AuthenticationRoutes{Controller: controller}
}

func (routes *AuthenticationRoutes) CreateRoutes(e *echo.Echo) {
	e.POST("/auth/signin", routes.Controller.SignIn)
	e.POST("/auth/refresh", routes.Controller.Refresh)
	if config.Configuration.Security.EnableLocalRegister {
		e.POST("/auth/register", routes.Controller.Register)
	}
}
