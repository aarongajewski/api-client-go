// Code generated by go-swagger; DO NOT EDIT.

package severities

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

// NewPatchV1SeveritiesSeveritySlugParams creates a new PatchV1SeveritiesSeveritySlugParams object
// with the default values initialized.
func NewPatchV1SeveritiesSeveritySlugParams() *PatchV1SeveritiesSeveritySlugParams {
	var ()
	return &PatchV1SeveritiesSeveritySlugParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchV1SeveritiesSeveritySlugParamsWithTimeout creates a new PatchV1SeveritiesSeveritySlugParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchV1SeveritiesSeveritySlugParamsWithTimeout(timeout time.Duration) *PatchV1SeveritiesSeveritySlugParams {
	var ()
	return &PatchV1SeveritiesSeveritySlugParams{

		timeout: timeout,
	}
}

// NewPatchV1SeveritiesSeveritySlugParamsWithContext creates a new PatchV1SeveritiesSeveritySlugParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchV1SeveritiesSeveritySlugParamsWithContext(ctx context.Context) *PatchV1SeveritiesSeveritySlugParams {
	var ()
	return &PatchV1SeveritiesSeveritySlugParams{

		Context: ctx,
	}
}

// NewPatchV1SeveritiesSeveritySlugParamsWithHTTPClient creates a new PatchV1SeveritiesSeveritySlugParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchV1SeveritiesSeveritySlugParamsWithHTTPClient(client *http.Client) *PatchV1SeveritiesSeveritySlugParams {
	var ()
	return &PatchV1SeveritiesSeveritySlugParams{
		HTTPClient: client,
	}
}

/*PatchV1SeveritiesSeveritySlugParams contains all the parameters to send to the API endpoint
for the patch v1 severities severity slug operation typically these are written to a http.Request
*/
type PatchV1SeveritiesSeveritySlugParams struct {

	/*Description*/
	Description *string
	/*SeveritySlug*/
	SeveritySlug string
	/*Slug*/
	Slug *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch v1 severities severity slug params
func (o *PatchV1SeveritiesSeveritySlugParams) WithTimeout(timeout time.Duration) *PatchV1SeveritiesSeveritySlugParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch v1 severities severity slug params
func (o *PatchV1SeveritiesSeveritySlugParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch v1 severities severity slug params
func (o *PatchV1SeveritiesSeveritySlugParams) WithContext(ctx context.Context) *PatchV1SeveritiesSeveritySlugParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch v1 severities severity slug params
func (o *PatchV1SeveritiesSeveritySlugParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch v1 severities severity slug params
func (o *PatchV1SeveritiesSeveritySlugParams) WithHTTPClient(client *http.Client) *PatchV1SeveritiesSeveritySlugParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch v1 severities severity slug params
func (o *PatchV1SeveritiesSeveritySlugParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDescription adds the description to the patch v1 severities severity slug params
func (o *PatchV1SeveritiesSeveritySlugParams) WithDescription(description *string) *PatchV1SeveritiesSeveritySlugParams {
	o.SetDescription(description)
	return o
}

// SetDescription adds the description to the patch v1 severities severity slug params
func (o *PatchV1SeveritiesSeveritySlugParams) SetDescription(description *string) {
	o.Description = description
}

// WithSeveritySlug adds the severitySlug to the patch v1 severities severity slug params
func (o *PatchV1SeveritiesSeveritySlugParams) WithSeveritySlug(severitySlug string) *PatchV1SeveritiesSeveritySlugParams {
	o.SetSeveritySlug(severitySlug)
	return o
}

// SetSeveritySlug adds the severitySlug to the patch v1 severities severity slug params
func (o *PatchV1SeveritiesSeveritySlugParams) SetSeveritySlug(severitySlug string) {
	o.SeveritySlug = severitySlug
}

// WithSlug adds the slug to the patch v1 severities severity slug params
func (o *PatchV1SeveritiesSeveritySlugParams) WithSlug(slug *string) *PatchV1SeveritiesSeveritySlugParams {
	o.SetSlug(slug)
	return o
}

// SetSlug adds the slug to the patch v1 severities severity slug params
func (o *PatchV1SeveritiesSeveritySlugParams) SetSlug(slug *string) {
	o.Slug = slug
}

// WriteToRequest writes these params to a swagger request
func (o *PatchV1SeveritiesSeveritySlugParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Description != nil {

		// form param description
		var frDescription string
		if o.Description != nil {
			frDescription = *o.Description
		}
		fDescription := frDescription
		if fDescription != "" {
			if err := r.SetFormParam("description", fDescription); err != nil {
				return err
			}
		}

	}

	// path param severity_slug
	if err := r.SetPathParam("severity_slug", o.SeveritySlug); err != nil {
		return err
	}

	if o.Slug != nil {

		// form param slug
		var frSlug string
		if o.Slug != nil {
			frSlug = *o.Slug
		}
		fSlug := frSlug
		if fSlug != "" {
			if err := r.SetFormParam("slug", fSlug); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
