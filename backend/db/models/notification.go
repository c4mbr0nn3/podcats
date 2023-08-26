package models

import (
	"time"

	"gorm.io/gorm"
)

type Notification struct {
	gorm.Model
	UserId          int
	PodcastId       int
	PodcastItemId   int
	PublicationDate time.Time `gorm:"type:datetime"`
	Message         string
	IsRead          bool
}
