// Code generated by go-swagger; DO NOT EDIT.

package cluster

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// DeleteWorkloadClusterHandlerFunc turns a function with the right signature into a delete workload cluster handler
type DeleteWorkloadClusterHandlerFunc func(DeleteWorkloadClusterParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteWorkloadClusterHandlerFunc) Handle(params DeleteWorkloadClusterParams) middleware.Responder {
	return fn(params)
}

// DeleteWorkloadClusterHandler interface for that can handle valid delete workload cluster params
type DeleteWorkloadClusterHandler interface {
	Handle(DeleteWorkloadClusterParams) middleware.Responder
}

// NewDeleteWorkloadCluster creates a new http.Handler for the delete workload cluster operation
func NewDeleteWorkloadCluster(ctx *middleware.Context, handler DeleteWorkloadClusterHandler) *DeleteWorkloadCluster {
	return &DeleteWorkloadCluster{Context: ctx, Handler: handler}
}

/*DeleteWorkloadCluster swagger:route DELETE /api/management/{managementClusterName}/cluster/{clusterName} cluster deleteWorkloadCluster

Delete a workload cluster.

*/
type DeleteWorkloadCluster struct {
	Context *middleware.Context
	Handler DeleteWorkloadClusterHandler
}

func (o *DeleteWorkloadCluster) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteWorkloadClusterParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
