// Code generated by go-swagger; DO NOT EDIT.

package pxc_clusters

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

// GetPXCClusterCredentialsReader is a Reader for the GetPXCClusterCredentials structure.
type GetPXCClusterCredentialsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPXCClusterCredentialsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPXCClusterCredentialsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetPXCClusterCredentialsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetPXCClusterCredentialsOK creates a GetPXCClusterCredentialsOK with default headers values
func NewGetPXCClusterCredentialsOK() *GetPXCClusterCredentialsOK {
	return &GetPXCClusterCredentialsOK{}
}

/* GetPXCClusterCredentialsOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetPXCClusterCredentialsOK struct {
	Payload *GetPXCClusterCredentialsOKBody
}

// IsSuccess returns true when this get Pxc cluster credentials Ok response has a 2xx status code
func (o *GetPXCClusterCredentialsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get Pxc cluster credentials Ok response has a 3xx status code
func (o *GetPXCClusterCredentialsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Pxc cluster credentials Ok response has a 4xx status code
func (o *GetPXCClusterCredentialsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get Pxc cluster credentials Ok response has a 5xx status code
func (o *GetPXCClusterCredentialsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get Pxc cluster credentials Ok response a status code equal to that given
func (o *GetPXCClusterCredentialsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetPXCClusterCredentialsOK) Error() string {
	return fmt.Sprintf("[POST /v1/management/DBaaS/PXCClusters/GetCredentials][%d] getPxcClusterCredentialsOk  %+v", 200, o.Payload)
}

func (o *GetPXCClusterCredentialsOK) String() string {
	return fmt.Sprintf("[POST /v1/management/DBaaS/PXCClusters/GetCredentials][%d] getPxcClusterCredentialsOk  %+v", 200, o.Payload)
}

func (o *GetPXCClusterCredentialsOK) GetPayload() *GetPXCClusterCredentialsOKBody {
	return o.Payload
}

func (o *GetPXCClusterCredentialsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(GetPXCClusterCredentialsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPXCClusterCredentialsDefault creates a GetPXCClusterCredentialsDefault with default headers values
func NewGetPXCClusterCredentialsDefault(code int) *GetPXCClusterCredentialsDefault {
	return &GetPXCClusterCredentialsDefault{
		_statusCode: code,
	}
}

/* GetPXCClusterCredentialsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type GetPXCClusterCredentialsDefault struct {
	_statusCode int

	Payload *GetPXCClusterCredentialsDefaultBody
}

// Code gets the status code for the get PXC cluster credentials default response
func (o *GetPXCClusterCredentialsDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this get PXC cluster credentials default response has a 2xx status code
func (o *GetPXCClusterCredentialsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get PXC cluster credentials default response has a 3xx status code
func (o *GetPXCClusterCredentialsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get PXC cluster credentials default response has a 4xx status code
func (o *GetPXCClusterCredentialsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get PXC cluster credentials default response has a 5xx status code
func (o *GetPXCClusterCredentialsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get PXC cluster credentials default response a status code equal to that given
func (o *GetPXCClusterCredentialsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GetPXCClusterCredentialsDefault) Error() string {
	return fmt.Sprintf("[POST /v1/management/DBaaS/PXCClusters/GetCredentials][%d] GetPXCClusterCredentials default  %+v", o._statusCode, o.Payload)
}

func (o *GetPXCClusterCredentialsDefault) String() string {
	return fmt.Sprintf("[POST /v1/management/DBaaS/PXCClusters/GetCredentials][%d] GetPXCClusterCredentials default  %+v", o._statusCode, o.Payload)
}

func (o *GetPXCClusterCredentialsDefault) GetPayload() *GetPXCClusterCredentialsDefaultBody {
	return o.Payload
}

func (o *GetPXCClusterCredentialsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(GetPXCClusterCredentialsDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetPXCClusterCredentialsBody get PXC cluster credentials body
swagger:model GetPXCClusterCredentialsBody
*/
type GetPXCClusterCredentialsBody struct {
	// Kubernetes cluster name.
	KubernetesClusterName string `json:"kubernetes_cluster_name,omitempty"`

	// PXC cluster name.
	Name string `json:"name,omitempty"`
}

