// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cmh2166/elag18apis/taquito/generated/models"
)

// UpdateResourceOKCode is the HTTP code returned for type UpdateResourceOK
const UpdateResourceOKCode int = 200

/*UpdateResourceOK TACO resource metadata updated & processing started.

swagger:response updateResourceOK
*/
type UpdateResourceOK struct {

	/*
	  In: Body
	*/
	Payload models.ResourceResponse `json:"body,omitempty"`
}

// NewUpdateResourceOK creates UpdateResourceOK with default headers values
func NewUpdateResourceOK() *UpdateResourceOK {
	return &UpdateResourceOK{}
}

// WithPayload adds the payload to the update resource o k response
func (o *UpdateResourceOK) WithPayload(payload models.ResourceResponse) *UpdateResourceOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update resource o k response
func (o *UpdateResourceOK) SetPayload(payload models.ResourceResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateResourceOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// UpdateResourceUnauthorizedCode is the HTTP code returned for type UpdateResourceUnauthorized
const UpdateResourceUnauthorizedCode int = 401

/*UpdateResourceUnauthorized You are not authorized to update a resource in TACO.

swagger:response updateResourceUnauthorized
*/
type UpdateResourceUnauthorized struct {
}

// NewUpdateResourceUnauthorized creates UpdateResourceUnauthorized with default headers values
func NewUpdateResourceUnauthorized() *UpdateResourceUnauthorized {
	return &UpdateResourceUnauthorized{}
}

// WriteResponse to the client
func (o *UpdateResourceUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// UpdateResourceNotFoundCode is the HTTP code returned for type UpdateResourceNotFound
const UpdateResourceNotFoundCode int = 404

/*UpdateResourceNotFound Invalid ID supplied

swagger:response updateResourceNotFound
*/
type UpdateResourceNotFound struct {
}

// NewUpdateResourceNotFound creates UpdateResourceNotFound with default headers values
func NewUpdateResourceNotFound() *UpdateResourceNotFound {
	return &UpdateResourceNotFound{}
}

// WriteResponse to the client
func (o *UpdateResourceNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// UpdateResourceUnsupportedMediaTypeCode is the HTTP code returned for type UpdateResourceUnsupportedMediaType
const UpdateResourceUnsupportedMediaTypeCode int = 415

/*UpdateResourceUnsupportedMediaType Unsupported resource type provided. TACO resources should be handed over as JSON or JSON-LD.

swagger:response updateResourceUnsupportedMediaType
*/
type UpdateResourceUnsupportedMediaType struct {
}

// NewUpdateResourceUnsupportedMediaType creates UpdateResourceUnsupportedMediaType with default headers values
func NewUpdateResourceUnsupportedMediaType() *UpdateResourceUnsupportedMediaType {
	return &UpdateResourceUnsupportedMediaType{}
}

// WriteResponse to the client
func (o *UpdateResourceUnsupportedMediaType) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(415)
}

// UpdateResourceUnprocessableEntityCode is the HTTP code returned for type UpdateResourceUnprocessableEntity
const UpdateResourceUnprocessableEntityCode int = 422

/*UpdateResourceUnprocessableEntity The resource JSON provided had an unspecified or unsupported field, or is otherwise unprocessable by TACO.

swagger:response updateResourceUnprocessableEntity
*/
type UpdateResourceUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewUpdateResourceUnprocessableEntity creates UpdateResourceUnprocessableEntity with default headers values
func NewUpdateResourceUnprocessableEntity() *UpdateResourceUnprocessableEntity {
	return &UpdateResourceUnprocessableEntity{}
}

// WithPayload adds the payload to the update resource unprocessable entity response
func (o *UpdateResourceUnprocessableEntity) WithPayload(payload *models.ErrorResponse) *UpdateResourceUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update resource unprocessable entity response
func (o *UpdateResourceUnprocessableEntity) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateResourceUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateResourceInternalServerErrorCode is the HTTP code returned for type UpdateResourceInternalServerError
const UpdateResourceInternalServerErrorCode int = 500

/*UpdateResourceInternalServerError This resource could be updated at this time by TACO.

swagger:response updateResourceInternalServerError
*/
type UpdateResourceInternalServerError struct {
}

// NewUpdateResourceInternalServerError creates UpdateResourceInternalServerError with default headers values
func NewUpdateResourceInternalServerError() *UpdateResourceInternalServerError {
	return &UpdateResourceInternalServerError{}
}

// WriteResponse to the client
func (o *UpdateResourceInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
