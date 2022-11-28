package controllers

import (
	"example/hello/db"
	"example/hello/db/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mmcdole/gofeed"
)

type PodcastsController struct{}

type PodcastImportDto struct {
	PodcastUrl string `json:"podcastUrl"`
}

func (c PodcastsController) GetAllPodcasts(ctx *gin.Context) {
	var podcasts []models.Podcast
	result := db.GetDb().Find(&podcasts)
	if result.Error != nil {
		ctx.AbortWithError(http.StatusBadRequest, result.Error)
	}
	ctx.JSON(http.StatusOK, gin.H{"data": podcasts})
}

func (c PodcastsController) GetPodcastById(ctx *gin.Context) {
	podcastId := ctx.Param("id")
	var podcast models.Podcast
	result := db.GetDb().Preload("PodcastItems").First(&podcast, podcastId)
	if result.Error != nil {
		ctx.AbortWithError(http.StatusBadRequest, result.Error)
	}
	ctx.JSON(http.StatusOK, gin.H{"data": podcast})
}

func (c PodcastsController) ImportPodcast(ctx *gin.Context) {
	body := PodcastImportDto{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL(body.PodcastUrl)
	podcast := models.Podcast{
		Title:   feed.Title,
		Summary: feed.Description,
		Author:  feed.Author.Name,
		Image:   feed.Image.URL,
		URL:     feed.FeedLink,
	}

	var podcastItemsArray []models.PodcastItem
	for _, item := range feed.Items {
		podcastItemsArray = append(podcastItemsArray, models.PodcastItem{
			Title:           item.Title,
			Summary:         item.Description,
			PubblicatonDate: *item.PublishedParsed,
			FileURL:         item.Enclosures[0].URL,
			GUID:            item.GUID,
			Image:           item.Image.URL,
		})
	}
	podcast.PodcastItems = podcastItemsArray
	tx := db.GetDb().Create(&podcast)
	if tx.Error != nil {
		ctx.AbortWithError(http.StatusBadRequest, tx.Error)
	}
	ctx.Status(http.StatusOK)
}

func (c PodcastsController) GetPodcastFake(ctx *gin.Context) {
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL("https://www.spreaker.com/show/3262415/episodes/feed")
	ctx.JSON(http.StatusOK, gin.H{"data": feed})
}
