// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SegmentDebugLog segment debug log
//
// swagger:model segmentDebugLog
type SegmentDebugLog struct {

	// msg
	Msg string `json:"msg,omitempty"`

	// segment ID
	// Minimum: 1
	SegmentID int64 `json:"segmentID,omitempty"`
}

// Validate validates this segment debug log
func (m *SegmentDebugLog) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSegmentID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SegmentDebugLog) validateSegmentID(formats strfmt.Registry) error {
	if swag.IsZero(m.SegmentID) { // not required
		return nil
	}

	if err := validate.MinimumInt("segmentID", "body", m.SegmentID, 1, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this segment debug log based on context it is used
func (m *SegmentDebugLog) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SegmentDebugLog) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SegmentDebugLog) UnmarshalBinary(b []byte) error {
	var res SegmentDebugLog
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
