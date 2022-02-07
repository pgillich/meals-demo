// Code generated by go-swagger; DO NOT EDIT.

package meal

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// CreateMealHandlerFunc turns a function with the right signature into a create meal handler
type CreateMealHandlerFunc func(CreateMealParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateMealHandlerFunc) Handle(params CreateMealParams) middleware.Responder {
	return fn(params)
}

// CreateMealHandler interface for that can handle valid create meal params
type CreateMealHandler interface {
	Handle(CreateMealParams) middleware.Responder
}

// NewCreateMeal creates a new http.Handler for the create meal operation
func NewCreateMeal(ctx *middleware.Context, handler CreateMealHandler) *CreateMeal {
	return &CreateMeal{Context: ctx, Handler: handler}
}

/* CreateMeal swagger:route POST /meal/{id} meal createMeal

Create a new meal

*/
type CreateMeal struct {
	Context *middleware.Context
	Handler CreateMealHandler
}

func (o *CreateMeal) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewCreateMealParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
