// Code generated by go-swagger; DO NOT EDIT.

package object_details

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

// FingerprintAndPlaceholdersCountByQueryIDReader is a Reader for the FingerprintAndPlaceholdersCountByQueryID structure.
type FingerprintAndPlaceholdersCountByQueryIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FingerprintAndPlaceholdersCountByQueryIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFingerprintAndPlaceholdersCountByQueryIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFingerprintAndPlaceholdersCountByQueryIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFingerprintAndPlaceholdersCountByQueryIDOK creates a FingerprintAndPlaceholdersCountByQueryIDOK with default headers values
func NewFingerprintAndPlaceholdersCountByQueryIDOK() *FingerprintAndPlaceholdersCountByQueryIDOK {
	return &FingerprintAndPlaceholdersCountByQueryIDOK{}
}

/*
FingerprintAndPlaceholdersCountByQueryIDOK describes a response with status code 200, with default header values.

A successful response.
*/
type FingerprintAndPlaceholdersCountByQueryIDOK struct {
	Payload *FingerprintAndPlaceholdersCountByQueryIDOKBody
}

func (o *FingerprintAndPlaceholdersCountByQueryIDOK) Error() string {
	return fmt.Sprintf("[POST /v0/qan/ObjectDetails/FingerprintAndPlaceholdersCountByQueryID][%d] fingerprintAndPlaceholdersCountByQueryIdOk  %+v", 200, o.Payload)
}

func (o *FingerprintAndPlaceholdersCountByQueryIDOK) GetPayload() *FingerprintAndPlaceholdersCountByQueryIDOKBody {
	return o.Payload
}

func (o *FingerprintAndPlaceholdersCountByQueryIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(FingerprintAndPlaceholdersCountByQueryIDOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFingerprintAndPlaceholdersCountByQueryIDDefault creates a FingerprintAndPlaceholdersCountByQueryIDDefault with default headers values
func NewFingerprintAndPlaceholdersCountByQueryIDDefault(code int) *FingerprintAndPlaceholdersCountByQueryIDDefault {
	return &FingerprintAndPlaceholdersCountByQueryIDDefault{
		_statusCode: code,
	}
}

/*
FingerprintAndPlaceholdersCountByQueryIDDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type FingerprintAndPlaceholdersCountByQueryIDDefault struct {
	_statusCode int

	Payload *FingerprintAndPlaceholdersCountByQueryIDDefaultBody
}

// Code gets the status code for the fingerprint and placeholders count by query ID default response
func (o *FingerprintAndPlaceholdersCountByQueryIDDefault) Code() int {
	return o._statusCode
}

func (o *FingerprintAndPlaceholdersCountByQueryIDDefault) Error() string {
	return fmt.Sprintf("[POST /v0/qan/ObjectDetails/FingerprintAndPlaceholdersCountByQueryID][%d] FingerprintAndPlaceholdersCountByQueryID default  %+v", o._statusCode, o.Payload)
}

func (o *FingerprintAndPlaceholdersCountByQueryIDDefault) GetPayload() *FingerprintAndPlaceholdersCountByQueryIDDefaultBody {
	return o.Payload
}

func (o *FingerprintAndPlaceholdersCountByQueryIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(FingerprintAndPlaceholdersCountByQueryIDDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
FingerprintAndPlaceholdersCountByQueryIDBody FingerprintAndPlaceholdersCountByQueryIDRequest gets fingerprint and placeholders count for given query ID.
swagger:model FingerprintAndPlaceholdersCountByQueryIDBody
*/
type FingerprintAndPlaceholdersCountByQueryIDBody struct {
	// serviceid
	Serviceid string `json:"serviceid,omitempty"`

	// query id
	QueryID string `json:"query_id,omitempty"`
}

// Validate validates this fingerprint and placeholders count by query ID body
func (o *FingerprintAndPlaceholdersCountByQueryIDBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this fingerprint and placeholders count by query ID body based on context it is used
func (o *FingerprintAndPlaceholdersCountByQueryIDBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *FingerprintAndPlaceholdersCountByQueryIDBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *FingerprintAndPlaceholdersCountByQueryIDBody) UnmarshalBinary(b []byte) error {
	var res FingerprintAndPlaceholdersCountByQueryIDBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
FingerprintAndPlaceholdersCountByQueryIDDefaultBody fingerprint and placeholders count by query ID default body
swagger:model FingerprintAndPlaceholdersCountByQueryIDDefaultBody
*/
type FingerprintAndPlaceholdersCountByQueryIDDefaultBody struct {
	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*FingerprintAndPlaceholdersCountByQueryIDDefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this fingerprint and placeholders count by query ID default body
func (o *FingerprintAndPlaceholdersCountByQueryIDDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *FingerprintAndPlaceholdersCountByQueryIDDefaultBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("FingerprintAndPlaceholdersCountByQueryID default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("FingerprintAndPlaceholdersCountByQueryID default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this fingerprint and placeholders count by query ID default body based on the context it is used
func (o *FingerprintAndPlaceholdersCountByQueryIDDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *FingerprintAndPlaceholdersCountByQueryIDDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.Details); i++ {
		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("FingerprintAndPlaceholdersCountByQueryID default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("FingerprintAndPlaceholdersCountByQueryID default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *FingerprintAndPlaceholdersCountByQueryIDDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *FingerprintAndPlaceholdersCountByQueryIDDefaultBody) UnmarshalBinary(b []byte) error {
	var res FingerprintAndPlaceholdersCountByQueryIDDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
FingerprintAndPlaceholdersCountByQueryIDDefaultBodyDetailsItems0 fingerprint and placeholders count by query ID default body details items0
swagger:model FingerprintAndPlaceholdersCountByQueryIDDefaultBodyDetailsItems0
*/
type FingerprintAndPlaceholdersCountByQueryIDDefaultBodyDetailsItems0 struct {
	// at type
	AtType string `json:"@type,omitempty"`
}

// Validate validates this fingerprint and placeholders count by query ID default body details items0
func (o *FingerprintAndPlaceholdersCountByQueryIDDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this fingerprint and placeholders count by query ID default body details items0 based on context it is used
func (o *FingerprintAndPlaceholdersCountByQueryIDDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *FingerprintAndPlaceholdersCountByQueryIDDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *FingerprintAndPlaceholdersCountByQueryIDDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res FingerprintAndPlaceholdersCountByQueryIDDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
FingerprintAndPlaceholdersCountByQueryIDOKBody FingerprintAndPlaceholdersCountByQueryIDReply is fingerprint and placeholders count for given query ID.
swagger:model FingerprintAndPlaceholdersCountByQueryIDOKBody
*/
type FingerprintAndPlaceholdersCountByQueryIDOKBody struct {
	// placeholders count
	PlaceholdersCount int64 `json:"placeholders_count,omitempty"`

	// fingerprint
	Fingerprint string `json:"fingerprint,omitempty"`
}

// Validate validates this fingerprint and placeholders count by query ID OK body
func (o *FingerprintAndPlaceholdersCountByQueryIDOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this fingerprint and placeholders count by query ID OK body based on context it is used
func (o *FingerprintAndPlaceholdersCountByQueryIDOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *FingerprintAndPlaceholdersCountByQueryIDOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *FingerprintAndPlaceholdersCountByQueryIDOKBody) UnmarshalBinary(b []byte) error {
	var res FingerprintAndPlaceholdersCountByQueryIDOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
