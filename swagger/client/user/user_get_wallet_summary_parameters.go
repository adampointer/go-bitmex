// Code generated by go-swagger; DO NOT EDIT.

package user

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

// NewUserGetWalletSummaryParams creates a new UserGetWalletSummaryParams object
// with the default values initialized.
func NewUserGetWalletSummaryParams() *UserGetWalletSummaryParams {
	var (
		currencyDefault = string("XBt")
	)
	return &UserGetWalletSummaryParams{
		Currency: &currencyDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewUserGetWalletSummaryParamsWithTimeout creates a new UserGetWalletSummaryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUserGetWalletSummaryParamsWithTimeout(timeout time.Duration) *UserGetWalletSummaryParams {
	var (
		currencyDefault = string("XBt")
	)
	return &UserGetWalletSummaryParams{
		Currency: &currencyDefault,

		timeout: timeout,
	}
}

// NewUserGetWalletSummaryParamsWithContext creates a new UserGetWalletSummaryParams object
// with the default values initialized, and the ability to set a context for a request
func NewUserGetWalletSummaryParamsWithContext(ctx context.Context) *UserGetWalletSummaryParams {
	var (
		currencyDefault = string("XBt")
	)
	return &UserGetWalletSummaryParams{
		Currency: &currencyDefault,

		Context: ctx,
	}
}

// NewUserGetWalletSummaryParamsWithHTTPClient creates a new UserGetWalletSummaryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUserGetWalletSummaryParamsWithHTTPClient(client *http.Client) *UserGetWalletSummaryParams {
	var (
		currencyDefault = string("XBt")
	)
	return &UserGetWalletSummaryParams{
		Currency:   &currencyDefault,
		HTTPClient: client,
	}
}

/*UserGetWalletSummaryParams contains all the parameters to send to the API endpoint
for the user get wallet summary operation typically these are written to a http.Request
*/
type UserGetWalletSummaryParams struct {

	/*Currency*/
	Currency *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the user get wallet summary params
func (o *UserGetWalletSummaryParams) WithTimeout(timeout time.Duration) *UserGetWalletSummaryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the user get wallet summary params
func (o *UserGetWalletSummaryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the user get wallet summary params
func (o *UserGetWalletSummaryParams) WithContext(ctx context.Context) *UserGetWalletSummaryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the user get wallet summary params
func (o *UserGetWalletSummaryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the user get wallet summary params
func (o *UserGetWalletSummaryParams) WithHTTPClient(client *http.Client) *UserGetWalletSummaryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the user get wallet summary params
func (o *UserGetWalletSummaryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCurrency adds the currency to the user get wallet summary params
func (o *UserGetWalletSummaryParams) WithCurrency(currency *string) *UserGetWalletSummaryParams {
	o.SetCurrency(currency)
	return o
}

// SetCurrency adds the currency to the user get wallet summary params
func (o *UserGetWalletSummaryParams) SetCurrency(currency *string) {
	o.Currency = currency
}

// WriteToRequest writes these params to a swagger request
func (o *UserGetWalletSummaryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Currency != nil {

		// query param currency
		var qrCurrency string
		if o.Currency != nil {
			qrCurrency = *o.Currency
		}
		qCurrency := qrCurrency
		if qCurrency != "" {
			if err := r.SetQueryParam("currency", qCurrency); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}