// Code generated by go-swagger; DO NOT EDIT.

package auth

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

// NewPostAuthAccountsAddressMultisignParams creates a new PostAuthAccountsAddressMultisignParams object
// with the default values initialized.
func NewPostAuthAccountsAddressMultisignParams() *PostAuthAccountsAddressMultisignParams {
	var ()
	return &PostAuthAccountsAddressMultisignParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostAuthAccountsAddressMultisignParamsWithTimeout creates a new PostAuthAccountsAddressMultisignParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostAuthAccountsAddressMultisignParamsWithTimeout(timeout time.Duration) *PostAuthAccountsAddressMultisignParams {
	var ()
	return &PostAuthAccountsAddressMultisignParams{

		timeout: timeout,
	}
}

// NewPostAuthAccountsAddressMultisignParamsWithContext creates a new PostAuthAccountsAddressMultisignParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostAuthAccountsAddressMultisignParamsWithContext(ctx context.Context) *PostAuthAccountsAddressMultisignParams {
	var ()
	return &PostAuthAccountsAddressMultisignParams{

		Context: ctx,
	}
}

// NewPostAuthAccountsAddressMultisignParamsWithHTTPClient creates a new PostAuthAccountsAddressMultisignParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostAuthAccountsAddressMultisignParamsWithHTTPClient(client *http.Client) *PostAuthAccountsAddressMultisignParams {
	var ()
	return &PostAuthAccountsAddressMultisignParams{
		HTTPClient: client,
	}
}

/*PostAuthAccountsAddressMultisignParams contains all the parameters to send to the API endpoint
for the post auth accounts address multisign operation typically these are written to a http.Request
*/
type PostAuthAccountsAddressMultisignParams struct {

	/*Address
	  Account address

	*/
	Address string
	/*MultisigReq
	  multisign request information; pubkey is optional param in case multisig account never used before

	*/
	MultisigReq *models.MultiSignReq

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post auth accounts address multisign params
func (o *PostAuthAccountsAddressMultisignParams) WithTimeout(timeout time.Duration) *PostAuthAccountsAddressMultisignParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post auth accounts address multisign params
func (o *PostAuthAccountsAddressMultisignParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post auth accounts address multisign params
func (o *PostAuthAccountsAddressMultisignParams) WithContext(ctx context.Context) *PostAuthAccountsAddressMultisignParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post auth accounts address multisign params
func (o *PostAuthAccountsAddressMultisignParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post auth accounts address multisign params
func (o *PostAuthAccountsAddressMultisignParams) WithHTTPClient(client *http.Client) *PostAuthAccountsAddressMultisignParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post auth accounts address multisign params
func (o *PostAuthAccountsAddressMultisignParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAddress adds the address to the post auth accounts address multisign params
func (o *PostAuthAccountsAddressMultisignParams) WithAddress(address string) *PostAuthAccountsAddressMultisignParams {
	o.SetAddress(address)
	return o
}

// SetAddress adds the address to the post auth accounts address multisign params
func (o *PostAuthAccountsAddressMultisignParams) SetAddress(address string) {
	o.Address = address
}

// WithMultisigReq adds the multisigReq to the post auth accounts address multisign params
func (o *PostAuthAccountsAddressMultisignParams) WithMultisigReq(multisigReq *models.MultiSignReq) *PostAuthAccountsAddressMultisignParams {
	o.SetMultisigReq(multisigReq)
	return o
}

// SetMultisigReq adds the multisigReq to the post auth accounts address multisign params
func (o *PostAuthAccountsAddressMultisignParams) SetMultisigReq(multisigReq *models.MultiSignReq) {
	o.MultisigReq = multisigReq
}

// WriteToRequest writes these params to a swagger request
func (o *PostAuthAccountsAddressMultisignParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param address
	if err := r.SetPathParam("address", o.Address); err != nil {
		return err
	}

	if o.MultisigReq != nil {
		if err := r.SetBodyParam(o.MultisigReq); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
