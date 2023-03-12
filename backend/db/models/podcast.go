package models

import (
	"time"

	"gorm.io/gorm"
)

type Podcast struct {
	gorm.Model
	Title         string
	Summary       string `gorm:"type:text"`
	Author        string
	Image         string
	ShowURL       string
	FeedURL       string
	PodcastItems  []PodcastItem
	UserId        int
	User          User
	EpisodesCount int       `gorm:"-"`
	PlayedCount   int       `gorm:"-"`
	LatestEpisode time.Time `gorm:"-"`
}
