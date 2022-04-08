// Code generated by go-swagger; DO NOT EDIT.

package vsphere

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/vmware-tanzu/community-edition/cli/cmd/plugin/ui/server/models"
)

// GetVSphereNetworksOKCode is the HTTP code returned for type GetVSphereNetworksOK
const GetVSphereNetworksOKCode int = 200

/*GetVSphereNetworksOK Successful retrieval of vSphere networks

swagger:response getVSphereNetworksOK
*/
type GetVSphereNetworksOK struct {

	/*a list of vpshere networks
	  In: Body
	*/
	Payload []*models.VSphereNetwork `json:"body,omitempty"`
}

// NewGetVSphereNetworksOK creates GetVSphereNetworksOK with default headers values
func NewGetVSphereNetworksOK() *GetVSphereNetworksOK {

	return &GetVSphereNetworksOK{}
}

// WithPayload adds the payload to the get v sphere networks o k response
func (o *GetVSphereNetworksOK) WithPayload(payload []*models.VSphereNetwork) *GetVSphereNetworksOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get v sphere networks o k response
func (o *GetVSphereNetworksOK) SetPayload(payload []*models.VSphereNetwork) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetVSphereNetworksOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.VSphereNetwork, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetVSphereNetworksBadRequestCode is the HTTP code returned for type GetVSphereNetworksBadRequest
const GetVSphereNetworksBadRequestCode int = 400

/*GetVSphereNetworksBadRequest Bad request

swagger:response getVSphereNetworksBadRequest
*/
type GetVSphereNetworksBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetVSphereNetworksBadRequest creates GetVSphereNetworksBadRequest with default headers values
func NewGetVSphereNetworksBadRequest() *GetVSphereNetworksBadRequest {

	return &GetVSphereNetworksBadRequest{}
}

// WithPayload adds the payload to the get v sphere networks bad request response
func (o *GetVSphereNetworksBadRequest) WithPayload(payload *models.Error) *GetVSphereNetworksBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get v sphere networks bad request response
func (o *GetVSphereNetworksBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetVSphereNetworksBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetVSphereNetworksUnauthorizedCode is the HTTP code returned for type GetVSphereNetworksUnauthorized
const GetVSphereNetworksUnauthorizedCode int = 401

/*GetVSphereNetworksUnauthorized Incorrect credentials

swagger:response getVSphereNetworksUnauthorized
*/
type GetVSphereNetworksUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetVSphereNetworksUnauthorized creates GetVSphereNetworksUnauthorized with default headers values
func NewGetVSphereNetworksUnauthorized() *GetVSphereNetworksUnauthorized {

	return &GetVSphereNetworksUnauthorized{}
}

// WithPayload adds the payload to the get v sphere networks unauthorized response
func (o *GetVSphereNetworksUnauthorized) WithPayload(payload *models.Error) *GetVSphereNetworksUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get v sphere networks unauthorized response
func (o *GetVSphereNetworksUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetVSphereNetworksUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetVSphereNetworksInternalServerErrorCode is the HTTP code returned for type GetVSphereNetworksInternalServerError
const GetVSphereNetworksInternalServerErrorCode int = 500

/*GetVSphereNetworksInternalServerError Internal server error

swagger:response getVSphereNetworksInternalServerError
*/
type GetVSphereNetworksInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetVSphereNetworksInternalServerError creates GetVSphereNetworksInternalServerError with default headers values
func NewGetVSphereNetworksInternalServerError() *GetVSphereNetworksInternalServerError {

	return &GetVSphereNetworksInternalServerError{}
}

// WithPayload adds the payload to the get v sphere networks internal server error response
func (o *GetVSphereNetworksInternalServerError) WithPayload(payload *models.Error) *GetVSphereNetworksInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get v sphere networks internal server error response
func (o *GetVSphereNetworksInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetVSphereNetworksInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}