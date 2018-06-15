package main

import (
	"BBS-Example/api/bbs-example-server/app"
	"github.com/goadesign/goa"
)

// HomeController implements the home resource.
type HomeController struct {
	*goa.Controller
}

// NewHomeController creates a home controller.
func NewHomeController(service *goa.Service) *HomeController {
	return &HomeController{Controller: service.NewController("HomeController")}
}

// Home runs the home action.
func (c *HomeController) Home(ctx *app.HomeHomeContext) error {
	return ctx.OK([]byte("Home."))
}
