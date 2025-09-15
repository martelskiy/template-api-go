package route

import "net/http"

type Method string

const (
	MethodGET     Method = "GET"
	MethodPOST    Method = "POST"
	MethodPUT     Method = "PUT"
	MethodDELETE  Method = "DELETE"
	MethodOPTIONS Method = "OPTIONS"
)

type Route struct {
	name    string
	method  Method
	handler func(responseWriter http.ResponseWriter, request *http.Request)
}

func NewRoute(name string, method Method, handler func(responseWriter http.ResponseWriter, request *http.Request)) Route {
	return Route{
		name:    name,
		handler: handler,
		method:  method,
	}
}

func (r *Route) Name() string {
	return r.name
}

func (r *Route) Method() Method {
	return r.method
}

func (r *Route) Handler() func(responseWriter http.ResponseWriter, request *http.Request) {
	return r.handler
}
