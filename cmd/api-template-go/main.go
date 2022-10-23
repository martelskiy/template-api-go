package main

import (
	"context"

	"github.com/api-template-go/internal/healthcheck"
	"github.com/api-template-go/internal/shared/api/host"
	"github.com/api-template-go/internal/shared/api/route"
	"github.com/api-template-go/internal/shared/lifecycle"
	"github.com/api-template-go/internal/shared/logger"
)

// @title           API Swagger
// @version         1.0
func main() {
	context := context.Background()

	log := logger.Initialize()
	log.Info("logger initialized")
	router := route.NewRouter()
	router.
		WithAPIDocumentation().
		WithRoute(route.NewRoute("/status", healthcheck.GetStatus))

	port := "8080"
	host := host.NewHost(port, router)
	err := host.RunAsync()
	if err != nil {
		lifecycle.StopApplication("error running web host")
	}
	log.Infof("web host is running at port: '%s'", port)

	lifecycle.ListenForApplicationShutDown(func() {
		log.Info("terminating the web host")
		host.Terminate(context)
		log.Info("disposing logger")
		logger.Dispose()
	})
}
