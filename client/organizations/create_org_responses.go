// Code generated by go-swagger; DO NOT EDIT.

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	apiserver_params "github.com/cloudbase/garm/apiserver/params"
	garm_params "github.com/cloudbase/garm/params"
)

// CreateOrgReader is a Reader for the CreateOrg structure.
type CreateOrgReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateOrgReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateOrgOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateOrgDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateOrgOK creates a CreateOrgOK with default headers values
func NewCreateOrgOK() *CreateOrgOK {
	return &CreateOrgOK{}
}

/*
CreateOrgOK describes a response with status code 200, with default header values.

Organization
*/
type CreateOrgOK struct {
	Payload garm_params.Organization
}

// IsSuccess returns true when this create org o k response has a 2xx status code
func (o *CreateOrgOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create org o k response has a 3xx status code
func (o *CreateOrgOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create org o k response has a 4xx status code
func (o *CreateOrgOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create org o k response has a 5xx status code
func (o *CreateOrgOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create org o k response a status code equal to that given
func (o *CreateOrgOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create org o k response
func (o *CreateOrgOK) Code() int {
	return 200
}

func (o *CreateOrgOK) Error() string {
	return fmt.Sprintf("[POST /organizations][%d] createOrgOK  %+v", 200, o.Payload)
}

func (o *CreateOrgOK) String() string {
	return fmt.Sprintf("[POST /organizations][%d] createOrgOK  %+v", 200, o.Payload)
}

func (o *CreateOrgOK) GetPayload() garm_params.Organization {
	return o.Payload
}

func (o *CreateOrgOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateOrgDefault creates a CreateOrgDefault with default headers values
func NewCreateOrgDefault(code int) *CreateOrgDefault {
	return &CreateOrgDefault{
		_statusCode: code,
	}
}

/*
CreateOrgDefault describes a response with status code -1, with default header values.

APIErrorResponse
*/
type CreateOrgDefault struct {
	_statusCode int

	Payload apiserver_params.APIErrorResponse
}

// IsSuccess returns true when this create org default response has a 2xx status code
func (o *CreateOrgDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create org default response has a 3xx status code
func (o *CreateOrgDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create org default response has a 4xx status code
func (o *CreateOrgDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create org default response has a 5xx status code
func (o *CreateOrgDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create org default response a status code equal to that given
func (o *CreateOrgDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create org default response
func (o *CreateOrgDefault) Code() int {
	return o._statusCode
}

func (o *CreateOrgDefault) Error() string {
	return fmt.Sprintf("[POST /organizations][%d] CreateOrg default  %+v", o._statusCode, o.Payload)
}

func (o *CreateOrgDefault) String() string {
	return fmt.Sprintf("[POST /organizations][%d] CreateOrg default  %+v", o._statusCode, o.Payload)
}

func (o *CreateOrgDefault) GetPayload() apiserver_params.APIErrorResponse {
	return o.Payload
}

func (o *CreateOrgDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
