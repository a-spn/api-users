package user_controller

import (
	authentication_model "api-users/authentication/model"
	"api-users/config"

	"github.com/labstack/echo/v4"
)

func (controller UserController) Whoami(c echo.Context) error {
	jwtClaims := c.Get(config.JwtContextKey).(*authentication_model.JwtClaims)
	return c.JSON(200, jwtClaims)
}
