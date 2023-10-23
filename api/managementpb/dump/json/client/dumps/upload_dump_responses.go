// Code generated by go-swagger; DO NOT EDIT.

package dumps

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

// UploadDumpReader is a Reader for the UploadDump structure.
type UploadDumpReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UploadDumpReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUploadDumpOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUploadDumpDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUploadDumpOK creates a UploadDumpOK with default headers values
func NewUploadDumpOK() *UploadDumpOK {
	return &UploadDumpOK{}
}

/*
UploadDumpOK describes a response with status code 200, with default header values.

A successful response.
*/
type UploadDumpOK struct {
	Payload interface{}
}

func (o *UploadDumpOK) Error() string {
	return fmt.Sprintf("[POST /v1/management/dump/Dumps/Upload][%d] uploadDumpOk  %+v", 200, o.Payload)
}

func (o *UploadDumpOK) GetPayload() interface{} {
	return o.Payload
}

func (o *UploadDumpOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUploadDumpDefault creates a UploadDumpDefault with default headers values
func NewUploadDumpDefault(code int) *UploadDumpDefault {
	return &UploadDumpDefault{
		_statusCode: code,
	}
}

/*
UploadDumpDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type UploadDumpDefault struct {
	_statusCode int

	Payload *UploadDumpDefaultBody
}

// Code gets the status code for the upload dump default response
func (o *UploadDumpDefault) Code() int {
	return o._statusCode
}

func (o *UploadDumpDefault) Error() string {
	return fmt.Sprintf("[POST /v1/management/dump/Dumps/Upload][%d] UploadDump default  %+v", o._statusCode, o.Payload)
}

func (o *UploadDumpDefault) GetPayload() *UploadDumpDefaultBody {
	return o.Payload
}

func (o *UploadDumpDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(UploadDumpDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
UploadDumpBody upload dump body
swagger:model UploadDumpBody
*/
type UploadDumpBody struct {
	// dump ids
	DumpIds []string `json:"dump_ids"`

	// ftp parameters
	FtpParameters *UploadDumpParamsBodyFtpParameters `json:"ftp_parameters,omitempty"`
}

// Validate validates this upload dump body
func (o *UploadDumpBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateFtpParameters(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UploadDumpBody) validateFtpParameters(formats strfmt.Registry) error {
	if swag.IsZero(o.FtpParameters) { // not required
		return nil
	}

	if o.FtpParameters != nil {
		if err := o.FtpParameters.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "ftp_parameters")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "ftp_parameters")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this upload dump body based on the context it is used
func (o *UploadDumpBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateFtpParameters(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UploadDumpBody) contextValidateFtpParameters(ctx context.Context, formats strfmt.Registry) error {
	if o.FtpParameters != nil {
		if err := o.FtpParameters.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "ftp_parameters")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "ftp_parameters")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UploadDumpBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UploadDumpBody) UnmarshalBinary(b []byte) error {
	var res UploadDumpBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UploadDumpDefaultBody upload dump default body
swagger:model UploadDumpDefaultBody
*/
type UploadDumpDefaultBody struct {
	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*UploadDumpDefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this upload dump default body
func (o *UploadDumpDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UploadDumpDefaultBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("UploadDump default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("UploadDump default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this upload dump default body based on the context it is used
func (o *UploadDumpDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UploadDumpDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.Details); i++ {
		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("UploadDump default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("UploadDump default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UploadDumpDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UploadDumpDefaultBody) UnmarshalBinary(b []byte) error {
	var res UploadDumpDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UploadDumpDefaultBodyDetailsItems0 upload dump default body details items0
swagger:model UploadDumpDefaultBodyDetailsItems0
*/
type UploadDumpDefaultBodyDetailsItems0 struct {
	// at type
	AtType string `json:"@type,omitempty"`
}

// Validate validates this upload dump default body details items0
func (o *UploadDumpDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this upload dump default body details items0 based on context it is used
func (o *UploadDumpDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UploadDumpDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UploadDumpDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res UploadDumpDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UploadDumpParamsBodyFtpParameters upload dump params body ftp parameters
swagger:model UploadDumpParamsBodyFtpParameters
*/
type UploadDumpParamsBodyFtpParameters struct {
	// address
	Address string `json:"address,omitempty"`

	// user
	User string `json:"user,omitempty"`

	// password
	Password string `json:"password,omitempty"`
}

// Validate validates this upload dump params body ftp parameters
func (o *UploadDumpParamsBodyFtpParameters) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this upload dump params body ftp parameters based on context it is used
func (o *UploadDumpParamsBodyFtpParameters) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UploadDumpParamsBodyFtpParameters) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UploadDumpParamsBodyFtpParameters) UnmarshalBinary(b []byte) error {
	var res UploadDumpParamsBodyFtpParameters
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
