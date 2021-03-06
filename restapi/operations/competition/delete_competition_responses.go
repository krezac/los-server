// Code generated by go-swagger; DO NOT EDIT.

package competition

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// DeleteCompetitionBadRequestCode is the HTTP code returned for type DeleteCompetitionBadRequest
const DeleteCompetitionBadRequestCode int = 400

/*DeleteCompetitionBadRequest Invalid competition supplied

swagger:response deleteCompetitionBadRequest
*/
type DeleteCompetitionBadRequest struct {
}

// NewDeleteCompetitionBadRequest creates DeleteCompetitionBadRequest with default headers values
func NewDeleteCompetitionBadRequest() *DeleteCompetitionBadRequest {

	return &DeleteCompetitionBadRequest{}
}

// WriteResponse to the client
func (o *DeleteCompetitionBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// DeleteCompetitionNotFoundCode is the HTTP code returned for type DeleteCompetitionNotFound
const DeleteCompetitionNotFoundCode int = 404

/*DeleteCompetitionNotFound Competition not found

swagger:response deleteCompetitionNotFound
*/
type DeleteCompetitionNotFound struct {
}

// NewDeleteCompetitionNotFound creates DeleteCompetitionNotFound with default headers values
func NewDeleteCompetitionNotFound() *DeleteCompetitionNotFound {

	return &DeleteCompetitionNotFound{}
}

// WriteResponse to the client
func (o *DeleteCompetitionNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}
