// Code generated by go-swagger; DO NOT EDIT.

package saved_searches

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

// NewGetV1SavedSearchesResourceTypeSavedSearchIDParams creates a new GetV1SavedSearchesResourceTypeSavedSearchIDParams object
// with the default values initialized.
func NewGetV1SavedSearchesResourceTypeSavedSearchIDParams() *GetV1SavedSearchesResourceTypeSavedSearchIDParams {
	var ()
	return &GetV1SavedSearchesResourceTypeSavedSearchIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1SavedSearchesResourceTypeSavedSearchIDParamsWithTimeout creates a new GetV1SavedSearchesResourceTypeSavedSearchIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetV1SavedSearchesResourceTypeSavedSearchIDParamsWithTimeout(timeout time.Duration) *GetV1SavedSearchesResourceTypeSavedSearchIDParams {
	var ()
	return &GetV1SavedSearchesResourceTypeSavedSearchIDParams{

		timeout: timeout,
	}
}

// NewGetV1SavedSearchesResourceTypeSavedSearchIDParamsWithContext creates a new GetV1SavedSearchesResourceTypeSavedSearchIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetV1SavedSearchesResourceTypeSavedSearchIDParamsWithContext(ctx context.Context) *GetV1SavedSearchesResourceTypeSavedSearchIDParams {
	var ()
	return &GetV1SavedSearchesResourceTypeSavedSearchIDParams{

		Context: ctx,
	}
}

// NewGetV1SavedSearchesResourceTypeSavedSearchIDParamsWithHTTPClient creates a new GetV1SavedSearchesResourceTypeSavedSearchIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetV1SavedSearchesResourceTypeSavedSearchIDParamsWithHTTPClient(client *http.Client) *GetV1SavedSearchesResourceTypeSavedSearchIDParams {
	var ()
	return &GetV1SavedSearchesResourceTypeSavedSearchIDParams{
		HTTPClient: client,
	}
}

/*GetV1SavedSearchesResourceTypeSavedSearchIDParams contains all the parameters to send to the API endpoint
for the get v1 saved searches resource type saved search Id operation typically these are written to a http.Request
*/
type GetV1SavedSearchesResourceTypeSavedSearchIDParams struct {

	/*ResourceType*/
	ResourceType string
	/*SavedSearchID
	  ID of a saved search

	*/
	SavedSearchID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get v1 saved searches resource type saved search Id params
func (o *GetV1SavedSearchesResourceTypeSavedSearchIDParams) WithTimeout(timeout time.Duration) *GetV1SavedSearchesResourceTypeSavedSearchIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 saved searches resource type saved search Id params
func (o *GetV1SavedSearchesResourceTypeSavedSearchIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 saved searches resource type saved search Id params
func (o *GetV1SavedSearchesResourceTypeSavedSearchIDParams) WithContext(ctx context.Context) *GetV1SavedSearchesResourceTypeSavedSearchIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 saved searches resource type saved search Id params
func (o *GetV1SavedSearchesResourceTypeSavedSearchIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 saved searches resource type saved search Id params
func (o *GetV1SavedSearchesResourceTypeSavedSearchIDParams) WithHTTPClient(client *http.Client) *GetV1SavedSearchesResourceTypeSavedSearchIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 saved searches resource type saved search Id params
func (o *GetV1SavedSearchesResourceTypeSavedSearchIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithResourceType adds the resourceType to the get v1 saved searches resource type saved search Id params
func (o *GetV1SavedSearchesResourceTypeSavedSearchIDParams) WithResourceType(resourceType string) *GetV1SavedSearchesResourceTypeSavedSearchIDParams {
	o.SetResourceType(resourceType)
	return o
}

// SetResourceType adds the resourceType to the get v1 saved searches resource type saved search Id params
func (o *GetV1SavedSearchesResourceTypeSavedSearchIDParams) SetResourceType(resourceType string) {
	o.ResourceType = resourceType
}

// WithSavedSearchID adds the savedSearchID to the get v1 saved searches resource type saved search Id params
func (o *GetV1SavedSearchesResourceTypeSavedSearchIDParams) WithSavedSearchID(savedSearchID string) *GetV1SavedSearchesResourceTypeSavedSearchIDParams {
	o.SetSavedSearchID(savedSearchID)
	return o
}

// SetSavedSearchID adds the savedSearchId to the get v1 saved searches resource type saved search Id params
func (o *GetV1SavedSearchesResourceTypeSavedSearchIDParams) SetSavedSearchID(savedSearchID string) {
	o.SavedSearchID = savedSearchID
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1SavedSearchesResourceTypeSavedSearchIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param resource_type
	if err := r.SetPathParam("resource_type", o.ResourceType); err != nil {
		return err
	}

	// path param saved_search_id
	if err := r.SetPathParam("saved_search_id", o.SavedSearchID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
