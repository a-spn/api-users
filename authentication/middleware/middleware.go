package authentication_middleware

import (
	authentication_service "api-users/authentication/service"
	"api-users/config"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type m map[string]interface{}

type AuthenticationMiddlewares struct {
	AuthService *authentication_service.AuthenticationService
}

func NewAuthenticationMiddlewares(authService *authentication_service.AuthenticationService) *AuthenticationMiddlewares {
	return &AuthenticationMiddlewares{AuthService: authService}
}

func (middleware AuthenticationMiddlewares) JwtAuthentication(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		AuthorizationHeader := c.Request().Header["Authorization"]
		if len(AuthorizationHeader) == 0 {
			return c.JSON(http.StatusBadRequest, m{"msg": "missing jwt authentication token", "tip": "Provide it with the 'Authorization' Header, in the format 'Bearer <your_token>'"})
		}
		parsedHeader := strings.Split(AuthorizationHeader[0], " ")
		if parsedHeader[0] != "Bearer" {
			return c.JSON(http.StatusBadRequest, m{"msg": "missing jwt authentication token", "tip": "Provide it with the 'Authorization' Header, in the format 'Bearer <your_token>'"})

		}
		claims, err := middleware.AuthService.DecodeJwt(parsedHeader[1], &config.Configuration.JWT.AccessToken.SignKey.PublicKey)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, m{"msg": err.Error()})
		}
		c.Set(config.JwtContextKey, claims)
		return next(c)
	}
}
