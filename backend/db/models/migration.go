package models

import (
	"time"

	"gorm.io/gorm"
)

type Migration struct {
	gorm.Model
	Date *time.Time
	Name string
}
