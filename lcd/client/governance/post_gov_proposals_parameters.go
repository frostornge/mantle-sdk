// Code generated by go-swagger; DO NOT EDIT.

package governance

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

// NewPostGovProposalsParams creates a new PostGovProposalsParams object
// with the default values initialized.
func NewPostGovProposalsParams() *PostGovProposalsParams {
	var ()
	return &PostGovProposalsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostGovProposalsParamsWithTimeout creates a new PostGovProposalsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostGovProposalsParamsWithTimeout(timeout time.Duration) *PostGovProposalsParams {
	var ()
	return &PostGovProposalsParams{

		timeout: timeout,
	}
}

// NewPostGovProposalsParamsWithContext creates a new PostGovProposalsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostGovProposalsParamsWithContext(ctx context.Context) *PostGovProposalsParams {
	var ()
	return &PostGovProposalsParams{

		Context: ctx,
	}
}

// NewPostGovProposalsParamsWithHTTPClient creates a new PostGovProposalsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostGovProposalsParamsWithHTTPClient(client *http.Client) *PostGovProposalsParams {
	var ()
	return &PostGovProposalsParams{
		HTTPClient: client,
	}
}

/*PostGovProposalsParams contains all the parameters to send to the API endpoint
for the post gov proposals operation typically these are written to a http.Request
*/
type PostGovProposalsParams struct {

	/*PostProposalBody
	  valid value of `"proposal_type"` can be `"text"`, `"parameter_change"`, `"software_upgrade"`

	*/
	PostProposalBody PostGovProposalsBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post gov proposals params
func (o *PostGovProposalsParams) WithTimeout(timeout time.Duration) *PostGovProposalsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post gov proposals params
func (o *PostGovProposalsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post gov proposals params
func (o *PostGovProposalsParams) WithContext(ctx context.Context) *PostGovProposalsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post gov proposals params
func (o *PostGovProposalsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post gov proposals params
func (o *PostGovProposalsParams) WithHTTPClient(client *http.Client) *PostGovProposalsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post gov proposals params
func (o *PostGovProposalsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPostProposalBody adds the postProposalBody to the post gov proposals params
func (o *PostGovProposalsParams) WithPostProposalBody(postProposalBody PostGovProposalsBody) *PostGovProposalsParams {
	o.SetPostProposalBody(postProposalBody)
	return o
}

// SetPostProposalBody adds the postProposalBody to the post gov proposals params
func (o *PostGovProposalsParams) SetPostProposalBody(postProposalBody PostGovProposalsBody) {
	o.PostProposalBody = postProposalBody
}

// WriteToRequest writes these params to a swagger request
func (o *PostGovProposalsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.PostProposalBody); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}