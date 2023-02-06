package route

import "net/http"

type Route struct {
	name    string
	handler func(responseWriter http.ResponseWriter, request *http.Request)
}

func NewRoute(name string, handler func(responseWriter http.ResponseWriter, request *http.Request)) Route {
	return Route{
		name:    name,
		handler: handler,
	}
}

func (r *Route) Name() string {
	return r.name
}

func (r *Route) Handler() func(responseWriter http.ResponseWriter, request *http.Request) {
	return r.handler
}
