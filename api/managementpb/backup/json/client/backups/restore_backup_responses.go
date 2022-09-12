// Code generated by go-swagger; DO NOT EDIT.

package backups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RestoreBackupReader is a Reader for the RestoreBackup structure.
type RestoreBackupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RestoreBackupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRestoreBackupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewRestoreBackupDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRestoreBackupOK creates a RestoreBackupOK with default headers values
func NewRestoreBackupOK() *RestoreBackupOK {
	return &RestoreBackupOK{}
}

/* RestoreBackupOK describes a response with status code 200, with default header values.

A successful response.
*/
type RestoreBackupOK struct {
	Payload *RestoreBackupOKBody
}

// IsSuccess returns true when this restore backup Ok response has a 2xx status code
func (o *RestoreBackupOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this restore backup Ok response has a 3xx status code
func (o *RestoreBackupOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this restore backup Ok response has a 4xx status code
func (o *RestoreBackupOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this restore backup Ok response has a 5xx status code
func (o *RestoreBackupOK) IsServerError() bool {
	return false
}

// IsCode returns true when this restore backup Ok response a status code equal to that given
func (o *RestoreBackupOK) IsCode(code int) bool {
	return code == 200
}

func (o *RestoreBackupOK) Error() string {
	return fmt.Sprintf("[POST /v1/management/backup/Backups/Restore][%d] restoreBackupOk  %+v", 200, o.Payload)
}

func (o *RestoreBackupOK) String() string {
	return fmt.Sprintf("[POST /v1/management/backup/Backups/Restore][%d] restoreBackupOk  %+v", 200, o.Payload)
}

func (o *RestoreBackupOK) GetPayload() *RestoreBackupOKBody {
	return o.Payload
}

func (o *RestoreBackupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(RestoreBackupOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRestoreBackupDefault creates a RestoreBackupDefault with default headers values
func NewRestoreBackupDefault(code int) *RestoreBackupDefault {
	return &RestoreBackupDefault{
		_statusCode: code,
	}
}

/* RestoreBackupDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type RestoreBackupDefault struct {
	_statusCode int

	Payload *RestoreBackupDefaultBody
}

// Code gets the status code for the restore backup default response
func (o *RestoreBackupDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this restore backup default response has a 2xx status code
func (o *RestoreBackupDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this restore backup default response has a 3xx status code
func (o *RestoreBackupDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this restore backup default response has a 4xx status code
func (o *RestoreBackupDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this restore backup default response has a 5xx status code
func (o *RestoreBackupDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this restore backup default response a status code equal to that given
func (o *RestoreBackupDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *RestoreBackupDefault) Error() string {
	return fmt.Sprintf("[POST /v1/management/backup/Backups/Restore][%d] RestoreBackup default  %+v", o._statusCode, o.Payload)
}

func (o *RestoreBackupDefault) String() string {
	return fmt.Sprintf("[POST /v1/management/backup/Backups/Restore][%d] RestoreBackup default  %+v", o._statusCode, o.Payload)
}

func (o *RestoreBackupDefault) GetPayload() *RestoreBackupDefaultBody {
	return o.Payload
}

func (o *RestoreBackupDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(RestoreBackupDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*RestoreBackupBody restore backup body
swagger:model RestoreBackupBody
*/
type RestoreBackupBody struct {
	// Service identifier where backup should be restored.
	ServiceID string `json:"service_id,omitempty"`

	// Artifact id to restore.
	ArtifactID string `json:"artifact_id,omitempty"`
}

// Validate validates this restore backup body
func (o *RestoreBackupBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this restore backup body based on context it is used
func (o *RestoreBackupBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *RestoreBackupBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RestoreBackupBody) UnmarshalBinary(b []byte) error {
	var res RestoreBackupBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*RestoreBackupDefaultBody restore backup default body
swagger:model RestoreBackupDefaultBody
*/
type RestoreBackupDefaultBody struct {
	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*RestoreBackupDefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this restore backup default body
func (o *RestoreBackupDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *RestoreBackupDefaultBody) validateDetails(formats strfmt.Registry) error {
	if swag.IsZero(o.Details) { // not required
		return nil
	}

	for i := 0; i < len(o.Details); i++ {
		if swag.IsZero(o.Details[i]) { // not required
			continue
		}

		if o.Details[i] != nil {
			if err := o.Details[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("RestoreBackup default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("RestoreBackup default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this restore backup default body based on the context it is used
func (o *RestoreBackupDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *RestoreBackupDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.Details); i++ {
		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("RestoreBackup default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("RestoreBackup default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *RestoreBackupDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RestoreBackupDefaultBody) UnmarshalBinary(b []byte) error {
	var res RestoreBackupDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*RestoreBackupDefaultBodyDetailsItems0 restore backup default body details items0
swagger:model RestoreBackupDefaultBodyDetailsItems0
*/
type RestoreBackupDefaultBodyDetailsItems0 struct {
	// at type
	AtType string `json:"@type,omitempty"`
}

// Validate validates this restore backup default body details items0
func (o *RestoreBackupDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this restore backup default body details items0 based on context it is used
func (o *RestoreBackupDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *RestoreBackupDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RestoreBackupDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res RestoreBackupDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*RestoreBackupOKBody restore backup OK body
swagger:model RestoreBackupOKBody
*/
type RestoreBackupOKBody struct {
	// Unique restore identifier.
	RestoreID string `json:"restore_id,omitempty"`
}

// Validate validates this restore backup OK body
func (o *RestoreBackupOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this restore backup OK body based on context it is used
func (o *RestoreBackupOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *RestoreBackupOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RestoreBackupOKBody) UnmarshalBinary(b []byte) error {
	var res RestoreBackupOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
