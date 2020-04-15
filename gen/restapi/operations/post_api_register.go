// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PostAPIRegisterHandlerFunc turns a function with the right signature into a post API register handler
type PostAPIRegisterHandlerFunc func(PostAPIRegisterParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostAPIRegisterHandlerFunc) Handle(params PostAPIRegisterParams) middleware.Responder {
	return fn(params)
}

// PostAPIRegisterHandler interface for that can handle valid post API register params
type PostAPIRegisterHandler interface {
	Handle(PostAPIRegisterParams) middleware.Responder
}

// NewPostAPIRegister creates a new http.Handler for the post API register operation
func NewPostAPIRegister(ctx *middleware.Context, handler PostAPIRegisterHandler) *PostAPIRegister {
	return &PostAPIRegister{Context: ctx, Handler: handler}
}

/*PostAPIRegister swagger:route POST /api/register postApiRegister

PostAPIRegister post API register API

*/
type PostAPIRegister struct {
	Context *middleware.Context
	Handler PostAPIRegisterHandler
}

func (o *PostAPIRegister) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPostAPIRegisterParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}