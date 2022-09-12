// Code generated by go-swagger; DO NOT EDIT.

package server

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// LogsReader is a Reader for the Logs structure.
type LogsReader struct {
	formats strfmt.Registry
	writer  io.Writer
}

// ReadResponse reads a server response into the received o.
func (o *LogsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLogsOK(o.writer)
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewLogsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewLogsOK creates a LogsOK with default headers values
func NewLogsOK(writer io.Writer) *LogsOK {
	return &LogsOK{
		Payload: writer,
	}
}

/* LogsOK describes a response with status code 200, with default header values.

A successful response.
*/
type LogsOK struct {
	Payload io.Writer
}

// IsSuccess returns true when this logs Ok response has a 2xx status code
func (o *LogsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this logs Ok response has a 3xx status code
func (o *LogsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this logs Ok response has a 4xx status code
func (o *LogsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this logs Ok response has a 5xx status code
func (o *LogsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this logs Ok response a status code equal to that given
func (o *LogsOK) IsCode(code int) bool {
	return code == 200
}

func (o *LogsOK) Error() string {
	return fmt.Sprintf("[GET /logs.zip][%d] logsOk  %+v", 200, o.Payload)
}

func (o *LogsOK) String() string {
	return fmt.Sprintf("[GET /logs.zip][%d] logsOk  %+v", 200, o.Payload)
}

func (o *LogsOK) GetPayload() io.Writer {
	return o.Payload
}

func (o *LogsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLogsDefault creates a LogsDefault with default headers values
func NewLogsDefault(code int) *LogsDefault {
	return &LogsDefault{
		_statusCode: code,
	}
}

/* LogsDefault describes a response with status code -1, with default header values.

An error response.
*/
type LogsDefault struct {
	_statusCode int

	Payload *LogsDefaultBody
}

// Code gets the status code for the logs default response
func (o *LogsDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this logs default response has a 2xx status code
func (o *LogsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this logs default response has a 3xx status code
func (o *LogsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this logs default response has a 4xx status code
func (o *LogsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this logs default response has a 5xx status code
func (o *LogsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this logs default response a status code equal to that given
func (o *LogsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *LogsDefault) Error() string {
	return fmt.Sprintf("[GET /logs.zip][%d] Logs default  %+v", o._statusCode, o.Payload)
}

func (o *LogsDefault) String() string {
	return fmt.Sprintf("[GET /logs.zip][%d] Logs default  %+v", o._statusCode, o.Payload)
}

func (o *LogsDefault) GetPayload() *LogsDefaultBody {
	return o.Payload
}

func (o *LogsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(LogsDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*LogsDefaultBody ErrorResponse is a message returned on HTTP error.
swagger:model LogsDefaultBody
*/
type LogsDefaultBody struct {
	// code
	Code int32 `json:"code,omitempty"`

	// error
	Error string `json:"error,omitempty"`

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this logs default body
func (o *LogsDefaultBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this logs default body based on context it is used
func (o *LogsDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *LogsDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LogsDefaultBody) UnmarshalBinary(b []byte) error {
	var res LogsDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
