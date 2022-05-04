// Code generated by go-swagger; DO NOT EDIT.

package cluster

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteWorkloadClusterParams creates a new DeleteWorkloadClusterParams object
// no default values defined in spec.
func NewDeleteWorkloadClusterParams() DeleteWorkloadClusterParams {

	return DeleteWorkloadClusterParams{}
}

// DeleteWorkloadClusterParams contains all the bound params for the delete workload cluster operation
// typically these are obtained from a http.Request
//
// swagger:parameters deleteWorkloadCluster
type DeleteWorkloadClusterParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*The name of the workload cluster.
	  Required: true
	  In: path
	*/
	ClusterName string
	/*The namespace of the workload cluster.
	  In: header
	*/
	ClusterNamespace *string
	/*The name of the management cluster.
	  Required: true
	  In: path
	*/
	ManagementClusterName string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewDeleteWorkloadClusterParams() beforehand.
func (o *DeleteWorkloadClusterParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rClusterName, rhkClusterName, _ := route.Params.GetOK("clusterName")
	if err := o.bindClusterName(rClusterName, rhkClusterName, route.Formats); err != nil {
		res = append(res, err)
	}

	if err := o.bindClusterNamespace(r.Header[http.CanonicalHeaderKey("clusterNamespace")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	rManagementClusterName, rhkManagementClusterName, _ := route.Params.GetOK("managementClusterName")
	if err := o.bindManagementClusterName(rManagementClusterName, rhkManagementClusterName, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindClusterName binds and validates parameter ClusterName from path.
func (o *DeleteWorkloadClusterParams) bindClusterName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.ClusterName = raw

	return nil
}

// bindClusterNamespace binds and validates parameter ClusterNamespace from header.
func (o *DeleteWorkloadClusterParams) bindClusterNamespace(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false

	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.ClusterNamespace = &raw

	return nil
}

// bindManagementClusterName binds and validates parameter ManagementClusterName from path.
func (o *DeleteWorkloadClusterParams) bindManagementClusterName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.ManagementClusterName = raw

	return nil
}
