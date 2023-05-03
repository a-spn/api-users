package user_service

import (
	authorization_service "api-users/authorization/service"
	user_dao "api-users/user/dao"
)

type UserService struct {
	DAO                  *user_dao.UserDAO
	AuthorizationService *authorization_service.AuthorizationService
}

func NewUserService(userDAO *user_dao.UserDAO, authorizationService *authorization_service.AuthorizationService) *UserService {
	return &UserService{DAO: userDAO, AuthorizationService: authorizationService}
}
