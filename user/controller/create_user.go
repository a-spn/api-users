package user_controller

import (
	authorization_model "api-users/authorization/model"
	authorization_service "api-users/authorization/service"
	"api-users/config"
	user_dao "api-users/user/dao"
	user_model "api-users/user/model"
	user_service "api-users/user/service"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (controller UserController) CreateUser(c echo.Context) error {
	var user user_model.User
	err := json.NewDecoder(c.Request().Body).Decode(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, m{"msg": config.JsonParsingFailMessage})
	}
	err = controller.UserService.CreateUser(user, c.Get(config.RbacAuthorizationContextKey).(authorization_model.AuthorizationContext))
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
		case authorization_service.ErrUnauthorizedOperation:
			return c.NoContent(http.StatusForbidden)
		default:
			c.NoContent(http.StatusInternalServerError)
		}
	}
	return c.JSON(http.StatusCreated, m{"msg": config.RessourceSuccesfullyCreatedMessage})
}
