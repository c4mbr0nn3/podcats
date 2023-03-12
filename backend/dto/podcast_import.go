package dto

type PodcastImportDto struct {
	PodcastUrl string `json:"podcastUrl" binding:"required"`
}
