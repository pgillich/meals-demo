// Code generated by go-swagger; DO NOT EDIT.

package meal

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetIngredientsHandlerFunc turns a function with the right signature into a get ingredients handler
type GetIngredientsHandlerFunc func(GetIngredientsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetIngredientsHandlerFunc) Handle(params GetIngredientsParams) middleware.Responder {
	return fn(params)
}

// GetIngredientsHandler interface for that can handle valid get ingredients params
type GetIngredientsHandler interface {
	Handle(GetIngredientsParams) middleware.Responder
}

// NewGetIngredients creates a new http.Handler for the get ingredients operation
func NewGetIngredients(ctx *middleware.Context, handler GetIngredientsHandler) *GetIngredients {
	return &GetIngredients{Context: ctx, Handler: handler}
}

/* GetIngredients swagger:route GET /ingredients meal getIngredients

Get all ingredients

ll ingredients are stored

*/
type GetIngredients struct {
	Context *middleware.Context
	Handler GetIngredientsHandler
}

func (o *GetIngredients) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetIngredientsParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
