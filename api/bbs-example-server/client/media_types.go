// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "bbs-example-server": Application Media Types
//
// Command:
// $ goagen
// --design=BBS-Example/api/bbs-example-server/design
// --out=$(GOPATH)/src/BBS-Example/api/bbs-example-server
// --version=v1.3.1

package client

import (
	"net/http"
)

// User post. (default view)
//
// Identifier: application/json; view=default
type JSON struct {
	// messages
	Messages []string `form:"messages,omitempty" json:"messages,omitempty" yaml:"messages,omitempty" xml:"messages,omitempty"`
}

// DecodeJSON decodes the JSON instance encoded in resp body.
func (c *Client) DecodeJSON(resp *http.Response) (*JSON, error) {
	var decoded JSON
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}
