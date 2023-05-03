package main

import (
	authorization_middleware "api-users/authorization/middleware"
	authorization_service "api-users/authorization/service"
	"api-users/config"
	user_controller "api-users/user/controller"
	user_dao "api-users/user/dao"
	user_model "api-users/user/model"
	user_route "api-users/user/route"
	user_service "api-users/user/service"

	authentication_controller "api-users/authentication/controller"
	authentication_middleware "api-users/authentication/middleware"
	authentication_route "api-users/authentication/route"
	authentication_service "api-users/authentication/service"

	"fmt"
	"os"
	"strconv"

	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
)

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.PathFlag{
				Name:        "config",
				Aliases:     []string{"c"},
				Usage:       "path to YAML config file",
				Value:       "./ressources/config.yml",
				Destination: &config.ConfigPath,
			},
		},
		Action: func(c *cli.Context) error {
			config.ParseConfig()
			config.InitConfig()
			//Init Db
			config.Configuration.MySQL.InitDbConnection()
			config.Db.AutoMigrate(&user_model.User{})
			//Init API
			config.InitPrometheusExporter()
			echoApp := config.Configuration.Api.InitAPI()
			//Init DAO
			userDAO := user_dao.NewUserDAO(config.Db)
			//Init services
			authorizationService := authorization_service.NewAuthorizationService()
			userService := user_service.NewUserService(userDAO, authorizationService)
			authenticationService := authentication_service.NewAuthenticationService(userService)
			//Init controllers
			userController := user_controller.NewUserController(userService)
			authController := authentication_controller.NewAuthenticationController(authenticationService)
			//Init middlewares
			authenticationMiddleware := authentication_middleware.NewAuthenticationMiddlewares(authenticationService)
			authorizationMiddleware := authorization_middleware.NewAuthorizationMiddlewares(authorizationService)
			//Init routes
			userRoutes := user_route.NewUserRoutes(userController, authenticationMiddleware, authorizationMiddleware)
			authRoutes := authentication_route.NewAuthenticationRoutes(authController)
			// Create routes
			userRoutes.CreateRoutes(echoApp)
			authRoutes.CreateRoutes(echoApp)
			//Run
			config.Logger.Info("starting http server", zap.Int("port", config.Configuration.Api.Port))
			err := echoApp.Start(fmt.Sprintf(":%s", strconv.Itoa(config.Configuration.Api.Port)))
			return err
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		config.Logger.Fatal("can't run app", zap.Error(err))
	}
}
