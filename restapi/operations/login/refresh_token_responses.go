// Code generated by go-swagger; DO NOT EDIT.

package login

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	models "github.com/krezac/los-server/models"
)

// RefreshTokenOKCode is the HTTP code returned for type RefreshTokenOK
const RefreshTokenOKCode int = 200

/*RefreshTokenOK successful operation

swagger:response refreshTokenOK
*/
type RefreshTokenOK struct {
	/*date in UTC when token expires

	 */
	XExpiresAfter strfmt.DateTime `json:"X-Expires-After"`
	/*calls per hour allowed by the user

	 */
	XRateLimit int32 `json:"X-Rate-Limit"`

	/*
	  In: Body
	*/
	Payload *models.LoginResponse `json:"body,omitempty"`
}

// NewRefreshTokenOK creates RefreshTokenOK with default headers values
func NewRefreshTokenOK() *RefreshTokenOK {

	return &RefreshTokenOK{}
}

// WithXExpiresAfter adds the xExpiresAfter to the refresh token o k response
func (o *RefreshTokenOK) WithXExpiresAfter(xExpiresAfter strfmt.DateTime) *RefreshTokenOK {
	o.XExpiresAfter = xExpiresAfter
	return o
}

// SetXExpiresAfter sets the xExpiresAfter to the refresh token o k response
func (o *RefreshTokenOK) SetXExpiresAfter(xExpiresAfter strfmt.DateTime) {
	o.XExpiresAfter = xExpiresAfter
}

// WithXRateLimit adds the xRateLimit to the refresh token o k response
func (o *RefreshTokenOK) WithXRateLimit(xRateLimit int32) *RefreshTokenOK {
	o.XRateLimit = xRateLimit
	return o
}

// SetXRateLimit sets the xRateLimit to the refresh token o k response
func (o *RefreshTokenOK) SetXRateLimit(xRateLimit int32) {
	o.XRateLimit = xRateLimit
}

// WithPayload adds the payload to the refresh token o k response
func (o *RefreshTokenOK) WithPayload(payload *models.LoginResponse) *RefreshTokenOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the refresh token o k response
func (o *RefreshTokenOK) SetPayload(payload *models.LoginResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RefreshTokenOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header X-Expires-After

	xExpiresAfter := o.XExpiresAfter.String()
	if xExpiresAfter != "" {
		rw.Header().Set("X-Expires-After", xExpiresAfter)
	}

	// response header X-Rate-Limit

	xRateLimit := swag.FormatInt32(o.XRateLimit)
	if xRateLimit != "" {
		rw.Header().Set("X-Rate-Limit", xRateLimit)
	}

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RefreshTokenUnauthorizedCode is the HTTP code returned for type RefreshTokenUnauthorized
const RefreshTokenUnauthorizedCode int = 401

/*RefreshTokenUnauthorized Invalid token provided

swagger:response refreshTokenUnauthorized
*/
type RefreshTokenUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.APIResponse `json:"body,omitempty"`
}

// NewRefreshTokenUnauthorized creates RefreshTokenUnauthorized with default headers values
func NewRefreshTokenUnauthorized() *RefreshTokenUnauthorized {

	return &RefreshTokenUnauthorized{}
}

// WithPayload adds the payload to the refresh token unauthorized response
func (o *RefreshTokenUnauthorized) WithPayload(payload *models.APIResponse) *RefreshTokenUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the refresh token unauthorized response
func (o *RefreshTokenUnauthorized) SetPayload(payload *models.APIResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RefreshTokenUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RefreshTokenInternalServerErrorCode is the HTTP code returned for type RefreshTokenInternalServerError
const RefreshTokenInternalServerErrorCode int = 500

/*RefreshTokenInternalServerError Invalidating of old token failed

swagger:response refreshTokenInternalServerError
*/
type RefreshTokenInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.APIResponse `json:"body,omitempty"`
}

// NewRefreshTokenInternalServerError creates RefreshTokenInternalServerError with default headers values
func NewRefreshTokenInternalServerError() *RefreshTokenInternalServerError {

	return &RefreshTokenInternalServerError{}
}

// WithPayload adds the payload to the refresh token internal server error response
func (o *RefreshTokenInternalServerError) WithPayload(payload *models.APIResponse) *RefreshTokenInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the refresh token internal server error response
func (o *RefreshTokenInternalServerError) SetPayload(payload *models.APIResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RefreshTokenInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
