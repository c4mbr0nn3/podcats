package models

import (
	"gorm.io/gorm"
)

type Podcast struct {
	gorm.Model
	Title         string
	Summary       string `gorm:"type:text"`
	Author        string
	Image         string
	URL           string
	PodcastItems  []PodcastItem
	UserId        int
	User          User
	EpisodesCount int `gorm:"-"`
	PlayedCount   int `gorm:"-"`
}