// Validate validates this get PXC cluster credentials body
func (o *GetPXCClusterCredentialsBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get PXC cluster credentials body based on context it is used
func (o *GetPXCClusterCredentialsBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetPXCClusterCredentialsBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetPXCClusterCredentialsBody) UnmarshalBinary(b []byte) error {
	var res GetPXCClusterCredentialsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetPXCClusterCredentialsDefaultBody get PXC cluster credentials default body
swagger:model GetPXCClusterCredentialsDefaultBody
*/
type GetPXCClusterCredentialsDefaultBody struct {
	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*GetPXCClusterCredentialsDefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this get PXC cluster credentials default body
func (o *GetPXCClusterCredentialsDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetPXCClusterCredentialsDefaultBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("GetPXCClusterCredentials default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("GetPXCClusterCredentials default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get PXC cluster credentials default body based on the context it is used
func (o *GetPXCClusterCredentialsDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetPXCClusterCredentialsDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.Details); i++ {
		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("GetPXCClusterCredentials default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("GetPXCClusterCredentials default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetPXCClusterCredentialsDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetPXCClusterCredentialsDefaultBody) UnmarshalBinary(b []byte) error {
	var res GetPXCClusterCredentialsDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetPXCClusterCredentialsDefaultBodyDetailsItems0 get PXC cluster credentials default body details items0
swagger:model GetPXCClusterCredentialsDefaultBodyDetailsItems0
*/
type GetPXCClusterCredentialsDefaultBodyDetailsItems0 struct {
	// at type
	AtType string `json:"@type,omitempty"`
}

// Validate validates this get PXC cluster credentials default body details items0
func (o *GetPXCClusterCredentialsDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get PXC cluster credentials default body details items0 based on context it is used
func (o *GetPXCClusterCredentialsDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetPXCClusterCredentialsDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetPXCClusterCredentialsDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res GetPXCClusterCredentialsDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetPXCClusterCredentialsOKBody get PXC cluster credentials OK body
swagger:model GetPXCClusterCredentialsOKBody
*/
type GetPXCClusterCredentialsOKBody struct {
	// connection credentials
	ConnectionCredentials *GetPXCClusterCredentialsOKBodyConnectionCredentials `json:"connection_credentials,omitempty"`
}

// Validate validates this get PXC cluster credentials OK body
func (o *GetPXCClusterCredentialsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateConnectionCredentials(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetPXCClusterCredentialsOKBody) validateConnectionCredentials(formats strfmt.Registry) error {
	if swag.IsZero(o.ConnectionCredentials) { // not required
		return nil
	}

	if o.ConnectionCredentials != nil {
		if err := o.ConnectionCredentials.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getPxcClusterCredentialsOk" + "." + "connection_credentials")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getPxcClusterCredentialsOk" + "." + "connection_credentials")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get PXC cluster credentials OK body based on the context it is used
func (o *GetPXCClusterCredentialsOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateConnectionCredentials(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetPXCClusterCredentialsOKBody) contextValidateConnectionCredentials(ctx context.Context, formats strfmt.Registry) error {
	if o.ConnectionCredentials != nil {
		if err := o.ConnectionCredentials.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getPxcClusterCredentialsOk" + "." + "connection_credentials")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getPxcClusterCredentialsOk" + "." + "connection_credentials")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetPXCClusterCredentialsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetPXCClusterCredentialsOKBody) UnmarshalBinary(b []byte) error {
	var res GetPXCClusterCredentialsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetPXCClusterCredentialsOKBodyConnectionCredentials PXCClusterConnectionCredentials is cluster connection credentials.
swagger:model GetPXCClusterCredentialsOKBodyConnectionCredentials
*/
type GetPXCClusterCredentialsOKBodyConnectionCredentials struct {
	// PXC username.
	Username string `json:"username,omitempty"`

	// PXC password.
	Password string `json:"password,omitempty"`

	// PXC host.
	Host string `json:"host,omitempty"`

	// PXC port.
	Port int32 `json:"port,omitempty"`
}

// Validate validates this get PXC cluster credentials OK body connection credentials
func (o *GetPXCClusterCredentialsOKBodyConnectionCredentials) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get PXC cluster credentials OK body connection credentials based on context it is used
func (o *GetPXCClusterCredentialsOKBodyConnectionCredentials) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetPXCClusterCredentialsOKBodyConnectionCredentials) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetPXCClusterCredentialsOKBodyConnectionCredentials) UnmarshalBinary(b []byte) error {
	var res GetPXCClusterCredentialsOKBodyConnectionCredentials
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
