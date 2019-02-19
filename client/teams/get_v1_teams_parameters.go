// Code generated by go-swagger; DO NOT EDIT.

package teams

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetV1TeamsParams creates a new GetV1TeamsParams object
// with the default values initialized.
func NewGetV1TeamsParams() *GetV1TeamsParams {
	var ()
	return &GetV1TeamsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1TeamsParamsWithTimeout creates a new GetV1TeamsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetV1TeamsParamsWithTimeout(timeout time.Duration) *GetV1TeamsParams {
	var ()
	return &GetV1TeamsParams{

		timeout: timeout,
	}
}

// NewGetV1TeamsParamsWithContext creates a new GetV1TeamsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetV1TeamsParamsWithContext(ctx context.Context) *GetV1TeamsParams {
	var ()
	return &GetV1TeamsParams{

		Context: ctx,
	}
}

// NewGetV1TeamsParamsWithHTTPClient creates a new GetV1TeamsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetV1TeamsParamsWithHTTPClient(client *http.Client) *GetV1TeamsParams {
	var ()
	return &GetV1TeamsParams{
		HTTPClient: client,
	}
}

/*GetV1TeamsParams contains all the parameters to send to the API endpoint
for the get v1 teams operation typically these are written to a http.Request
*/
type GetV1TeamsParams struct {

	/*Page*/
	Page *int32
	/*PerPage*/
	PerPage *int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get v1 teams params
func (o *GetV1TeamsParams) WithTimeout(timeout time.Duration) *GetV1TeamsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 teams params
func (o *GetV1TeamsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 teams params
func (o *GetV1TeamsParams) WithContext(ctx context.Context) *GetV1TeamsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 teams params
func (o *GetV1TeamsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 teams params
func (o *GetV1TeamsParams) WithHTTPClient(client *http.Client) *GetV1TeamsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 teams params
func (o *GetV1TeamsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPage adds the page to the get v1 teams params
func (o *GetV1TeamsParams) WithPage(page *int32) *GetV1TeamsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get v1 teams params
func (o *GetV1TeamsParams) SetPage(page *int32) {
	o.Page = page
}

// WithPerPage adds the perPage to the get v1 teams params
func (o *GetV1TeamsParams) WithPerPage(perPage *int32) *GetV1TeamsParams {
	o.SetPerPage(perPage)
	return o
}

// SetPerPage adds the perPage to the get v1 teams params
func (o *GetV1TeamsParams) SetPerPage(perPage *int32) {
	o.PerPage = perPage
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1TeamsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Page != nil {

		// query param page
		var qrPage int32
		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt32(qrPage)
		if qPage != "" {
			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}

	}

	if o.PerPage != nil {

		// query param per_page
		var qrPerPage int32
		if o.PerPage != nil {
			qrPerPage = *o.PerPage
		}
		qPerPage := swag.FormatInt32(qrPerPage)
		if qPerPage != "" {
			if err := r.SetQueryParam("per_page", qPerPage); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
