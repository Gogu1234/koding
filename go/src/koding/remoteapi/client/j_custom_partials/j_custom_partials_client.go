package j_custom_partials

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new j custom partials API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for j custom partials API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
PostRemoteAPIJCustomPartialsCreate post remote API j custom partials create API
*/
func (a *Client) PostRemoteAPIJCustomPartialsCreate(params *PostRemoteAPIJCustomPartialsCreateParams) (*PostRemoteAPIJCustomPartialsCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostRemoteAPIJCustomPartialsCreateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostRemoteAPIJCustomPartialsCreate",
		Method:             "POST",
		PathPattern:        "/remote.api/JCustomPartials.create",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostRemoteAPIJCustomPartialsCreateReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostRemoteAPIJCustomPartialsCreateOK), nil

}

/*
PostRemoteAPIJCustomPartialsRemoveID post remote API j custom partials remove ID API
*/
func (a *Client) PostRemoteAPIJCustomPartialsRemoveID(params *PostRemoteAPIJCustomPartialsRemoveIDParams) (*PostRemoteAPIJCustomPartialsRemoveIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostRemoteAPIJCustomPartialsRemoveIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostRemoteAPIJCustomPartialsRemoveID",
		Method:             "POST",
		PathPattern:        "/remote.api/JCustomPartials.remove/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostRemoteAPIJCustomPartialsRemoveIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostRemoteAPIJCustomPartialsRemoveIDOK), nil

}

/*
PostRemoteAPIJCustomPartialsSome post remote API j custom partials some API
*/
func (a *Client) PostRemoteAPIJCustomPartialsSome(params *PostRemoteAPIJCustomPartialsSomeParams) (*PostRemoteAPIJCustomPartialsSomeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostRemoteAPIJCustomPartialsSomeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostRemoteAPIJCustomPartialsSome",
		Method:             "POST",
		PathPattern:        "/remote.api/JCustomPartials.some",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostRemoteAPIJCustomPartialsSomeReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostRemoteAPIJCustomPartialsSomeOK), nil

}

/*
PostRemoteAPIJCustomPartialsUpdateID post remote API j custom partials update ID API
*/
func (a *Client) PostRemoteAPIJCustomPartialsUpdateID(params *PostRemoteAPIJCustomPartialsUpdateIDParams) (*PostRemoteAPIJCustomPartialsUpdateIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostRemoteAPIJCustomPartialsUpdateIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostRemoteAPIJCustomPartialsUpdateID",
		Method:             "POST",
		PathPattern:        "/remote.api/JCustomPartials.update/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostRemoteAPIJCustomPartialsUpdateIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostRemoteAPIJCustomPartialsUpdateIDOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
