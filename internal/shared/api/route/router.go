package route

import (
	"github.com/gorilla/mux"
	_ "github.com/martelskiy/template-api-go/api/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Router interface {
	WithAPIDocumentation() Router
	WithRoute(route Route) Router
	GetRouter() *mux.Router
}

type WebRouter struct {
	muxRouter *mux.Router
}

func NewRouter() *WebRouter {
	muxRouter := mux.NewRouter().StrictSlash(true)
	return &WebRouter{
		muxRouter: muxRouter,
	}
}

func (r *WebRouter) WithAPIDocumentation() Router {
	r.muxRouter.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
	return r
}

func (r *WebRouter) WithRoute(route Route) Router {
	r.muxRouter.HandleFunc(route.name, route.handler)
	return r
}

func (r *WebRouter) GetRouter() *mux.Router {
	return r.muxRouter
}
