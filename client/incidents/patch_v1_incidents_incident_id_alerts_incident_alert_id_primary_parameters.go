// Code generated by go-swagger; DO NOT EDIT.

package incidents

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

	models "github.com/firehydrant/api-client-go/models"
)

// NewPatchV1IncidentsIncidentIDAlertsIncidentAlertIDPrimaryParams creates a new PatchV1IncidentsIncidentIDAlertsIncidentAlertIDPrimaryParams object
// with the default values initialized.
func NewPatchV1IncidentsIncidentIDAlertsIncidentAlertIDPrimaryParams() *PatchV1IncidentsIncidentIDAlertsIncidentAlertIDPrimaryParams {
	var ()
	return &PatchV1IncidentsIncidentIDAlertsIncidentAlertIDPrimaryParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchV1IncidentsIncidentIDAlertsIncidentAlertIDPrimaryParamsWithTimeout creates a new PatchV1IncidentsIncidentIDAlertsIncidentAlertIDPrimaryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchV1IncidentsIncidentIDAlertsIncidentAlertIDPrimaryParamsWithTimeout(timeout time.Duration) *PatchV1IncidentsIncidentIDAlertsIncidentAlertIDPrimaryParams {
	var ()
	return &PatchV1IncidentsIncidentIDAlertsIncidentAlertIDPrimaryParams{

		timeout: timeout,
	}
}

// NewPatchV1IncidentsIncidentIDAlertsIncidentAlertIDPrimaryParamsWithContext creates a new PatchV1IncidentsIncidentIDAlertsIncidentAlertIDPrimaryParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchV1IncidentsIncidentIDAlertsIncidentAlertIDPrimaryParamsWithContext(ctx context.Context) *PatchV1IncidentsIncidentIDAlertsIncidentAlertIDPrimaryParams {
	var ()
	return &PatchV1IncidentsIncidentIDAlertsIncidentAlertIDPrimaryParams{

		Context: ctx,
	}
}

// NewPatchV1IncidentsIncidentIDAlertsIncidentAlertIDPrimaryParamsWithHTTPClient creates a new PatchV1IncidentsIncidentIDAlertsIncidentAlertIDPrimaryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchV1IncidentsIncidentIDAlertsIncidentAlertIDPrimaryParamsWithHTTPClient(client *http.Client) *PatchV1IncidentsIncidentIDAlertsIncidentAlertIDPrimaryParams {
	var ()
	return &PatchV1IncidentsIncidentIDAlertsIncidentAlertIDPrimaryParams{
		HTTPClient: client,
	}
}

