//go:generate goagen bootstrap -d BBS-Example/api/bbs-example-server/design

package main

import (
	"BBS-Example/api/bbs-example-server/app"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
)

func main() {
	// Create service
	service := goa.New("bbs-example-server")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "home" controller
	c := NewHomeController(service)
	app.MountHomeController(service, c)
	// Mount "ping" controller
	c2 := NewPingController(service)
	app.MountPingController(service, c2)

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}

}
