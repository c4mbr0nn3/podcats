package models

import (
	"time"

	"gorm.io/gorm"
)

type PodcastItem struct {
	gorm.Model
	Title            string
	Summary          string `gorm:"type:text"`
	EpisodeType      string
	Duration         int
	LastPlayPosition int `gorm:"default:0"`
	PublicationDate  time.Time
	FileURL          string
	GUID             string
	Image            string
	IsPlayed         bool `gorm:"default:false"`
	BookmarkDate     time.Time
	PodcastID        int
}
