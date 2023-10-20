// Code generated by go-swagger; DO NOT EDIT.

package agents

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

// AddProxySQLExporterReader is a Reader for the AddProxySQLExporter structure.
type AddProxySQLExporterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddProxySQLExporterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddProxySQLExporterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAddProxySQLExporterDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddProxySQLExporterOK creates a AddProxySQLExporterOK with default headers values
func NewAddProxySQLExporterOK() *AddProxySQLExporterOK {
	return &AddProxySQLExporterOK{}
}

/*
AddProxySQLExporterOK describes a response with status code 200, with default header values.

A successful response.
*/
type AddProxySQLExporterOK struct {
	Payload *AddProxySQLExporterOKBody
}

func (o *AddProxySQLExporterOK) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Agents/AddProxySQLExporter][%d] addProxySqlExporterOk  %+v", 200, o.Payload)
}

func (o *AddProxySQLExporterOK) GetPayload() *AddProxySQLExporterOKBody {
	return o.Payload
}

func (o *AddProxySQLExporterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(AddProxySQLExporterOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddProxySQLExporterDefault creates a AddProxySQLExporterDefault with default headers values
func NewAddProxySQLExporterDefault(code int) *AddProxySQLExporterDefault {
	return &AddProxySQLExporterDefault{
		_statusCode: code,
	}
}

/*
AddProxySQLExporterDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type AddProxySQLExporterDefault struct {
	_statusCode int

	Payload *AddProxySQLExporterDefaultBody
}

// Code gets the status code for the add proxy SQL exporter default response
func (o *AddProxySQLExporterDefault) Code() int {
	return o._statusCode
}

func (o *AddProxySQLExporterDefault) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Agents/AddProxySQLExporter][%d] AddProxySQLExporter default  %+v", o._statusCode, o.Payload)
}

func (o *AddProxySQLExporterDefault) GetPayload() *AddProxySQLExporterDefaultBody {
	return o.Payload
}

func (o *AddProxySQLExporterDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(AddProxySQLExporterDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
AddProxySQLExporterBody add proxy SQL exporter body
swagger:model AddProxySQLExporterBody
*/
type AddProxySQLExporterBody struct {
	// The pmm-agent identifier which runs this instance.
	PMMAgentID string `json:"pmm_agent_id,omitempty"`

	// Service identifier.
	ServiceID string `json:"service_id,omitempty"`

	// ProxySQL username for scraping metrics.
	Username string `json:"username,omitempty"`

	// ProxySQL password for scraping metrics.
	Password string `json:"password,omitempty"`

	// Use TLS for database connections.
	TLS bool `json:"tls,omitempty"`

	// Skip TLS certificate and hostname validation.
	TLSSkipVerify bool `json:"tls_skip_verify,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// Skip connection check.
	SkipConnectionCheck bool `json:"skip_connection_check,omitempty"`

	// Enables push metrics mode for exporter.
	PushMetrics bool `json:"push_metrics,omitempty"`

	// List of collector names to disable in this exporter.
	DisableCollectors []string `json:"disable_collectors"`

	// Custom password for exporter endpoint /metrics.
	AgentPassword string `json:"agent_password,omitempty"`

	// Log level for exporters
	// Enum: [auto fatal error warn info debug]
	LogLevel *string `json:"log_level,omitempty"`
}

// Validate validates this add proxy SQL exporter body
func (o *AddProxySQLExporterBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateLogLevel(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var addProxySqlExporterBodyTypeLogLevelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["auto","fatal","error","warn","info","debug"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		addProxySqlExporterBodyTypeLogLevelPropEnum = append(addProxySqlExporterBodyTypeLogLevelPropEnum, v)
	}
}

const (

	// AddProxySQLExporterBodyLogLevelAuto captures enum value "auto"
	AddProxySQLExporterBodyLogLevelAuto string = "auto"

	// AddProxySQLExporterBodyLogLevelFatal captures enum value "fatal"
	AddProxySQLExporterBodyLogLevelFatal string = "fatal"

	// AddProxySQLExporterBodyLogLevelError captures enum value "error"
	AddProxySQLExporterBodyLogLevelError string = "error"

	// AddProxySQLExporterBodyLogLevelWarn captures enum value "warn"
	AddProxySQLExporterBodyLogLevelWarn string = "warn"

	// AddProxySQLExporterBodyLogLevelInfo captures enum value "info"
	AddProxySQLExporterBodyLogLevelInfo string = "info"

	// AddProxySQLExporterBodyLogLevelDebug captures enum value "debug"
	AddProxySQLExporterBodyLogLevelDebug string = "debug"
)

// prop value enum
func (o *AddProxySQLExporterBody) validateLogLevelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, addProxySqlExporterBodyTypeLogLevelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *AddProxySQLExporterBody) validateLogLevel(formats strfmt.Registry) error {
	if swag.IsZero(o.LogLevel) { // not required
		return nil
	}

	// value enum
	if err := o.validateLogLevelEnum("body"+"."+"log_level", "body", *o.LogLevel); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this add proxy SQL exporter body based on context it is used
func (o *AddProxySQLExporterBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddProxySQLExporterBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddProxySQLExporterBody) UnmarshalBinary(b []byte) error {
	var res AddProxySQLExporterBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
AddProxySQLExporterDefaultBody add proxy SQL exporter default body
swagger:model AddProxySQLExporterDefaultBody
*/
type AddProxySQLExporterDefaultBody struct {
	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*AddProxySQLExporterDefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this add proxy SQL exporter default body
func (o *AddProxySQLExporterDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddProxySQLExporterDefaultBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("AddProxySQLExporter default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("AddProxySQLExporter default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this add proxy SQL exporter default body based on the context it is used
func (o *AddProxySQLExporterDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddProxySQLExporterDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.Details); i++ {
		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("AddProxySQLExporter default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("AddProxySQLExporter default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddProxySQLExporterDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddProxySQLExporterDefaultBody) UnmarshalBinary(b []byte) error {
	var res AddProxySQLExporterDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
AddProxySQLExporterDefaultBodyDetailsItems0 add proxy SQL exporter default body details items0
swagger:model AddProxySQLExporterDefaultBodyDetailsItems0
*/
type AddProxySQLExporterDefaultBodyDetailsItems0 struct {
	// at type
	AtType string `json:"@type,omitempty"`
}

// Validate validates this add proxy SQL exporter default body details items0
func (o *AddProxySQLExporterDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this add proxy SQL exporter default body details items0 based on context it is used
func (o *AddProxySQLExporterDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddProxySQLExporterDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddProxySQLExporterDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res AddProxySQLExporterDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
AddProxySQLExporterOKBody add proxy SQL exporter OK body
swagger:model AddProxySQLExporterOKBody
*/
type AddProxySQLExporterOKBody struct {
	// proxysql exporter
	ProxysqlExporter *AddProxySQLExporterOKBodyProxysqlExporter `json:"proxysql_exporter,omitempty"`
}

// Validate validates this add proxy SQL exporter OK body
func (o *AddProxySQLExporterOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateProxysqlExporter(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddProxySQLExporterOKBody) validateProxysqlExporter(formats strfmt.Registry) error {
	if swag.IsZero(o.ProxysqlExporter) { // not required
		return nil
	}

	if o.ProxysqlExporter != nil {
		if err := o.ProxysqlExporter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("addProxySqlExporterOk" + "." + "proxysql_exporter")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("addProxySqlExporterOk" + "." + "proxysql_exporter")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this add proxy SQL exporter OK body based on the context it is used
func (o *AddProxySQLExporterOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateProxysqlExporter(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddProxySQLExporterOKBody) contextValidateProxysqlExporter(ctx context.Context, formats strfmt.Registry) error {
	if o.ProxysqlExporter != nil {
		if err := o.ProxysqlExporter.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("addProxySqlExporterOk" + "." + "proxysql_exporter")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("addProxySqlExporterOk" + "." + "proxysql_exporter")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddProxySQLExporterOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddProxySQLExporterOKBody) UnmarshalBinary(b []byte) error {
	var res AddProxySQLExporterOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
AddProxySQLExporterOKBodyProxysqlExporter ProxySQLExporter runs on Generic or Container Node and exposes ProxySQL Service metrics.
swagger:model AddProxySQLExporterOKBodyProxysqlExporter
*/
type AddProxySQLExporterOKBodyProxysqlExporter struct {
	// Unique randomly generated instance identifier.
	AgentID string `json:"agent_id,omitempty"`

	// The pmm-agent identifier which runs this instance.
	PMMAgentID string `json:"pmm_agent_id,omitempty"`

	// Desired Agent status: enabled (false) or disabled (true).
	Disabled bool `json:"disabled,omitempty"`

	// Service identifier.
	ServiceID string `json:"service_id,omitempty"`

	// ProxySQL username for scraping metrics.
	Username string `json:"username,omitempty"`

	// Use TLS for database connections.
	TLS bool `json:"tls,omitempty"`

	// Skip TLS certificate and hostname validation.
	TLSSkipVerify bool `json:"tls_skip_verify,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// True if exporter uses push metrics mode.
	PushMetricsEnabled bool `json:"push_metrics_enabled,omitempty"`

	// List of disabled collector names.
	//
	// Status fields below.
	DisabledCollectors []string `json:"disabled_collectors"`

	// AgentStatus represents actual Agent status.
	//
	//  - STARTING: Agent is starting.
	//  - RUNNING: Agent is running.
	//  - WAITING: Agent encountered error and will be restarted automatically soon.
	//  - STOPPING: Agent is stopping.
	//  - DONE: Agent finished.
	//  - UNKNOWN: Agent is not connected, we don't know anything about it's state.
	// Enum: [AGENT_STATUS_INVALID STARTING RUNNING WAITING STOPPING DONE UNKNOWN]
	Status *string `json:"status,omitempty"`

	// Listen port for scraping metrics.
	ListenPort int64 `json:"listen_port,omitempty"`

	// Path to exec process.
	ProcessExecPath string `json:"process_exec_path,omitempty"`

	// Log level for exporters
	// Enum: [auto fatal error warn info debug]
	LogLevel *string `json:"log_level,omitempty"`

	// Optionally expose the exporter process on 0.0.0.0
	ExposeExporter bool `json:"expose_exporter,omitempty"`
}

// Validate validates this add proxy SQL exporter OK body proxysql exporter
func (o *AddProxySQLExporterOKBodyProxysqlExporter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateLogLevel(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var addProxySqlExporterOkBodyProxysqlExporterTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AGENT_STATUS_INVALID","STARTING","RUNNING","WAITING","STOPPING","DONE","UNKNOWN"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		addProxySqlExporterOkBodyProxysqlExporterTypeStatusPropEnum = append(addProxySqlExporterOkBodyProxysqlExporterTypeStatusPropEnum, v)
	}
}

const (

	// AddProxySQLExporterOKBodyProxysqlExporterStatusAGENTSTATUSINVALID captures enum value "AGENT_STATUS_INVALID"
	AddProxySQLExporterOKBodyProxysqlExporterStatusAGENTSTATUSINVALID string = "AGENT_STATUS_INVALID"

	// AddProxySQLExporterOKBodyProxysqlExporterStatusSTARTING captures enum value "STARTING"
	AddProxySQLExporterOKBodyProxysqlExporterStatusSTARTING string = "STARTING"

	// AddProxySQLExporterOKBodyProxysqlExporterStatusRUNNING captures enum value "RUNNING"
	AddProxySQLExporterOKBodyProxysqlExporterStatusRUNNING string = "RUNNING"

	// AddProxySQLExporterOKBodyProxysqlExporterStatusWAITING captures enum value "WAITING"
	AddProxySQLExporterOKBodyProxysqlExporterStatusWAITING string = "WAITING"

	// AddProxySQLExporterOKBodyProxysqlExporterStatusSTOPPING captures enum value "STOPPING"
	AddProxySQLExporterOKBodyProxysqlExporterStatusSTOPPING string = "STOPPING"

	// AddProxySQLExporterOKBodyProxysqlExporterStatusDONE captures enum value "DONE"
	AddProxySQLExporterOKBodyProxysqlExporterStatusDONE string = "DONE"

	// AddProxySQLExporterOKBodyProxysqlExporterStatusUNKNOWN captures enum value "UNKNOWN"
	AddProxySQLExporterOKBodyProxysqlExporterStatusUNKNOWN string = "UNKNOWN"
)

// prop value enum
func (o *AddProxySQLExporterOKBodyProxysqlExporter) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, addProxySqlExporterOkBodyProxysqlExporterTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *AddProxySQLExporterOKBodyProxysqlExporter) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(o.Status) { // not required
		return nil
	}

	// value enum
	if err := o.validateStatusEnum("addProxySqlExporterOk"+"."+"proxysql_exporter"+"."+"status", "body", *o.Status); err != nil {
		return err
	}

	return nil
}

