// Code generated by go-swagger; DO NOT EDIT.

package teams

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/firehydrant/api-client-go/models"
)

// PostV1TeamsReader is a Reader for the PostV1Teams structure.
type PostV1TeamsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostV1TeamsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostV1TeamsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostV1TeamsCreated creates a PostV1TeamsCreated with default headers values
func NewPostV1TeamsCreated() *PostV1TeamsCreated {
	return &PostV1TeamsCreated{}
}

/*PostV1TeamsCreated handles this case with default header values.

Create a team
*/
type PostV1TeamsCreated struct {
	Payload *models.TeamEntity
}

func (o *PostV1TeamsCreated) Error() string {
	return fmt.Sprintf("[POST /v1/teams][%d] postV1TeamsCreated  %+v", 201, o.Payload)
}

func (o *PostV1TeamsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TeamEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
