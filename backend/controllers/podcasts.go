package controllers

import (
	"example/hello/db"
	"example/hello/db/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mmcdole/gofeed"
)

type PodcastsController struct{}

func (c PodcastsController) ImportPodcast(ctx *gin.Context) {
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL("https://www.spreaker.com/show/3262415/episodes/feed")
	podcast := models.Podcast{
		Title:   feed.Title,
		Summary: feed.Description,
		Author:  feed.Author.Name,
		Image:   feed.Image.URL,
		URL:     feed.FeedLink,
	}
	tx := db.GetDb().Create(&podcast)
	fmt.Println(tx.Error)
	ctx.Status(http.StatusOK)
}
