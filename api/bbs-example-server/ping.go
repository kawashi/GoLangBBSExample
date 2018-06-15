package main

import (
	"BBS-Example/api/bbs-example-server/app"
	"github.com/goadesign/goa"
)

// PingController implements the ping resource.
type PingController struct {
	*goa.Controller
}

// NewPingController creates a ping controller.
func NewPingController(service *goa.Service) *PingController {
	return &PingController{Controller: service.NewController("PingController")}
}

// Ping runs the ping action.
func (c *PingController) Ping(ctx *app.PingPingContext) error {
	return ctx.OK([]byte("Pong."))
}
