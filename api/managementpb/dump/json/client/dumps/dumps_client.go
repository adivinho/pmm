// Code generated by go-swagger; DO NOT EDIT.

package dumps

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new dumps API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for dumps API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteDump(params *DeleteDumpParams, opts ...ClientOption) (*DeleteDumpOK, error)

	GetDumpLogs(params *GetDumpLogsParams, opts ...ClientOption) (*GetDumpLogsOK, error)

	ListDumps(params *ListDumpsParams, opts ...ClientOption) (*ListDumpsOK, error)

	StartDump(params *StartDumpParams, opts ...ClientOption) (*StartDumpOK, error)

	UploadDump(params *UploadDumpParams, opts ...ClientOption) (*UploadDumpOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
DeleteDump deletes dump deletes specified pmm dump
*/
func (a *Client) DeleteDump(params *DeleteDumpParams, opts ...ClientOption) (*DeleteDumpOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteDumpParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteDump",
		Method:             "POST",
		PathPattern:        "/v1/management/dump/Dumps/Delete",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteDumpReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteDumpOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteDumpDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetDumpLogs gets logs returns logs from pmm dump tool
*/
func (a *Client) GetDumpLogs(params *GetDumpLogsParams, opts ...ClientOption) (*GetDumpLogsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDumpLogsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetDumpLogs",
		Method:             "POST",
		PathPattern:        "/v1/management/dump/Dumps/GetLogs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetDumpLogsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetDumpLogsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetDumpLogsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListDumps lists dumps returns a list of all pmm dumps
*/
func (a *Client) ListDumps(params *ListDumpsParams, opts ...ClientOption) (*ListDumpsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListDumpsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ListDumps",
		Method:             "POST",
		PathPattern:        "/v1/management/dump/Dumps/List",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListDumpsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListDumpsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListDumpsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
	StartDump starts dump request creates pmm dump

	Could return the Error message in the details containing specific ErrorCode indicating failure reason:

ERROR_CODE_XTRABACKUP_NOT_INSTALLED - xtrabackup is not installed on the service
ERROR_CODE_INVALID_XTRABACKUP - different versions of xtrabackup and xbcloud
ERROR_CODE_INCOMPATIBLE_XTRABACKUP - xtrabackup is not compatible with MySQL for taking a backup
*/
func (a *Client) StartDump(params *StartDumpParams, opts ...ClientOption) (*StartDumpOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStartDumpParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "StartDump",
		Method:             "POST",
		PathPattern:        "/v1/management/dump/Dumps/Start",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &StartDumpReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*StartDumpOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*StartDumpDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UploadDump gets logs returns logs from pmm dump tool
*/
func (a *Client) UploadDump(params *UploadDumpParams, opts ...ClientOption) (*UploadDumpOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUploadDumpParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UploadDump",
		Method:             "POST",
		PathPattern:        "/v1/management/dump/Dumps/Upload",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UploadDumpReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UploadDumpOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UploadDumpDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
