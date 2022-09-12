// Code generated by go-swagger; DO NOT EDIT.

package components

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

// GetPSMDBComponentsReader is a Reader for the GetPSMDBComponents structure.
type GetPSMDBComponentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPSMDBComponentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPSMDBComponentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetPSMDBComponentsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetPSMDBComponentsOK creates a GetPSMDBComponentsOK with default headers values
func NewGetPSMDBComponentsOK() *GetPSMDBComponentsOK {
	return &GetPSMDBComponentsOK{}
}

/* GetPSMDBComponentsOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetPSMDBComponentsOK struct {
	Payload *GetPSMDBComponentsOKBody
}

// IsSuccess returns true when this get Psmdb components Ok response has a 2xx status code
func (o *GetPSMDBComponentsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get Psmdb components Ok response has a 3xx status code
func (o *GetPSMDBComponentsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Psmdb components Ok response has a 4xx status code
func (o *GetPSMDBComponentsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get Psmdb components Ok response has a 5xx status code
func (o *GetPSMDBComponentsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get Psmdb components Ok response a status code equal to that given
func (o *GetPSMDBComponentsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetPSMDBComponentsOK) Error() string {
	return fmt.Sprintf("[POST /v1/management/DBaaS/Components/GetPSMDB][%d] getPsmdbComponentsOk  %+v", 200, o.Payload)
}

func (o *GetPSMDBComponentsOK) String() string {
	return fmt.Sprintf("[POST /v1/management/DBaaS/Components/GetPSMDB][%d] getPsmdbComponentsOk  %+v", 200, o.Payload)
}

func (o *GetPSMDBComponentsOK) GetPayload() *GetPSMDBComponentsOKBody {
	return o.Payload
}

func (o *GetPSMDBComponentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(GetPSMDBComponentsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPSMDBComponentsDefault creates a GetPSMDBComponentsDefault with default headers values
func NewGetPSMDBComponentsDefault(code int) *GetPSMDBComponentsDefault {
	return &GetPSMDBComponentsDefault{
		_statusCode: code,
	}
}

/* GetPSMDBComponentsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type GetPSMDBComponentsDefault struct {
	_statusCode int

	Payload *GetPSMDBComponentsDefaultBody
}

// Code gets the status code for the get PSMDB components default response
func (o *GetPSMDBComponentsDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this get PSMDB components default response has a 2xx status code
func (o *GetPSMDBComponentsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get PSMDB components default response has a 3xx status code
func (o *GetPSMDBComponentsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get PSMDB components default response has a 4xx status code
func (o *GetPSMDBComponentsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get PSMDB components default response has a 5xx status code
func (o *GetPSMDBComponentsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get PSMDB components default response a status code equal to that given
func (o *GetPSMDBComponentsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GetPSMDBComponentsDefault) Error() string {
	return fmt.Sprintf("[POST /v1/management/DBaaS/Components/GetPSMDB][%d] GetPSMDBComponents default  %+v", o._statusCode, o.Payload)
}

func (o *GetPSMDBComponentsDefault) String() string {
	return fmt.Sprintf("[POST /v1/management/DBaaS/Components/GetPSMDB][%d] GetPSMDBComponents default  %+v", o._statusCode, o.Payload)
}

func (o *GetPSMDBComponentsDefault) GetPayload() *GetPSMDBComponentsDefaultBody {
	return o.Payload
}

func (o *GetPSMDBComponentsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(GetPSMDBComponentsDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetPSMDBComponentsBody get PSMDB components body
swagger:model GetPSMDBComponentsBody
*/
type GetPSMDBComponentsBody struct {
	// Kubernetes cluster name.
	KubernetesClusterName string `json:"kubernetes_cluster_name,omitempty"`

	// Version of DB.
	DBVersion string `json:"db_version,omitempty"`
}

