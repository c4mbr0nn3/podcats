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
	LastPlayPosition int       `gorm:"default:0"`
	PublicationDate  time.Time `gorm:"type:datetime"`
	FileURL          string
	GUID             string
	Image            string
	IsPlayed         bool `gorm:"default:false"`
	BookmarkDate     *time.Time
	PodcastID        int
	UserID           int
}
