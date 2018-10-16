// Code generated by go-swagger; DO NOT EDIT.

package competition

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

/*CreateCompetitionDefault successful operation

swagger:response createCompetitionDefault
*/
type CreateCompetitionDefault struct {
	_statusCode int
}

// NewCreateCompetitionDefault creates CreateCompetitionDefault with default headers values
func NewCreateCompetitionDefault(code int) *CreateCompetitionDefault {
	if code <= 0 {
		code = 500
	}

	return &CreateCompetitionDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the create competition default response
func (o *CreateCompetitionDefault) WithStatusCode(code int) *CreateCompetitionDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the create competition default response
func (o *CreateCompetitionDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WriteResponse to the client
func (o *CreateCompetitionDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(o._statusCode)
}