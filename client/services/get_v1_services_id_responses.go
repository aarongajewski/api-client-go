// Code generated by go-swagger; DO NOT EDIT.

package services

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/firehydrant/api-client-go/models"
)

// GetV1ServicesIDReader is a Reader for the GetV1ServicesID structure.
type GetV1ServicesIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1ServicesIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetV1ServicesIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetV1ServicesIDOK creates a GetV1ServicesIDOK with default headers values
func NewGetV1ServicesIDOK() *GetV1ServicesIDOK {
	return &GetV1ServicesIDOK{}
}

/*GetV1ServicesIDOK handles this case with default header values.

Retrieve a single service
*/
type GetV1ServicesIDOK struct {
	Payload *models.ServiceEntity
}

func (o *GetV1ServicesIDOK) Error() string {
	return fmt.Sprintf("[GET /v1/services/{id}][%d] getV1ServicesIdOK  %+v", 200, o.Payload)
}

func (o *GetV1ServicesIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
