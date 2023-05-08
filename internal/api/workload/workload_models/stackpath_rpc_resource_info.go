// Code generated by go-swagger; DO NOT EDIT.

package workload_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// StackpathRPCResourceInfo stackpath rpc resource info
//
// swagger:model stackpath.rpc.ResourceInfo
type StackpathRPCResourceInfo struct {

	// description
	Description string `json:"description,omitempty"`

	// owner
	Owner string `json:"owner,omitempty"`

	// resource name
	ResourceName string `json:"resourceName,omitempty"`

	// resource type
	ResourceType string `json:"resourceType,omitempty"`
}

// AtType gets the at type of this subtype
func (m *StackpathRPCResourceInfo) AtType() string {
	return "stackpath.rpc.ResourceInfo"
}

// SetAtType sets the at type of this subtype
func (m *StackpathRPCResourceInfo) SetAtType(val string) {
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *StackpathRPCResourceInfo) UnmarshalJSON(raw []byte) error {
	var data struct {

		// description
		Description string `json:"description,omitempty"`

		// owner
		Owner string `json:"owner,omitempty"`

		// resource name
		ResourceName string `json:"resourceName,omitempty"`

		// resource type
		ResourceType string `json:"resourceType,omitempty"`
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var base struct {
		/* Just the base type fields. Used for unmashalling polymorphic types.*/

		AtType string `json:"@type"`
	}
	buf = bytes.NewBuffer(raw)
	dec = json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&base); err != nil {
		return err
	}

	var result StackpathRPCResourceInfo

	if base.AtType != result.AtType() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid @type value: %q", base.AtType)
	}

	result.Description = data.Description
	result.Owner = data.Owner
	result.ResourceName = data.ResourceName
	result.ResourceType = data.ResourceType

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m StackpathRPCResourceInfo) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// description
		Description string `json:"description,omitempty"`

		// owner
		Owner string `json:"owner,omitempty"`

		// resource name
		ResourceName string `json:"resourceName,omitempty"`

		// resource type
		ResourceType string `json:"resourceType,omitempty"`
	}{

		Description: m.Description,

		Owner: m.Owner,

		ResourceName: m.ResourceName,

		ResourceType: m.ResourceType,
	})
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		AtType string `json:"@type"`
	}{

		AtType: m.AtType(),
	})
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this stackpath rpc resource info
func (m *StackpathRPCResourceInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this stackpath rpc resource info based on the context it is used
func (m *StackpathRPCResourceInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *StackpathRPCResourceInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StackpathRPCResourceInfo) UnmarshalBinary(b []byte) error {
	var res StackpathRPCResourceInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
