package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("bbs-example-server", func() {
	Title("BBS Example Server.")
	Description("A teaser for goa.")
	Host("localhost:8080")
	Scheme("http")
})

var _ = Resource("ping", func() {
	Action("ping", func() {
		Routing(GET("/ping"))
		Description("PingPong.")
		Response(OK, "text/plain")
	})
})

var _ = Resource("home", func() {
	Action("home", func() {
		Routing(GET("/home"))
		Description("home.")
		Response(OK, UserPostMedia)
	})
})

var UserPostMedia = MediaType("application/json", func() {
	Description("User post.")
	Attributes(func() {
		Attribute("message", String, "message", func() {
			Example("message")
		})
	})
	View("default", func() {
		Attribute("message")
	})
})
