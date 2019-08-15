// Code generated by go-swagger; DO NOT EDIT.

package severity_matrix

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetV1SeverityMatrixParams creates a new GetV1SeverityMatrixParams object
// with the default values initialized.
func NewGetV1SeverityMatrixParams() *GetV1SeverityMatrixParams {

	return &GetV1SeverityMatrixParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1SeverityMatrixParamsWithTimeout creates a new GetV1SeverityMatrixParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetV1SeverityMatrixParamsWithTimeout(timeout time.Duration) *GetV1SeverityMatrixParams {

	return &GetV1SeverityMatrixParams{

		timeout: timeout,
	}
}

// NewGetV1SeverityMatrixParamsWithContext creates a new GetV1SeverityMatrixParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetV1SeverityMatrixParamsWithContext(ctx context.Context) *GetV1SeverityMatrixParams {

	return &GetV1SeverityMatrixParams{

		Context: ctx,
	}
}

// NewGetV1SeverityMatrixParamsWithHTTPClient creates a new GetV1SeverityMatrixParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetV1SeverityMatrixParamsWithHTTPClient(client *http.Client) *GetV1SeverityMatrixParams {

	return &GetV1SeverityMatrixParams{
		HTTPClient: client,
	}
}

/*GetV1SeverityMatrixParams contains all the parameters to send to the API endpoint
for the get v1 severity matrix operation typically these are written to a http.Request
*/
type GetV1SeverityMatrixParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get v1 severity matrix params
func (o *GetV1SeverityMatrixParams) WithTimeout(timeout time.Duration) *GetV1SeverityMatrixParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 severity matrix params
func (o *GetV1SeverityMatrixParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 severity matrix params
func (o *GetV1SeverityMatrixParams) WithContext(ctx context.Context) *GetV1SeverityMatrixParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 severity matrix params
func (o *GetV1SeverityMatrixParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 severity matrix params
func (o *GetV1SeverityMatrixParams) WithHTTPClient(client *http.Client) *GetV1SeverityMatrixParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 severity matrix params
func (o *GetV1SeverityMatrixParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1SeverityMatrixParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
