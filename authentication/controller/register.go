package authentication_controller

import (
	"api-users/config"
	user_dao "api-users/user/dao"
	user_model "api-users/user/model"
	user_service "api-users/user/service"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func (controller *AuthenticationController) Register(c echo.Context) error {
	var newUser user_model.User
	err := json.NewDecoder(c.Request().Body).Decode(&newUser)
	if err != nil {
		config.Logger.Info("failed to parse JSON body", zap.Error(err))
		return c.JSON(400, m{"msg": config.JsonParsingFailMessage})
	}
	err = controller.AuthenticationService.Register(newUser)
	if err != nil {
		switch err {
		case user_service.ErrorPasswordContainSpace:
			return c.JSON(http.StatusBadRequest, m{"msg": err.Error()})
		case user_service.ErrorPasswordRequirements:
			return c.JSON(http.StatusBadRequest, m{"msg": err.Error()})
		case user_service.ErrorInvalidRole:
			return c.JSON(http.StatusBadRequest, m{"msg": err.Error()})
		case user_service.ErrorInvalidEmailAdress:
			return c.JSON(http.StatusBadRequest, m{"msg": err.Error()})
		case user_service.ErrorInvalidUsername:
			return c.JSON(http.StatusBadRequest, m{"msg": err.Error()})
		case user_dao.ErrDuplicateKeyEntry:
			return c.JSON(http.StatusConflict, m{"msg": "the fields 'ID','email' and 'username' must be unique"})
		default:
			c.NoContent(http.StatusInternalServerError)
		}
	}
	return c.NoContent(200)
}
