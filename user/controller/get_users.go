package user_controller

import (
	authorization_model "api-users/authorization/model"
	"api-users/config"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (controller *UserController) GetUsers(c echo.Context) error {
	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil || limit == -1 {
		limit = config.DEFAULT_API_OUTPUT_SIZE
	}
	if limit > config.LIMIT_API_OUTPUT_SIZE {
		limit = config.LIMIT_API_OUTPUT_SIZE
	}
	offset, err := strconv.Atoi(c.QueryParam("offset"))
	if err != nil {
		offset = 0
	}
	includeDeletedUsers, err := strconv.ParseBool(c.QueryParam("include_deleted"))
	if err != nil {
		includeDeletedUsers = false
	}
	users, err := controller.UserService.GetAPIUsers(limit, offset, includeDeletedUsers, c.Get(config.RbacAuthorizationContextKey).(authorization_model.AuthorizationContext))
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, users)
}
