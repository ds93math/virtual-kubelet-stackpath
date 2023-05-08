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

// CreateWorkloadReader is a Reader for the CreateWorkload structure.
type CreateWorkloadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateWorkloadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateWorkloadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateWorkloadUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateWorkloadInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateWorkloadDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateWorkloadOK creates a CreateWorkloadOK with default headers values
func NewCreateWorkloadOK() *CreateWorkloadOK {
	return &CreateWorkloadOK{}
}

/*
CreateWorkloadOK describes a response with status code 200, with default header values.

CreateWorkloadOK create workload o k
*/
type CreateWorkloadOK struct {
	Payload *workload_models.V1CreateWorkloadResponse
}

// IsSuccess returns true when this create workload o k response has a 2xx status code
func (o *CreateWorkloadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create workload o k response has a 3xx status code
func (o *CreateWorkloadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create workload o k response has a 4xx status code
func (o *CreateWorkloadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create workload o k response has a 5xx status code
func (o *CreateWorkloadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create workload o k response a status code equal to that given
func (o *CreateWorkloadOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create workload o k response
func (o *CreateWorkloadOK) Code() int {
	return 200
}

func (o *CreateWorkloadOK) Error() string {
	return fmt.Sprintf("[POST /workload/v1/stacks/{stack_id}/workloads][%d] createWorkloadOK  %+v", 200, o.Payload)
}

func (o *CreateWorkloadOK) String() string {
	return fmt.Sprintf("[POST /workload/v1/stacks/{stack_id}/workloads][%d] createWorkloadOK  %+v", 200, o.Payload)
}

func (o *CreateWorkloadOK) GetPayload() *workload_models.V1CreateWorkloadResponse {
	return o.Payload
}

func (o *CreateWorkloadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(workload_models.V1CreateWorkloadResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateWorkloadUnauthorized creates a CreateWorkloadUnauthorized with default headers values
func NewCreateWorkloadUnauthorized() *CreateWorkloadUnauthorized {
	return &CreateWorkloadUnauthorized{}
}

/*
CreateWorkloadUnauthorized describes a response with status code 401, with default header values.

Returned when an unauthorized request is attempted.
*/
type CreateWorkloadUnauthorized struct {
	Payload *workload_models.StackpathapiStatus
}

// IsSuccess returns true when this create workload unauthorized response has a 2xx status code
func (o *CreateWorkloadUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create workload unauthorized response has a 3xx status code
func (o *CreateWorkloadUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create workload unauthorized response has a 4xx status code
func (o *CreateWorkloadUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create workload unauthorized response has a 5xx status code
func (o *CreateWorkloadUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create workload unauthorized response a status code equal to that given
func (o *CreateWorkloadUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the create workload unauthorized response
func (o *CreateWorkloadUnauthorized) Code() int {
	return 401
}

func (o *CreateWorkloadUnauthorized) Error() string {
	return fmt.Sprintf("[POST /workload/v1/stacks/{stack_id}/workloads][%d] createWorkloadUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateWorkloadUnauthorized) String() string {
	return fmt.Sprintf("[POST /workload/v1/stacks/{stack_id}/workloads][%d] createWorkloadUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateWorkloadUnauthorized) GetPayload() *workload_models.StackpathapiStatus {
	return o.Payload
}

func (o *CreateWorkloadUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(workload_models.StackpathapiStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateWorkloadInternalServerError creates a CreateWorkloadInternalServerError with default headers values
func NewCreateWorkloadInternalServerError() *CreateWorkloadInternalServerError {
	return &CreateWorkloadInternalServerError{}
}

/*
CreateWorkloadInternalServerError describes a response with status code 500, with default header values.

Internal server error.
*/
type CreateWorkloadInternalServerError struct {
	Payload *workload_models.StackpathapiStatus
}

// IsSuccess returns true when this create workload internal server error response has a 2xx status code
func (o *CreateWorkloadInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create workload internal server error response has a 3xx status code
func (o *CreateWorkloadInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create workload internal server error response has a 4xx status code
func (o *CreateWorkloadInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create workload internal server error response has a 5xx status code
func (o *CreateWorkloadInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create workload internal server error response a status code equal to that given
func (o *CreateWorkloadInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the create workload internal server error response
func (o *CreateWorkloadInternalServerError) Code() int {
	return 500
}

func (o *CreateWorkloadInternalServerError) Error() string {
	return fmt.Sprintf("[POST /workload/v1/stacks/{stack_id}/workloads][%d] createWorkloadInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateWorkloadInternalServerError) String() string {
	return fmt.Sprintf("[POST /workload/v1/stacks/{stack_id}/workloads][%d] createWorkloadInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateWorkloadInternalServerError) GetPayload() *workload_models.StackpathapiStatus {
	return o.Payload
}

func (o *CreateWorkloadInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(workload_models.StackpathapiStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateWorkloadDefault creates a CreateWorkloadDefault with default headers values
func NewCreateWorkloadDefault(code int) *CreateWorkloadDefault {
	return &CreateWorkloadDefault{
		_statusCode: code,
	}
}

/*
CreateWorkloadDefault describes a response with status code -1, with default header values.

Default error structure.
*/
type CreateWorkloadDefault struct {
	_statusCode int

	Payload *workload_models.StackpathapiStatus
}

// IsSuccess returns true when this create workload default response has a 2xx status code
func (o *CreateWorkloadDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create workload default response has a 3xx status code
func (o *CreateWorkloadDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create workload default response has a 4xx status code
func (o *CreateWorkloadDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create workload default response has a 5xx status code
func (o *CreateWorkloadDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create workload default response a status code equal to that given
func (o *CreateWorkloadDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create workload default response
func (o *CreateWorkloadDefault) Code() int {
	return o._statusCode
}

func (o *CreateWorkloadDefault) Error() string {
	return fmt.Sprintf("[POST /workload/v1/stacks/{stack_id}/workloads][%d] CreateWorkload default  %+v", o._statusCode, o.Payload)
}

func (o *CreateWorkloadDefault) String() string {
	return fmt.Sprintf("[POST /workload/v1/stacks/{stack_id}/workloads][%d] CreateWorkload default  %+v", o._statusCode, o.Payload)
}

func (o *CreateWorkloadDefault) GetPayload() *workload_models.StackpathapiStatus {
	return o.Payload
}

func (o *CreateWorkloadDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(workload_models.StackpathapiStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
