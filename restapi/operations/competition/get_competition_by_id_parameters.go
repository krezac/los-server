// Code generated by go-swagger; DO NOT EDIT.

package competition

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewGetCompetitionByIDParams creates a new GetCompetitionByIDParams object
// no default values defined in spec.
func NewGetCompetitionByIDParams() GetCompetitionByIDParams {

	return GetCompetitionByIDParams{}
}

// GetCompetitionByIDParams contains all the bound params for the get competition by Id operation
// typically these are obtained from a http.Request
//
// swagger:parameters getCompetitionById
type GetCompetitionByIDParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*ID of competition to return
	  Required: true
	  In: path
	*/
	CompetitionID int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetCompetitionByIDParams() beforehand.
func (o *GetCompetitionByIDParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rCompetitionID, rhkCompetitionID, _ := route.Params.GetOK("competitionId")
	if err := o.bindCompetitionID(rCompetitionID, rhkCompetitionID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindCompetitionID binds and validates parameter CompetitionID from path.
func (o *GetCompetitionByIDParams) bindCompetitionID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("competitionId", "path", "int64", raw)
	}
	o.CompetitionID = value

	return nil
}
