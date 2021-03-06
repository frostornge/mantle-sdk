// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// VoteReq vote req
//
// swagger:model VoteReq
type VoteReq struct {

	// base req
	BaseReq *BaseReq `json:"base_req,omitempty"`

	// proof exchange rate of Luna in denom currency was used to make prevote hash; initial prevote does not require this field
	ExchangeRate float64 `json:"exchange_rate,omitempty"`

	// proof salt was used to make prevote hash; initial prevote does not require this field
	Salt string `json:"salt,omitempty"`

	// validator
	Validator ValidatorAddress `json:"validator,omitempty"`
}

// Validate validates this vote req
func (m *VoteReq) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBaseReq(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValidator(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VoteReq) validateBaseReq(formats strfmt.Registry) error {

	if swag.IsZero(m.BaseReq) { // not required
		return nil
	}

	if m.BaseReq != nil {
		if err := m.BaseReq.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("base_req")
			}
			return err
		}
	}

	return nil
}

func (m *VoteReq) validateValidator(formats strfmt.Registry) error {

	if swag.IsZero(m.Validator) { // not required
		return nil
	}

	if err := m.Validator.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("validator")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VoteReq) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VoteReq) UnmarshalBinary(b []byte) error {
	var res VoteReq
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
