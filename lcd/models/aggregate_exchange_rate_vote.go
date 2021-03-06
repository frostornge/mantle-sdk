// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AggregateExchangeRateVote aggregate exchange rate vote
//
// swagger:model AggregateExchangeRateVote
type AggregateExchangeRateVote struct {

	// exchange rates
	ExchangeRates []*DecCoin `json:"exchange_rates"`

	// voter
	Voter ValidatorAddress `json:"voter,omitempty"`
}

// Validate validates this aggregate exchange rate vote
func (m *AggregateExchangeRateVote) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExchangeRates(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVoter(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AggregateExchangeRateVote) validateExchangeRates(formats strfmt.Registry) error {

	if swag.IsZero(m.ExchangeRates) { // not required
		return nil
	}

	for i := 0; i < len(m.ExchangeRates); i++ {
		if swag.IsZero(m.ExchangeRates[i]) { // not required
			continue
		}

		if m.ExchangeRates[i] != nil {
			if err := m.ExchangeRates[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("exchange_rates" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AggregateExchangeRateVote) validateVoter(formats strfmt.Registry) error {

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
func (m *AggregateExchangeRateVote) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AggregateExchangeRateVote) UnmarshalBinary(b []byte) error {
	var res AggregateExchangeRateVote
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
