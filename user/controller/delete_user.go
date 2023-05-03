package user_controller

import (
	authorization_model "api-users/authorization/model"
	"api-users/config"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func (controller UserController) DeleteUser(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		config.Logger.Info("Invalid user input", zap.String("input", c.Param("id")), zap.Error(err))
		return c.JSON(http.StatusBadRequest, m{"msg": config.InvalidInputParameterMessage})
	}
	err = controller.UserService.DeleteUser(uint(id), c.Get(config.RbacAuthorizationContextKey).(authorization_model.AuthorizationContext))
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return c.NoContent(http.StatusNotFound)
		default:
			return c.NoContent(http.StatusInternalServerError)
		}

	}
	return c.NoContent(http.StatusNoContent)
}
