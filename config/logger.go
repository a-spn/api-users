package config

import (
	"log"
	"strings"

	auth_model "api-users/authentication/model"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

var (
	Logger *zap.Logger
)

func InitLogger() {
	var err error
	loggerConfig := zap.NewProductionConfig()
	loggerConfig.DisableStacktrace = true
	loggerConfig.DisableCaller = false
	Logger, err = loggerConfig.Build()
	if err != nil {
		log.Fatalf("failed to initialize zap logger: %v", err)
	}
	defer Logger.Sync()
	Logger.Info("logger successfuly started")
}

func LoggerMiddleware(e *echo.Echo) {
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		Skipper: func(c echo.Context) bool {
			if strings.HasPrefix(c.Path(), "/metrics") || strings.HasPrefix(c.Path(), "/status") {
				return true
			}
			return false
		},
		LogRemoteIP:  true,
		LogRoutePath: true,
		LogURI:       true,
		LogMethod:    true,
		LogStatus:    true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			jwtContent := c.Get(JwtContextKey)
			if jwtContent != nil {
				jwtParsed := jwtContent.(*auth_model.JwtClaims)
				Logger.Info("request",
					zap.String("remote_ip", v.RemoteIP),
					zap.String("method", v.Method),
					zap.String("log_route_path", v.RoutePath),
					zap.Int("status", v.Status),
					zap.String("url", v.URI),
					zap.Uint("user_id", jwtParsed.UserID),
					zap.String("user_role", jwtParsed.Role),
				)
				return nil
			}
			Logger.Info("request",
				zap.String("remote_ip", v.RemoteIP),
				zap.String("method", v.Method),
				zap.String("log_route_path", v.RoutePath),
				zap.Int("status", v.Status),
				zap.String("url", v.URI),
			)
			return nil
		},
	}))
}
