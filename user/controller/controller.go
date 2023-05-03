package user_controller

import user_service "api-users/user/service"

type m map[string]interface{}

type UserController struct {
	UserService *user_service.UserService
}

func NewUserController(userService *user_service.UserService) *UserController {
	return &UserController{UserService: userService}
}
