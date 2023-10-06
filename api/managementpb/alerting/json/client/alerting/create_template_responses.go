// Code generated by go-swagger; DO NOT EDIT.

package alerting

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

// CreateTemplateReader is a Reader for the CreateTemplate structure.
type CreateTemplateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateTemplateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateTemplateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateTemplateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateTemplateOK creates a CreateTemplateOK with default headers values
func NewCreateTemplateOK() *CreateTemplateOK {
	return &CreateTemplateOK{}
}

/*
CreateTemplateOK describes a response with status code 200, with default header values.

A successful response.
*/
type CreateTemplateOK struct {
	Payload interface{}
}

func (o *CreateTemplateOK) Error() string {
	return fmt.Sprintf("[POST /v1/management/alerting/Templates/Create][%d] createTemplateOk  %+v", 200, o.Payload)
}

func (o *CreateTemplateOK) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateTemplateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTemplateDefault creates a CreateTemplateDefault with default headers values
func NewCreateTemplateDefault(code int) *CreateTemplateDefault {
	return &CreateTemplateDefault{
		_statusCode: code,
	}
}

/*
CreateTemplateDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type CreateTemplateDefault struct {
	_statusCode int

	Payload *CreateTemplateDefaultBody
}

// Code gets the status code for the create template default response
func (o *CreateTemplateDefault) Code() int {
	return o._statusCode
}

func (o *CreateTemplateDefault) Error() string {
	return fmt.Sprintf("[POST /v1/management/alerting/Templates/Create][%d] CreateTemplate default  %+v", o._statusCode, o.Payload)
}

func (o *CreateTemplateDefault) GetPayload() *CreateTemplateDefaultBody {
	return o.Payload
}

func (o *CreateTemplateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(CreateTemplateDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
CreateTemplateBody create template body
swagger:model CreateTemplateBody
*/
type CreateTemplateBody struct {
	// YAML template file content.
	Yaml string `json:"yaml,omitempty"`
}

// Validate validates this create template body
func (o *CreateTemplateBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this create template body based on context it is used
func (o *CreateTemplateBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateTemplateBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateTemplateBody) UnmarshalBinary(b []byte) error {
	var res CreateTemplateBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
CreateTemplateDefaultBody create template default body
swagger:model CreateTemplateDefaultBody
*/
type CreateTemplateDefaultBody struct {
	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*CreateTemplateDefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this create template default body
func (o *CreateTemplateDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateTemplateDefaultBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("CreateTemplate default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("CreateTemplate default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this create template default body based on the context it is used
func (o *CreateTemplateDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateTemplateDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.Details); i++ {
		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("CreateTemplate default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("CreateTemplate default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateTemplateDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateTemplateDefaultBody) UnmarshalBinary(b []byte) error {
	var res CreateTemplateDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
CreateTemplateDefaultBodyDetailsItems0 create template default body details items0
swagger:model CreateTemplateDefaultBodyDetailsItems0
*/
type CreateTemplateDefaultBodyDetailsItems0 struct {
	// at type
	AtType string `json:"@type,omitempty"`
}

// Validate validates this create template default body details items0
func (o *CreateTemplateDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this create template default body details items0 based on context it is used
func (o *CreateTemplateDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateTemplateDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateTemplateDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res CreateTemplateDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
