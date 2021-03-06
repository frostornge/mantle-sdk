// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AggregateExchangeRatePrevote aggregate exchange rate prevote
//
// swagger:model AggregateExchangeRatePrevote
type AggregateExchangeRatePrevote struct {

	// hash
	Hash string `json:"hash,omitempty"`

	// submit block
	SubmitBlock float64 `json:"submit_block,omitempty"`

	// voter
	Voter ValidatorAddress `json:"voter,omitempty"`
}

// Validate validates this aggregate exchange rate prevote
func (m *AggregateExchangeRatePrevote) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVoter(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AggregateExchangeRatePrevote) validateVoter(formats strfmt.Registry) error {

	if swag.IsZero(m.Voter) { // not required
		return nil
	}

	if err := m.Voter.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("voter")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AggregateExchangeRatePrevote) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AggregateExchangeRatePrevote) UnmarshalBinary(b []byte) error {
	var res AggregateExchangeRatePrevote
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
