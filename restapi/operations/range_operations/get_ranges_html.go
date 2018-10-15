// Code generated by go-swagger; DO NOT EDIT.

package range_operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetRangesHTMLHandlerFunc turns a function with the right signature into a get ranges Html handler
type GetRangesHTMLHandlerFunc func(GetRangesHTMLParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetRangesHTMLHandlerFunc) Handle(params GetRangesHTMLParams) middleware.Responder {
	return fn(params)
}

// GetRangesHTMLHandler interface for that can handle valid get ranges Html params
type GetRangesHTMLHandler interface {
	Handle(GetRangesHTMLParams) middleware.Responder
}

// NewGetRangesHTML creates a new http.Handler for the get ranges Html operation
func NewGetRangesHTML(ctx *middleware.Context, handler GetRangesHTMLHandler) *GetRangesHTML {
	return &GetRangesHTML{Context: ctx, Handler: handler}
}

/*GetRangesHTML swagger:route GET /ranges/html range getRangesHtml

List shooting ranges in the system as HTML page

Returns a list of shooting ranges as HTML page

*/
type GetRangesHTML struct {
	Context *middleware.Context
	Handler GetRangesHTMLHandler
}

func (o *GetRangesHTML) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetRangesHTMLParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
