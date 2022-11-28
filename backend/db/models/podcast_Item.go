package models

import (
	"time"

	"gorm.io/gorm"
)

type PodcastItem struct {
	gorm.Model
	Title           string
	Summary         string `gorm:"type:text"`
	EpisodeType     string
	Duration        int
	PubblicatonDate time.Time
	FileURL         string
	GUID            string
	Image           string
	IsPlayed        bool `gorm:"default:false"`
	BookmarkDate    time.Time
	PodcastID       int
}
