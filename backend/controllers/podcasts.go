package controllers

import (
	"example/hello/db"
	"example/hello/db/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mmcdole/gofeed"

	md "github.com/JohannesKaufmann/html-to-markdown"
)

var Converter = md.NewConverter("", true, nil)

type PodcastsController struct{}

type PodcastImportDto struct {
	PodcastUrl string `json:"podcastUrl" binding:"required"`
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
	result := db.GetDb().Preload("PodcastItems").Find(&podcast, podcastId)
	if result.Error != nil {
		ctx.AbortWithError(http.StatusBadRequest, result.Error)
	}
	ctx.JSON(http.StatusOK, gin.H{"data": podcast})
}

func (c PodcastsController) GetPodcastItemById(ctx *gin.Context) {
	itemId := ctx.Param("itemId")
	var podcastItem models.PodcastItem
	result := db.GetDb().Find(&podcastItem, itemId)
	if result.Error != nil {
		ctx.AbortWithError(http.StatusBadRequest, result.Error)
	}
	ctx.JSON(http.StatusOK, gin.H{"data": podcastItem})
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
		Summary: convertHtmlToMarkdown(ctx, feed.Description),
		Author:  feed.Author.Name,
		Image:   feed.Image.URL,
		URL:     feed.FeedLink,
	}

	var podcastItemsArray []models.PodcastItem
	for _, item := range feed.Items {
		podcastItemsArray = append(podcastItemsArray, models.PodcastItem{
			Title:           item.Title,
			Summary:         convertHtmlToMarkdown(ctx, item.Description),
			PublicationDate: *item.PublishedParsed,
			FileURL:         item.Enclosures[0].URL,
			Duration:        convertStringToInt(ctx, item.ITunesExt.Duration),
			GUID:            item.GUID,
			Image:           getSafeImageURL(item),
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
	feed, _ := fp.ParseURL("https://www.spreaker.com/show/2952600/episodes/feed")
	ctx.JSON(http.StatusOK, gin.H{"data": feed})
}

func getSafeImageURL(item *gofeed.Item) string {
	if item.Image == nil {
		return ""
	}
	return item.Image.URL
}

func convertHtmlToMarkdown(ctx *gin.Context, toConvert string) string {
	markdown, err := Converter.ConvertString(toConvert)
	if err != nil {
		log.Fatal(err)
		ctx.AbortWithError(http.StatusBadRequest, err)
	}
	return markdown
}

func convertStringToInt(ctx *gin.Context, toConvert string) int {
	converted, err := strconv.Atoi(toConvert)
	if err != nil {
		log.Fatal(err)
		ctx.AbortWithError(http.StatusBadRequest, err)
	}
	return converted
}
