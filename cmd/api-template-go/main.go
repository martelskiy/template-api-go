package main

import (
	"context"

	"github.com/api-template-go/config"
	"github.com/api-template-go/internal/healthcheck"
	"github.com/api-template-go/internal/shared/api/host"
	"github.com/api-template-go/internal/shared/api/route"
	"github.com/api-template-go/internal/shared/lifecycle"
	"github.com/api-template-go/internal/shared/logger"
	"github.com/api-template-go/internal/shared/parser/json"
)

const appConfigPath = "config/appconfig.json"

// @title           API Swagger
// @version         1.0
func main() {
	context := context.Background()
	log := logger.Initialize()
	log.Info("logger initialized")

	parser := json.NewParser[config.AppConfig]()
	configuration, err := parser.Parse(appConfigPath)
	if err != nil {
		lifecycle.StopApplication("error initialization application configurations")
	}

	router := route.NewRouter()
	router.
		WithAPIDocumentation().
		WithRoute(route.NewRoute("/status", healthcheck.GetStatus))

	host := host.NewHost(configuration.Api.Port, router)
	err = host.RunAsync()
	if err != nil {
		lifecycle.StopApplication("error running web host")
	}
	log.Infof("web host is running at port: '%s'", configuration.Api.Port)

	lifecycle.ListenForApplicationShutDown(func() {
		log.Info("terminating the web host")
		host.Terminate(context)
		log.Info("disposing logger")
		logger.Dispose()
	})
}
