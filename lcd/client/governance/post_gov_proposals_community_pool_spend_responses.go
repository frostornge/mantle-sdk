// Code generated by go-swagger; DO NOT EDIT.

package governance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/terra-project/mantle-sdk/lcd/models"
)

// PostGovProposalsCommunityPoolSpendReader is a Reader for the PostGovProposalsCommunityPoolSpend structure.
type PostGovProposalsCommunityPoolSpendReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostGovProposalsCommunityPoolSpendReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostGovProposalsCommunityPoolSpendOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostGovProposalsCommunityPoolSpendBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostGovProposalsCommunityPoolSpendInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostGovProposalsCommunityPoolSpendOK creates a PostGovProposalsCommunityPoolSpendOK with default headers values
func NewPostGovProposalsCommunityPoolSpendOK() *PostGovProposalsCommunityPoolSpendOK {
	return &PostGovProposalsCommunityPoolSpendOK{}
}

/*PostGovProposalsCommunityPoolSpendOK handles this case with default header values.

The transaction was successfully generated
*/
type PostGovProposalsCommunityPoolSpendOK struct {
	Payload *models.StdTx
}

func (o *PostGovProposalsCommunityPoolSpendOK) Error() string {
	return fmt.Sprintf("[POST /gov/proposals/community_pool_spend][%d] postGovProposalsCommunityPoolSpendOK  %+v", 200, o.Payload)
}

func (o *PostGovProposalsCommunityPoolSpendOK) GetPayload() *models.StdTx {
	return o.Payload
}

func (o *PostGovProposalsCommunityPoolSpendOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StdTx)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGovProposalsCommunityPoolSpendBadRequest creates a PostGovProposalsCommunityPoolSpendBadRequest with default headers values
func NewPostGovProposalsCommunityPoolSpendBadRequest() *PostGovProposalsCommunityPoolSpendBadRequest {
	return &PostGovProposalsCommunityPoolSpendBadRequest{}
}

/*PostGovProposalsCommunityPoolSpendBadRequest handles this case with default header values.

Invalid proposal body
*/
type PostGovProposalsCommunityPoolSpendBadRequest struct {
}

func (o *PostGovProposalsCommunityPoolSpendBadRequest) Error() string {
	return fmt.Sprintf("[POST /gov/proposals/community_pool_spend][%d] postGovProposalsCommunityPoolSpendBadRequest ", 400)
}

func (o *PostGovProposalsCommunityPoolSpendBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostGovProposalsCommunityPoolSpendInternalServerError creates a PostGovProposalsCommunityPoolSpendInternalServerError with default headers values
func NewPostGovProposalsCommunityPoolSpendInternalServerError() *PostGovProposalsCommunityPoolSpendInternalServerError {
	return &PostGovProposalsCommunityPoolSpendInternalServerError{}
}

/*PostGovProposalsCommunityPoolSpendInternalServerError handles this case with default header values.

Internal Server Error
*/
type PostGovProposalsCommunityPoolSpendInternalServerError struct {
}

func (o *PostGovProposalsCommunityPoolSpendInternalServerError) Error() string {
	return fmt.Sprintf("[POST /gov/proposals/community_pool_spend][%d] postGovProposalsCommunityPoolSpendInternalServerError ", 500)
}

func (o *PostGovProposalsCommunityPoolSpendInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*PostGovProposalsCommunityPoolSpendBody post gov proposals community pool spend body
swagger:model PostGovProposalsCommunityPoolSpendBody
*/
type PostGovProposalsCommunityPoolSpendBody struct {

	// amount
	Amount []*models.Coin `json:"amount"`

	// base req
	BaseReq *models.BaseReq `json:"base_req,omitempty"`

	// deposit
	Deposit []*models.Coin `json:"deposit"`

	// description
	Description string `json:"description,omitempty"`

	// proposer
	Proposer models.Address `json:"proposer,omitempty"`

	// recipient
	Recipient models.Address `json:"recipient,omitempty"`

	// title
	Title string `json:"title,omitempty"`
}

// Validate validates this post gov proposals community pool spend body
func (o *PostGovProposalsCommunityPoolSpendBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateBaseReq(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateDeposit(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateProposer(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRecipient(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostGovProposalsCommunityPoolSpendBody) validateAmount(formats strfmt.Registry) error {

	if swag.IsZero(o.Amount) { // not required
		return nil
	}

	for i := 0; i < len(o.Amount); i++ {
		if swag.IsZero(o.Amount[i]) { // not required
			continue
		}

		if o.Amount[i] != nil {
			if err := o.Amount[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("post_proposal_body" + "." + "amount" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *PostGovProposalsCommunityPoolSpendBody) validateBaseReq(formats strfmt.Registry) error {

	if swag.IsZero(o.BaseReq) { // not required
		return nil
	}

	if o.BaseReq != nil {
		if err := o.BaseReq.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("post_proposal_body" + "." + "base_req")
			}
			return err
		}
	}

	return nil
}

func (o *PostGovProposalsCommunityPoolSpendBody) validateDeposit(formats strfmt.Registry) error {

	if swag.IsZero(o.Deposit) { // not required
		return nil
	}

	for i := 0; i < len(o.Deposit); i++ {
		if swag.IsZero(o.Deposit[i]) { // not required
			continue
		}

		if o.Deposit[i] != nil {
			if err := o.Deposit[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("post_proposal_body" + "." + "deposit" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *PostGovProposalsCommunityPoolSpendBody) validateProposer(formats strfmt.Registry) error {

	if swag.IsZero(o.Proposer) { // not required
		return nil
	}

	if err := o.Proposer.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("post_proposal_body" + "." + "proposer")
		}
		return err
	}

	return nil
}

func (o *PostGovProposalsCommunityPoolSpendBody) validateRecipient(formats strfmt.Registry) error {

	if swag.IsZero(o.Recipient) { // not required
		return nil
	}

	if err := o.Recipient.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("post_proposal_body" + "." + "recipient")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostGovProposalsCommunityPoolSpendBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostGovProposalsCommunityPoolSpendBody) UnmarshalBinary(b []byte) error {
	var res PostGovProposalsCommunityPoolSpendBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
