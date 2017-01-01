package social_message

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

// NewPostRemoteAPISocialMessageListRepliesParams creates a new PostRemoteAPISocialMessageListRepliesParams object
// with the default values initialized.
func NewPostRemoteAPISocialMessageListRepliesParams() *PostRemoteAPISocialMessageListRepliesParams {
	var ()
	return &PostRemoteAPISocialMessageListRepliesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRemoteAPISocialMessageListRepliesParamsWithTimeout creates a new PostRemoteAPISocialMessageListRepliesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRemoteAPISocialMessageListRepliesParamsWithTimeout(timeout time.Duration) *PostRemoteAPISocialMessageListRepliesParams {
	var ()
	return &PostRemoteAPISocialMessageListRepliesParams{

		timeout: timeout,
	}
}

// NewPostRemoteAPISocialMessageListRepliesParamsWithContext creates a new PostRemoteAPISocialMessageListRepliesParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRemoteAPISocialMessageListRepliesParamsWithContext(ctx context.Context) *PostRemoteAPISocialMessageListRepliesParams {
	var ()
	return &PostRemoteAPISocialMessageListRepliesParams{

		Context: ctx,
	}
}

/*PostRemoteAPISocialMessageListRepliesParams contains all the parameters to send to the API endpoint
for the post remote API social message list replies operation typically these are written to a http.Request
*/
type PostRemoteAPISocialMessageListRepliesParams struct {

	/*Body
	  body of the request

	*/
	Body *models.DefaultSelector

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post remote API social message list replies params
func (o *PostRemoteAPISocialMessageListRepliesParams) WithTimeout(timeout time.Duration) *PostRemoteAPISocialMessageListRepliesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post remote API social message list replies params
func (o *PostRemoteAPISocialMessageListRepliesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post remote API social message list replies params
func (o *PostRemoteAPISocialMessageListRepliesParams) WithContext(ctx context.Context) *PostRemoteAPISocialMessageListRepliesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post remote API social message list replies params
func (o *PostRemoteAPISocialMessageListRepliesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the post remote API social message list replies params
func (o *PostRemoteAPISocialMessageListRepliesParams) WithBody(body *models.DefaultSelector) *PostRemoteAPISocialMessageListRepliesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post remote API social message list replies params
func (o *PostRemoteAPISocialMessageListRepliesParams) SetBody(body *models.DefaultSelector) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostRemoteAPISocialMessageListRepliesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