// Validate validates this get PSMDB components body
func (o *GetPSMDBComponentsBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get PSMDB components body based on context it is used
func (o *GetPSMDBComponentsBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetPSMDBComponentsBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetPSMDBComponentsBody) UnmarshalBinary(b []byte) error {
	var res GetPSMDBComponentsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetPSMDBComponentsDefaultBody get PSMDB components default body
swagger:model GetPSMDBComponentsDefaultBody
*/
type GetPSMDBComponentsDefaultBody struct {
	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*GetPSMDBComponentsDefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this get PSMDB components default body
func (o *GetPSMDBComponentsDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetPSMDBComponentsDefaultBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("GetPSMDBComponents default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("GetPSMDBComponents default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get PSMDB components default body based on the context it is used
func (o *GetPSMDBComponentsDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetPSMDBComponentsDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.Details); i++ {
		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("GetPSMDBComponents default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("GetPSMDBComponents default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetPSMDBComponentsDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetPSMDBComponentsDefaultBody) UnmarshalBinary(b []byte) error {
	var res GetPSMDBComponentsDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetPSMDBComponentsDefaultBodyDetailsItems0 get PSMDB components default body details items0
swagger:model GetPSMDBComponentsDefaultBodyDetailsItems0
*/
type GetPSMDBComponentsDefaultBodyDetailsItems0 struct {
	// at type
	AtType string `json:"@type,omitempty"`
}

// Validate validates this get PSMDB components default body details items0
func (o *GetPSMDBComponentsDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get PSMDB components default body details items0 based on context it is used
func (o *GetPSMDBComponentsDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetPSMDBComponentsDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetPSMDBComponentsDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res GetPSMDBComponentsDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetPSMDBComponentsOKBody get PSMDB components OK body
swagger:model GetPSMDBComponentsOKBody
*/
type GetPSMDBComponentsOKBody struct {
	// versions
	Versions []*GetPSMDBComponentsOKBodyVersionsItems0 `json:"versions"`
}

// Validate validates this get PSMDB components OK body
func (o *GetPSMDBComponentsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateVersions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetPSMDBComponentsOKBody) validateVersions(formats strfmt.Registry) error {
	if swag.IsZero(o.Versions) { // not required
		return nil
	}

	for i := 0; i < len(o.Versions); i++ {
		if swag.IsZero(o.Versions[i]) { // not required
			continue
		}

		if o.Versions[i] != nil {
			if err := o.Versions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getPsmdbComponentsOk" + "." + "versions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getPsmdbComponentsOk" + "." + "versions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get PSMDB components OK body based on the context it is used
func (o *GetPSMDBComponentsOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateVersions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetPSMDBComponentsOKBody) contextValidateVersions(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.Versions); i++ {
		if o.Versions[i] != nil {
			if err := o.Versions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getPsmdbComponentsOk" + "." + "versions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getPsmdbComponentsOk" + "." + "versions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetPSMDBComponentsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetPSMDBComponentsOKBody) UnmarshalBinary(b []byte) error {
	var res GetPSMDBComponentsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetPSMDBComponentsOKBodyVersionsItems0 OperatorVersion contains information about operator and components matrix.
swagger:model GetPSMDBComponentsOKBodyVersionsItems0
*/
type GetPSMDBComponentsOKBodyVersionsItems0 struct {
	// product
	Product string `json:"product,omitempty"`

	// operator
	Operator string `json:"operator,omitempty"`

	// matrix
	Matrix *GetPSMDBComponentsOKBodyVersionsItems0Matrix `json:"matrix,omitempty"`
}

// Validate validates this get PSMDB components OK body versions items0
func (o *GetPSMDBComponentsOKBodyVersionsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMatrix(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetPSMDBComponentsOKBodyVersionsItems0) validateMatrix(formats strfmt.Registry) error {
	if swag.IsZero(o.Matrix) { // not required
		return nil
	}

	if o.Matrix != nil {
		if err := o.Matrix.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("matrix")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("matrix")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get PSMDB components OK body versions items0 based on the context it is used
func (o *GetPSMDBComponentsOKBodyVersionsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateMatrix(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetPSMDBComponentsOKBodyVersionsItems0) contextValidateMatrix(ctx context.Context, formats strfmt.Registry) error {
	if o.Matrix != nil {
		if err := o.Matrix.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("matrix")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("matrix")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetPSMDBComponentsOKBodyVersionsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetPSMDBComponentsOKBodyVersionsItems0) UnmarshalBinary(b []byte) error {
	var res GetPSMDBComponentsOKBodyVersionsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetPSMDBComponentsOKBodyVersionsItems0Matrix Matrix contains all available components.
swagger:model GetPSMDBComponentsOKBodyVersionsItems0Matrix
*/
type GetPSMDBComponentsOKBodyVersionsItems0Matrix struct {
	// mongod
	Mongod map[string]GetPSMDBComponentsOKBodyVersionsItems0MatrixMongodAnon `json:"mongod,omitempty"`

	// pxc
	PXC map[string]GetPSMDBComponentsOKBodyVersionsItems0MatrixPXCAnon `json:"pxc,omitempty"`

	// pmm
	PMM map[string]GetPSMDBComponentsOKBodyVersionsItems0MatrixPMMAnon `json:"pmm,omitempty"`

	// proxysql
	Proxysql map[string]GetPSMDBComponentsOKBodyVersionsItems0MatrixProxysqlAnon `json:"proxysql,omitempty"`

	// haproxy
	Haproxy map[string]GetPSMDBComponentsOKBodyVersionsItems0MatrixHaproxyAnon `json:"haproxy,omitempty"`

	// backup
	Backup map[string]GetPSMDBComponentsOKBodyVersionsItems0MatrixBackupAnon `json:"backup,omitempty"`

	// operator
	Operator map[string]GetPSMDBComponentsOKBodyVersionsItems0MatrixOperatorAnon `json:"operator,omitempty"`

	// log collector
	LogCollector map[string]GetPSMDBComponentsOKBodyVersionsItems0MatrixLogCollectorAnon `json:"log_collector,omitempty"`
}

// Validate validates this get PSMDB components OK body versions items0 matrix
func (o *GetPSMDBComponentsOKBodyVersionsItems0Matrix) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMongod(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePXC(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePMM(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateProxysql(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateHaproxy(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateBackup(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateOperator(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateLogCollector(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetPSMDBComponentsOKBodyVersionsItems0Matrix) validateMongod(formats strfmt.Registry) error {
	if swag.IsZero(o.Mongod) { // not required
		return nil
	}

	for k := range o.Mongod {

		if swag.IsZero(o.Mongod[k]) { // not required
			continue
		}
		if val, ok := o.Mongod[k]; ok {
			if err := val.Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("matrix" + "." + "mongod" + "." + k)
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("matrix" + "." + "mongod" + "." + k)
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetPSMDBComponentsOKBodyVersionsItems0Matrix) validatePXC(formats strfmt.Registry) error {
	if swag.IsZero(o.PXC) { // not required
		return nil
	}

	for k := range o.PXC {

		if swag.IsZero(o.PXC[k]) { // not required
			continue
		}
		if val, ok := o.PXC[k]; ok {
			if err := val.Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("matrix" + "." + "pxc" + "." + k)
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("matrix" + "." + "pxc" + "." + k)
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetPSMDBComponentsOKBodyVersionsItems0Matrix) validatePMM(formats strfmt.Registry) error {
	if swag.IsZero(o.PMM) { // not required
		return nil
	}

	for k := range o.PMM {

		if swag.IsZero(o.PMM[k]) { // not required
			continue
		}
		if val, ok := o.PMM[k]; ok {
			if err := val.Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("matrix" + "." + "pmm" + "." + k)
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("matrix" + "." + "pmm" + "." + k)
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetPSMDBComponentsOKBodyVersionsItems0Matrix) validateProxysql(formats strfmt.Registry) error {
	if swag.IsZero(o.Proxysql) { // not required
		return nil
	}

	for k := range o.Proxysql {

		if swag.IsZero(o.Proxysql[k]) { // not required
			continue
		}
		if val, ok := o.Proxysql[k]; ok {
			if err := val.Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("matrix" + "." + "proxysql" + "." + k)
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("matrix" + "." + "proxysql" + "." + k)
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetPSMDBComponentsOKBodyVersionsItems0Matrix) validateHaproxy(formats strfmt.Registry) error {
	if swag.IsZero(o.Haproxy) { // not required
		return nil
	}

	for k := range o.Haproxy {

		if swag.IsZero(o.Haproxy[k]) { // not required
			continue
		}
		if val, ok := o.Haproxy[k]; ok {
			if err := val.Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("matrix" + "." + "haproxy" + "." + k)
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("matrix" + "." + "haproxy" + "." + k)
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetPSMDBComponentsOKBodyVersionsItems0Matrix) validateBackup(formats strfmt.Registry) error {
	if swag.IsZero(o.Backup) { // not required
		return nil
	}

	for k := range o.Backup {

		if swag.IsZero(o.Backup[k]) { // not required
			continue
		}
		if val, ok := o.Backup[k]; ok {
			if err := val.Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("matrix" + "." + "backup" + "." + k)
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("matrix" + "." + "backup" + "." + k)
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetPSMDBComponentsOKBodyVersionsItems0Matrix) validateOperator(formats strfmt.Registry) error {
	if swag.IsZero(o.Operator) { // not required
		return nil
	}

	for k := range o.Operator {

		if swag.IsZero(o.Operator[k]) { // not required
			continue
		}
		if val, ok := o.Operator[k]; ok {
			if err := val.Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("matrix" + "." + "operator" + "." + k)
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("matrix" + "." + "operator" + "." + k)
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetPSMDBComponentsOKBodyVersionsItems0Matrix) validateLogCollector(formats strfmt.Registry) error {
	if swag.IsZero(o.LogCollector) { // not required
		return nil
	}

	for k := range o.LogCollector {

		if swag.IsZero(o.LogCollector[k]) { // not required
			continue
		}
		if val, ok := o.LogCollector[k]; ok {
			if err := val.Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("matrix" + "." + "log_collector" + "." + k)
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("matrix" + "." + "log_collector" + "." + k)
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get PSMDB components OK body versions items0 matrix based on the context it is used
func (o *GetPSMDBComponentsOKBodyVersionsItems0Matrix) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateMongod(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidatePXC(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidatePMM(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateProxysql(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateHaproxy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateBackup(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateOperator(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateLogCollector(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetPSMDBComponentsOKBodyVersionsItems0Matrix) contextValidateMongod(ctx context.Context, formats strfmt.Registry) error {
	for k := range o.Mongod {
		if val, ok := o.Mongod[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}
	}

	return nil
}

func (o *GetPSMDBComponentsOKBodyVersionsItems0Matrix) contextValidatePXC(ctx context.Context, formats strfmt.Registry) error {
	for k := range o.PXC {
		if val, ok := o.PXC[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}
	}

	return nil
}

func (o *GetPSMDBComponentsOKBodyVersionsItems0Matrix) contextValidatePMM(ctx context.Context, formats strfmt.Registry) error {
	for k := range o.PMM {
		if val, ok := o.PMM[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}
	}

	return nil
}

func (o *GetPSMDBComponentsOKBodyVersionsItems0Matrix) contextValidateProxysql(ctx context.Context, formats strfmt.Registry) error {
	for k := range o.Proxysql {
		if val, ok := o.Proxysql[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}
	}

	return nil
}

func (o *GetPSMDBComponentsOKBodyVersionsItems0Matrix) contextValidateHaproxy(ctx context.Context, formats strfmt.Registry) error {
	for k := range o.Haproxy {
		if val, ok := o.Haproxy[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}
	}

	return nil
}

func (o *GetPSMDBComponentsOKBodyVersionsItems0Matrix) contextValidateBackup(ctx context.Context, formats strfmt.Registry) error {
	for k := range o.Backup {
		if val, ok := o.Backup[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}
	}

	return nil
}

func (o *GetPSMDBComponentsOKBodyVersionsItems0Matrix) contextValidateOperator(ctx context.Context, formats strfmt.Registry) error {
	for k := range o.Operator {
		if val, ok := o.Operator[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}
	}

	return nil
}

func (o *GetPSMDBComponentsOKBodyVersionsItems0Matrix) contextValidateLogCollector(ctx context.Context, formats strfmt.Registry) error {
	for k := range o.LogCollector {
		if val, ok := o.LogCollector[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetPSMDBComponentsOKBodyVersionsItems0Matrix) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetPSMDBComponentsOKBodyVersionsItems0Matrix) UnmarshalBinary(b []byte) error {
	var res GetPSMDBComponentsOKBodyVersionsItems0Matrix
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetPSMDBComponentsOKBodyVersionsItems0MatrixBackupAnon Component contains information about component.
swagger:model GetPSMDBComponentsOKBodyVersionsItems0MatrixBackupAnon
*/
type GetPSMDBComponentsOKBodyVersionsItems0MatrixBackupAnon struct {
	// image path
	ImagePath string `json:"image_path,omitempty"`

	// image hash
	ImageHash string `json:"image_hash,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// critical
	Critical bool `json:"critical,omitempty"`

	// default
	Default bool `json:"default,omitempty"`

	// disabled
	Disabled bool `json:"disabled,omitempty"`
}

// Validate validates this get PSMDB components OK body versions items0 matrix backup anon
func (o *GetPSMDBComponentsOKBodyVersionsItems0MatrixBackupAnon) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get PSMDB components OK body versions items0 matrix backup anon based on context it is used
func (o *GetPSMDBComponentsOKBodyVersionsItems0MatrixBackupAnon) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetPSMDBComponentsOKBodyVersionsItems0MatrixBackupAnon) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetPSMDBComponentsOKBodyVersionsItems0MatrixBackupAnon) UnmarshalBinary(b []byte) error {
	var res GetPSMDBComponentsOKBodyVersionsItems0MatrixBackupAnon
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetPSMDBComponentsOKBodyVersionsItems0MatrixHaproxyAnon Component contains information about component.
swagger:model GetPSMDBComponentsOKBodyVersionsItems0MatrixHaproxyAnon
*/
type GetPSMDBComponentsOKBodyVersionsItems0MatrixHaproxyAnon struct {
	// image path
	ImagePath string `json:"image_path,omitempty"`

	// image hash
	ImageHash string `json:"image_hash,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// critical
	Critical bool `json:"critical,omitempty"`

	// default
	Default bool `json:"default,omitempty"`

	// disabled
	Disabled bool `json:"disabled,omitempty"`
}

// Validate validates this get PSMDB components OK body versions items0 matrix haproxy anon
func (o *GetPSMDBComponentsOKBodyVersionsItems0MatrixHaproxyAnon) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get PSMDB components OK body versions items0 matrix haproxy anon based on context it is used
func (o *GetPSMDBComponentsOKBodyVersionsItems0MatrixHaproxyAnon) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetPSMDBComponentsOKBodyVersionsItems0MatrixHaproxyAnon) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetPSMDBComponentsOKBodyVersionsItems0MatrixHaproxyAnon) UnmarshalBinary(b []byte) error {
	var res GetPSMDBComponentsOKBodyVersionsItems0MatrixHaproxyAnon
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetPSMDBComponentsOKBodyVersionsItems0MatrixLogCollectorAnon Component contains information about component.
swagger:model GetPSMDBComponentsOKBodyVersionsItems0MatrixLogCollectorAnon
*/
type GetPSMDBComponentsOKBodyVersionsItems0MatrixLogCollectorAnon struct {
	// image path
	ImagePath string `json:"image_path,omitempty"`

	// image hash
	ImageHash string `json:"image_hash,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// critical
	Critical bool `json:"critical,omitempty"`

	// default
	Default bool `json:"default,omitempty"`

	// disabled
	Disabled bool `json:"disabled,omitempty"`
}

// Validate validates this get PSMDB components OK body versions items0 matrix log collector anon
func (o *GetPSMDBComponentsOKBodyVersionsItems0MatrixLogCollectorAnon) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get PSMDB components OK body versions items0 matrix log collector anon based on context it is used
func (o *GetPSMDBComponentsOKBodyVersionsItems0MatrixLogCollectorAnon) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetPSMDBComponentsOKBodyVersionsItems0MatrixLogCollectorAnon) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetPSMDBComponentsOKBodyVersionsItems0MatrixLogCollectorAnon) UnmarshalBinary(b []byte) error {
	var res GetPSMDBComponentsOKBodyVersionsItems0MatrixLogCollectorAnon
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetPSMDBComponentsOKBodyVersionsItems0MatrixMongodAnon Component contains information about component.
swagger:model GetPSMDBComponentsOKBodyVersionsItems0MatrixMongodAnon
*/
type GetPSMDBComponentsOKBodyVersionsItems0MatrixMongodAnon struct {
	// image path
	ImagePath string `json:"image_path,omitempty"`

	// image hash
	ImageHash string `json:"image_hash,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// critical
	Critical bool `json:"critical,omitempty"`

	// default
	Default bool `json:"default,omitempty"`

	// disabled
	Disabled bool `json:"disabled,omitempty"`
}

// Validate validates this get PSMDB components OK body versions items0 matrix mongod anon
func (o *GetPSMDBComponentsOKBodyVersionsItems0MatrixMongodAnon) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get PSMDB components OK body versions items0 matrix mongod anon based on context it is used
func (o *GetPSMDBComponentsOKBodyVersionsItems0MatrixMongodAnon) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetPSMDBComponentsOKBodyVersionsItems0MatrixMongodAnon) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetPSMDBComponentsOKBodyVersionsItems0MatrixMongodAnon) UnmarshalBinary(b []byte) error {
	var res GetPSMDBComponentsOKBodyVersionsItems0MatrixMongodAnon
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetPSMDBComponentsOKBodyVersionsItems0MatrixOperatorAnon Component contains information about component.
swagger:model GetPSMDBComponentsOKBodyVersionsItems0MatrixOperatorAnon
*/
type GetPSMDBComponentsOKBodyVersionsItems0MatrixOperatorAnon struct {
	// image path
	ImagePath string `json:"image_path,omitempty"`

	// image hash
	ImageHash string `json:"image_hash,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// critical
	Critical bool `json:"critical,omitempty"`

	// default
	Default bool `json:"default,omitempty"`

	// disabled
	Disabled bool `json:"disabled,omitempty"`
}

// Validate validates this get PSMDB components OK body versions items0 matrix operator anon
func (o *GetPSMDBComponentsOKBodyVersionsItems0MatrixOperatorAnon) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get PSMDB components OK body versions items0 matrix operator anon based on context it is used
func (o *GetPSMDBComponentsOKBodyVersionsItems0MatrixOperatorAnon) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetPSMDBComponentsOKBodyVersionsItems0MatrixOperatorAnon) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetPSMDBComponentsOKBodyVersionsItems0MatrixOperatorAnon) UnmarshalBinary(b []byte) error {
	var res GetPSMDBComponentsOKBodyVersionsItems0MatrixOperatorAnon
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetPSMDBComponentsOKBodyVersionsItems0MatrixPMMAnon Component contains information about component.
swagger:model GetPSMDBComponentsOKBodyVersionsItems0MatrixPMMAnon
*/
type GetPSMDBComponentsOKBodyVersionsItems0MatrixPMMAnon struct {
	// image path
	ImagePath string `json:"image_path,omitempty"`

	// image hash
	ImageHash string `json:"image_hash,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// critical
	Critical bool `json:"critical,omitempty"`

	// default
	Default bool `json:"default,omitempty"`

	// disabled
	Disabled bool `json:"disabled,omitempty"`
}

// Validate validates this get PSMDB components OK body versions items0 matrix PMM anon
func (o *GetPSMDBComponentsOKBodyVersionsItems0MatrixPMMAnon) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get PSMDB components OK body versions items0 matrix PMM anon based on context it is used
func (o *GetPSMDBComponentsOKBodyVersionsItems0MatrixPMMAnon) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetPSMDBComponentsOKBodyVersionsItems0MatrixPMMAnon) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetPSMDBComponentsOKBodyVersionsItems0MatrixPMMAnon) UnmarshalBinary(b []byte) error {
	var res GetPSMDBComponentsOKBodyVersionsItems0MatrixPMMAnon
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetPSMDBComponentsOKBodyVersionsItems0MatrixPXCAnon Component contains information about component.
swagger:model GetPSMDBComponentsOKBodyVersionsItems0MatrixPXCAnon
*/
type GetPSMDBComponentsOKBodyVersionsItems0MatrixPXCAnon struct {
	// image path
	ImagePath string `json:"image_path,omitempty"`

	// image hash
	ImageHash string `json:"image_hash,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// critical
	Critical bool `json:"critical,omitempty"`

	// default
	Default bool `json:"default,omitempty"`

	// disabled
	Disabled bool `json:"disabled,omitempty"`
}

// Validate validates this get PSMDB components OK body versions items0 matrix PXC anon
func (o *GetPSMDBComponentsOKBodyVersionsItems0MatrixPXCAnon) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get PSMDB components OK body versions items0 matrix PXC anon based on context it is used
func (o *GetPSMDBComponentsOKBodyVersionsItems0MatrixPXCAnon) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetPSMDBComponentsOKBodyVersionsItems0MatrixPXCAnon) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetPSMDBComponentsOKBodyVersionsItems0MatrixPXCAnon) UnmarshalBinary(b []byte) error {
	var res GetPSMDBComponentsOKBodyVersionsItems0MatrixPXCAnon
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetPSMDBComponentsOKBodyVersionsItems0MatrixProxysqlAnon Component contains information about component.
swagger:model GetPSMDBComponentsOKBodyVersionsItems0MatrixProxysqlAnon
*/
type GetPSMDBComponentsOKBodyVersionsItems0MatrixProxysqlAnon struct {
	// image path
	ImagePath string `json:"image_path,omitempty"`

	// image hash
	ImageHash string `json:"image_hash,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// critical
	Critical bool `json:"critical,omitempty"`

	// default
	Default bool `json:"default,omitempty"`

	// disabled
	Disabled bool `json:"disabled,omitempty"`
}

// Validate validates this get PSMDB components OK body versions items0 matrix proxysql anon
func (o *GetPSMDBComponentsOKBodyVersionsItems0MatrixProxysqlAnon) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get PSMDB components OK body versions items0 matrix proxysql anon based on context it is used
func (o *GetPSMDBComponentsOKBodyVersionsItems0MatrixProxysqlAnon) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetPSMDBComponentsOKBodyVersionsItems0MatrixProxysqlAnon) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetPSMDBComponentsOKBodyVersionsItems0MatrixProxysqlAnon) UnmarshalBinary(b []byte) error {
	var res GetPSMDBComponentsOKBodyVersionsItems0MatrixProxysqlAnon
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
