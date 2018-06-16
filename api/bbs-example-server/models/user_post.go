package models

import "github.com/jinzhu/gorm"

type UserPost struct {
	gorm.Model
	Message string
}