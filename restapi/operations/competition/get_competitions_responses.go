// Code generated by go-swagger; DO NOT EDIT.

package competition

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/krezac/los-server/models"
)

// GetCompetitionsOKCode is the HTTP code returned for type GetCompetitionsOK
const GetCompetitionsOKCode int = 200

/*GetCompetitionsOK successful operation

swagger:response getCompetitionsOK
*/
type GetCompetitionsOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Competition `json:"body,omitempty"`
}

// NewGetCompetitionsOK creates GetCompetitionsOK with default headers values
func NewGetCompetitionsOK() *GetCompetitionsOK {

	return &GetCompetitionsOK{}
}

// WithPayload adds the payload to the get competitions o k response
func (o *GetCompetitionsOK) WithPayload(payload []*models.Competition) *GetCompetitionsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get competitions o k response
func (o *GetCompetitionsOK) SetPayload(payload []*models.Competition) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCompetitionsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make([]*models.Competition, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// GetCompetitionsInternalServerErrorCode is the HTTP code returned for type GetCompetitionsInternalServerError
const GetCompetitionsInternalServerErrorCode int = 500

/*GetCompetitionsInternalServerError Retrieving list of competitions failed

swagger:response getCompetitionsInternalServerError
*/
type GetCompetitionsInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.APIResponse `json:"body,omitempty"`
}

// NewGetCompetitionsInternalServerError creates GetCompetitionsInternalServerError with default headers values
func NewGetCompetitionsInternalServerError() *GetCompetitionsInternalServerError {

	return &GetCompetitionsInternalServerError{}
}

// WithPayload adds the payload to the get competitions internal server error response
func (o *GetCompetitionsInternalServerError) WithPayload(payload *models.APIResponse) *GetCompetitionsInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get competitions internal server error response
func (o *GetCompetitionsInternalServerError) SetPayload(payload *models.APIResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCompetitionsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}