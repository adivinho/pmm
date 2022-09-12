// Code generated by go-swagger; DO NOT EDIT.

package server

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

// VersionReader is a Reader for the Version structure.
type VersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewVersionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewVersionOK creates a VersionOK with default headers values
func NewVersionOK() *VersionOK {
	return &VersionOK{}
}

/* VersionOK describes a response with status code 200, with default header values.

A successful response.
*/
type VersionOK struct {
	Payload *VersionOKBody
}

// IsSuccess returns true when this version Ok response has a 2xx status code
func (o *VersionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this version Ok response has a 3xx status code
func (o *VersionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this version Ok response has a 4xx status code
func (o *VersionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this version Ok response has a 5xx status code
func (o *VersionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this version Ok response a status code equal to that given
func (o *VersionOK) IsCode(code int) bool {
	return code == 200
}

func (o *VersionOK) Error() string {
	return fmt.Sprintf("[GET /v1/version][%d] versionOk  %+v", 200, o.Payload)
}

func (o *VersionOK) String() string {
	return fmt.Sprintf("[GET /v1/version][%d] versionOk  %+v", 200, o.Payload)
}

func (o *VersionOK) GetPayload() *VersionOKBody {
	return o.Payload
}

func (o *VersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(VersionOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVersionDefault creates a VersionDefault with default headers values
func NewVersionDefault(code int) *VersionDefault {
	return &VersionDefault{
		_statusCode: code,
	}
}

/* VersionDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type VersionDefault struct {
	_statusCode int

	Payload *VersionDefaultBody
}

// Code gets the status code for the version default response
func (o *VersionDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this version default response has a 2xx status code
func (o *VersionDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this version default response has a 3xx status code
func (o *VersionDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this version default response has a 4xx status code
func (o *VersionDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this version default response has a 5xx status code
func (o *VersionDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this version default response a status code equal to that given
func (o *VersionDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *VersionDefault) Error() string {
	return fmt.Sprintf("[GET /v1/version][%d] Version default  %+v", o._statusCode, o.Payload)
}

func (o *VersionDefault) String() string {
	return fmt.Sprintf("[GET /v1/version][%d] Version default  %+v", o._statusCode, o.Payload)
}

func (o *VersionDefault) GetPayload() *VersionDefaultBody {
	return o.Payload
}

func (o *VersionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(VersionDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*VersionDefaultBody version default body
swagger:model VersionDefaultBody
*/
type VersionDefaultBody struct {
	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*VersionDefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this version default body
func (o *VersionDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *VersionDefaultBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("Version default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Version default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this version default body based on the context it is used
func (o *VersionDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *VersionDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.Details); i++ {
		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Version default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Version default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *VersionDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *VersionDefaultBody) UnmarshalBinary(b []byte) error {
	var res VersionDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*VersionDefaultBodyDetailsItems0 `Any` contains an arbitrary serialized protocol buffer message along with a
// URL that describes the type of the serialized message.
//
// Protobuf library provides support to pack/unpack Any values in the form
// of utility functions or additional generated methods of the Any type.
//
// Example 1: Pack and unpack a message in C++.
//
//     Foo foo = ...;
//     Any any;
//     any.PackFrom(foo);
//     ...
//     if (any.UnpackTo(&foo)) {
//       ...
//     }
//
// Example 2: Pack and unpack a message in Java.
//
//     Foo foo = ...;
//     Any any = Any.pack(foo);
//     ...
//     if (any.is(Foo.class)) {
//       foo = any.unpack(Foo.class);
//     }
//
// Example 3: Pack and unpack a message in Python.
//
//     foo = Foo(...)
//     any = Any()
//     any.Pack(foo)
//     ...
//     if any.Is(Foo.DESCRIPTOR):
//       any.Unpack(foo)
//       ...
//
// Example 4: Pack and unpack a message in Go
//
//      foo := &pb.Foo{...}
//      any, err := anypb.New(foo)
//      if err != nil {
//        ...
//      }
//      ...
//      foo := &pb.Foo{}
//      if err := any.UnmarshalTo(foo); err != nil {
//        ...
//      }
//
// The pack methods provided by protobuf library will by default use
// 'type.googleapis.com/full.type.name' as the type URL and the unpack
// methods only use the fully qualified type name after the last '/'
// in the type URL, for example "foo.bar.com/x/y.z" will yield type
// name "y.z".
//
//
// JSON
//
// The JSON representation of an `Any` value uses the regular
// representation of the deserialized, embedded message, with an
// additional field `@type` which contains the type URL. Example:
//
//     package google.profile;
//     message Person {
//       string first_name = 1;
//       string last_name = 2;
//     }
//
//     {
//       "@type": "type.googleapis.com/google.profile.Person",
//       "firstName": <string>,
//       "lastName": <string>
//     }
//
// If the embedded message type is well-known and has a custom JSON
// representation, that representation will be embedded adding a field
// `value` which holds the custom JSON in addition to the `@type`
// field. Example (for message [google.protobuf.Duration][]):
//
//     {
//       "@type": "type.googleapis.com/google.protobuf.Duration",
//       "value": "1.212s"
//     }
swagger:model VersionDefaultBodyDetailsItems0
*/
type VersionDefaultBodyDetailsItems0 struct {
	// A URL/resource name that uniquely identifies the type of the serialized
	// protocol buffer message. This string must contain at least
	// one "/" character. The last segment of the URL's path must represent
	// the fully qualified name of the type (as in
	// `path/google.protobuf.Duration`). The name should be in a canonical form
	// (e.g., leading "." is not accepted).
	//
	// In practice, teams usually precompile into the binary all types that they
	// expect it to use in the context of Any. However, for URLs which use the
	// scheme `http`, `https`, or no scheme, one can optionally set up a type
	// server that maps type URLs to message definitions as follows:
	//
	// * If no scheme is provided, `https` is assumed.
	// * An HTTP GET on the URL must yield a [google.protobuf.Type][]
	//   value in binary format, or produce an error.
	// * Applications are allowed to cache lookup results based on the
	//   URL, or have them precompiled into a binary to avoid any
	//   lookup. Therefore, binary compatibility needs to be preserved
	//   on changes to types. (Use versioned type names to manage
	//   breaking changes.)
	//
	// Note: this functionality is not currently available in the official
	// protobuf release, and it is not used for type URLs beginning with
	// type.googleapis.com.
	//
	// Schemes other than `http`, `https` (or the empty scheme) might be
	// used with implementation specific semantics.
	AtType string `json:"@type,omitempty"`
}

// Validate validates this version default body details items0
func (o *VersionDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this version default body details items0 based on context it is used
func (o *VersionDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *VersionDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *VersionDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res VersionDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*VersionOKBody version OK body
swagger:model VersionOKBody
*/
type VersionOKBody struct {
	// PMM Server version.
	Version string `json:"version,omitempty"`

	// DistributionMethod defines PMM Server distribution method: Docker image, OVF/OVA, or AMI.
	// Enum: [DISTRIBUTION_METHOD_INVALID DOCKER OVF AMI AZURE DO]
	DistributionMethod *string `json:"distribution_method,omitempty"`

	// managed
	Managed *VersionOKBodyManaged `json:"managed,omitempty"`

	// server
	Server *VersionOKBodyServer `json:"server,omitempty"`
}

// Validate validates this version OK body
func (o *VersionOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDistributionMethod(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateManaged(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateServer(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var versionOkBodyTypeDistributionMethodPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DISTRIBUTION_METHOD_INVALID","DOCKER","OVF","AMI","AZURE","DO"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		versionOkBodyTypeDistributionMethodPropEnum = append(versionOkBodyTypeDistributionMethodPropEnum, v)
	}
}

const (

	// VersionOKBodyDistributionMethodDISTRIBUTIONMETHODINVALID captures enum value "DISTRIBUTION_METHOD_INVALID"
	VersionOKBodyDistributionMethodDISTRIBUTIONMETHODINVALID string = "DISTRIBUTION_METHOD_INVALID"

	// VersionOKBodyDistributionMethodDOCKER captures enum value "DOCKER"
	VersionOKBodyDistributionMethodDOCKER string = "DOCKER"

	// VersionOKBodyDistributionMethodOVF captures enum value "OVF"
	VersionOKBodyDistributionMethodOVF string = "OVF"

	// VersionOKBodyDistributionMethodAMI captures enum value "AMI"
	VersionOKBodyDistributionMethodAMI string = "AMI"

	// VersionOKBodyDistributionMethodAZURE captures enum value "AZURE"
	VersionOKBodyDistributionMethodAZURE string = "AZURE"

	// VersionOKBodyDistributionMethodDO captures enum value "DO"
	VersionOKBodyDistributionMethodDO string = "DO"
)

// prop value enum
func (o *VersionOKBody) validateDistributionMethodEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, versionOkBodyTypeDistributionMethodPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *VersionOKBody) validateDistributionMethod(formats strfmt.Registry) error {
	if swag.IsZero(o.DistributionMethod) { // not required
		return nil
	}

	// value enum
	if err := o.validateDistributionMethodEnum("versionOk"+"."+"distribution_method", "body", *o.DistributionMethod); err != nil {
		return err
	}

	return nil
}

func (o *VersionOKBody) validateManaged(formats strfmt.Registry) error {
	if swag.IsZero(o.Managed) { // not required
		return nil
	}

	if o.Managed != nil {
		if err := o.Managed.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("versionOk" + "." + "managed")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("versionOk" + "." + "managed")
			}
			return err
		}
	}

	return nil
}

func (o *VersionOKBody) validateServer(formats strfmt.Registry) error {
	if swag.IsZero(o.Server) { // not required
		return nil
	}

	if o.Server != nil {
		if err := o.Server.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("versionOk" + "." + "server")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("versionOk" + "." + "server")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this version OK body based on the context it is used
func (o *VersionOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateManaged(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateServer(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *VersionOKBody) contextValidateManaged(ctx context.Context, formats strfmt.Registry) error {
	if o.Managed != nil {
		if err := o.Managed.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("versionOk" + "." + "managed")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("versionOk" + "." + "managed")
			}
			return err
		}
	}

	return nil
}

func (o *VersionOKBody) contextValidateServer(ctx context.Context, formats strfmt.Registry) error {
	if o.Server != nil {
		if err := o.Server.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("versionOk" + "." + "server")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("versionOk" + "." + "server")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *VersionOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *VersionOKBody) UnmarshalBinary(b []byte) error {
	var res VersionOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*VersionOKBodyManaged VersionInfo describes component version, or PMM Server as a whole.
swagger:model VersionOKBodyManaged
*/
type VersionOKBodyManaged struct {
	// User-visible version.
	Version string `json:"version,omitempty"`

	// Full version for debugging.
	FullVersion string `json:"full_version,omitempty"`

	// Build or release date.
	// Format: date-time
	Timestamp strfmt.DateTime `json:"timestamp,omitempty"`
}

// Validate validates this version OK body managed
func (o *VersionOKBodyManaged) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *VersionOKBodyManaged) validateTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(o.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("versionOk"+"."+"managed"+"."+"timestamp", "body", "date-time", o.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this version OK body managed based on context it is used
func (o *VersionOKBodyManaged) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *VersionOKBodyManaged) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *VersionOKBodyManaged) UnmarshalBinary(b []byte) error {
	var res VersionOKBodyManaged
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*VersionOKBodyServer VersionInfo describes component version, or PMM Server as a whole.
swagger:model VersionOKBodyServer
*/
type VersionOKBodyServer struct {
	// User-visible version.
	Version string `json:"version,omitempty"`

	// Full version for debugging.
	FullVersion string `json:"full_version,omitempty"`

	// Build or release date.
	// Format: date-time
	Timestamp strfmt.DateTime `json:"timestamp,omitempty"`
}

// Validate validates this version OK body server
func (o *VersionOKBodyServer) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *VersionOKBodyServer) validateTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(o.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("versionOk"+"."+"server"+"."+"timestamp", "body", "date-time", o.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this version OK body server based on context it is used
func (o *VersionOKBodyServer) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *VersionOKBodyServer) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *VersionOKBodyServer) UnmarshalBinary(b []byte) error {
	var res VersionOKBodyServer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
