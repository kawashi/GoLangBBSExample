package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("bbs-example-server", func() {
	Title("BBS Example Server.")
	Description("A teaser for goa.")
	Host("localhost:8888")
	Scheme("http")
})

var _ = Resource("user_post", func() {
	BasePath("/user_posts")
	Action("index", func() {
		Description("Fetch user posts.")
		Routing(GET("/"))
		Response(OK, UserPostMedia)
	})

	Action("create", func() {
		Description("Create user post.")
		Payload(UserPostPayload)
		Routing(POST("/"))
		Response(OK)
	})
})

var UserPostPayload = Type("UserPostPayload", func() {
	Attribute("message", String, "User post message.")
})

var _ = Resource("ping", func() {
	Action("ping", func() {
		Routing(GET("/ping"))
		Description("PingPong.")
		Response(OK, "text/plain")
	})
})

var UserPostMedia = MediaType("application/json", func() {
	Description("User post.")
	Attributes(func() {
		Attribute("messages", ArrayOf(String), "messages", func() {
			Example([]string{"messages1", "message2"})
		})
	})
	View("default", func() {
		Attribute("messages")
	})
})
