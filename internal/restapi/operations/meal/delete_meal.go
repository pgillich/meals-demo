// Code generated by go-swagger; DO NOT EDIT.

package meal

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/pgillich/meals-demo/internal/models"
)

// DeleteMealHandlerFunc turns a function with the right signature into a delete meal handler
type DeleteMealHandlerFunc func(DeleteMealParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteMealHandlerFunc) Handle(params DeleteMealParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// DeleteMealHandler interface for that can handle valid delete meal params
type DeleteMealHandler interface {
	Handle(DeleteMealParams, *models.User) middleware.Responder
}

// NewDeleteMeal creates a new http.Handler for the delete meal operation
func NewDeleteMeal(ctx *middleware.Context, handler DeleteMealHandler) *DeleteMeal {
	return &DeleteMeal{Context: ctx, Handler: handler}
}

/* DeleteMeal swagger:route DELETE /meal/{id} meal deleteMeal

Deletes a meal

*/
type DeleteMeal struct {
	Context *middleware.Context
	Handler DeleteMealHandler
}

func (o *DeleteMeal) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDeleteMealParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal *models.User
	if uprinc != nil {
		principal = uprinc.(*models.User) // this is really a models.User, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
