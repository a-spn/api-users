package user_service

import (
	authorization_model "api-users/authorization/model"
	user_model "api-users/user/model"
)

func (service *UserService) CreateUser(user user_model.User, authorizationContext authorization_model.AuthorizationContext) (err error) {
	if err = service.ValidateUser(user, false); err != nil {
		return err
	}
	authorizationContext.ObjectRole = user.Role
	if err = service.AuthorizationService.IsAuthorized(authorizationContext); err != nil {
		return err
	}
	user.PasswordHash, err = service.HashPassword(user.Password)
	if err != nil {
		return err
	}

	if err = service.DAO.CreateUser(user); err != nil {
		return err
	}
	return nil
}
