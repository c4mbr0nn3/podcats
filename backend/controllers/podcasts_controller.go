package controllers

import (
	"example/hello/db"
	"example/hello/db/models"
	"example/hello/utils"
	"math"
	"net/http"
	"sort"
	"strconv"
	"time"

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

func (c PodcastsController) GetPodcastById(ctx *gin.Context) {
	podcastId := ctx.Param("id")
	var podcast models.Podcast
	result := db.GetDb().Find(&podcast, podcastId)
	if result.Error != nil {
		ctx.AbortWithError(http.StatusBadRequest, result.Error)
		return
	}

	ctx.JSON(http.StatusOK, podcast)
}

func (c PodcastsController) DeletePodcastById(ctx *gin.Context) {
	podcastId := ctx.Param("id")
	var podcast models.Podcast
	result := db.GetDb().Find(&podcast, podcastId)
	if result.Error != nil {
		ctx.AbortWithError(http.StatusBadRequest, result.Error)
		return
	}
	var podcastItems []models.PodcastItem
	podcastItemsResult := db.GetDb().Where("podcast_id = ?", podcastId).Find(&podcastItems)
	if podcastItemsResult.Error != nil {
		ctx.AbortWithError(http.StatusBadRequest, podcastItemsResult.Error)
		return
	}
	db.GetDb().Delete(&podcast)
	db.GetDb().Delete(&podcastItems)
	ctx.Status(http.StatusOK)
}

func (c PodcastsController) SetAllPlayed(ctx *gin.Context) {
	podcastId := ctx.Param("id")
	podcastItemsResult := db.GetDb().
		Where("podcast_id = ? AND is_played = ?", podcastId, false).
		Updates(models.PodcastItem{IsPlayed: true})
	if podcastItemsResult.Error != nil {
		ctx.AbortWithError(http.StatusBadRequest, podcastItemsResult.Error)
		return
	}
	ctx.Status(http.StatusOK)
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

func (c PodcastsController) GetFavouritesPodcastItems(ctx *gin.Context) {
	pageString := ctx.DefaultQuery("page", "1")
	page, err := strconv.Atoi(pageString)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var podcastItemsCount int64
	countQuery := db.GetDb().Table("podcast_items").
		Where("bookmark_date IS NOT NULL").
		Count(&podcastItemsCount)
	if countQuery.Error != nil {
		ctx.AbortWithError(http.StatusInternalServerError, countQuery.Error)
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
	podcastItemsQuery := db.GetDb().
		Where("bookmark_date IS NOT NULL").
		Limit(podcastItemsPerPage).
		Offset(offset).
		Order("bookmark_date desc").
		Find(&podcastItems)
	if podcastItemsQuery.Error != nil {
		ctx.AbortWithError(http.StatusBadRequest, podcastItemsQuery.Error)
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

func (c PodcastsController) GetPodcastItemsByPodcastId(ctx *gin.Context) {
	podcastId := ctx.Param("id")

	pageString := ctx.DefaultQuery("page", "1")
	page, err := strconv.Atoi(pageString)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var podcastItemsCount int64
	countQuery := db.GetDb().
		Table("podcast_items").
		Where("podcast_id = ?", podcastId).
		Count(&podcastItemsCount)
	if countQuery.Error != nil {
		ctx.AbortWithError(http.StatusInternalServerError, countQuery.Error)
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
	podcastItemsQuery := db.GetDb().
		Limit(podcastItemsPerPage).
		Offset(offset).
		Order("publication_date desc").
		Where("podcast_id = ?", podcastId).
		Find(&podcastItems)
	if podcastItemsQuery.Error != nil {
		ctx.AbortWithError(http.StatusBadRequest, podcastItemsQuery.Error)
		return
	}

	var nextPage int
	if page < pageCount {
		nextPage = page + 1
	}

	sort.Slice(podcastItems[:], func(i, j int) bool {
		return podcastItems[i].PublicationDate.After(podcastItems[j].PublicationDate)
	})

	ctx.JSON(http.StatusOK, gin.H{
		"podcastItems": podcastItems,
		"pageCount":    pageCount,
		"thisPage":     page,
		"nextPage":     nextPage,
	})
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

func (c PodcastsController) SetPodcastItemCompleted(ctx *gin.Context) {
	itemId := ctx.Param("itemId")
	var podcastItem models.PodcastItem
	result := db.GetDb().Find(&podcastItem, itemId)
	if result.Error != nil {
		ctx.AbortWithError(http.StatusBadRequest, result.Error)
		return
	}

	podcastItem.IsPlayed = true
	db.GetDb().Save(&podcastItem)
	ctx.Status(http.StatusOK)
}

func (c PodcastsController) SwitchPodcastItemPlayedStatus(ctx *gin.Context) {
	itemId := ctx.Param("itemId")
	var podcastItem models.PodcastItem
	result := db.GetDb().Find(&podcastItem, itemId)
	if result.Error != nil {
		ctx.AbortWithError(http.StatusBadRequest, result.Error)
		return
	}

	podcastItem.IsPlayed = !podcastItem.IsPlayed
	db.GetDb().Save(&podcastItem)
	ctx.Status(http.StatusOK)
}

func (c PodcastsController) SwitchPodcastItemFavouriteStatus(ctx *gin.Context) {
	itemId := ctx.Param("itemId")
	var podcastItem models.PodcastItem
	result := db.GetDb().Find(&podcastItem, itemId)
	if result.Error != nil {
		ctx.AbortWithError(http.StatusBadRequest, result.Error)
		return
	}

	if podcastItem.BookmarkDate != nil {
		podcastItem.BookmarkDate = nil
	} else {
		nowTime := time.Now()
		podcastItem.BookmarkDate = &nowTime
	}
	db.GetDb().Save(&podcastItem)
	ctx.JSON(http.StatusOK, podcastItem.BookmarkDate)
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
		Summary: utils.ConvertHtmlToMarkdown(feed.Description),
		Author:  feed.Author.Name,
		Image:   feed.Image.URL,
		ShowURL: feed.Link,
		FeedURL: feed.FeedLink,
	}

	var podcastItemsArray []models.PodcastItem
	for _, item := range feed.Items {
		podcastItemsArray = append(podcastItemsArray, models.PodcastItem{
			Title:           item.Title,
			Summary:         utils.ConvertHtmlToMarkdown(item.Description),
			PublicationDate: *item.PublishedParsed,
			FileURL:         utils.GetSafeFileURL(item),
			GUID:            item.GUID,
			Image:           utils.GetSafeImageURL(item),
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