var addProxySqlExporterOkBodyProxysqlExporterTypeLogLevelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["auto","fatal","error","warn","info","debug"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		addProxySqlExporterOkBodyProxysqlExporterTypeLogLevelPropEnum = append(addProxySqlExporterOkBodyProxysqlExporterTypeLogLevelPropEnum, v)
	}
}

const (

	// AddProxySQLExporterOKBodyProxysqlExporterLogLevelAuto captures enum value "auto"
	AddProxySQLExporterOKBodyProxysqlExporterLogLevelAuto string = "auto"

	// AddProxySQLExporterOKBodyProxysqlExporterLogLevelFatal captures enum value "fatal"
	AddProxySQLExporterOKBodyProxysqlExporterLogLevelFatal string = "fatal"

	// AddProxySQLExporterOKBodyProxysqlExporterLogLevelError captures enum value "error"
	AddProxySQLExporterOKBodyProxysqlExporterLogLevelError string = "error"

	// AddProxySQLExporterOKBodyProxysqlExporterLogLevelWarn captures enum value "warn"
	AddProxySQLExporterOKBodyProxysqlExporterLogLevelWarn string = "warn"

	// AddProxySQLExporterOKBodyProxysqlExporterLogLevelInfo captures enum value "info"
	AddProxySQLExporterOKBodyProxysqlExporterLogLevelInfo string = "info"

	// AddProxySQLExporterOKBodyProxysqlExporterLogLevelDebug captures enum value "debug"
	AddProxySQLExporterOKBodyProxysqlExporterLogLevelDebug string = "debug"
)

// prop value enum
func (o *AddProxySQLExporterOKBodyProxysqlExporter) validateLogLevelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, addProxySqlExporterOkBodyProxysqlExporterTypeLogLevelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *AddProxySQLExporterOKBodyProxysqlExporter) validateLogLevel(formats strfmt.Registry) error {
	if swag.IsZero(o.LogLevel) { // not required
		return nil
	}

	// value enum
	if err := o.validateLogLevelEnum("addProxySqlExporterOk"+"."+"proxysql_exporter"+"."+"log_level", "body", *o.LogLevel); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this add proxy SQL exporter OK body proxysql exporter based on context it is used
func (o *AddProxySQLExporterOKBodyProxysqlExporter) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddProxySQLExporterOKBodyProxysqlExporter) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddProxySQLExporterOKBodyProxysqlExporter) UnmarshalBinary(b []byte) error {
	var res AddProxySQLExporterOKBodyProxysqlExporter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
