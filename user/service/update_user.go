package user_service

import (
	authorization_model "api-users/authorization/model"
	user_model "api-users/user/model"
)

func (service *UserService) UpdateUser(user user_model.User, authorizationContext authorization_model.AuthorizationContext) (err error) {
	currentUser, err := service.DAO.GetUserByID(user.ID)
	if err != nil {
		return err
	}
	authorizationContext.ObjectID = currentUser.ID
	authorizationContext.ObjectRole = currentUser.Role
	if err = service.AuthorizationService.IsAuthorized(authorizationContext); err != nil { //Check if authorized to modify the user.
		return err
	}
	if err = service.ValidateUser(user, true); err != nil {
		return err
	}
	authorizationContext.ObjectRole = user.Role
	if err = service.AuthorizationService.IsAuthorized(authorizationContext); err != nil { //Check if allowed to give this new role to the user.
		return err
	}
	if user.Password != "" {
		user.PasswordHash, err = service.HashPassword(user.Password)
		if err != nil {
			return err
		}
	}
	return service.DAO.UpdateUser(user)
}
