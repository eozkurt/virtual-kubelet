package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// ContainerGetStateResponse container get state response
// swagger:model ContainerGetStateResponse
type ContainerGetStateResponse struct {

	// handle
	// Required: true
	Handle string `json:"handle"`

	// state
	// Required: true
	State string `json:"state"`
}

// Validate validates this container get state response
func (m *ContainerGetStateResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHandle(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContainerGetStateResponse) validateHandle(formats strfmt.Registry) error {

	if err := validate.RequiredString("handle", "body", string(m.Handle)); err != nil {
		return err
	}

	return nil
}

var containerGetStateResponseTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["RUNNING","STOPPED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		containerGetStateResponseTypeStatePropEnum = append(containerGetStateResponseTypeStatePropEnum, v)
	}
}

const (
	// ContainerGetStateResponseStateRUNNING captures enum value "RUNNING"
	ContainerGetStateResponseStateRUNNING string = "RUNNING"
	// ContainerGetStateResponseStateSTOPPED captures enum value "STOPPED"
	ContainerGetStateResponseStateSTOPPED string = "STOPPED"
)

// prop value enum
func (m *ContainerGetStateResponse) validateStateEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, containerGetStateResponseTypeStatePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ContainerGetStateResponse) validateState(formats strfmt.Registry) error {

	if err := validate.RequiredString("state", "body", string(m.State)); err != nil {
		return err
	}

	// value enum
	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}