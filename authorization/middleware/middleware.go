package authorization_middleware

import (
	authentication_model "api-users/authentication/model"
	authorization_model "api-users/authorization/model"
	authorization_service "api-users/authorization/service"
	"api-users/config"

	"github.com/labstack/echo/v4"
)

type AuthorizationMiddlewares struct {
	AuthorizationService *authorization_service.AuthorizationService
}

func NewAuthorizationMiddlewares(authorizationService *authorization_service.AuthorizationService) *AuthorizationMiddlewares {
	return &AuthorizationMiddlewares{
		AuthorizationService: authorizationService,
	}
}

func (middleware AuthorizationMiddlewares) RbacAuthorization(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		jwtClaims := c.Get(config.JwtContextKey).(*authentication_model.JwtClaims)
		authorizationContext := authorization_model.AuthorizationContext{
			SubjectID:   jwtClaims.UserID,
			SubjectRole: jwtClaims.Role,
			ObjectID:    0,
			ObjectRole:  "",
			Method:      c.Request().Method,
		}
		c.Set(config.RbacAuthorizationContextKey, authorizationContext)
		return next(c)
	}
}
