// Code generated by go-swagger; DO NOT EDIT.

package run_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	run_model "github.com/kubeflow/pipelines/backend/api/v1/go_http_client/run_model"
)

// DeleteRunReader is a Reader for the DeleteRun structure.
type DeleteRunReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteRunReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteRunOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewDeleteRunDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteRunOK creates a DeleteRunOK with default headers values
func NewDeleteRunOK() *DeleteRunOK {
	return &DeleteRunOK{}
}

/*DeleteRunOK handles this case with default header values.

A successful response.
*/
type DeleteRunOK struct {
	Payload interface{}
}

func (o *DeleteRunOK) Error() string {
	return fmt.Sprintf("[DELETE /apis/v1/runs/{id}][%d] deleteRunOK  %+v", 200, o.Payload)
}

func (o *DeleteRunOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRunDefault creates a DeleteRunDefault with default headers values
func NewDeleteRunDefault(code int) *DeleteRunDefault {
	return &DeleteRunDefault{
		_statusCode: code,
	}
}

/*DeleteRunDefault handles this case with default header values.

DeleteRunDefault delete run default
*/
type DeleteRunDefault struct {
	_statusCode int

	Payload *run_model.V1Status
}

// Code gets the status code for the delete run default response
func (o *DeleteRunDefault) Code() int {
	return o._statusCode
}

func (o *DeleteRunDefault) Error() string {
	return fmt.Sprintf("[DELETE /apis/v1/runs/{id}][%d] DeleteRun default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteRunDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(run_model.V1Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}