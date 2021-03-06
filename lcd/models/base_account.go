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

// BaseAccount base account
//
// swagger:model BaseAccount
type BaseAccount struct {

	// account number
	AccountNumber string `json:"account_number,omitempty"`

	// address
	Address string `json:"address,omitempty"`

	// coins
	Coins []*Coin `json:"coins"`

	// public key
	PublicKey string `json:"public_key,omitempty"`

	// sequence
	Sequence string `json:"sequence,omitempty"`
}

// Validate validates this base account
func (m *BaseAccount) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCoins(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BaseAccount) validateCoins(formats strfmt.Registry) error {

	if swag.IsZero(m.Coins) { // not required
		return nil
	}

	for i := 0; i < len(m.Coins); i++ {
		if swag.IsZero(m.Coins[i]) { // not required
			continue
		}

		if m.Coins[i] != nil {
			if err := m.Coins[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("coins" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *BaseAccount) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BaseAccount) UnmarshalBinary(b []byte) error {
	var res BaseAccount
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
