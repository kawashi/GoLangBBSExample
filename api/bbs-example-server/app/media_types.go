// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "bbs-example-server": Application Media Types
//
// Command:
// $ goagen
// --design=BBS-Example/api/bbs-example-server/design
// --out=$(GOPATH)/src/BBS-Example/api/bbs-example-server
// --version=v1.3.1

package app

// User post. (default view)
//
// Identifier: application/json; view=default
type JSON struct {
	// messages
	Messages []string `form:"messages,omitempty" json:"messages,omitempty" yaml:"messages,omitempty" xml:"messages,omitempty"`
}
