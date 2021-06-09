// THIS FILE IS AUTO GENERATED BY GK-CLI DO NOT EDIT!!
package endpoint

import (
	service "bproject/pkg/service"
	endpoint "github.com/go-kit/kit/endpoint"
)

// Endpoints collects all of the endpoints that compose a profile service. It's
// meant to be used as a helper struct, to collect all of the endpoints into a
// single parameter.
type Endpoints struct {
	AddUserEndpoint endpoint.Endpoint
	GetUserEndpoint endpoint.Endpoint
}

// New returns a Endpoints struct that wraps the provided service, and wires in all of the
// expected endpoint middlewares
func New(s service.BProjectService, mdw map[string][]endpoint.Middleware) Endpoints {
	eps := Endpoints{
		AddUserEndpoint: MakeAddUserEndpoint(s),
		GetUserEndpoint: MakeGetUserEndpoint(s),
	}
	for _, m := range mdw["AddUser"] {
		eps.AddUserEndpoint = m(eps.AddUserEndpoint)
	}
	for _, m := range mdw["GetUser"] {
		eps.GetUserEndpoint = m(eps.GetUserEndpoint)
	}
	return eps
}