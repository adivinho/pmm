// Code generated by go-swagger; DO NOT EDIT.

package db_clusters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RestartDBClusterReader is a Reader for the RestartDBCluster structure.
type RestartDBClusterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RestartDBClusterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRestartDBClusterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewRestartDBClusterDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRestartDBClusterOK creates a RestartDBClusterOK with default headers values
func NewRestartDBClusterOK() *RestartDBClusterOK {
	return &RestartDBClusterOK{}
}

/* RestartDBClusterOK describes a response with status code 200, with default header values.

A successful response.
*/
type RestartDBClusterOK struct {
	Payload interface{}
}

// IsSuccess returns true when this restart Db cluster Ok response has a 2xx status code
func (o *RestartDBClusterOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this restart Db cluster Ok response has a 3xx status code
func (o *RestartDBClusterOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this restart Db cluster Ok response has a 4xx status code
func (o *RestartDBClusterOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this restart Db cluster Ok response has a 5xx status code
func (o *RestartDBClusterOK) IsServerError() bool {
	return false
}

// IsCode returns true when this restart Db cluster Ok response a status code equal to that given
func (o *RestartDBClusterOK) IsCode(code int) bool {
	return code == 200
}

func (o *RestartDBClusterOK) Error() string {
	return fmt.Sprintf("[POST /v1/management/DBaaS/DBClusters/Restart][%d] restartDbClusterOk  %+v", 200, o.Payload)
}

func (o *RestartDBClusterOK) String() string {
	return fmt.Sprintf("[POST /v1/management/DBaaS/DBClusters/Restart][%d] restartDbClusterOk  %+v", 200, o.Payload)
}

func (o *RestartDBClusterOK) GetPayload() interface{} {
	return o.Payload
}

func (o *RestartDBClusterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRestartDBClusterDefault creates a RestartDBClusterDefault with default headers values
func NewRestartDBClusterDefault(code int) *RestartDBClusterDefault {
	return &RestartDBClusterDefault{
		_statusCode: code,
	}
}

/* RestartDBClusterDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type RestartDBClusterDefault struct {
	_statusCode int

	Payload *RestartDBClusterDefaultBody
}

// Code gets the status code for the restart DB cluster default response
func (o *RestartDBClusterDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this restart DB cluster default response has a 2xx status code
func (o *RestartDBClusterDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this restart DB cluster default response has a 3xx status code
func (o *RestartDBClusterDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this restart DB cluster default response has a 4xx status code
func (o *RestartDBClusterDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this restart DB cluster default response has a 5xx status code
func (o *RestartDBClusterDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this restart DB cluster default response a status code equal to that given
func (o *RestartDBClusterDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *RestartDBClusterDefault) Error() string {
	return fmt.Sprintf("[POST /v1/management/DBaaS/DBClusters/Restart][%d] RestartDBCluster default  %+v", o._statusCode, o.Payload)
}

func (o *RestartDBClusterDefault) String() string {
	return fmt.Sprintf("[POST /v1/management/DBaaS/DBClusters/Restart][%d] RestartDBCluster default  %+v", o._statusCode, o.Payload)
}

func (o *RestartDBClusterDefault) GetPayload() *RestartDBClusterDefaultBody {
	return o.Payload
}

func (o *RestartDBClusterDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(RestartDBClusterDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*RestartDBClusterBody restart DB cluster body
swagger:model RestartDBClusterBody
*/
type RestartDBClusterBody struct {
	// Kubernetes cluster name.
	KubernetesClusterName string `json:"kubernetes_cluster_name,omitempty"`

	// PXC cluster name.
	Name string `json:"name,omitempty"`

	// DBClusterType represents database cluster type.
	//
	//  - DB_CLUSTER_TYPE_INVALID: DB_CLUSTER_TYPE_INVALID represents unknown cluster type.
	//  - DB_CLUSTER_TYPE_PXC: DB_CLUSTER_TYPE_PXC represents pxc cluster type.
	//  - DB_CLUSTER_TYPE_PSMDB: DB_CLUSTER_TYPE_PSMDB represents psmdb cluster type.
	// Enum: [DB_CLUSTER_TYPE_INVALID DB_CLUSTER_TYPE_PXC DB_CLUSTER_TYPE_PSMDB]
	ClusterType *string `json:"cluster_type,omitempty"`
}

// Validate validates this restart DB cluster body
func (o *RestartDBClusterBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateClusterType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var restartDbClusterBodyTypeClusterTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DB_CLUSTER_TYPE_INVALID","DB_CLUSTER_TYPE_PXC","DB_CLUSTER_TYPE_PSMDB"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		restartDbClusterBodyTypeClusterTypePropEnum = append(restartDbClusterBodyTypeClusterTypePropEnum, v)
	}
}

const (

	// RestartDBClusterBodyClusterTypeDBCLUSTERTYPEINVALID captures enum value "DB_CLUSTER_TYPE_INVALID"
	RestartDBClusterBodyClusterTypeDBCLUSTERTYPEINVALID string = "DB_CLUSTER_TYPE_INVALID"

	// RestartDBClusterBodyClusterTypeDBCLUSTERTYPEPXC captures enum value "DB_CLUSTER_TYPE_PXC"
	RestartDBClusterBodyClusterTypeDBCLUSTERTYPEPXC string = "DB_CLUSTER_TYPE_PXC"

	// RestartDBClusterBodyClusterTypeDBCLUSTERTYPEPSMDB captures enum value "DB_CLUSTER_TYPE_PSMDB"
	RestartDBClusterBodyClusterTypeDBCLUSTERTYPEPSMDB string = "DB_CLUSTER_TYPE_PSMDB"
)

// prop value enum
func (o *RestartDBClusterBody) validateClusterTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, restartDbClusterBodyTypeClusterTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *RestartDBClusterBody) validateClusterType(formats strfmt.Registry) error {
	if swag.IsZero(o.ClusterType) { // not required
		return nil
	}

	// value enum
	if err := o.validateClusterTypeEnum("body"+"."+"cluster_type", "body", *o.ClusterType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this restart DB cluster body based on context it is used
func (o *RestartDBClusterBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *RestartDBClusterBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RestartDBClusterBody) UnmarshalBinary(b []byte) error {
	var res RestartDBClusterBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*RestartDBClusterDefaultBody restart DB cluster default body
swagger:model RestartDBClusterDefaultBody
*/
type RestartDBClusterDefaultBody struct {
	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*RestartDBClusterDefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this restart DB cluster default body
func (o *RestartDBClusterDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *RestartDBClusterDefaultBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("RestartDBCluster default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("RestartDBCluster default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this restart DB cluster default body based on the context it is used
func (o *RestartDBClusterDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *RestartDBClusterDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.Details); i++ {
		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("RestartDBCluster default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("RestartDBCluster default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *RestartDBClusterDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RestartDBClusterDefaultBody) UnmarshalBinary(b []byte) error {
	var res RestartDBClusterDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*RestartDBClusterDefaultBodyDetailsItems0 restart DB cluster default body details items0
swagger:model RestartDBClusterDefaultBodyDetailsItems0
*/
type RestartDBClusterDefaultBodyDetailsItems0 struct {
	// at type
	AtType string `json:"@type,omitempty"`
}

// Validate validates this restart DB cluster default body details items0
func (o *RestartDBClusterDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this restart DB cluster default body details items0 based on context it is used
func (o *RestartDBClusterDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *RestartDBClusterDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RestartDBClusterDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res RestartDBClusterDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
