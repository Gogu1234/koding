package j_url_alias

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new j url alias API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for j url alias API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
PostRemoteAPIJURLAliasCreate post remote API j URL alias create API
*/
func (a *Client) PostRemoteAPIJURLAliasCreate(params *PostRemoteAPIJURLAliasCreateParams) (*PostRemoteAPIJURLAliasCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostRemoteAPIJURLAliasCreateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostRemoteAPIJURLAliasCreate",
		Method:             "POST",
		PathPattern:        "/remote.api/JUrlAlias.create",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostRemoteAPIJURLAliasCreateReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostRemoteAPIJURLAliasCreateOK), nil

}

/*
PostRemoteAPIJURLAliasResolve Method JUrlAlias.resolve
*/
func (a *Client) PostRemoteAPIJURLAliasResolve(params *PostRemoteAPIJURLAliasResolveParams) (*PostRemoteAPIJURLAliasResolveOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostRemoteAPIJURLAliasResolveParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostRemoteAPIJURLAliasResolve",
		Method:             "POST",
		PathPattern:        "/remote.api/JUrlAlias.resolve",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostRemoteAPIJURLAliasResolveReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostRemoteAPIJURLAliasResolveOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
