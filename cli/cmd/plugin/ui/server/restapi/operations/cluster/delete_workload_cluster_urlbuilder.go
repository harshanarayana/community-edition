// Code generated by go-swagger; DO NOT EDIT.

package cluster

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"
)

// DeleteWorkloadClusterURL generates an URL for the delete workload cluster operation
type DeleteWorkloadClusterURL struct {
	ClusterName           string
	ManagementClusterName string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *DeleteWorkloadClusterURL) WithBasePath(bp string) *DeleteWorkloadClusterURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *DeleteWorkloadClusterURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *DeleteWorkloadClusterURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/api/management/{managementClusterName}/cluster/{clusterName}"

	clusterName := o.ClusterName
	if clusterName != "" {
		_path = strings.Replace(_path, "{clusterName}", clusterName, -1)
	} else {
		return nil, errors.New("clusterName is required on DeleteWorkloadClusterURL")
	}

	managementClusterName := o.ManagementClusterName
	if managementClusterName != "" {
		_path = strings.Replace(_path, "{managementClusterName}", managementClusterName, -1)
	} else {
		return nil, errors.New("managementClusterName is required on DeleteWorkloadClusterURL")
	}

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/"
	}
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *DeleteWorkloadClusterURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *DeleteWorkloadClusterURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *DeleteWorkloadClusterURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on DeleteWorkloadClusterURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on DeleteWorkloadClusterURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *DeleteWorkloadClusterURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
