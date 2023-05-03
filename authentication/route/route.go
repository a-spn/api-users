package authentication_routes

import (
	authentication_controller "api-users/authentication/controller"

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
}
