package config

import (
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

var (
	PrometheusExporter *prometheus.Prometheus
)

func InitPrometheusExporter() {
	echoPrometheus := echo.New()
	echoPrometheus.HideBanner = true
	echoPrometheus.Logger.SetOutput(ioutil.Discard)
	PrometheusExporter = prometheus.NewPrometheus("echo", nil)
	PrometheusExporter.SetMetricsPath(echoPrometheus)
	go func() {
		Logger.Info("starting prometheus endpoint", zap.Int("port", Configuration.Api.PrometheusExporterPort), zap.String("path", PrometheusExporter.MetricsPath))
		err := echoPrometheus.Start(fmt.Sprintf(":%s", strconv.Itoa(Configuration.Api.PrometheusExporterPort)))
		if err != nil {
			Logger.Fatal("failed to start prometheus endpoint", zap.Int("port", Configuration.Api.PrometheusExporterPort), zap.String("path", PrometheusExporter.MetricsPath), zap.Error(err))
		}
		Logger.Info("prometheus endpoint successfuly started")
	}()
}

func MetricsMiddleware(e *echo.Echo) {
	e.Use(PrometheusExporter.HandlerFunc)
}
