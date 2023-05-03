package user_service

import authorization_model "api-users/authorization/model"

func (service *UserService) DeleteUser(id uint, authorizationContext authorization_model.AuthorizationContext) (err error) {
	user, err := service.DAO.GetAPIUserByIDUnscopped(id)
	if err != nil || user.DeletedAt.Valid {
		return err
	}
	authorizationContext.ObjectID = user.ID
	authorizationContext.ObjectRole = user.Role
	if err = service.AuthorizationService.IsAuthorized(authorizationContext); err != nil {
		return err
	}
	return service.DAO.DeleteUser(id)
}
