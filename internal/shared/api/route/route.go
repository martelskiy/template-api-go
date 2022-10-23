package route

import "net/http"

type route struct {
	name    string
	handler func(responseWriter http.ResponseWriter, request *http.Request)
}

func NewRoute(name string, handler func(responseWriter http.ResponseWriter, request *http.Request)) route {
	return route{
		name:    name,
		handler: handler,
	}
}

func (r *route) GetName() string {
	return r.name
}

func (r *route) GetHandler() func(responseWriter http.ResponseWriter, request *http.Request) {
	return r.handler
}
