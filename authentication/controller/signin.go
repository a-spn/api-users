package authentication_controller

import (
	authentication_model "api-users/authentication/model"
	"api-users/config"
	"encoding/json"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func (controller *AuthenticationController) SignIn(c echo.Context) error {
	var auth authentication_model.AuthLoginPassword
	err := json.NewDecoder(c.Request().Body).Decode(&auth)
	if err != nil {
		config.Logger.Info("failed to parse JSON body", zap.Error(err))
		return c.JSON(400, m{"msg": config.JsonParsingFailMessage})
	}
	if auth.Login == "" || auth.Password == "" {
		config.Logger.Info("missing login or password", zap.Error(err))
		return c.JSON(400, m{"msg": "missing login or password"})
	}
	tokens, authorized, err := controller.AuthenticationService.SignIn(auth)
	if !authorized {
		return c.JSON(401, m{"msg": config.AuthenticationFailedMessage})
	}
	c.Set(config.JwtContextKey, tokens.AccessToken)
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return c.JSON(401, m{"msg": config.AuthenticationFailedMessage})
		default:
			return c.JSON(500, m{"msg": config.InternalServerErrorMessage})
		}

	}

	return c.JSON(200, tokens)
}
