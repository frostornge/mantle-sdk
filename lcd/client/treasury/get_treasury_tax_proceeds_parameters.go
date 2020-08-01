// Code generated by go-swagger; DO NOT EDIT.

package treasury

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

// NewGetTreasuryTaxProceedsParams creates a new GetTreasuryTaxProceedsParams object
// with the default values initialized.
func NewGetTreasuryTaxProceedsParams() *GetTreasuryTaxProceedsParams {

	return &GetTreasuryTaxProceedsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetTreasuryTaxProceedsParamsWithTimeout creates a new GetTreasuryTaxProceedsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetTreasuryTaxProceedsParamsWithTimeout(timeout time.Duration) *GetTreasuryTaxProceedsParams {

	return &GetTreasuryTaxProceedsParams{

		timeout: timeout,
	}
}

// NewGetTreasuryTaxProceedsParamsWithContext creates a new GetTreasuryTaxProceedsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetTreasuryTaxProceedsParamsWithContext(ctx context.Context) *GetTreasuryTaxProceedsParams {

	return &GetTreasuryTaxProceedsParams{

		Context: ctx,
	}
}

// NewGetTreasuryTaxProceedsParamsWithHTTPClient creates a new GetTreasuryTaxProceedsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetTreasuryTaxProceedsParamsWithHTTPClient(client *http.Client) *GetTreasuryTaxProceedsParams {

	return &GetTreasuryTaxProceedsParams{
		HTTPClient: client,
	}
}

/*GetTreasuryTaxProceedsParams contains all the parameters to send to the API endpoint
for the get treasury tax proceeds operation typically these are written to a http.Request
*/
type GetTreasuryTaxProceedsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get treasury tax proceeds params
func (o *GetTreasuryTaxProceedsParams) WithTimeout(timeout time.Duration) *GetTreasuryTaxProceedsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get treasury tax proceeds params
func (o *GetTreasuryTaxProceedsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get treasury tax proceeds params
func (o *GetTreasuryTaxProceedsParams) WithContext(ctx context.Context) *GetTreasuryTaxProceedsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get treasury tax proceeds params
func (o *GetTreasuryTaxProceedsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get treasury tax proceeds params
func (o *GetTreasuryTaxProceedsParams) WithHTTPClient(client *http.Client) *GetTreasuryTaxProceedsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get treasury tax proceeds params
func (o *GetTreasuryTaxProceedsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetTreasuryTaxProceedsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}