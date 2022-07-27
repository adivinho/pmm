// Code generated by go-swagger; DO NOT EDIT.

package metrics_names

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetMetricsNamesParams creates a new GetMetricsNamesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetMetricsNamesParams() *GetMetricsNamesParams {
	return &GetMetricsNamesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetMetricsNamesParamsWithTimeout creates a new GetMetricsNamesParams object
// with the ability to set a timeout on a request.
func NewGetMetricsNamesParamsWithTimeout(timeout time.Duration) *GetMetricsNamesParams {
	return &GetMetricsNamesParams{
		timeout: timeout,
	}
}

// NewGetMetricsNamesParamsWithContext creates a new GetMetricsNamesParams object
// with the ability to set a context for a request.
func NewGetMetricsNamesParamsWithContext(ctx context.Context) *GetMetricsNamesParams {
	return &GetMetricsNamesParams{
		Context: ctx,
	}
}

// NewGetMetricsNamesParamsWithHTTPClient creates a new GetMetricsNamesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetMetricsNamesParamsWithHTTPClient(client *http.Client) *GetMetricsNamesParams {
	return &GetMetricsNamesParams{
		HTTPClient: client,
	}
}

/* GetMetricsNamesParams contains all the parameters to send to the API endpoint
   for the get metrics names operation.

   Typically these are written to a http.Request.
*/
type GetMetricsNamesParams struct {
	/* Body.

	   MetricsNamesRequest is emty.
	*/
	Body interface{}

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get metrics names params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetMetricsNamesParams) WithDefaults() *GetMetricsNamesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get metrics names params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetMetricsNamesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get metrics names params
func (o *GetMetricsNamesParams) WithTimeout(timeout time.Duration) *GetMetricsNamesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get metrics names params
func (o *GetMetricsNamesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get metrics names params
func (o *GetMetricsNamesParams) WithContext(ctx context.Context) *GetMetricsNamesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get metrics names params
func (o *GetMetricsNamesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get metrics names params
func (o *GetMetricsNamesParams) WithHTTPClient(client *http.Client) *GetMetricsNamesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get metrics names params
func (o *GetMetricsNamesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the get metrics names params
func (o *GetMetricsNamesParams) WithBody(body interface{}) *GetMetricsNamesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the get metrics names params
func (o *GetMetricsNamesParams) SetBody(body interface{}) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *GetMetricsNamesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {
	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
