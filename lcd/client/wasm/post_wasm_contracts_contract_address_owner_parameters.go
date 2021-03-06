// Code generated by go-swagger; DO NOT EDIT.

package wasm

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

	"github.com/terra-project/mantle-sdk/lcd/models"
)

// NewPostWasmContractsContractAddressOwnerParams creates a new PostWasmContractsContractAddressOwnerParams object
// with the default values initialized.
func NewPostWasmContractsContractAddressOwnerParams() *PostWasmContractsContractAddressOwnerParams {
	var ()
	return &PostWasmContractsContractAddressOwnerParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostWasmContractsContractAddressOwnerParamsWithTimeout creates a new PostWasmContractsContractAddressOwnerParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostWasmContractsContractAddressOwnerParamsWithTimeout(timeout time.Duration) *PostWasmContractsContractAddressOwnerParams {
	var ()
	return &PostWasmContractsContractAddressOwnerParams{

		timeout: timeout,
	}
}

// NewPostWasmContractsContractAddressOwnerParamsWithContext creates a new PostWasmContractsContractAddressOwnerParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostWasmContractsContractAddressOwnerParamsWithContext(ctx context.Context) *PostWasmContractsContractAddressOwnerParams {
	var ()
	return &PostWasmContractsContractAddressOwnerParams{

		Context: ctx,
	}
}

// NewPostWasmContractsContractAddressOwnerParamsWithHTTPClient creates a new PostWasmContractsContractAddressOwnerParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostWasmContractsContractAddressOwnerParamsWithHTTPClient(client *http.Client) *PostWasmContractsContractAddressOwnerParams {
	var ()
	return &PostWasmContractsContractAddressOwnerParams{
		HTTPClient: client,
	}
}

/*PostWasmContractsContractAddressOwnerParams contains all the parameters to send to the API endpoint
for the post wasm contracts contract address owner operation typically these are written to a http.Request
*/
type PostWasmContractsContractAddressOwnerParams struct {

	/*ContractAddress
	  contract address you want to update owner

	*/
	ContractAddress string
	/*UpdateContractOwnerRequestBody*/
	UpdateContractOwnerRequestBody *models.UpdateContractOwnerReq

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post wasm contracts contract address owner params
func (o *PostWasmContractsContractAddressOwnerParams) WithTimeout(timeout time.Duration) *PostWasmContractsContractAddressOwnerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post wasm contracts contract address owner params
func (o *PostWasmContractsContractAddressOwnerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post wasm contracts contract address owner params
func (o *PostWasmContractsContractAddressOwnerParams) WithContext(ctx context.Context) *PostWasmContractsContractAddressOwnerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post wasm contracts contract address owner params
func (o *PostWasmContractsContractAddressOwnerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post wasm contracts contract address owner params
func (o *PostWasmContractsContractAddressOwnerParams) WithHTTPClient(client *http.Client) *PostWasmContractsContractAddressOwnerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post wasm contracts contract address owner params
func (o *PostWasmContractsContractAddressOwnerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContractAddress adds the contractAddress to the post wasm contracts contract address owner params
func (o *PostWasmContractsContractAddressOwnerParams) WithContractAddress(contractAddress string) *PostWasmContractsContractAddressOwnerParams {
	o.SetContractAddress(contractAddress)
	return o
}

// SetContractAddress adds the contractAddress to the post wasm contracts contract address owner params
func (o *PostWasmContractsContractAddressOwnerParams) SetContractAddress(contractAddress string) {
	o.ContractAddress = contractAddress
}

// WithUpdateContractOwnerRequestBody adds the updateContractOwnerRequestBody to the post wasm contracts contract address owner params
func (o *PostWasmContractsContractAddressOwnerParams) WithUpdateContractOwnerRequestBody(updateContractOwnerRequestBody *models.UpdateContractOwnerReq) *PostWasmContractsContractAddressOwnerParams {
	o.SetUpdateContractOwnerRequestBody(updateContractOwnerRequestBody)
	return o
}

// SetUpdateContractOwnerRequestBody adds the updateContractOwnerRequestBody to the post wasm contracts contract address owner params
func (o *PostWasmContractsContractAddressOwnerParams) SetUpdateContractOwnerRequestBody(updateContractOwnerRequestBody *models.UpdateContractOwnerReq) {
	o.UpdateContractOwnerRequestBody = updateContractOwnerRequestBody
}

// WriteToRequest writes these params to a swagger request
func (o *PostWasmContractsContractAddressOwnerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param contractAddress
	if err := r.SetPathParam("contractAddress", o.ContractAddress); err != nil {
		return err
	}

	if o.UpdateContractOwnerRequestBody != nil {
		if err := r.SetBodyParam(o.UpdateContractOwnerRequestBody); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
