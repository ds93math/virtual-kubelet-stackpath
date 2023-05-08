// Code generated by go-swagger; DO NOT EDIT.

package workloads

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/stackpath/vk-stackpath-provider/internal/api/workload/workload_models"
)

// GetWorkloadsReader is a Reader for the GetWorkloads structure.
type GetWorkloadsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWorkloadsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWorkloadsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetWorkloadsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetWorkloadsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetWorkloadsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetWorkloadsOK creates a GetWorkloadsOK with default headers values
func NewGetWorkloadsOK() *GetWorkloadsOK {
	return &GetWorkloadsOK{}
}

/*
GetWorkloadsOK describes a response with status code 200, with default header values.

GetWorkloadsOK get workloads o k
*/
type GetWorkloadsOK struct {
	Payload *workload_models.V1GetWorkloadsResponse
}

// IsSuccess returns true when this get workloads o k response has a 2xx status code
func (o *GetWorkloadsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get workloads o k response has a 3xx status code
func (o *GetWorkloadsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workloads o k response has a 4xx status code
func (o *GetWorkloadsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get workloads o k response has a 5xx status code
func (o *GetWorkloadsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get workloads o k response a status code equal to that given
func (o *GetWorkloadsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get workloads o k response
func (o *GetWorkloadsOK) Code() int {
	return 200
}

func (o *GetWorkloadsOK) Error() string {
	return fmt.Sprintf("[GET /workload/v1/stacks/{stack_id}/workloads][%d] getWorkloadsOK  %+v", 200, o.Payload)
}

func (o *GetWorkloadsOK) String() string {
	return fmt.Sprintf("[GET /workload/v1/stacks/{stack_id}/workloads][%d] getWorkloadsOK  %+v", 200, o.Payload)
}

func (o *GetWorkloadsOK) GetPayload() *workload_models.V1GetWorkloadsResponse {
	return o.Payload
}

func (o *GetWorkloadsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(workload_models.V1GetWorkloadsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkloadsUnauthorized creates a GetWorkloadsUnauthorized with default headers values
func NewGetWorkloadsUnauthorized() *GetWorkloadsUnauthorized {
	return &GetWorkloadsUnauthorized{}
}

/*
GetWorkloadsUnauthorized describes a response with status code 401, with default header values.

Returned when an unauthorized request is attempted.
*/
type GetWorkloadsUnauthorized struct {
	Payload *workload_models.StackpathapiStatus
}

// IsSuccess returns true when this get workloads unauthorized response has a 2xx status code
func (o *GetWorkloadsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workloads unauthorized response has a 3xx status code
func (o *GetWorkloadsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workloads unauthorized response has a 4xx status code
func (o *GetWorkloadsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workloads unauthorized response has a 5xx status code
func (o *GetWorkloadsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get workloads unauthorized response a status code equal to that given
func (o *GetWorkloadsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get workloads unauthorized response
func (o *GetWorkloadsUnauthorized) Code() int {
	return 401
}

func (o *GetWorkloadsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /workload/v1/stacks/{stack_id}/workloads][%d] getWorkloadsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetWorkloadsUnauthorized) String() string {
	return fmt.Sprintf("[GET /workload/v1/stacks/{stack_id}/workloads][%d] getWorkloadsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetWorkloadsUnauthorized) GetPayload() *workload_models.StackpathapiStatus {
	return o.Payload
}

func (o *GetWorkloadsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(workload_models.StackpathapiStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkloadsInternalServerError creates a GetWorkloadsInternalServerError with default headers values
func NewGetWorkloadsInternalServerError() *GetWorkloadsInternalServerError {
	return &GetWorkloadsInternalServerError{}
}

/*
GetWorkloadsInternalServerError describes a response with status code 500, with default header values.

Internal server error.
*/
type GetWorkloadsInternalServerError struct {
	Payload *workload_models.StackpathapiStatus
}

// IsSuccess returns true when this get workloads internal server error response has a 2xx status code
func (o *GetWorkloadsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workloads internal server error response has a 3xx status code
func (o *GetWorkloadsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workloads internal server error response has a 4xx status code
func (o *GetWorkloadsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get workloads internal server error response has a 5xx status code
func (o *GetWorkloadsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get workloads internal server error response a status code equal to that given
func (o *GetWorkloadsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get workloads internal server error response
func (o *GetWorkloadsInternalServerError) Code() int {
	return 500
}

func (o *GetWorkloadsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /workload/v1/stacks/{stack_id}/workloads][%d] getWorkloadsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetWorkloadsInternalServerError) String() string {
	return fmt.Sprintf("[GET /workload/v1/stacks/{stack_id}/workloads][%d] getWorkloadsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetWorkloadsInternalServerError) GetPayload() *workload_models.StackpathapiStatus {
	return o.Payload
}

func (o *GetWorkloadsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(workload_models.StackpathapiStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkloadsDefault creates a GetWorkloadsDefault with default headers values
func NewGetWorkloadsDefault(code int) *GetWorkloadsDefault {
	return &GetWorkloadsDefault{
		_statusCode: code,
	}
}

/*
GetWorkloadsDefault describes a response with status code -1, with default header values.

Default error structure.
*/
type GetWorkloadsDefault struct {
	_statusCode int

	Payload *workload_models.StackpathapiStatus
}

// IsSuccess returns true when this get workloads default response has a 2xx status code
func (o *GetWorkloadsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get workloads default response has a 3xx status code
func (o *GetWorkloadsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get workloads default response has a 4xx status code
func (o *GetWorkloadsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get workloads default response has a 5xx status code
func (o *GetWorkloadsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get workloads default response a status code equal to that given
func (o *GetWorkloadsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get workloads default response
func (o *GetWorkloadsDefault) Code() int {
	return o._statusCode
}

func (o *GetWorkloadsDefault) Error() string {
	return fmt.Sprintf("[GET /workload/v1/stacks/{stack_id}/workloads][%d] GetWorkloads default  %+v", o._statusCode, o.Payload)
}

func (o *GetWorkloadsDefault) String() string {
	return fmt.Sprintf("[GET /workload/v1/stacks/{stack_id}/workloads][%d] GetWorkloads default  %+v", o._statusCode, o.Payload)
}

func (o *GetWorkloadsDefault) GetPayload() *workload_models.StackpathapiStatus {
	return o.Payload
}

func (o *GetWorkloadsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(workload_models.StackpathapiStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
