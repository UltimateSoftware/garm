// Code generated by go-swagger; DO NOT EDIT.

package enterprises

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

// ListEnterpriseInstancesReader is a Reader for the ListEnterpriseInstances structure.
type ListEnterpriseInstancesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListEnterpriseInstancesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListEnterpriseInstancesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListEnterpriseInstancesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListEnterpriseInstancesOK creates a ListEnterpriseInstancesOK with default headers values
func NewListEnterpriseInstancesOK() *ListEnterpriseInstancesOK {
	return &ListEnterpriseInstancesOK{}
}

/*
ListEnterpriseInstancesOK describes a response with status code 200, with default header values.

Instances
*/
type ListEnterpriseInstancesOK struct {
	Payload garm_params.Instances
}

// IsSuccess returns true when this list enterprise instances o k response has a 2xx status code
func (o *ListEnterpriseInstancesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list enterprise instances o k response has a 3xx status code
func (o *ListEnterpriseInstancesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list enterprise instances o k response has a 4xx status code
func (o *ListEnterpriseInstancesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list enterprise instances o k response has a 5xx status code
func (o *ListEnterpriseInstancesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list enterprise instances o k response a status code equal to that given
func (o *ListEnterpriseInstancesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list enterprise instances o k response
func (o *ListEnterpriseInstancesOK) Code() int {
	return 200
}

func (o *ListEnterpriseInstancesOK) Error() string {
	return fmt.Sprintf("[GET /enterprises/{enterpriseID}/instances][%d] listEnterpriseInstancesOK  %+v", 200, o.Payload)
}

func (o *ListEnterpriseInstancesOK) String() string {
	return fmt.Sprintf("[GET /enterprises/{enterpriseID}/instances][%d] listEnterpriseInstancesOK  %+v", 200, o.Payload)
}

func (o *ListEnterpriseInstancesOK) GetPayload() garm_params.Instances {
	return o.Payload
}

func (o *ListEnterpriseInstancesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListEnterpriseInstancesDefault creates a ListEnterpriseInstancesDefault with default headers values
func NewListEnterpriseInstancesDefault(code int) *ListEnterpriseInstancesDefault {
	return &ListEnterpriseInstancesDefault{
		_statusCode: code,
	}
}

/*
ListEnterpriseInstancesDefault describes a response with status code -1, with default header values.

APIErrorResponse
*/
type ListEnterpriseInstancesDefault struct {
	_statusCode int

	Payload apiserver_params.APIErrorResponse
}

// IsSuccess returns true when this list enterprise instances default response has a 2xx status code
func (o *ListEnterpriseInstancesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list enterprise instances default response has a 3xx status code
func (o *ListEnterpriseInstancesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list enterprise instances default response has a 4xx status code
func (o *ListEnterpriseInstancesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list enterprise instances default response has a 5xx status code
func (o *ListEnterpriseInstancesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list enterprise instances default response a status code equal to that given
func (o *ListEnterpriseInstancesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list enterprise instances default response
func (o *ListEnterpriseInstancesDefault) Code() int {
	return o._statusCode
}

func (o *ListEnterpriseInstancesDefault) Error() string {
	return fmt.Sprintf("[GET /enterprises/{enterpriseID}/instances][%d] ListEnterpriseInstances default  %+v", o._statusCode, o.Payload)
}

func (o *ListEnterpriseInstancesDefault) String() string {
	return fmt.Sprintf("[GET /enterprises/{enterpriseID}/instances][%d] ListEnterpriseInstances default  %+v", o._statusCode, o.Payload)
}

func (o *ListEnterpriseInstancesDefault) GetPayload() apiserver_params.APIErrorResponse {
	return o.Payload
}

func (o *ListEnterpriseInstancesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
