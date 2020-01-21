// Code generated by go-swagger; DO NOT EDIT.

package deployment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cloustone/pandas/models"
)

// GetDeploymentsOKCode is the HTTP code returned for type GetDeploymentsOK
const GetDeploymentsOKCode int = 200

/*GetDeploymentsOK successfully operation

swagger:response getDeploymentsOK
*/
type GetDeploymentsOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Deployment `json:"body,omitempty"`
}

// NewGetDeploymentsOK creates GetDeploymentsOK with default headers values
func NewGetDeploymentsOK() *GetDeploymentsOK {

	return &GetDeploymentsOK{}
}

// WithPayload adds the payload to the get deployments o k response
func (o *GetDeploymentsOK) WithPayload(payload []*models.Deployment) *GetDeploymentsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get deployments o k response
func (o *GetDeploymentsOK) SetPayload(payload []*models.Deployment) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetDeploymentsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Deployment, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetDeploymentsBadRequestCode is the HTTP code returned for type GetDeploymentsBadRequest
const GetDeploymentsBadRequestCode int = 400

/*GetDeploymentsBadRequest Internal server rooro

swagger:response getDeploymentsBadRequest
*/
type GetDeploymentsBadRequest struct {
}

// NewGetDeploymentsBadRequest creates GetDeploymentsBadRequest with default headers values
func NewGetDeploymentsBadRequest() *GetDeploymentsBadRequest {

	return &GetDeploymentsBadRequest{}
}

// WriteResponse to the client
func (o *GetDeploymentsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

/*GetDeploymentsDefault failed operation

swagger:response getDeploymentsDefault
*/
type GetDeploymentsDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *GetDeploymentsDefaultBody `json:"body,omitempty"`
}

// NewGetDeploymentsDefault creates GetDeploymentsDefault with default headers values
func NewGetDeploymentsDefault(code int) *GetDeploymentsDefault {
	if code <= 0 {
		code = 500
	}

	return &GetDeploymentsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get deployments default response
func (o *GetDeploymentsDefault) WithStatusCode(code int) *GetDeploymentsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get deployments default response
func (o *GetDeploymentsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get deployments default response
func (o *GetDeploymentsDefault) WithPayload(payload *GetDeploymentsDefaultBody) *GetDeploymentsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get deployments default response
func (o *GetDeploymentsDefault) SetPayload(payload *GetDeploymentsDefaultBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetDeploymentsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
