package main

import (
	"BBS-Example/api/bbs-example-server/app"
	"github.com/goadesign/goa"
	"BBS-Example/api/bbs-example-server/models"
)

// UserPostController implements the user_post resource.
type UserPostController struct {
	*goa.Controller
}

// NewUserPostController creates a user_post controller.
func NewUserPostController(service *goa.Service) *UserPostController {
	return &UserPostController{Controller: service.NewController("UserPostController")}
}

// Create runs the create action.
func (c *UserPostController) Create(ctx *app.CreateUserPostContext) error {
	// UserPostController_Create: start_implement

	// Put your logic here

	return nil
	// UserPostController_Create: end_implement
}

// Index runs the index action.
func (c *UserPostController) Index(ctx *app.IndexUserPostContext) error {
	userPosts := &[]models.UserPost{}
	db.Find(&userPosts)
	
	return ctx.OK(UserPostsToJSON(userPosts))
}

func UserPostsToJSON(userPosts *[]models.UserPost) *app.JSON {
	res := &app.JSON{}

	for _, value := range *userPosts {
		res.Messages = append(res.Messages, value.Message)
	}

	return res
}