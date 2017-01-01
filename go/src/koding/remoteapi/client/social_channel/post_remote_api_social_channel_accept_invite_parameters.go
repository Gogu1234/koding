package social_channel

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

	"koding/remoteapi/models"
)

// NewPostRemoteAPISocialChannelAcceptInviteParams creates a new PostRemoteAPISocialChannelAcceptInviteParams object
// with the default values initialized.
func NewPostRemoteAPISocialChannelAcceptInviteParams() *PostRemoteAPISocialChannelAcceptInviteParams {
	var ()
	return &PostRemoteAPISocialChannelAcceptInviteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRemoteAPISocialChannelAcceptInviteParamsWithTimeout creates a new PostRemoteAPISocialChannelAcceptInviteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRemoteAPISocialChannelAcceptInviteParamsWithTimeout(timeout time.Duration) *PostRemoteAPISocialChannelAcceptInviteParams {
	var ()
	return &PostRemoteAPISocialChannelAcceptInviteParams{

		timeout: timeout,
	}
}

// NewPostRemoteAPISocialChannelAcceptInviteParamsWithContext creates a new PostRemoteAPISocialChannelAcceptInviteParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRemoteAPISocialChannelAcceptInviteParamsWithContext(ctx context.Context) *PostRemoteAPISocialChannelAcceptInviteParams {
	var ()
	return &PostRemoteAPISocialChannelAcceptInviteParams{

		Context: ctx,
	}
}

/*PostRemoteAPISocialChannelAcceptInviteParams contains all the parameters to send to the API endpoint
for the post remote API social channel accept invite operation typically these are written to a http.Request
*/
type PostRemoteAPISocialChannelAcceptInviteParams struct {

	/*Body
	  body of the request

	*/
	Body *models.DefaultSelector

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post remote API social channel accept invite params
func (o *PostRemoteAPISocialChannelAcceptInviteParams) WithTimeout(timeout time.Duration) *PostRemoteAPISocialChannelAcceptInviteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post remote API social channel accept invite params
func (o *PostRemoteAPISocialChannelAcceptInviteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post remote API social channel accept invite params
func (o *PostRemoteAPISocialChannelAcceptInviteParams) WithContext(ctx context.Context) *PostRemoteAPISocialChannelAcceptInviteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post remote API social channel accept invite params
func (o *PostRemoteAPISocialChannelAcceptInviteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the post remote API social channel accept invite params
func (o *PostRemoteAPISocialChannelAcceptInviteParams) WithBody(body *models.DefaultSelector) *PostRemoteAPISocialChannelAcceptInviteParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post remote API social channel accept invite params
func (o *PostRemoteAPISocialChannelAcceptInviteParams) SetBody(body *models.DefaultSelector) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostRemoteAPISocialChannelAcceptInviteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.Body == nil {
		o.Body = new(models.DefaultSelector)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
