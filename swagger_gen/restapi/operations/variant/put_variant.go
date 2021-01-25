// Code generated by go-swagger; DO NOT EDIT.

package variant

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PutVariantHandlerFunc turns a function with the right signature into a put variant handler
type PutVariantHandlerFunc func(PutVariantParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PutVariantHandlerFunc) Handle(params PutVariantParams) middleware.Responder {
	return fn(params)
}

// PutVariantHandler interface for that can handle valid put variant params
type PutVariantHandler interface {
	Handle(PutVariantParams) middleware.Responder
}

// NewPutVariant creates a new http.Handler for the put variant operation
func NewPutVariant(ctx *middleware.Context, handler PutVariantHandler) *PutVariant {
	return &PutVariant{Context: ctx, Handler: handler}
}

/* PutVariant swagger:route PUT /flags/{flagID}/variants/{variantID} variant putVariant

PutVariant put variant API

*/
type PutVariant struct {
	Context *middleware.Context
	Handler PutVariantHandler
}

func (o *PutVariant) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPutVariantParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
