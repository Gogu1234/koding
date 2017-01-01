package j_user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewPostRemoteAPIJUserWhoamiParams creates a new PostRemoteAPIJUserWhoamiParams object
// with the default values initialized.
func NewPostRemoteAPIJUserWhoamiParams() *PostRemoteAPIJUserWhoamiParams {

	return &PostRemoteAPIJUserWhoamiParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRemoteAPIJUserWhoamiParamsWithTimeout creates a new PostRemoteAPIJUserWhoamiParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRemoteAPIJUserWhoamiParamsWithTimeout(timeout time.Duration) *PostRemoteAPIJUserWhoamiParams {

	return &PostRemoteAPIJUserWhoamiParams{

		timeout: timeout,
	}
}

// NewPostRemoteAPIJUserWhoamiParamsWithContext creates a new PostRemoteAPIJUserWhoamiParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRemoteAPIJUserWhoamiParamsWithContext(ctx context.Context) *PostRemoteAPIJUserWhoamiParams {

	return &PostRemoteAPIJUserWhoamiParams{

		Context: ctx,
	}
}

/*PostRemoteAPIJUserWhoamiParams contains all the parameters to send to the API endpoint
for the post remote API j user whoami operation typically these are written to a http.Request
*/
type PostRemoteAPIJUserWhoamiParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post remote API j user whoami params
func (o *PostRemoteAPIJUserWhoamiParams) WithTimeout(timeout time.Duration) *PostRemoteAPIJUserWhoamiParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post remote API j user whoami params
func (o *PostRemoteAPIJUserWhoamiParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post remote API j user whoami params
func (o *PostRemoteAPIJUserWhoamiParams) WithContext(ctx context.Context) *PostRemoteAPIJUserWhoamiParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post remote API j user whoami params
func (o *PostRemoteAPIJUserWhoamiParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WriteToRequest writes these params to a swagger request
func (o *PostRemoteAPIJUserWhoamiParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
