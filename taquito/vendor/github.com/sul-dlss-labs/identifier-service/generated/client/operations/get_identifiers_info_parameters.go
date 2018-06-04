// Code generated by go-swagger; DO NOT EDIT.

package operations

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

// NewGetIdentifiersInfoParams creates a new GetIdentifiersInfoParams object
// with the default values initialized.
func NewGetIdentifiersInfoParams() *GetIdentifiersInfoParams {

	return &GetIdentifiersInfoParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetIdentifiersInfoParamsWithTimeout creates a new GetIdentifiersInfoParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetIdentifiersInfoParamsWithTimeout(timeout time.Duration) *GetIdentifiersInfoParams {

	return &GetIdentifiersInfoParams{

		timeout: timeout,
	}
}

// NewGetIdentifiersInfoParamsWithContext creates a new GetIdentifiersInfoParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetIdentifiersInfoParamsWithContext(ctx context.Context) *GetIdentifiersInfoParams {

	return &GetIdentifiersInfoParams{

		Context: ctx,
	}
}

// NewGetIdentifiersInfoParamsWithHTTPClient creates a new GetIdentifiersInfoParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetIdentifiersInfoParamsWithHTTPClient(client *http.Client) *GetIdentifiersInfoParams {

	return &GetIdentifiersInfoParams{
		HTTPClient: client,
	}
}

/*GetIdentifiersInfoParams contains all the parameters to send to the API endpoint
for the get identifiers info operation typically these are written to a http.Request
*/
type GetIdentifiersInfoParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get identifiers info params
func (o *GetIdentifiersInfoParams) WithTimeout(timeout time.Duration) *GetIdentifiersInfoParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get identifiers info params
func (o *GetIdentifiersInfoParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get identifiers info params
func (o *GetIdentifiersInfoParams) WithContext(ctx context.Context) *GetIdentifiersInfoParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get identifiers info params
func (o *GetIdentifiersInfoParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get identifiers info params
func (o *GetIdentifiersInfoParams) WithHTTPClient(client *http.Client) *GetIdentifiersInfoParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get identifiers info params
func (o *GetIdentifiersInfoParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetIdentifiersInfoParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}