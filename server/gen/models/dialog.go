// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"encoding/json"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Dialog dialog
//
// swagger:model dialog
type Dialog struct {

	// author
	Author string `json:"author,omitempty"`

	// content
	Content string `json:"content,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// link
	Link string `json:"link,omitempty"`

	// source
	Source string `json:"source,omitempty"`

	// style
	Style string `json:"style,omitempty"`

	// title
	Title string `json:"title,omitempty"`

	// user id
	UserID int64 `json:"user_id,omitempty"`
}

// UnmarshalJSON unmarshals this object while disallowing additional properties from JSON
func (m *Dialog) UnmarshalJSON(data []byte) error {
	var props struct {

		// author
		Author string `json:"author,omitempty"`

		// content
		Content string `json:"content,omitempty"`

		// id
		ID int64 `json:"id,omitempty"`

		// link
		Link string `json:"link,omitempty"`

		// source
		Source string `json:"source,omitempty"`

		// style
		Style string `json:"style,omitempty"`

		// title
		Title string `json:"title,omitempty"`

		// user id
		UserID int64 `json:"user_id,omitempty"`
	}

	dec := json.NewDecoder(bytes.NewReader(data))
	dec.DisallowUnknownFields()
	if err := dec.Decode(&props); err != nil {
		return err
	}

	m.Author = props.Author
	m.Content = props.Content
	m.ID = props.ID
	m.Link = props.Link
	m.Source = props.Source
	m.Style = props.Style
	m.Title = props.Title
	m.UserID = props.UserID
	return nil
}

// Validate validates this dialog
func (m *Dialog) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Dialog) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Dialog) UnmarshalBinary(b []byte) error {
	var res Dialog
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
