// Code generated by go-swagger; DO NOT EDIT.

package range_operations

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

// NewGetRangesParams creates a new GetRangesParams object
// with the default values initialized.
func NewGetRangesParams() GetRangesParams {

	var (
		// initialize parameters with default values

		activeOnlyDefault = bool(true)
	)

	return GetRangesParams{
		ActiveOnly: &activeOnlyDefault,
	}
}

// GetRangesParams contains all the bound params for the get ranges operation
// typically these are obtained from a http.Request
//
// swagger:parameters getRanges
type GetRangesParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Return active only ranges
	  In: query
	  Default: true
	*/
	ActiveOnly *bool
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetRangesParams() beforehand.
func (o *GetRangesParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qActiveOnly, qhkActiveOnly, _ := qs.GetOK("activeOnly")
	if err := o.bindActiveOnly(qActiveOnly, qhkActiveOnly, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindActiveOnly binds and validates parameter ActiveOnly from query.
func (o *GetRangesParams) bindActiveOnly(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewGetRangesParams()
		return nil
	}

	value, err := swag.ConvertBool(raw)
	if err != nil {
		return errors.InvalidType("activeOnly", "query", "bool", raw)
	}
	o.ActiveOnly = &value

	return nil
}
