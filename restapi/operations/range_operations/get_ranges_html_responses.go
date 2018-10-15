// Code generated by go-swagger; DO NOT EDIT.

package range_operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/krezac/los-server/models"
)

// GetRangesHTMLOKCode is the HTTP code returned for type GetRangesHTMLOK
const GetRangesHTMLOKCode int = 200

/*GetRangesHTMLOK successful operation

swagger:response getRangesHtmlOK
*/
type GetRangesHTMLOK struct {

	/*
	  In: Body
	*/
	Payload *models.HTMLResponse `json:"body,omitempty"`
}

// NewGetRangesHTMLOK creates GetRangesHTMLOK with default headers values
func NewGetRangesHTMLOK() *GetRangesHTMLOK {

	return &GetRangesHTMLOK{}
}

// WithPayload adds the payload to the get ranges Html o k response
func (o *GetRangesHTMLOK) WithPayload(payload *models.HTMLResponse) *GetRangesHTMLOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get ranges Html o k response
func (o *GetRangesHTMLOK) SetPayload(payload *models.HTMLResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRangesHTMLOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetRangesHTMLInternalServerErrorCode is the HTTP code returned for type GetRangesHTMLInternalServerError
const GetRangesHTMLInternalServerErrorCode int = 500

/*GetRangesHTMLInternalServerError Retrieving list of ranges failed

swagger:response getRangesHtmlInternalServerError
*/
type GetRangesHTMLInternalServerError struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewGetRangesHTMLInternalServerError creates GetRangesHTMLInternalServerError with default headers values
func NewGetRangesHTMLInternalServerError() *GetRangesHTMLInternalServerError {

	return &GetRangesHTMLInternalServerError{}
}

// WithPayload adds the payload to the get ranges Html internal server error response
func (o *GetRangesHTMLInternalServerError) WithPayload(payload string) *GetRangesHTMLInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get ranges Html internal server error response
func (o *GetRangesHTMLInternalServerError) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRangesHTMLInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}
