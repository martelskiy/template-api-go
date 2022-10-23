package route

import "github.com/gorilla/mux"

type IRouter interface {
	WithAPIDocumentation() IRouter
	WithRoute(route route) IRouter
	GetRouter() *mux.Router
}
