package host

import (
	"context"
	"net/http"

	"github.com/martelskiy/template-api-go/internal/shared/api/route"
	"github.com/martelskiy/template-api-go/internal/shared/logger"
	"go.uber.org/zap"
)

const hostName = "127.0.0.1"

type Host interface {
	RunAsync() error
	Terminate(ctx context.Context) error
}

type WebHost struct {
	router route.Router
	server http.Server
	logger *zap.SugaredLogger
}

func New(port string, router route.Router) *WebHost {
	return &WebHost{
		router: router,
		logger: logger.Get(),
		server: http.Server{
			Addr:    hostName + ":" + port,
			Handler: router.GetRouter(),
		},
	}
}

func (h *WebHost) RunAsync() error {
	go func() {
		if err := h.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			h.logger.With("error", err.Error()).Panic("error running web host")
		}
	}()
	return nil
}

func (h *WebHost) Terminate(ctx context.Context) error {
	return h.server.Shutdown(ctx)
}
