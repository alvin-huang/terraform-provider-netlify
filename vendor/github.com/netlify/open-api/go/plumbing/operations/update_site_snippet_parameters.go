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

	models "github.com/netlify/open-api/go/models"
)

// NewUpdateSiteSnippetParams creates a new UpdateSiteSnippetParams object
// with the default values initialized.
func NewUpdateSiteSnippetParams() *UpdateSiteSnippetParams {
	var ()
	return &UpdateSiteSnippetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateSiteSnippetParamsWithTimeout creates a new UpdateSiteSnippetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateSiteSnippetParamsWithTimeout(timeout time.Duration) *UpdateSiteSnippetParams {
	var ()
	return &UpdateSiteSnippetParams{

		timeout: timeout,
	}
}

// NewUpdateSiteSnippetParamsWithContext creates a new UpdateSiteSnippetParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateSiteSnippetParamsWithContext(ctx context.Context) *UpdateSiteSnippetParams {
	var ()
	return &UpdateSiteSnippetParams{

		Context: ctx,
	}
}

// NewUpdateSiteSnippetParamsWithHTTPClient creates a new UpdateSiteSnippetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateSiteSnippetParamsWithHTTPClient(client *http.Client) *UpdateSiteSnippetParams {
	var ()
	return &UpdateSiteSnippetParams{
		HTTPClient: client,
	}
}

/*UpdateSiteSnippetParams contains all the parameters to send to the API endpoint
for the update site snippet operation typically these are written to a http.Request
*/
type UpdateSiteSnippetParams struct {

	/*SiteID*/
	SiteID string
	/*Snippet*/
	Snippet *models.Snippet
	/*SnippetID*/
	SnippetID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update site snippet params
func (o *UpdateSiteSnippetParams) WithTimeout(timeout time.Duration) *UpdateSiteSnippetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update site snippet params
func (o *UpdateSiteSnippetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update site snippet params
func (o *UpdateSiteSnippetParams) WithContext(ctx context.Context) *UpdateSiteSnippetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update site snippet params
func (o *UpdateSiteSnippetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update site snippet params
func (o *UpdateSiteSnippetParams) WithHTTPClient(client *http.Client) *UpdateSiteSnippetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update site snippet params
func (o *UpdateSiteSnippetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSiteID adds the siteID to the update site snippet params
func (o *UpdateSiteSnippetParams) WithSiteID(siteID string) *UpdateSiteSnippetParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the update site snippet params
func (o *UpdateSiteSnippetParams) SetSiteID(siteID string) {
	o.SiteID = siteID
}

// WithSnippet adds the snippet to the update site snippet params
func (o *UpdateSiteSnippetParams) WithSnippet(snippet *models.Snippet) *UpdateSiteSnippetParams {
	o.SetSnippet(snippet)
	return o
}

// SetSnippet adds the snippet to the update site snippet params
func (o *UpdateSiteSnippetParams) SetSnippet(snippet *models.Snippet) {
	o.Snippet = snippet
}

// WithSnippetID adds the snippetID to the update site snippet params
func (o *UpdateSiteSnippetParams) WithSnippetID(snippetID string) *UpdateSiteSnippetParams {
	o.SetSnippetID(snippetID)
	return o
}

// SetSnippetID adds the snippetId to the update site snippet params
func (o *UpdateSiteSnippetParams) SetSnippetID(snippetID string) {
	o.SnippetID = snippetID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateSiteSnippetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param site_id
	if err := r.SetPathParam("site_id", o.SiteID); err != nil {
		return err
	}

	if o.Snippet != nil {
		if err := r.SetBodyParam(o.Snippet); err != nil {
			return err
		}
	}

	// path param snippet_id
	if err := r.SetPathParam("snippet_id", o.SnippetID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}