// THIS FILE IS AUTO GENERATED BY GK-CLI DO NOT EDIT!!
package endpoint

import (
	endpoint "github.com/go-kit/kit/endpoint"
	service "hello/pkg/service"
)

// Endpoints collects all of the endpoints that compose a profile service. It's
// meant to be used as a helper struct, to collect all of the endpoints into a
// single parameter.
type Endpoints struct {
	SayHiEndpoint          endpoint.Endpoint
	MakeADateEndpoint      endpoint.Endpoint
	UpdateUserInfoEndpoint endpoint.Endpoint
}

// New returns a Endpoints struct that wraps the provided service, and wires in all of the
// expected endpoint middlewares
func New(s service.HelloService, mdw map[string][]endpoint.Middleware) Endpoints {
	eps := Endpoints{
		MakeADateEndpoint:      MakeMakeADateEndpoint(s),
		SayHiEndpoint:          MakeSayHiEndpoint(s),
		UpdateUserInfoEndpoint: MakeUpdateUserInfoEndpoint(s),
	}
	for _, m := range mdw["SayHi"] {
		eps.SayHiEndpoint = m(eps.SayHiEndpoint)
	}
	for _, m := range mdw["MakeADate"] {
		eps.MakeADateEndpoint = m(eps.MakeADateEndpoint)
	}
	for _, m := range mdw["UpdateUserInfo"] {
		eps.UpdateUserInfoEndpoint = m(eps.UpdateUserInfoEndpoint)
	}
	return eps
}
