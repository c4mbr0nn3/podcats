package controllers

import (
	"example/hello/db"
	"example/hello/db/models"
	"log"
	"math"
	"net/http"
	"sort"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mmcdole/gofeed"

	md "github.com/JohannesKaufmann/html-to-markdown"
)

var Converter = md.NewConverter("", true, nil)

type PodcastsController struct{}

type PodcastStats struct {
	PodcastID uint
	IsPlayed  bool
	Count     int
}

type PodcastStatsKey struct {
	PodcastID uint
	IsPlayed  bool
}

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
	statsMap := getPodcastsStats()

	var toReturn []models.Podcast
	for _, podcast := range podcasts {
		podcast.PlayedCount = statsMap[PodcastStatsKey{podcast.ID, true}].Count
		podcast.EpisodesCount = statsMap[PodcastStatsKey{podcast.ID, true}].Count + statsMap[PodcastStatsKey{podcast.ID, false}].Count
		toReturn = append(toReturn, podcast)
	}

	ctx.JSON(http.StatusOK, toReturn)
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
		"podcastItems": podcastItems,
		"pageCount":    pageCount,
		"thisPage":     page,
		"nextPage":     nextPage,
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

	podcastItems := podcast.PodcastItems
	sort.Slice(podcastItems[:], func(i, j int) bool {
		return podcastItems[i].PublicationDate.After(podcastItems[j].PublicationDate)
	})

	ctx.JSON(http.StatusOK, podcast)
}

func (c PodcastsController) GetPodcastItemById(ctx *gin.Context) {
	itemId := ctx.Param("itemId")
	var podcastItem models.PodcastItem
	result := db.GetDb().Find(&podcastItem, itemId)
	if result.Error != nil {
		ctx.AbortWithError(http.StatusBadRequest, result.Error)
		return
	}
	ctx.JSON(http.StatusOK, podcastItem)
}

func (c PodcastsController) UpdatePodcastItemPlayedTime(ctx *gin.Context) {
	timePlayedString := ctx.DefaultQuery("time", "0")
	timePlayed, err := strconv.Atoi(timePlayedString)
	if err != nil || timePlayed < 0 {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	itemId := ctx.Param("itemId")
	var podcastItem models.PodcastItem
	result := db.GetDb().Find(&podcastItem, itemId)
	if result.Error != nil {
		ctx.AbortWithError(http.StatusBadRequest, result.Error)
		return
	}

	podcastItem.LastPlayPosition = timePlayed
	db.GetDb().Save(&podcastItem)

	ctx.Status(http.StatusOK)
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
			FileURL:         getSafeFileURL(item),
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
	url := ctx.DefaultQuery("url", "")
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL(url)
	ctx.JSON(http.StatusOK, feed)
}

func getSafeFileURL(item *gofeed.Item) string {
	if item.Enclosures == nil || len(item.Enclosures) == 0 {
		return ""
	}
	return item.Enclosures[0].URL
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

func getPodcastsStats() map[PodcastStatsKey]PodcastStats {
	var stats []PodcastStats
	db.GetDb().Model(&models.PodcastItem{}).
		Select("podcast_id, is_played, count(is_played) as count").
		Group("podcast_id, is_played").
		Find(&stats)

	statsMap := map[PodcastStatsKey]PodcastStats{}
	for _, v := range stats {
		statsMap[PodcastStatsKey{v.PodcastID, v.IsPlayed}] = v
	}
	return statsMap
}

/*func getPodcastLastEpisode(statsMap map[PodcastStatsKey]PodcastStats, podcastID uint) time.Time {
	if statsMap[PodcastStatsKey{podcastID, true}].LastPippo.After(statsMap[PodcastStatsKey{podcastID, false}].LastPippo) {
		return statsMap[PodcastStatsKey{podcastID, true}].LastPippo
	}
	return statsMap[PodcastStatsKey{podcastID, false}].LastPippo
}*/
