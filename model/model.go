package model

import "github.com/jinzhu/gorm"

// User table
type User struct {
	gorm.Model
	Name  string
	Email string
}
