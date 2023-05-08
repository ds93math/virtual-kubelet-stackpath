// Code generated by go-swagger; DO NOT EDIT.

package workload_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1ContainerStatus The status of a container in a workload
//
// swagger:model v1ContainerStatus
type V1ContainerStatus struct {

	// A unique value that identifies a container
	ContainerID string `json:"containerId,omitempty"`

	// The date a container terminated
	// Format: date-time
	FinishedAt strfmt.DateTime `json:"finishedAt,omitempty"`

	// A container status' name
	Name string `json:"name,omitempty"`

	// phase
	Phase *V1ContainerStatusContainerPhase `json:"phase,omitempty"`

	// Whether or not a container is running and ready for service
	Ready bool `json:"ready,omitempty"`

	// How many times a container has restarted since it was first started
	RestartCount int32 `json:"restartCount,omitempty"`

	// running
	Running *ContainerStatusRunning `json:"running,omitempty"`

	// The date a container started
	// Format: date-time
	StartedAt strfmt.DateTime `json:"startedAt,omitempty"`

	// terminated
	Terminated *ContainerStatusTerminated `json:"terminated,omitempty"`

	// waiting
	Waiting *ContainerStatusWaiting `json:"waiting,omitempty"`
}

// Validate validates this v1 container status
func (m *V1ContainerStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFinishedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePhase(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRunning(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTerminated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWaiting(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ContainerStatus) validateFinishedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.FinishedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("finishedAt", "body", "date-time", m.FinishedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V1ContainerStatus) validatePhase(formats strfmt.Registry) error {
	if swag.IsZero(m.Phase) { // not required
		return nil
	}

	if m.Phase != nil {
		if err := m.Phase.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("phase")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("phase")
			}
			return err
		}
	}

	return nil
}

func (m *V1ContainerStatus) validateRunning(formats strfmt.Registry) error {
	if swag.IsZero(m.Running) { // not required
		return nil
	}

	if m.Running != nil {
		if err := m.Running.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("running")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("running")
			}
			return err
		}
	}

	return nil
}

func (m *V1ContainerStatus) validateStartedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.StartedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("startedAt", "body", "date-time", m.StartedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V1ContainerStatus) validateTerminated(formats strfmt.Registry) error {
	if swag.IsZero(m.Terminated) { // not required
		return nil
	}

	if m.Terminated != nil {
		if err := m.Terminated.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("terminated")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("terminated")
			}
			return err
		}
	}

	return nil
}

func (m *V1ContainerStatus) validateWaiting(formats strfmt.Registry) error {
	if swag.IsZero(m.Waiting) { // not required
		return nil
	}

	if m.Waiting != nil {
		if err := m.Waiting.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("waiting")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("waiting")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 container status based on the context it is used
func (m *V1ContainerStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePhase(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRunning(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTerminated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWaiting(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ContainerStatus) contextValidatePhase(ctx context.Context, formats strfmt.Registry) error {

	if m.Phase != nil {
		if err := m.Phase.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("phase")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("phase")
			}
			return err
		}
	}

	return nil
}

func (m *V1ContainerStatus) contextValidateRunning(ctx context.Context, formats strfmt.Registry) error {

	if m.Running != nil {
		if err := m.Running.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("running")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("running")
			}
			return err
		}
	}

	return nil
}

func (m *V1ContainerStatus) contextValidateTerminated(ctx context.Context, formats strfmt.Registry) error {

	if m.Terminated != nil {
		if err := m.Terminated.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("terminated")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("terminated")
			}
			return err
		}
	}

	return nil
}

func (m *V1ContainerStatus) contextValidateWaiting(ctx context.Context, formats strfmt.Registry) error {

	if m.Waiting != nil {
		if err := m.Waiting.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("waiting")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("waiting")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1ContainerStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ContainerStatus) UnmarshalBinary(b []byte) error {
	var res V1ContainerStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
