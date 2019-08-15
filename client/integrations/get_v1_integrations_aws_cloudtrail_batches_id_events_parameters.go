// Code generated by go-swagger; DO NOT EDIT.

package integrations

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

// NewGetV1IntegrationsAwsCloudtrailBatchesIDEventsParams creates a new GetV1IntegrationsAwsCloudtrailBatchesIDEventsParams object
// with the default values initialized.
func NewGetV1IntegrationsAwsCloudtrailBatchesIDEventsParams() *GetV1IntegrationsAwsCloudtrailBatchesIDEventsParams {
	var ()
	return &GetV1IntegrationsAwsCloudtrailBatchesIDEventsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1IntegrationsAwsCloudtrailBatchesIDEventsParamsWithTimeout creates a new GetV1IntegrationsAwsCloudtrailBatchesIDEventsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetV1IntegrationsAwsCloudtrailBatchesIDEventsParamsWithTimeout(timeout time.Duration) *GetV1IntegrationsAwsCloudtrailBatchesIDEventsParams {
	var ()
	return &GetV1IntegrationsAwsCloudtrailBatchesIDEventsParams{

		timeout: timeout,
	}
}

// NewGetV1IntegrationsAwsCloudtrailBatchesIDEventsParamsWithContext creates a new GetV1IntegrationsAwsCloudtrailBatchesIDEventsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetV1IntegrationsAwsCloudtrailBatchesIDEventsParamsWithContext(ctx context.Context) *GetV1IntegrationsAwsCloudtrailBatchesIDEventsParams {
	var ()
	return &GetV1IntegrationsAwsCloudtrailBatchesIDEventsParams{

		Context: ctx,
	}
}

// NewGetV1IntegrationsAwsCloudtrailBatchesIDEventsParamsWithHTTPClient creates a new GetV1IntegrationsAwsCloudtrailBatchesIDEventsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetV1IntegrationsAwsCloudtrailBatchesIDEventsParamsWithHTTPClient(client *http.Client) *GetV1IntegrationsAwsCloudtrailBatchesIDEventsParams {
	var ()
	return &GetV1IntegrationsAwsCloudtrailBatchesIDEventsParams{
		HTTPClient: client,
	}
}

/*GetV1IntegrationsAwsCloudtrailBatchesIDEventsParams contains all the parameters to send to the API endpoint
for the get v1 integrations aws cloudtrail batches Id events operation typically these are written to a http.Request
*/
type GetV1IntegrationsAwsCloudtrailBatchesIDEventsParams struct {

	/*ID
	  Batch UUID

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get v1 integrations aws cloudtrail batches Id events params
func (o *GetV1IntegrationsAwsCloudtrailBatchesIDEventsParams) WithTimeout(timeout time.Duration) *GetV1IntegrationsAwsCloudtrailBatchesIDEventsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 integrations aws cloudtrail batches Id events params
func (o *GetV1IntegrationsAwsCloudtrailBatchesIDEventsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 integrations aws cloudtrail batches Id events params
func (o *GetV1IntegrationsAwsCloudtrailBatchesIDEventsParams) WithContext(ctx context.Context) *GetV1IntegrationsAwsCloudtrailBatchesIDEventsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 integrations aws cloudtrail batches Id events params
func (o *GetV1IntegrationsAwsCloudtrailBatchesIDEventsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 integrations aws cloudtrail batches Id events params
func (o *GetV1IntegrationsAwsCloudtrailBatchesIDEventsParams) WithHTTPClient(client *http.Client) *GetV1IntegrationsAwsCloudtrailBatchesIDEventsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 integrations aws cloudtrail batches Id events params
func (o *GetV1IntegrationsAwsCloudtrailBatchesIDEventsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get v1 integrations aws cloudtrail batches Id events params
func (o *GetV1IntegrationsAwsCloudtrailBatchesIDEventsParams) WithID(id string) *GetV1IntegrationsAwsCloudtrailBatchesIDEventsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get v1 integrations aws cloudtrail batches Id events params
func (o *GetV1IntegrationsAwsCloudtrailBatchesIDEventsParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1IntegrationsAwsCloudtrailBatchesIDEventsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
