// Code generated by go-swagger; DO NOT EDIT.

package competition

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewGetCompetitionsHTMLParams creates a new GetCompetitionsHTMLParams object
// with the default values initialized.
func NewGetCompetitionsHTMLParams() GetCompetitionsHTMLParams {

	var (
		// initialize parameters with default values

		activeOnlyDefault = bool(true)
	)

	return GetCompetitionsHTMLParams{
		ActiveOnly: &activeOnlyDefault,
	}
}

// GetCompetitionsHTMLParams contains all the bound params for the get competitions Html operation
// typically these are obtained from a http.Request
//
// swagger:parameters getCompetitionsHtml
type GetCompetitionsHTMLParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Return active only competitions
	  In: query
	  Default: true
	*/
	ActiveOnly *bool
	/*ID of shooting range to read the competitions for
	  Required: true
	  In: path
	*/
	RangeID int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetCompetitionsHTMLParams() beforehand.
func (o *GetCompetitionsHTMLParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qActiveOnly, qhkActiveOnly, _ := qs.GetOK("activeOnly")
	if err := o.bindActiveOnly(qActiveOnly, qhkActiveOnly, route.Formats); err != nil {
		res = append(res, err)
	}

	rRangeID, rhkRangeID, _ := route.Params.GetOK("rangeId")
	if err := o.bindRangeID(rRangeID, rhkRangeID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindActiveOnly binds and validates parameter ActiveOnly from query.
func (o *GetCompetitionsHTMLParams) bindActiveOnly(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewGetCompetitionsHTMLParams()
		return nil
	}

	value, err := swag.ConvertBool(raw)
	if err != nil {
		return errors.InvalidType("activeOnly", "query", "bool", raw)
	}
	o.ActiveOnly = &value

	return nil
}

// bindRangeID binds and validates parameter RangeID from path.
func (o *GetCompetitionsHTMLParams) bindRangeID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("rangeId", "path", "int64", raw)
	}
	o.RangeID = value

	return nil
}
