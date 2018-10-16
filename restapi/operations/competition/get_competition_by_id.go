// Code generated by go-swagger; DO NOT EDIT.

package competition

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetCompetitionByIDHandlerFunc turns a function with the right signature into a get competition by Id handler
type GetCompetitionByIDHandlerFunc func(GetCompetitionByIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetCompetitionByIDHandlerFunc) Handle(params GetCompetitionByIDParams) middleware.Responder {
	return fn(params)
}

// GetCompetitionByIDHandler interface for that can handle valid get competition by Id params
type GetCompetitionByIDHandler interface {
	Handle(GetCompetitionByIDParams) middleware.Responder
}

// NewGetCompetitionByID creates a new http.Handler for the get competition by Id operation
func NewGetCompetitionByID(ctx *middleware.Context, handler GetCompetitionByIDHandler) *GetCompetitionByID {
	return &GetCompetitionByID{Context: ctx, Handler: handler}
}

/*GetCompetitionByID swagger:route GET /competitions/{competitionId} competition getCompetitionById

Find competition by ID

Returns a single competition

*/
type GetCompetitionByID struct {
	Context *middleware.Context
	Handler GetCompetitionByIDHandler
}

func (o *GetCompetitionByID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetCompetitionByIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}