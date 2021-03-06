// Code generated by go-swagger; DO NOT EDIT.

package carrier_wifi_gateways

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

// NewGetCwfNetworkIDGatewaysGatewayIDCarrierWifiParams creates a new GetCwfNetworkIDGatewaysGatewayIDCarrierWifiParams object
// with the default values initialized.
func NewGetCwfNetworkIDGatewaysGatewayIDCarrierWifiParams() *GetCwfNetworkIDGatewaysGatewayIDCarrierWifiParams {
	var ()
	return &GetCwfNetworkIDGatewaysGatewayIDCarrierWifiParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCwfNetworkIDGatewaysGatewayIDCarrierWifiParamsWithTimeout creates a new GetCwfNetworkIDGatewaysGatewayIDCarrierWifiParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCwfNetworkIDGatewaysGatewayIDCarrierWifiParamsWithTimeout(timeout time.Duration) *GetCwfNetworkIDGatewaysGatewayIDCarrierWifiParams {
	var ()
	return &GetCwfNetworkIDGatewaysGatewayIDCarrierWifiParams{

		timeout: timeout,
	}
}

// NewGetCwfNetworkIDGatewaysGatewayIDCarrierWifiParamsWithContext creates a new GetCwfNetworkIDGatewaysGatewayIDCarrierWifiParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCwfNetworkIDGatewaysGatewayIDCarrierWifiParamsWithContext(ctx context.Context) *GetCwfNetworkIDGatewaysGatewayIDCarrierWifiParams {
	var ()
	return &GetCwfNetworkIDGatewaysGatewayIDCarrierWifiParams{

		Context: ctx,
	}
}

// NewGetCwfNetworkIDGatewaysGatewayIDCarrierWifiParamsWithHTTPClient creates a new GetCwfNetworkIDGatewaysGatewayIDCarrierWifiParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCwfNetworkIDGatewaysGatewayIDCarrierWifiParamsWithHTTPClient(client *http.Client) *GetCwfNetworkIDGatewaysGatewayIDCarrierWifiParams {
	var ()
	return &GetCwfNetworkIDGatewaysGatewayIDCarrierWifiParams{
		HTTPClient: client,
	}
}

/*GetCwfNetworkIDGatewaysGatewayIDCarrierWifiParams contains all the parameters to send to the API endpoint
for the get cwf network ID gateways gateway ID carrier wifi operation typically these are written to a http.Request
*/
type GetCwfNetworkIDGatewaysGatewayIDCarrierWifiParams struct {

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

// WithTimeout adds the timeout to the get cwf network ID gateways gateway ID carrier wifi params
func (o *GetCwfNetworkIDGatewaysGatewayIDCarrierWifiParams) WithTimeout(timeout time.Duration) *GetCwfNetworkIDGatewaysGatewayIDCarrierWifiParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get cwf network ID gateways gateway ID carrier wifi params
func (o *GetCwfNetworkIDGatewaysGatewayIDCarrierWifiParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get cwf network ID gateways gateway ID carrier wifi params
func (o *GetCwfNetworkIDGatewaysGatewayIDCarrierWifiParams) WithContext(ctx context.Context) *GetCwfNetworkIDGatewaysGatewayIDCarrierWifiParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get cwf network ID gateways gateway ID carrier wifi params
func (o *GetCwfNetworkIDGatewaysGatewayIDCarrierWifiParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get cwf network ID gateways gateway ID carrier wifi params
func (o *GetCwfNetworkIDGatewaysGatewayIDCarrierWifiParams) WithHTTPClient(client *http.Client) *GetCwfNetworkIDGatewaysGatewayIDCarrierWifiParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get cwf network ID gateways gateway ID carrier wifi params
func (o *GetCwfNetworkIDGatewaysGatewayIDCarrierWifiParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGatewayID adds the gatewayID to the get cwf network ID gateways gateway ID carrier wifi params
func (o *GetCwfNetworkIDGatewaysGatewayIDCarrierWifiParams) WithGatewayID(gatewayID string) *GetCwfNetworkIDGatewaysGatewayIDCarrierWifiParams {
	o.SetGatewayID(gatewayID)
	return o
}

// SetGatewayID adds the gatewayId to the get cwf network ID gateways gateway ID carrier wifi params
func (o *GetCwfNetworkIDGatewaysGatewayIDCarrierWifiParams) SetGatewayID(gatewayID string) {
	o.GatewayID = gatewayID
}

// WithNetworkID adds the networkID to the get cwf network ID gateways gateway ID carrier wifi params
func (o *GetCwfNetworkIDGatewaysGatewayIDCarrierWifiParams) WithNetworkID(networkID string) *GetCwfNetworkIDGatewaysGatewayIDCarrierWifiParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get cwf network ID gateways gateway ID carrier wifi params
func (o *GetCwfNetworkIDGatewaysGatewayIDCarrierWifiParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *GetCwfNetworkIDGatewaysGatewayIDCarrierWifiParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
