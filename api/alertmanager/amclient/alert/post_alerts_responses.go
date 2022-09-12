// Code generated by go-swagger; DO NOT EDIT.

package alert

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PostAlertsReader is a Reader for the PostAlerts structure.
type PostAlertsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAlertsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostAlertsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostAlertsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostAlertsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostAlertsOK creates a PostAlertsOK with default headers values
func NewPostAlertsOK() *PostAlertsOK {
	return &PostAlertsOK{}
}

/* PostAlertsOK describes a response with status code 200, with default header values.

Create alerts response
*/
type PostAlertsOK struct{}

// IsSuccess returns true when this post alerts o k response has a 2xx status code
func (o *PostAlertsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post alerts o k response has a 3xx status code
func (o *PostAlertsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post alerts o k response has a 4xx status code
func (o *PostAlertsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post alerts o k response has a 5xx status code
func (o *PostAlertsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post alerts o k response a status code equal to that given
func (o *PostAlertsOK) IsCode(code int) bool {
	return code == 200
}

func (o *PostAlertsOK) Error() string {
	return fmt.Sprintf("[POST /alerts][%d] postAlertsOK ", 200)
}

func (o *PostAlertsOK) String() string {
	return fmt.Sprintf("[POST /alerts][%d] postAlertsOK ", 200)
}

func (o *PostAlertsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	return nil
}

// NewPostAlertsBadRequest creates a PostAlertsBadRequest with default headers values
func NewPostAlertsBadRequest() *PostAlertsBadRequest {
	return &PostAlertsBadRequest{}
}

/* PostAlertsBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type PostAlertsBadRequest struct {
	Payload string
}

// IsSuccess returns true when this post alerts bad request response has a 2xx status code
func (o *PostAlertsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post alerts bad request response has a 3xx status code
func (o *PostAlertsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post alerts bad request response has a 4xx status code
func (o *PostAlertsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post alerts bad request response has a 5xx status code
func (o *PostAlertsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post alerts bad request response a status code equal to that given
func (o *PostAlertsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostAlertsBadRequest) Error() string {
	return fmt.Sprintf("[POST /alerts][%d] postAlertsBadRequest  %+v", 400, o.Payload)
}

func (o *PostAlertsBadRequest) String() string {
	return fmt.Sprintf("[POST /alerts][%d] postAlertsBadRequest  %+v", 400, o.Payload)
}

func (o *PostAlertsBadRequest) GetPayload() string {
	return o.Payload
}

func (o *PostAlertsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAlertsInternalServerError creates a PostAlertsInternalServerError with default headers values
func NewPostAlertsInternalServerError() *PostAlertsInternalServerError {
	return &PostAlertsInternalServerError{}
}

/* PostAlertsInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type PostAlertsInternalServerError struct {
	Payload string
}

// IsSuccess returns true when this post alerts internal server error response has a 2xx status code
func (o *PostAlertsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post alerts internal server error response has a 3xx status code
func (o *PostAlertsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post alerts internal server error response has a 4xx status code
func (o *PostAlertsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post alerts internal server error response has a 5xx status code
func (o *PostAlertsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post alerts internal server error response a status code equal to that given
func (o *PostAlertsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostAlertsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /alerts][%d] postAlertsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostAlertsInternalServerError) String() string {
	return fmt.Sprintf("[POST /alerts][%d] postAlertsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostAlertsInternalServerError) GetPayload() string {
	return o.Payload
}

func (o *PostAlertsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
