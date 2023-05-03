package user_service

import (
	authorization_model "api-users/authorization/model"
	user_model "api-users/user/model"
)

func (service *UserService) GetAPIUserById(id uint, authorizationContext authorization_model.AuthorizationContext) (user *user_model.APIUser, err error) {
	user, err = service.DAO.GetAPIUserByIDUnscopped(id)
	if err != nil {
		return &user_model.APIUser{}, err
	}
	authorizationContext.ObjectID = user.ID
	authorizationContext.ObjectRole = user.Role
	return user, service.AuthorizationService.IsAuthorized(authorizationContext)
}

func (service *UserService) GetUserByUsername(username string) (*user_model.User, error) {
	return service.DAO.GetUserByUsername(username)
}

func (service *UserService) GetUserByEmail(email string) (*user_model.User, error) {
	return service.DAO.GetUserByEmail(email)
}

func (service *UserService) GetUserById(id uint) (*user_model.User, error) {
	return service.DAO.GetUserByID(id)
}
