package user_controller

import (
	authorization_model "api-users/authorization/model"
	"api-users/config"
	user_dao "api-users/user/dao"
	user_model "api-users/user/model"
	user_service "api-users/user/service"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func (controller UserController) UpdateUser(c echo.Context) error {
	var user user_model.User
	err := json.NewDecoder(c.Request().Body).Decode(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, m{"msg": config.JsonParsingFailMessage})
	}
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		config.Logger.Info("Invalid user input", zap.String("input", c.Param("id")), zap.Error(err))
		return c.JSON(http.StatusBadRequest, m{"msg": config.InvalidInputParameterMessage})
	}
	user.ID = uint(id)
	err = controller.UserService.UpdateUser(user, c.Get(config.RbacAuthorizationContextKey).(authorization_model.AuthorizationContext))
	if err != nil {
		if errors.Is(err, user_service.ErrorPasswordContainSpace) || errors.Is(err, user_service.ErrorPasswordRequirements) || errors.Is(err, user_service.ErrorInvalidRole) || errors.Is(err, user_service.ErrorInvalidRole) || errors.Is(err, user_service.ErrorInvalidEmailAdress) || errors.Is(err, user_service.ErrorInvalidUsername) {
			return c.JSON(http.StatusBadRequest, m{"msg": err})
		} else if errors.Is(err, user_dao.ErrDuplicateKeyEntry) {
			return c.JSON(http.StatusConflict, m{"msg": err})
		} else if errors.Is(err, user_dao.ErrUserDoesNotExist) {
			return c.NoContent(http.StatusNotFound)
		}
		return c.NoContent(http.StatusInternalServerError)
	}
	return c.NoContent(http.StatusOK)
}
