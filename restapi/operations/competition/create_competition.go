// Code generated by go-swagger; DO NOT EDIT.

package competition

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"

	models "github.com/krezac/los-server/models"
)

// CreateCompetitionHandlerFunc turns a function with the right signature into a create competition handler
type CreateCompetitionHandlerFunc func(CreateCompetitionParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateCompetitionHandlerFunc) Handle(params CreateCompetitionParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// CreateCompetitionHandler interface for that can handle valid create competition params
type CreateCompetitionHandler interface {
	Handle(CreateCompetitionParams, *models.Principal) middleware.Responder
}

// NewCreateCompetition creates a new http.Handler for the create competition operation
func NewCreateCompetition(ctx *middleware.Context, handler CreateCompetitionHandler) *CreateCompetition {
	return &CreateCompetition{Context: ctx, Handler: handler}
}

/*CreateCompetition swagger:route POST /ranges/{rangeId}/competition competition createCompetition

Create new competiton

This can only be done by admin. Only id needs to be probided for subobjects (type, category).

*/
type CreateCompetition struct {
	Context *middleware.Context
	Handler CreateCompetitionHandler
}

func (o *CreateCompetition) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCreateCompetitionParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal *models.Principal
	if uprinc != nil {
		principal = uprinc.(*models.Principal) // this is really a models.Principal, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