/*PatchV1IncidentsIncidentIDAlertsIncidentAlertIDPrimaryParams contains all the parameters to send to the API endpoint
for the patch v1 incidents incident Id alerts incident alert Id primary operation typically these are written to a http.Request
*/
type PatchV1IncidentsIncidentIDAlertsIncidentAlertIDPrimaryParams struct {

	/*V1IncidentsIncidentIDAlertsIncidentAlertIDPrimary*/
	V1IncidentsIncidentIDAlertsIncidentAlertIDPrimary *models.PatchV1IncidentsIncidentIDAlertsIncidentAlertIDPrimary
	/*IncidentAlertID*/
	IncidentAlertID string
	/*IncidentID*/
	IncidentID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch v1 incidents incident Id alerts incident alert Id primary params
func (o *PatchV1IncidentsIncidentIDAlertsIncidentAlertIDPrimaryParams) WithTimeout(timeout time.Duration) *PatchV1IncidentsIncidentIDAlertsIncidentAlertIDPrimaryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch v1 incidents incident Id alerts incident alert Id primary params
func (o *PatchV1IncidentsIncidentIDAlertsIncidentAlertIDPrimaryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch v1 incidents incident Id alerts incident alert Id primary params
func (o *PatchV1IncidentsIncidentIDAlertsIncidentAlertIDPrimaryParams) WithContext(ctx context.Context) *PatchV1IncidentsIncidentIDAlertsIncidentAlertIDPrimaryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch v1 incidents incident Id alerts incident alert Id primary params
func (o *PatchV1IncidentsIncidentIDAlertsIncidentAlertIDPrimaryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch v1 incidents incident Id alerts incident alert Id primary params
func (o *PatchV1IncidentsIncidentIDAlertsIncidentAlertIDPrimaryParams) WithHTTPClient(client *http.Client) *PatchV1IncidentsIncidentIDAlertsIncidentAlertIDPrimaryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch v1 incidents incident Id alerts incident alert Id primary params
func (o *PatchV1IncidentsIncidentIDAlertsIncidentAlertIDPrimaryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithV1IncidentsIncidentIDAlertsIncidentAlertIDPrimary adds the v1IncidentsIncidentIDAlertsIncidentAlertIDPrimary to the patch v1 incidents incident Id alerts incident alert Id primary params
func (o *PatchV1IncidentsIncidentIDAlertsIncidentAlertIDPrimaryParams) WithV1IncidentsIncidentIDAlertsIncidentAlertIDPrimary(v1IncidentsIncidentIDAlertsIncidentAlertIDPrimary *models.PatchV1IncidentsIncidentIDAlertsIncidentAlertIDPrimary) *PatchV1IncidentsIncidentIDAlertsIncidentAlertIDPrimaryParams {
	o.SetV1IncidentsIncidentIDAlertsIncidentAlertIDPrimary(v1IncidentsIncidentIDAlertsIncidentAlertIDPrimary)
	return o
}

// SetV1IncidentsIncidentIDAlertsIncidentAlertIDPrimary adds the v1IncidentsIncidentIdAlertsIncidentAlertIdPrimary to the patch v1 incidents incident Id alerts incident alert Id primary params
func (o *PatchV1IncidentsIncidentIDAlertsIncidentAlertIDPrimaryParams) SetV1IncidentsIncidentIDAlertsIncidentAlertIDPrimary(v1IncidentsIncidentIDAlertsIncidentAlertIDPrimary *models.PatchV1IncidentsIncidentIDAlertsIncidentAlertIDPrimary) {
	o.V1IncidentsIncidentIDAlertsIncidentAlertIDPrimary = v1IncidentsIncidentIDAlertsIncidentAlertIDPrimary
}

// WithIncidentAlertID adds the incidentAlertID to the patch v1 incidents incident Id alerts incident alert Id primary params
func (o *PatchV1IncidentsIncidentIDAlertsIncidentAlertIDPrimaryParams) WithIncidentAlertID(incidentAlertID string) *PatchV1IncidentsIncidentIDAlertsIncidentAlertIDPrimaryParams {
	o.SetIncidentAlertID(incidentAlertID)
	return o
}

// SetIncidentAlertID adds the incidentAlertId to the patch v1 incidents incident Id alerts incident alert Id primary params
func (o *PatchV1IncidentsIncidentIDAlertsIncidentAlertIDPrimaryParams) SetIncidentAlertID(incidentAlertID string) {
	o.IncidentAlertID = incidentAlertID
}

// WithIncidentID adds the incidentID to the patch v1 incidents incident Id alerts incident alert Id primary params
func (o *PatchV1IncidentsIncidentIDAlertsIncidentAlertIDPrimaryParams) WithIncidentID(incidentID string) *PatchV1IncidentsIncidentIDAlertsIncidentAlertIDPrimaryParams {
	o.SetIncidentID(incidentID)
	return o
}

// SetIncidentID adds the incidentId to the patch v1 incidents incident Id alerts incident alert Id primary params
func (o *PatchV1IncidentsIncidentIDAlertsIncidentAlertIDPrimaryParams) SetIncidentID(incidentID string) {
	o.IncidentID = incidentID
}

// WriteToRequest writes these params to a swagger request
func (o *PatchV1IncidentsIncidentIDAlertsIncidentAlertIDPrimaryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.V1IncidentsIncidentIDAlertsIncidentAlertIDPrimary != nil {
		if err := r.SetBodyParam(o.V1IncidentsIncidentIDAlertsIncidentAlertIDPrimary); err != nil {
			return err
		}
	}

	// path param incident_alert_id
	if err := r.SetPathParam("incident_alert_id", o.IncidentAlertID); err != nil {
		return err
	}

	// path param incident_id
	if err := r.SetPathParam("incident_id", o.IncidentID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
