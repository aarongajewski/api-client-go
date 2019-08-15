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

// PostV1IncidentsIncidentIDRoleAssignmentsReader is a Reader for the PostV1IncidentsIncidentIDRoleAssignments structure.
type PostV1IncidentsIncidentIDRoleAssignmentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostV1IncidentsIncidentIDRoleAssignmentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostV1IncidentsIncidentIDRoleAssignmentsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostV1IncidentsIncidentIDRoleAssignmentsCreated creates a PostV1IncidentsIncidentIDRoleAssignmentsCreated with default headers values
func NewPostV1IncidentsIncidentIDRoleAssignmentsCreated() *PostV1IncidentsIncidentIDRoleAssignmentsCreated {
	return &PostV1IncidentsIncidentIDRoleAssignmentsCreated{}
}

/*PostV1IncidentsIncidentIDRoleAssignmentsCreated handles this case with default header values.

Assign a role to a user for this incident. All tasks associated to the role will also automatically be attached
*/
type PostV1IncidentsIncidentIDRoleAssignmentsCreated struct {
	Payload *models.RoleAssignmentEntity
}

func (o *PostV1IncidentsIncidentIDRoleAssignmentsCreated) Error() string {
	return fmt.Sprintf("[POST /v1/incidents/{incident_id}/role_assignments][%d] postV1IncidentsIncidentIdRoleAssignmentsCreated  %+v", 201, o.Payload)
}

func (o *PostV1IncidentsIncidentIDRoleAssignmentsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RoleAssignmentEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
