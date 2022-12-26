package controllers

import (
	"example/hello/db"
	"example/hello/db/models"
	"log"
	"math"
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
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": podcasts})
}

func (c PodcastsController) GetLatestPodcastItems(ctx *gin.Context) {
	pageString := ctx.DefaultQuery("page", "1")
	page, err := strconv.Atoi(pageString)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var podcastItemsCount int64
	if err := db.GetDb().Table("podcast_items").Count(&podcastItemsCount).Error; err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	const podcastItemsPerPage = 10
	pageCount := int(math.Ceil(float64(podcastItemsCount) / float64(podcastItemsPerPage)))
	if pageCount == 0 {
		pageCount = 1
	}
	if page < 1 || page > pageCount {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	offset := (page - 1) * podcastItemsPerPage
	var podcastItems []models.PodcastItem
	result := db.GetDb().Limit(podcastItemsPerPage).Offset(offset).Order("publication_date desc").Find(&podcastItems)
	if result.Error != nil {
		ctx.AbortWithError(http.StatusBadRequest, result.Error)
		return
	}

	var nextPage int
	if page < pageCount {
		nextPage = page + 1
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":      podcastItems,
		"pageCount": pageCount,
		"thisPage":  page,
		"nextPage":  nextPage,
	})
}

func (c PodcastsController) GetPodcastById(ctx *gin.Context) {
	podcastId := ctx.Param("id")
	var podcast models.Podcast
	result := db.GetDb().Preload("PodcastItems").Find(&podcast, podcastId)
	if result.Error != nil {
		ctx.AbortWithError(http.StatusBadRequest, result.Error)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": podcast})
}

func (c PodcastsController) GetPodcastItemById(ctx *gin.Context) {
	itemId := ctx.Param("itemId")
	var podcastItem models.PodcastItem
	result := db.GetDb().Find(&podcastItem, itemId)
	if result.Error != nil {
		ctx.AbortWithError(http.StatusBadRequest, result.Error)
		return
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
		return
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
