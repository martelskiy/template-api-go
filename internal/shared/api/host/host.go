package host

import (
	"context"
	"net/http"

	"github.com/api-template-go/internal/shared/api/route"
	"github.com/api-template-go/internal/shared/logger"
	"go.uber.org/zap"
)

const hostName = "127.0.0.1"

type host struct {
	router route.IRouter
	server http.Server
	logger *zap.SugaredLogger
}

func NewHost(port string, router route.IRouter) IHost {
	return &host{
		router: router,
		logger: logger.Get(),
		server: http.Server{
			Addr:    hostName + ":" + port,
			Handler: router.GetRouter(),
		},
	}
}

func (h *host) RunAsync() error {
	go func() {
		if err := h.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			h.logger.With("error", err.Error()).Panic("error running web host")
		}
	}()
	return nil
}

func (h *host) Terminate(ctx context.Context) error {
	return h.server.Shutdown(ctx)
}
