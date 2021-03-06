// Code generated by go-swagger; DO NOT EDIT.

package federation_gateways

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetFegNetworkIDGatewaysGatewayIDFederationParams creates a new GetFegNetworkIDGatewaysGatewayIDFederationParams object
// with the default values initialized.
func NewGetFegNetworkIDGatewaysGatewayIDFederationParams() *GetFegNetworkIDGatewaysGatewayIDFederationParams {
	var ()
	return &GetFegNetworkIDGatewaysGatewayIDFederationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetFegNetworkIDGatewaysGatewayIDFederationParamsWithTimeout creates a new GetFegNetworkIDGatewaysGatewayIDFederationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetFegNetworkIDGatewaysGatewayIDFederationParamsWithTimeout(timeout time.Duration) *GetFegNetworkIDGatewaysGatewayIDFederationParams {
	var ()
	return &GetFegNetworkIDGatewaysGatewayIDFederationParams{

		timeout: timeout,
	}
}

// NewGetFegNetworkIDGatewaysGatewayIDFederationParamsWithContext creates a new GetFegNetworkIDGatewaysGatewayIDFederationParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetFegNetworkIDGatewaysGatewayIDFederationParamsWithContext(ctx context.Context) *GetFegNetworkIDGatewaysGatewayIDFederationParams {
	var ()
	return &GetFegNetworkIDGatewaysGatewayIDFederationParams{

		Context: ctx,
	}
}

// NewGetFegNetworkIDGatewaysGatewayIDFederationParamsWithHTTPClient creates a new GetFegNetworkIDGatewaysGatewayIDFederationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetFegNetworkIDGatewaysGatewayIDFederationParamsWithHTTPClient(client *http.Client) *GetFegNetworkIDGatewaysGatewayIDFederationParams {
	var ()
	return &GetFegNetworkIDGatewaysGatewayIDFederationParams{
		HTTPClient: client,
	}
}

/*GetFegNetworkIDGatewaysGatewayIDFederationParams contains all the parameters to send to the API endpoint
for the get feg network ID gateways gateway ID federation operation typically these are written to a http.Request
*/
type GetFegNetworkIDGatewaysGatewayIDFederationParams struct {

	/*GatewayID
	  Gateway ID

	*/
	GatewayID string
	/*NetworkID
	  Network ID

	*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get feg network ID gateways gateway ID federation params
func (o *GetFegNetworkIDGatewaysGatewayIDFederationParams) WithTimeout(timeout time.Duration) *GetFegNetworkIDGatewaysGatewayIDFederationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get feg network ID gateways gateway ID federation params
func (o *GetFegNetworkIDGatewaysGatewayIDFederationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get feg network ID gateways gateway ID federation params
func (o *GetFegNetworkIDGatewaysGatewayIDFederationParams) WithContext(ctx context.Context) *GetFegNetworkIDGatewaysGatewayIDFederationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get feg network ID gateways gateway ID federation params
func (o *GetFegNetworkIDGatewaysGatewayIDFederationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get feg network ID gateways gateway ID federation params
func (o *GetFegNetworkIDGatewaysGatewayIDFederationParams) WithHTTPClient(client *http.Client) *GetFegNetworkIDGatewaysGatewayIDFederationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get feg network ID gateways gateway ID federation params
func (o *GetFegNetworkIDGatewaysGatewayIDFederationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGatewayID adds the gatewayID to the get feg network ID gateways gateway ID federation params
func (o *GetFegNetworkIDGatewaysGatewayIDFederationParams) WithGatewayID(gatewayID string) *GetFegNetworkIDGatewaysGatewayIDFederationParams {
	o.SetGatewayID(gatewayID)
	return o
}

// SetGatewayID adds the gatewayId to the get feg network ID gateways gateway ID federation params
func (o *GetFegNetworkIDGatewaysGatewayIDFederationParams) SetGatewayID(gatewayID string) {
	o.GatewayID = gatewayID
}

// WithNetworkID adds the networkID to the get feg network ID gateways gateway ID federation params
func (o *GetFegNetworkIDGatewaysGatewayIDFederationParams) WithNetworkID(networkID string) *GetFegNetworkIDGatewaysGatewayIDFederationParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get feg network ID gateways gateway ID federation params
func (o *GetFegNetworkIDGatewaysGatewayIDFederationParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *GetFegNetworkIDGatewaysGatewayIDFederationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param gateway_id
	if err := r.SetPathParam("gateway_id", o.GatewayID); err != nil {
		return err
	}

	// path param network_id
	if err := r.SetPathParam("network_id", o.NetworkID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
