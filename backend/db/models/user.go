package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Surname  string `gorm:"type:text"`
	Email    string
	Password string
}
