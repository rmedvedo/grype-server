// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Portshift/grype-server/api/client/models"
)

// PostScanSBOMReader is a Reader for the PostScanSBOM structure.
type PostScanSBOMReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostScanSBOMReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostScanSBOMOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostScanSBOMDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostScanSBOMOK creates a PostScanSBOMOK with default headers values
func NewPostScanSBOMOK() *PostScanSBOMOK {
	return &PostScanSBOMOK{}
}

/* PostScanSBOMOK describes a response with status code 200, with default header values.

Successful scan result of the provided SBOM
*/
type PostScanSBOMOK struct {
	Payload *models.Vulnerabilities
}

func (o *PostScanSBOMOK) Error() string {
	return fmt.Sprintf("[POST /scanSBOM][%d] postScanSBOMOK  %+v", 200, o.Payload)
}
func (o *PostScanSBOMOK) GetPayload() *models.Vulnerabilities {
	return o.Payload
}

func (o *PostScanSBOMOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Vulnerabilities)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostScanSBOMDefault creates a PostScanSBOMDefault with default headers values
func NewPostScanSBOMDefault(code int) *PostScanSBOMDefault {
	return &PostScanSBOMDefault{
		_statusCode: code,
	}
}

/* PostScanSBOMDefault describes a response with status code -1, with default header values.

unknown error
*/
type PostScanSBOMDefault struct {
	_statusCode int

	Payload *models.APIResponse
}

// Code gets the status code for the post scan s b o m default response
func (o *PostScanSBOMDefault) Code() int {
	return o._statusCode
}

func (o *PostScanSBOMDefault) Error() string {
	return fmt.Sprintf("[POST /scanSBOM][%d] PostScanSBOM default  %+v", o._statusCode, o.Payload)
}
func (o *PostScanSBOMDefault) GetPayload() *models.APIResponse {
	return o.Payload
}

func (o *PostScanSBOMDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
