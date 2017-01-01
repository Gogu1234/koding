package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
)

// JSession j session
// swagger:model JSession
type JSession struct {

	// client IP
	ClientIP string `json:"clientIP,omitempty"`

	// client Id
	ClientID string `json:"clientId,omitempty"`

	// data
	Data interface{} `json:"data,omitempty"`

	// foreign auth
	ForeignAuth *JSessionForeignAuth `json:"foreignAuth,omitempty"`

	// foreign auth type
	ForeignAuthType string `json:"foreignAuthType,omitempty"`

	// group name
	GroupName string `json:"groupName,omitempty"`

	// guest Id
	GuestID float64 `json:"guestId,omitempty"`

	// guest session began
	GuestSessionBegan strfmt.Date `json:"guestSessionBegan,omitempty"`

	// impersonating
	Impersonating bool `json:"impersonating,omitempty"`

	// last access
	LastAccess strfmt.Date `json:"lastAccess,omitempty"`

	// ota token
	OtaToken string `json:"otaToken,omitempty"`

	// return Url
	ReturnURL string `json:"returnUrl,omitempty"`

	// session began
	SessionBegan strfmt.Date `json:"sessionBegan,omitempty"`

	// terminal Id
	TerminalID string `json:"terminalId,omitempty"`

	// username
	Username string `json:"username,omitempty"`
}

// Validate validates this j session
func (m *JSession) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateForeignAuth(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *JSession) validateForeignAuth(formats strfmt.Registry) error {

	if swag.IsZero(m.ForeignAuth) { // not required
		return nil
	}

	if m.ForeignAuth != nil {

		if err := m.ForeignAuth.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

// JSessionForeignAuth j session foreign auth
// swagger:model JSessionForeignAuth
type JSessionForeignAuth struct {

	// facebook
	Facebook interface{} `json:"facebook,omitempty"`

	// github
	Github interface{} `json:"github,omitempty"`

	// gitlab
	Gitlab interface{} `json:"gitlab,omitempty"`

	// linkedin
	Linkedin interface{} `json:"linkedin,omitempty"`
}

// Validate validates this j session foreign auth
func (m *JSessionForeignAuth) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
