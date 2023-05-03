package user_service

import (
	authorization_model "api-users/authorization/model"
	user_model "api-users/user/model"
)

func (service *UserService) GetAPIUsers(limit, offset int, includeDeletedUsers bool, authorizationContext authorization_model.AuthorizationContext) (users []user_model.APIUser, err error) {
	visibleRoles, err := service.AuthorizationService.ListVisibleRoles(authorizationContext)
	if err != nil {
		return []user_model.APIUser{}, err
	}
	if includeDeletedUsers {
		return service.DAO.GetAPIUsersUnscopped(limit, offset, visibleRoles)
	}
	return service.DAO.GetAPIUsers(limit, offset, visibleRoles)
}
