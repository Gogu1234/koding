package j_account

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

// NewPostRemoteAPIJAccountFetchRoleIDParams creates a new PostRemoteAPIJAccountFetchRoleIDParams object
// with the default values initialized.
func NewPostRemoteAPIJAccountFetchRoleIDParams() *PostRemoteAPIJAccountFetchRoleIDParams {
	var ()
	return &PostRemoteAPIJAccountFetchRoleIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRemoteAPIJAccountFetchRoleIDParamsWithTimeout creates a new PostRemoteAPIJAccountFetchRoleIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRemoteAPIJAccountFetchRoleIDParamsWithTimeout(timeout time.Duration) *PostRemoteAPIJAccountFetchRoleIDParams {
	var ()
	return &PostRemoteAPIJAccountFetchRoleIDParams{

		timeout: timeout,
	}
}

// NewPostRemoteAPIJAccountFetchRoleIDParamsWithContext creates a new PostRemoteAPIJAccountFetchRoleIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRemoteAPIJAccountFetchRoleIDParamsWithContext(ctx context.Context) *PostRemoteAPIJAccountFetchRoleIDParams {
	var ()
	return &PostRemoteAPIJAccountFetchRoleIDParams{

		Context: ctx,
	}
}

/*PostRemoteAPIJAccountFetchRoleIDParams contains all the parameters to send to the API endpoint
for the post remote API j account fetch role ID operation typically these are written to a http.Request
*/
type PostRemoteAPIJAccountFetchRoleIDParams struct {

	/*ID
	  Mongo ID of target instance

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post remote API j account fetch role ID params
func (o *PostRemoteAPIJAccountFetchRoleIDParams) WithTimeout(timeout time.Duration) *PostRemoteAPIJAccountFetchRoleIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post remote API j account fetch role ID params
func (o *PostRemoteAPIJAccountFetchRoleIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post remote API j account fetch role ID params
func (o *PostRemoteAPIJAccountFetchRoleIDParams) WithContext(ctx context.Context) *PostRemoteAPIJAccountFetchRoleIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post remote API j account fetch role ID params
func (o *PostRemoteAPIJAccountFetchRoleIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithID adds the id to the post remote API j account fetch role ID params
func (o *PostRemoteAPIJAccountFetchRoleIDParams) WithID(id string) *PostRemoteAPIJAccountFetchRoleIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post remote API j account fetch role ID params
func (o *PostRemoteAPIJAccountFetchRoleIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PostRemoteAPIJAccountFetchRoleIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
