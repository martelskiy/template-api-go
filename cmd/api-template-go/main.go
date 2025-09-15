package main

import (
	"context"
	"os"

	"github.com/martelskiy/template-api-go/config"
	"github.com/martelskiy/template-api-go/internal/healthcheck"
	"github.com/martelskiy/template-api-go/internal/shared/api/host"
	"github.com/martelskiy/template-api-go/internal/shared/api/route"
	"github.com/martelskiy/template-api-go/internal/shared/lifecycle"
	"github.com/martelskiy/template-api-go/internal/shared/logger"
	"github.com/martelskiy/template-api-go/internal/shared/parser/json"
)

const appConfigPath = "config/appconfig.json"

// @title           API Swagger
// @version         1.0
func main() {
	context := context.Background()
	log := logger.Get()
	log.Info("initializing configuration")

	parser := json.NewParser[config.AppConfig]()
	configuration, err := parser.Parse(appConfigPath)
	if err != nil {
		lifecycle.StopApplication("error initialization application configurations")
	}
	router := route.NewRouter()
	router.
		WithAPIDocumentation().
		WithRoute(route.NewRoute("/status", healthcheck.GetStatus))

	host := host.New(configuration.Api.Port, router)
	host.Serve()

	if err != nil {
		lifecycle.StopApplication("error running web host")
	}

	log.Infof("web host is running at port: '%s'", configuration.Api.Port)

	lifecycle.ListenForApplicationShutDown(func() {
		defer logger.Dispose()

		log.Info("terminating the web host")
		if err := host.Terminate(context); err != nil {
			log.Error("error terminating the host: '%v'", err)
		}

		log.Info("disposing logger")
	}, make(chan os.Signal, 1))
}
