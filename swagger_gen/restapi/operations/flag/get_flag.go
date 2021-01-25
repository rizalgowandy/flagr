// Code generated by go-swagger; DO NOT EDIT.

package flag

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetFlagHandlerFunc turns a function with the right signature into a get flag handler
type GetFlagHandlerFunc func(GetFlagParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetFlagHandlerFunc) Handle(params GetFlagParams) middleware.Responder {
	return fn(params)
}

// GetFlagHandler interface for that can handle valid get flag params
type GetFlagHandler interface {
	Handle(GetFlagParams) middleware.Responder
}

// NewGetFlag creates a new http.Handler for the get flag operation
func NewGetFlag(ctx *middleware.Context, handler GetFlagHandler) *GetFlag {
	return &GetFlag{Context: ctx, Handler: handler}
}

/* GetFlag swagger:route GET /flags/{flagID} flag getFlag

GetFlag get flag API

*/
type GetFlag struct {
	Context *middleware.Context
	Handler GetFlagHandler
}

func (o *GetFlag) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetFlagParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
