package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/docker/compliance/nlptooling/nlputil/textanalytics/client/operations"
)

// Default text analytics HTTP client.
var Default = NewHTTPClient(nil)

// NewHTTPClient creates a new text analytics HTTP client.
func NewHTTPClient(formats strfmt.Registry) *TextAnalytics {
	if formats == nil {
		formats = strfmt.Default
	}
	transport := httptransport.New("westus.api.cognitive.microsoft.com", "/text/analytics/v2.0", []string{"https"})
	return New(transport, formats)
}

// New creates a new text analytics client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *TextAnalytics {
	cli := new(TextAnalytics)
	cli.Transport = transport

	cli.Operations = operations.New(transport, formats)

	return cli
}

// TextAnalytics is a client for text analytics
type TextAnalytics struct {
	Operations *operations.Client

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *TextAnalytics) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport

	c.Operations.SetTransport(transport)

}
