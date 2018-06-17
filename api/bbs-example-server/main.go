//go:generate goagen bootstrap -d BBS-Example/api/bbs-example-server/design

package main

import (
	"BBS-Example/api/bbs-example-server/app"
	"BBS-Example/api/bbs-example-server/models"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"os"
)

var db *gorm.DB

func init() {
	var err error

	pass := os.Getenv("DB_PASS")
	db, err = gorm.Open("mysql", "root:" + pass + "@/bbs_example?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		log.Printf("%s", err)
		os.Exit(1)
	}
	db.LogMode(true)
}

func main() {
	db.AutoMigrate(&models.UserPost{})

	// Create service
	service := goa.New("bbs-example-server")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "user_post" controller
	c := NewUserPostController(service)
	app.MountUserPostController(service, c)
	// Mount "ping" controller
	c2 := NewPingController(service)
	app.MountPingController(service, c2)

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}

}
