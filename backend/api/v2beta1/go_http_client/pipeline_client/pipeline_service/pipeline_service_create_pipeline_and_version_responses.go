// Code generated by go-swagger; DO NOT EDIT.

package pipeline_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	pipeline_model "github.com/kubeflow/pipelines/backend/api/v2beta1/go_http_client/pipeline_model"
)

// PipelineServiceCreatePipelineAndVersionReader is a Reader for the PipelineServiceCreatePipelineAndVersion structure.
type PipelineServiceCreatePipelineAndVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PipelineServiceCreatePipelineAndVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPipelineServiceCreatePipelineAndVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPipelineServiceCreatePipelineAndVersionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPipelineServiceCreatePipelineAndVersionOK creates a PipelineServiceCreatePipelineAndVersionOK with default headers values
func NewPipelineServiceCreatePipelineAndVersionOK() *PipelineServiceCreatePipelineAndVersionOK {
	return &PipelineServiceCreatePipelineAndVersionOK{}
}

/*PipelineServiceCreatePipelineAndVersionOK handles this case with default header values.

A successful response.
*/
type PipelineServiceCreatePipelineAndVersionOK struct {
	Payload *pipeline_model.V2beta1Pipeline
}

func (o *PipelineServiceCreatePipelineAndVersionOK) Error() string {
	return fmt.Sprintf("[POST /apis/v2beta1/pipelines/create][%d] pipelineServiceCreatePipelineAndVersionOK  %+v", 200, o.Payload)
}

func (o *PipelineServiceCreatePipelineAndVersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(pipeline_model.V2beta1Pipeline)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPipelineServiceCreatePipelineAndVersionDefault creates a PipelineServiceCreatePipelineAndVersionDefault with default headers values
func NewPipelineServiceCreatePipelineAndVersionDefault(code int) *PipelineServiceCreatePipelineAndVersionDefault {
	return &PipelineServiceCreatePipelineAndVersionDefault{
		_statusCode: code,
	}
}

/*PipelineServiceCreatePipelineAndVersionDefault handles this case with default header values.

An unexpected error response.
*/
type PipelineServiceCreatePipelineAndVersionDefault struct {
	_statusCode int

	Payload *pipeline_model.RuntimeError
}

// Code gets the status code for the pipeline service create pipeline and version default response
func (o *PipelineServiceCreatePipelineAndVersionDefault) Code() int {
	return o._statusCode
}

func (o *PipelineServiceCreatePipelineAndVersionDefault) Error() string {
	return fmt.Sprintf("[POST /apis/v2beta1/pipelines/create][%d] PipelineService_CreatePipelineAndVersion default  %+v", o._statusCode, o.Payload)
}

func (o *PipelineServiceCreatePipelineAndVersionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(pipeline_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
