package j_reward

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

// NewPostRemoteAPIJRewardFetchEarnedAmountParams creates a new PostRemoteAPIJRewardFetchEarnedAmountParams object
// with the default values initialized.
func NewPostRemoteAPIJRewardFetchEarnedAmountParams() *PostRemoteAPIJRewardFetchEarnedAmountParams {
	var ()
	return &PostRemoteAPIJRewardFetchEarnedAmountParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRemoteAPIJRewardFetchEarnedAmountParamsWithTimeout creates a new PostRemoteAPIJRewardFetchEarnedAmountParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRemoteAPIJRewardFetchEarnedAmountParamsWithTimeout(timeout time.Duration) *PostRemoteAPIJRewardFetchEarnedAmountParams {
	var ()
	return &PostRemoteAPIJRewardFetchEarnedAmountParams{

		timeout: timeout,
	}
}

// NewPostRemoteAPIJRewardFetchEarnedAmountParamsWithContext creates a new PostRemoteAPIJRewardFetchEarnedAmountParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRemoteAPIJRewardFetchEarnedAmountParamsWithContext(ctx context.Context) *PostRemoteAPIJRewardFetchEarnedAmountParams {
	var ()
	return &PostRemoteAPIJRewardFetchEarnedAmountParams{

		Context: ctx,
	}
}

/*PostRemoteAPIJRewardFetchEarnedAmountParams contains all the parameters to send to the API endpoint
for the post remote API j reward fetch earned amount operation typically these are written to a http.Request
*/
type PostRemoteAPIJRewardFetchEarnedAmountParams struct {

	/*Body
	  body of the request

	*/
	Body *models.DefaultSelector

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post remote API j reward fetch earned amount params
func (o *PostRemoteAPIJRewardFetchEarnedAmountParams) WithTimeout(timeout time.Duration) *PostRemoteAPIJRewardFetchEarnedAmountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post remote API j reward fetch earned amount params
func (o *PostRemoteAPIJRewardFetchEarnedAmountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post remote API j reward fetch earned amount params
func (o *PostRemoteAPIJRewardFetchEarnedAmountParams) WithContext(ctx context.Context) *PostRemoteAPIJRewardFetchEarnedAmountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post remote API j reward fetch earned amount params
func (o *PostRemoteAPIJRewardFetchEarnedAmountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the post remote API j reward fetch earned amount params
func (o *PostRemoteAPIJRewardFetchEarnedAmountParams) WithBody(body *models.DefaultSelector) *PostRemoteAPIJRewardFetchEarnedAmountParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post remote API j reward fetch earned amount params
func (o *PostRemoteAPIJRewardFetchEarnedAmountParams) SetBody(body *models.DefaultSelector) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostRemoteAPIJRewardFetchEarnedAmountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
