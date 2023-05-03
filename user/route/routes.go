package user_route

import (
	authentication_middleware "api-users/authentication/middleware"
	authorization_middleware "api-users/authorization/middleware"
	user_controller "api-users/user/controller"

	"github.com/labstack/echo/v4"
)

type UserRoutes struct {
	UserController            *user_controller.UserController
	AuthenticationMiddlewares *authentication_middleware.AuthenticationMiddlewares
	AuthorizationMiddleware   *authorization_middleware.AuthorizationMiddlewares
}

func NewUserRoutes(controller *user_controller.UserController, authenticationMiddlewares *authentication_middleware.AuthenticationMiddlewares, authorizationMiddleware *authorization_middleware.AuthorizationMiddlewares) *UserRoutes {
	return &UserRoutes{UserController: controller, AuthenticationMiddlewares: authenticationMiddlewares, AuthorizationMiddleware: authorizationMiddleware}
}
func (routes *UserRoutes) CreateRoutes(e *echo.Echo) {
	r := e.Group("/users", routes.AuthenticationMiddlewares.JwtAuthentication, routes.AuthorizationMiddleware.RbacAuthorization) //middlewares.RBAC_Authorization
	r.POST("", routes.UserController.CreateUser)
	r.GET("", routes.UserController.GetUsers)
	r.GET("/:id", routes.UserController.GetUser)
	r.PUT("/:id", routes.UserController.UpdateUser)
	r.DELETE("/:id", routes.UserController.DeleteUser)
}
