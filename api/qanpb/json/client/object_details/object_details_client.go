// Code generated by go-swagger; DO NOT EDIT.

package object_details

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new object details API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for object details API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	FingerprintAndPlaceholdersCountByQueryID(params *FingerprintAndPlaceholdersCountByQueryIDParams, opts ...ClientOption) (*FingerprintAndPlaceholdersCountByQueryIDOK, error)

	GetHistogram(params *GetHistogramParams, opts ...ClientOption) (*GetHistogramOK, error)

	GetLabels(params *GetLabelsParams, opts ...ClientOption) (*GetLabelsOK, error)

	GetMetrics(params *GetMetricsParams, opts ...ClientOption) (*GetMetricsOK, error)

	GetQueryExample(params *GetQueryExampleParams, opts ...ClientOption) (*GetQueryExampleOK, error)

	GetQueryPlan(params *GetQueryPlanParams, opts ...ClientOption) (*GetQueryPlanOK, error)

	QueryExists(params *QueryExistsParams, opts ...ClientOption) (*QueryExistsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
FingerprintAndPlaceholdersCountByQueryID fingerprints and placeholders count by query ID get fingerprint and placeholders count for given query ID
*/
func (a *Client) FingerprintAndPlaceholdersCountByQueryID(params *FingerprintAndPlaceholdersCountByQueryIDParams, opts ...ClientOption) (*FingerprintAndPlaceholdersCountByQueryIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFingerprintAndPlaceholdersCountByQueryIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "FingerprintAndPlaceholdersCountByQueryID",
		Method:             "POST",
		PathPattern:        "/v0/qan/ObjectDetails/FingerprintAndPlaceholdersCountByQueryID",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &FingerprintAndPlaceholdersCountByQueryIDReader{formats: a.formats},
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
	success, ok := result.(*FingerprintAndPlaceholdersCountByQueryIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*FingerprintAndPlaceholdersCountByQueryIDDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetHistogram gets histogram gets histogram items for specific filtering
*/
func (a *Client) GetHistogram(params *GetHistogramParams, opts ...ClientOption) (*GetHistogramOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetHistogramParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetHistogram",
		Method:             "POST",
		PathPattern:        "/v0/qan/ObjectDetails/GetHistogram",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetHistogramReader{formats: a.formats},
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
	success, ok := result.(*GetHistogramOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetHistogramDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetLabels gets labels gets list of labels for object details
*/
func (a *Client) GetLabels(params *GetLabelsParams, opts ...ClientOption) (*GetLabelsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLabelsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetLabels",
		Method:             "POST",
		PathPattern:        "/v0/qan/ObjectDetails/GetLabels",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetLabelsReader{formats: a.formats},
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
	success, ok := result.(*GetLabelsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetLabelsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetMetrics gets metrics gets map of metrics for specific filtering
*/
func (a *Client) GetMetrics(params *GetMetricsParams, opts ...ClientOption) (*GetMetricsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetMetricsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetMetrics",
		Method:             "POST",
		PathPattern:        "/v0/qan/ObjectDetails/GetMetrics",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetMetricsReader{formats: a.formats},
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
	success, ok := result.(*GetMetricsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetMetricsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetQueryExample gets query example gets list of query examples
*/
func (a *Client) GetQueryExample(params *GetQueryExampleParams, opts ...ClientOption) (*GetQueryExampleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetQueryExampleParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetQueryExample",
		Method:             "POST",
		PathPattern:        "/v0/qan/ObjectDetails/GetQueryExample",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetQueryExampleReader{formats: a.formats},
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
	success, ok := result.(*GetQueryExampleOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetQueryExampleDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetQueryPlan gets query plan gets query plan and plan id for specific filtering
*/
func (a *Client) GetQueryPlan(params *GetQueryPlanParams, opts ...ClientOption) (*GetQueryPlanOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetQueryPlanParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetQueryPlan",
		Method:             "POST",
		PathPattern:        "/v0/qan/ObjectDetails/GetQueryPlan",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetQueryPlanReader{formats: a.formats},
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
	success, ok := result.(*GetQueryPlanOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetQueryPlanDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
QueryExists queries exists check if query exists in clickhouse
*/
func (a *Client) QueryExists(params *QueryExistsParams, opts ...ClientOption) (*QueryExistsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQueryExistsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "QueryExists",
		Method:             "POST",
		PathPattern:        "/v0/qan/ObjectDetails/QueryExists",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &QueryExistsReader{formats: a.formats},
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
	success, ok := result.(*QueryExistsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*QueryExistsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
