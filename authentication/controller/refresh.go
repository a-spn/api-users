package authentication_controller

import (
	authentication_model "api-users/authentication/model"
	"api-users/config"
	"encoding/json"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func (controller *AuthenticationController) Refresh(c echo.Context) error {
	var auth authentication_model.AuthTokens
	err := json.NewDecoder(c.Request().Body).Decode(&auth)
	if err != nil {
		config.Logger.Info("failed to parse JSON body", zap.Error(err))
		return c.JSON(400, m{"msg": config.JsonParsingFailMessage})
	}
	if auth.RefreshTokenString == "" {
		config.Logger.Info("missing refresh token", zap.Error(err))
		return c.JSON(400, m{"msg": "missing refresh token"})
	}
	tokens, authorized, err := controller.AuthenticationService.Refresh(auth)
	c.Set(config.JwtContextKey, tokens.AccessToken)
	if err != nil {
		return c.JSON(500, m{"msg": config.InternalServerErrorMessage})
	}
	if !authorized {
		return c.JSON(401, m{"msg": config.AuthenticationFailedMessage})
	}
	return c.JSON(200, tokens)
}
