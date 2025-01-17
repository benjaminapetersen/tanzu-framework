// Code generated by go-swagger; DO NOT EDIT.

package azure

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// SetAzureEndpointHandlerFunc turns a function with the right signature into a set azure endpoint handler
type SetAzureEndpointHandlerFunc func(SetAzureEndpointParams) middleware.Responder

// Handle executing the request and returning a response
func (fn SetAzureEndpointHandlerFunc) Handle(params SetAzureEndpointParams) middleware.Responder {
	return fn(params)
}

// SetAzureEndpointHandler interface for that can handle valid set azure endpoint params
type SetAzureEndpointHandler interface {
	Handle(SetAzureEndpointParams) middleware.Responder
}

// NewSetAzureEndpoint creates a new http.Handler for the set azure endpoint operation
func NewSetAzureEndpoint(ctx *middleware.Context, handler SetAzureEndpointHandler) *SetAzureEndpoint {
	return &SetAzureEndpoint{Context: ctx, Handler: handler}
}

/*
SetAzureEndpoint swagger:route POST /api/providers/azure azure setAzureEndpoint

Validate and set azure credentials
*/
type SetAzureEndpoint struct {
	Context *middleware.Context
	Handler SetAzureEndpointHandler
}

func (o *SetAzureEndpoint) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewSetAzureEndpointParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
