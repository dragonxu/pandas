// Code generated by go-swagger; DO NOT EDIT.

package rulechain

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cloustone/pandas/models"
)

// GetRuleChainsOKCode is the HTTP code returned for type GetRuleChainsOK
const GetRuleChainsOKCode int = 200

/*GetRuleChainsOK successfully operation

swagger:response getRuleChainsOK
*/
type GetRuleChainsOK struct {

	/*
	  In: Body
	*/
	Payload []*models.RuleChain `json:"body,omitempty"`
}

// NewGetRuleChainsOK creates GetRuleChainsOK with default headers values
func NewGetRuleChainsOK() *GetRuleChainsOK {

	return &GetRuleChainsOK{}
}

// WithPayload adds the payload to the get rule chains o k response
func (o *GetRuleChainsOK) WithPayload(payload []*models.RuleChain) *GetRuleChainsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get rule chains o k response
func (o *GetRuleChainsOK) SetPayload(payload []*models.RuleChain) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRuleChainsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.RuleChain, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetRuleChainsBadRequestCode is the HTTP code returned for type GetRuleChainsBadRequest
const GetRuleChainsBadRequestCode int = 400

/*GetRuleChainsBadRequest Bad request

swagger:response getRuleChainsBadRequest
*/
type GetRuleChainsBadRequest struct {
}

// NewGetRuleChainsBadRequest creates GetRuleChainsBadRequest with default headers values
func NewGetRuleChainsBadRequest() *GetRuleChainsBadRequest {

	return &GetRuleChainsBadRequest{}
}

// WriteResponse to the client
func (o *GetRuleChainsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// GetRuleChainsNotFoundCode is the HTTP code returned for type GetRuleChainsNotFound
const GetRuleChainsNotFoundCode int = 404

/*GetRuleChainsNotFound Not found

swagger:response getRuleChainsNotFound
*/
type GetRuleChainsNotFound struct {
}

// NewGetRuleChainsNotFound creates GetRuleChainsNotFound with default headers values
func NewGetRuleChainsNotFound() *GetRuleChainsNotFound {

	return &GetRuleChainsNotFound{}
}

// WriteResponse to the client
func (o *GetRuleChainsNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// GetRuleChainsInternalServerErrorCode is the HTTP code returned for type GetRuleChainsInternalServerError
const GetRuleChainsInternalServerErrorCode int = 500

/*GetRuleChainsInternalServerError Internal error

swagger:response getRuleChainsInternalServerError
*/
type GetRuleChainsInternalServerError struct {
}

// NewGetRuleChainsInternalServerError creates GetRuleChainsInternalServerError with default headers values
func NewGetRuleChainsInternalServerError() *GetRuleChainsInternalServerError {

	return &GetRuleChainsInternalServerError{}
}

// WriteResponse to the client
func (o *GetRuleChainsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
