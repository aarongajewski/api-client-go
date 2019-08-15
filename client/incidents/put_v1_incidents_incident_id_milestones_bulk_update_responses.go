// Code generated by go-swagger; DO NOT EDIT.

package incidents

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/firehydrant/api-client-go/models"
)

// PutV1IncidentsIncidentIDMilestonesBulkUpdateReader is a Reader for the PutV1IncidentsIncidentIDMilestonesBulkUpdate structure.
type PutV1IncidentsIncidentIDMilestonesBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutV1IncidentsIncidentIDMilestonesBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPutV1IncidentsIncidentIDMilestonesBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutV1IncidentsIncidentIDMilestonesBulkUpdateOK creates a PutV1IncidentsIncidentIDMilestonesBulkUpdateOK with default headers values
func NewPutV1IncidentsIncidentIDMilestonesBulkUpdateOK() *PutV1IncidentsIncidentIDMilestonesBulkUpdateOK {
	return &PutV1IncidentsIncidentIDMilestonesBulkUpdateOK{}
}

/*PutV1IncidentsIncidentIDMilestonesBulkUpdateOK handles this case with default header values.

Update a list of milestones on an incident
*/
type PutV1IncidentsIncidentIDMilestonesBulkUpdateOK struct {
	Payload *models.MilestoneEntityPaginated
}

func (o *PutV1IncidentsIncidentIDMilestonesBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /v1/incidents/{incident_id}/milestones/bulk_update][%d] putV1IncidentsIncidentIdMilestonesBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *PutV1IncidentsIncidentIDMilestonesBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MilestoneEntityPaginated)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
