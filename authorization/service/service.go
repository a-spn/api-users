package authorization_service

import (
	authorization_model "api-users/authorization/model"
	config "api-users/config"
	"errors"

	"go.uber.org/zap"
)

var (
	ErrUnauthorizedOperation = errors.New("unauthorized operation")
)

type AuthorizationService struct {
}

func NewAuthorizationService() *AuthorizationService {
	return &AuthorizationService{}
}

func (service *AuthorizationService) IsAuthorized(authorizationContext authorization_model.AuthorizationContext) (err error) {
	authorized, err := config.Configuration.Rbac.Enforcer.Enforce(authorizationContext.SubjectID, authorizationContext.SubjectRole, authorizationContext.ObjectID, authorizationContext.ObjectRole, authorizationContext.Method)
	if err != nil {
		config.Logger.Error("authorization check failed", zap.Error(err), zap.Uint("subject_id", authorizationContext.SubjectID), zap.Uint("object_id", authorizationContext.ObjectID), zap.String("subject_role", authorizationContext.SubjectRole), zap.String("object_role", authorizationContext.ObjectRole), zap.String("method", authorizationContext.Method))
	} else if !authorized {
		err = ErrUnauthorizedOperation
	}
	return err
}

func (service *AuthorizationService) ListVisibleRoles(authorizationContext authorization_model.AuthorizationContext) (visibleRoles []string, err error) {
	for _, role := range config.Configuration.Rbac.Enforcer.GetAllObjects() {
		authorizationContext.ObjectRole = role
		err = service.IsAuthorized(authorizationContext)
		if err != nil {
			if !errors.Is(err, ErrUnauthorizedOperation) {
				return []string{}, err
			}
			continue
		}
		visibleRoles = append(visibleRoles, role)
	}
	return visibleRoles, nil
}

func (service *AuthorizationService) ListAllRoles() (roles []string) {
	return config.Configuration.Rbac.Enforcer.GetAllObjects()
}
