package models

import (
	"time"

	"gorm.io/gorm"
)

type Podcast struct {
	gorm.Model
	Title        string
	Summary      string `gorm:"type:text"`
	Author       string
	Image        string
	URL          string
	LastEpisode  *time.Time
	PodcastItems []PodcastItem
	IsPaused     bool `gorm:"default:false"`
	UserId       int
	User         User
}
