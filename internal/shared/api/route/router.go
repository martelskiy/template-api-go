package route

import (
	_ "github.com/api-template-go/api/docs"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

type router struct {
	muxRouter *mux.Router
}

func NewRouter() IRouter {
	muxRouter := mux.NewRouter().StrictSlash(true)
	return &router{
		muxRouter: muxRouter,
	}
}

func (r *router) WithAPIDocumentation() IRouter {
	r.muxRouter.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
	return r
}

func (r *router) WithRoute(route route) IRouter {
	r.muxRouter.HandleFunc(route.name, route.handler)
	return r
}

func (r *router) GetRouter() *mux.Router {
	return r.muxRouter
}
