// Code generated by go-swagger; DO NOT EDIT.

package public

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewInitializeSelfServiceRecoveryViaAPIFlowParams creates a new InitializeSelfServiceRecoveryViaAPIFlowParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewInitializeSelfServiceRecoveryViaAPIFlowParams() *InitializeSelfServiceRecoveryViaAPIFlowParams {
	return &InitializeSelfServiceRecoveryViaAPIFlowParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewInitializeSelfServiceRecoveryViaAPIFlowParamsWithTimeout creates a new InitializeSelfServiceRecoveryViaAPIFlowParams object
// with the ability to set a timeout on a request.
func NewInitializeSelfServiceRecoveryViaAPIFlowParamsWithTimeout(timeout time.Duration) *InitializeSelfServiceRecoveryViaAPIFlowParams {
	return &InitializeSelfServiceRecoveryViaAPIFlowParams{
		timeout: timeout,
	}
}

// NewInitializeSelfServiceRecoveryViaAPIFlowParamsWithContext creates a new InitializeSelfServiceRecoveryViaAPIFlowParams object
// with the ability to set a context for a request.
func NewInitializeSelfServiceRecoveryViaAPIFlowParamsWithContext(ctx context.Context) *InitializeSelfServiceRecoveryViaAPIFlowParams {
	return &InitializeSelfServiceRecoveryViaAPIFlowParams{
		Context: ctx,
	}
}

// NewInitializeSelfServiceRecoveryViaAPIFlowParamsWithHTTPClient creates a new InitializeSelfServiceRecoveryViaAPIFlowParams object
// with the ability to set a custom HTTPClient for a request.
func NewInitializeSelfServiceRecoveryViaAPIFlowParamsWithHTTPClient(client *http.Client) *InitializeSelfServiceRecoveryViaAPIFlowParams {
	return &InitializeSelfServiceRecoveryViaAPIFlowParams{
		HTTPClient: client,
	}
}

/* InitializeSelfServiceRecoveryViaAPIFlowParams contains all the parameters to send to the API endpoint
   for the initialize self service recovery via API flow operation.

   Typically these are written to a http.Request.
*/
type InitializeSelfServiceRecoveryViaAPIFlowParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the initialize self service recovery via API flow params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InitializeSelfServiceRecoveryViaAPIFlowParams) WithDefaults() *InitializeSelfServiceRecoveryViaAPIFlowParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the initialize self service recovery via API flow params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InitializeSelfServiceRecoveryViaAPIFlowParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the initialize self service recovery via API flow params
func (o *InitializeSelfServiceRecoveryViaAPIFlowParams) WithTimeout(timeout time.Duration) *InitializeSelfServiceRecoveryViaAPIFlowParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the initialize self service recovery via API flow params
func (o *InitializeSelfServiceRecoveryViaAPIFlowParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the initialize self service recovery via API flow params
func (o *InitializeSelfServiceRecoveryViaAPIFlowParams) WithContext(ctx context.Context) *InitializeSelfServiceRecoveryViaAPIFlowParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the initialize self service recovery via API flow params
func (o *InitializeSelfServiceRecoveryViaAPIFlowParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the initialize self service recovery via API flow params
func (o *InitializeSelfServiceRecoveryViaAPIFlowParams) WithHTTPClient(client *http.Client) *InitializeSelfServiceRecoveryViaAPIFlowParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the initialize self service recovery via API flow params
func (o *InitializeSelfServiceRecoveryViaAPIFlowParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *InitializeSelfServiceRecoveryViaAPIFlowParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
