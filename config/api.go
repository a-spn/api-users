package config

import (
	"io/ioutil"

	"github.com/labstack/echo/v4"
)

const (
	DEFAULT_API_OUTPUT_SIZE = 20
	LIMIT_API_OUTPUT_SIZE   = 50
)

const (
	InternalServerErrorMessage         = "Internal server error"
	JsonParsingFailMessage             = "failed to parse JSON input"
	InvalidInputParameterMessage       = "invalid input parameter"
	AuthenticationFailedMessage        = "authentication failed"
	RessourceSuccesfullyCreatedMessage = "ressource succesfully created"
	RessourceSuccesfullyDeletedMessage = "ressource succesfully deleted"
)

type m map[string]interface{}

type ApiConfig struct {
	Port                   int `yaml:"port"`
	PrometheusExporterPort int `yaml:"prometheus_exporter_port"`
}

func (apiConfig *ApiConfig) SetupHeathCheck(e *echo.Echo) {
	e.GET("/status", func(c echo.Context) error { return c.JSON(200, m{"status": "up"}) })
}

func (apiConfig *ApiConfig) InitAPI() (newEchoApp *echo.Echo) {
	echoApp := echo.New()
	echoApp.HideBanner = true
	echoApp.Logger.SetOutput(ioutil.Discard)
	apiConfig.SetupHeathCheck(echoApp)
	LoggerMiddleware(echoApp)
	MetricsMiddleware(echoApp)
	return echoApp
}
