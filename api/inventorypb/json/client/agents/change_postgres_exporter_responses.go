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

// ChangePostgresExporterReader is a Reader for the ChangePostgresExporter structure.
type ChangePostgresExporterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChangePostgresExporterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewChangePostgresExporterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewChangePostgresExporterDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewChangePostgresExporterOK creates a ChangePostgresExporterOK with default headers values
func NewChangePostgresExporterOK() *ChangePostgresExporterOK {
	return &ChangePostgresExporterOK{}
}

/*
ChangePostgresExporterOK describes a response with status code 200, with default header values.

A successful response.
*/
type ChangePostgresExporterOK struct {
	Payload *ChangePostgresExporterOKBody
}

func (o *ChangePostgresExporterOK) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Agents/ChangePostgresExporter][%d] changePostgresExporterOk  %+v", 200, o.Payload)
}

func (o *ChangePostgresExporterOK) GetPayload() *ChangePostgresExporterOKBody {
	return o.Payload
}

func (o *ChangePostgresExporterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(ChangePostgresExporterOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewChangePostgresExporterDefault creates a ChangePostgresExporterDefault with default headers values
func NewChangePostgresExporterDefault(code int) *ChangePostgresExporterDefault {
	return &ChangePostgresExporterDefault{
		_statusCode: code,
	}
}

/*
ChangePostgresExporterDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ChangePostgresExporterDefault struct {
	_statusCode int

	Payload *ChangePostgresExporterDefaultBody
}

// Code gets the status code for the change postgres exporter default response
func (o *ChangePostgresExporterDefault) Code() int {
	return o._statusCode
}

func (o *ChangePostgresExporterDefault) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Agents/ChangePostgresExporter][%d] ChangePostgresExporter default  %+v", o._statusCode, o.Payload)
}

func (o *ChangePostgresExporterDefault) GetPayload() *ChangePostgresExporterDefaultBody {
	return o.Payload
}

func (o *ChangePostgresExporterDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(ChangePostgresExporterDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
ChangePostgresExporterBody change postgres exporter body
swagger:model ChangePostgresExporterBody
*/
type ChangePostgresExporterBody struct {
	// agent id
	AgentID string `json:"agent_id,omitempty"`

	// common
	Common *ChangePostgresExporterParamsBodyCommon `json:"common,omitempty"`
}

// Validate validates this change postgres exporter body
func (o *ChangePostgresExporterBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCommon(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangePostgresExporterBody) validateCommon(formats strfmt.Registry) error {
	if swag.IsZero(o.Common) { // not required
		return nil
	}

	if o.Common != nil {
		if err := o.Common.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "common")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "common")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this change postgres exporter body based on the context it is used
func (o *ChangePostgresExporterBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateCommon(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangePostgresExporterBody) contextValidateCommon(ctx context.Context, formats strfmt.Registry) error {
	if o.Common != nil {
		if err := o.Common.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "common")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "common")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ChangePostgresExporterBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangePostgresExporterBody) UnmarshalBinary(b []byte) error {
	var res ChangePostgresExporterBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ChangePostgresExporterDefaultBody change postgres exporter default body
swagger:model ChangePostgresExporterDefaultBody
*/
type ChangePostgresExporterDefaultBody struct {
	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*ChangePostgresExporterDefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this change postgres exporter default body
func (o *ChangePostgresExporterDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangePostgresExporterDefaultBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("ChangePostgresExporter default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ChangePostgresExporter default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this change postgres exporter default body based on the context it is used
func (o *ChangePostgresExporterDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangePostgresExporterDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.Details); i++ {
		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ChangePostgresExporter default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ChangePostgresExporter default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ChangePostgresExporterDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangePostgresExporterDefaultBody) UnmarshalBinary(b []byte) error {
	var res ChangePostgresExporterDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ChangePostgresExporterDefaultBodyDetailsItems0 change postgres exporter default body details items0
swagger:model ChangePostgresExporterDefaultBodyDetailsItems0
*/
type ChangePostgresExporterDefaultBodyDetailsItems0 struct {
	// at type
	AtType string `json:"@type,omitempty"`
}

// Validate validates this change postgres exporter default body details items0
func (o *ChangePostgresExporterDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this change postgres exporter default body details items0 based on context it is used
func (o *ChangePostgresExporterDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ChangePostgresExporterDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangePostgresExporterDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res ChangePostgresExporterDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ChangePostgresExporterOKBody change postgres exporter OK body
swagger:model ChangePostgresExporterOKBody
*/
type ChangePostgresExporterOKBody struct {
	// postgres exporter
	PostgresExporter *ChangePostgresExporterOKBodyPostgresExporter `json:"postgres_exporter,omitempty"`
}

// Validate validates this change postgres exporter OK body
func (o *ChangePostgresExporterOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validatePostgresExporter(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangePostgresExporterOKBody) validatePostgresExporter(formats strfmt.Registry) error {
	if swag.IsZero(o.PostgresExporter) { // not required
		return nil
	}

	if o.PostgresExporter != nil {
		if err := o.PostgresExporter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("changePostgresExporterOk" + "." + "postgres_exporter")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("changePostgresExporterOk" + "." + "postgres_exporter")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this change postgres exporter OK body based on the context it is used
func (o *ChangePostgresExporterOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidatePostgresExporter(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangePostgresExporterOKBody) contextValidatePostgresExporter(ctx context.Context, formats strfmt.Registry) error {
	if o.PostgresExporter != nil {
		if err := o.PostgresExporter.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("changePostgresExporterOk" + "." + "postgres_exporter")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("changePostgresExporterOk" + "." + "postgres_exporter")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ChangePostgresExporterOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangePostgresExporterOKBody) UnmarshalBinary(b []byte) error {
	var res ChangePostgresExporterOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ChangePostgresExporterOKBodyPostgresExporter PostgresExporter runs on Generic or Container Node and exposes PostgreSQL Service metrics.
swagger:model ChangePostgresExporterOKBodyPostgresExporter
*/
type ChangePostgresExporterOKBodyPostgresExporter struct {
	// Unique randomly generated instance identifier.
	AgentID string `json:"agent_id,omitempty"`

	// The pmm-agent identifier which runs this instance.
	PMMAgentID string `json:"pmm_agent_id,omitempty"`

	// Desired Agent status: enabled (false) or disabled (true).
	Disabled bool `json:"disabled,omitempty"`

	// Service identifier.
	ServiceID string `json:"service_id,omitempty"`

	// PostgreSQL username for scraping metrics.
	Username string `json:"username,omitempty"`

	// Use TLS for database connections.
	TLS bool `json:"tls,omitempty"`

	// Skip TLS certificate and hostname validation. Uses sslmode=required instead of verify-full.
	TLSSkipVerify bool `json:"tls_skip_verify,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// True if exporter uses push metrics mode.
	PushMetricsEnabled bool `json:"push_metrics_enabled,omitempty"`

	// List of disabled collector names.
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

// Validate validates this change postgres exporter OK body postgres exporter
func (o *ChangePostgresExporterOKBodyPostgresExporter) Validate(formats strfmt.Registry) error {
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

var changePostgresExporterOkBodyPostgresExporterTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AGENT_STATUS_INVALID","STARTING","RUNNING","WAITING","STOPPING","DONE","UNKNOWN"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		changePostgresExporterOkBodyPostgresExporterTypeStatusPropEnum = append(changePostgresExporterOkBodyPostgresExporterTypeStatusPropEnum, v)
	}
}

const (

	// ChangePostgresExporterOKBodyPostgresExporterStatusAGENTSTATUSINVALID captures enum value "AGENT_STATUS_INVALID"
	ChangePostgresExporterOKBodyPostgresExporterStatusAGENTSTATUSINVALID string = "AGENT_STATUS_INVALID"

	// ChangePostgresExporterOKBodyPostgresExporterStatusSTARTING captures enum value "STARTING"
	ChangePostgresExporterOKBodyPostgresExporterStatusSTARTING string = "STARTING"

	// ChangePostgresExporterOKBodyPostgresExporterStatusRUNNING captures enum value "RUNNING"
	ChangePostgresExporterOKBodyPostgresExporterStatusRUNNING string = "RUNNING"

	// ChangePostgresExporterOKBodyPostgresExporterStatusWAITING captures enum value "WAITING"
	ChangePostgresExporterOKBodyPostgresExporterStatusWAITING string = "WAITING"

	// ChangePostgresExporterOKBodyPostgresExporterStatusSTOPPING captures enum value "STOPPING"
	ChangePostgresExporterOKBodyPostgresExporterStatusSTOPPING string = "STOPPING"

	// ChangePostgresExporterOKBodyPostgresExporterStatusDONE captures enum value "DONE"
	ChangePostgresExporterOKBodyPostgresExporterStatusDONE string = "DONE"

	// ChangePostgresExporterOKBodyPostgresExporterStatusUNKNOWN captures enum value "UNKNOWN"
	ChangePostgresExporterOKBodyPostgresExporterStatusUNKNOWN string = "UNKNOWN"
)

// prop value enum
func (o *ChangePostgresExporterOKBodyPostgresExporter) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, changePostgresExporterOkBodyPostgresExporterTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *ChangePostgresExporterOKBodyPostgresExporter) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(o.Status) { // not required
		return nil
	}

	// value enum
	if err := o.validateStatusEnum("changePostgresExporterOk"+"."+"postgres_exporter"+"."+"status", "body", *o.Status); err != nil {
		return err
	}

	return nil
}

var changePostgresExporterOkBodyPostgresExporterTypeLogLevelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["auto","fatal","error","warn","info","debug"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		changePostgresExporterOkBodyPostgresExporterTypeLogLevelPropEnum = append(changePostgresExporterOkBodyPostgresExporterTypeLogLevelPropEnum, v)
	}
}

const (

	// ChangePostgresExporterOKBodyPostgresExporterLogLevelAuto captures enum value "auto"
	ChangePostgresExporterOKBodyPostgresExporterLogLevelAuto string = "auto"

	// ChangePostgresExporterOKBodyPostgresExporterLogLevelFatal captures enum value "fatal"
	ChangePostgresExporterOKBodyPostgresExporterLogLevelFatal string = "fatal"

	// ChangePostgresExporterOKBodyPostgresExporterLogLevelError captures enum value "error"
	ChangePostgresExporterOKBodyPostgresExporterLogLevelError string = "error"

	// ChangePostgresExporterOKBodyPostgresExporterLogLevelWarn captures enum value "warn"
	ChangePostgresExporterOKBodyPostgresExporterLogLevelWarn string = "warn"

	// ChangePostgresExporterOKBodyPostgresExporterLogLevelInfo captures enum value "info"
	ChangePostgresExporterOKBodyPostgresExporterLogLevelInfo string = "info"

	// ChangePostgresExporterOKBodyPostgresExporterLogLevelDebug captures enum value "debug"
	ChangePostgresExporterOKBodyPostgresExporterLogLevelDebug string = "debug"
)

// prop value enum
func (o *ChangePostgresExporterOKBodyPostgresExporter) validateLogLevelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, changePostgresExporterOkBodyPostgresExporterTypeLogLevelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *ChangePostgresExporterOKBodyPostgresExporter) validateLogLevel(formats strfmt.Registry) error {
	if swag.IsZero(o.LogLevel) { // not required
		return nil
	}

	// value enum
	if err := o.validateLogLevelEnum("changePostgresExporterOk"+"."+"postgres_exporter"+"."+"log_level", "body", *o.LogLevel); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this change postgres exporter OK body postgres exporter based on context it is used
func (o *ChangePostgresExporterOKBodyPostgresExporter) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ChangePostgresExporterOKBodyPostgresExporter) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangePostgresExporterOKBodyPostgresExporter) UnmarshalBinary(b []byte) error {
	var res ChangePostgresExporterOKBodyPostgresExporter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ChangePostgresExporterParamsBodyCommon ChangeCommonAgentParams contains parameters that can be changed for all Agents.
swagger:model ChangePostgresExporterParamsBodyCommon
*/
type ChangePostgresExporterParamsBodyCommon struct {
	// Enable this Agent. Can't be used with disabled.
	Enable bool `json:"enable,omitempty"`

	// Disable this Agent. Can't be used with enabled.
	Disable bool `json:"disable,omitempty"`

	// Replace all custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// Remove all custom user-assigned labels.
	RemoveCustomLabels bool `json:"remove_custom_labels,omitempty"`

	// Enables push metrics with vmagent, can't be used with disable_push_metrics.
	// Can't be used with agent version lower then 2.12 and unsupported agents.
	EnablePushMetrics bool `json:"enable_push_metrics,omitempty"`

	// Disables push metrics, pmm-server starts to pull it, can't be used with enable_push_metrics.
	DisablePushMetrics bool `json:"disable_push_metrics,omitempty"`
}

// Validate validates this change postgres exporter params body common
func (o *ChangePostgresExporterParamsBodyCommon) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this change postgres exporter params body common based on context it is used
func (o *ChangePostgresExporterParamsBodyCommon) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ChangePostgresExporterParamsBodyCommon) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangePostgresExporterParamsBodyCommon) UnmarshalBinary(b []byte) error {
	var res ChangePostgresExporterParamsBodyCommon
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
