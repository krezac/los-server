// Code generated by go-swagger; DO NOT EDIT.

package competition

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"

	"github.com/go-openapi/swag"
)

// GetCompetitionByIDURL generates an URL for the get competition by Id operation
type GetCompetitionByIDURL struct {
	CompetitionID int64

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetCompetitionByIDURL) WithBasePath(bp string) *GetCompetitionByIDURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetCompetitionByIDURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *GetCompetitionByIDURL) Build() (*url.URL, error) {
	var result url.URL

	var _path = "/competitions/{competitionId}"

	competitionID := swag.FormatInt64(o.CompetitionID)
	if competitionID != "" {
		_path = strings.Replace(_path, "{competitionId}", competitionID, -1)
	} else {
		return nil, errors.New("CompetitionID is required on GetCompetitionByIDURL")
	}

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/v1"
	}
	result.Path = golangswaggerpaths.Join(_basePath, _path)

	return &result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *GetCompetitionByIDURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *GetCompetitionByIDURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *GetCompetitionByIDURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on GetCompetitionByIDURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on GetCompetitionByIDURL")
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
func (o *GetCompetitionByIDURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}