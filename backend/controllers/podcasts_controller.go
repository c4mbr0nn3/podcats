package controllers

import (
	"example/hello/auth"
	"example/hello/db"
	"example/hello/db/models"
	"example/hello/dto"
	"example/hello/utils"
	"math"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mmcdole/gofeed"
)

type PodcastsController struct{}

func (c PodcastsController) GetAllPodcasts(ctx *gin.Context) {
	jwt, _ := auth.GetJwtClaims(ctx)
	userId := jwt.Id

	var podcasts []models.Podcast
	result := db.GetDb().Where("user_id = ?", userId).Find(&podcasts)
	if result.Error != nil {
		ctx.AbortWithError(http.StatusBadRequest, result.Error)
		return
	}
	podcastStatMap := getPodcastsStats()

	var toReturn []models.Podcast
	for _, podcast := range podcasts {
		stats := podcastStatMap[podcast.ID]
		counters := stats.Counters
		podcast.PlayedCount = counters[dto.PodcastStatsKey{PodcastID: podcast.ID, IsPlayed: true}].Count
		podcast.EpisodesCount = counters[dto.PodcastStatsKey{PodcastID: podcast.ID, IsPlayed: true}].Count + counters[dto.PodcastStatsKey{PodcastID: podcast.ID, IsPlayed: false}].Count
		podcast.LatestEpisode, _ = time.Parse("2006-01-02 15:04:05", strings.Split(stats.PodcastStats.LatestEpisode, "+")[0])
		toReturn = append(toReturn, podcast)
	}
	sort.Slice(toReturn[:], func(i int, j int) bool { return toReturn[i].LatestEpisode.After(toReturn[j].LatestEpisode) })

	ctx.JSON(http.StatusOK, toReturn)
}

func (c PodcastsController) GetPodcastById(ctx *gin.Context) {
	jwt, _ := auth.GetJwtClaims(ctx)
	userId := jwt.Id

	podcastId := ctx.Param("id")
	var podcast models.Podcast
	result := db.GetDb().Find(&podcast, podcastId)
	if result.Error != nil {
		ctx.AbortWithError(http.StatusBadRequest, result.Error)
		return
	}

	if podcast.UserId != userId {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	ctx.JSON(http.StatusOK, podcast)
}

func (c PodcastsController) GetPodcastItemsByPodcastId(ctx *gin.Context) {
	jwt, _ := auth.GetJwtClaims(ctx)
	userId := jwt.Id

	podcastId := ctx.Param("id")
	var podcast models.Podcast
	result := db.GetDb().Find(&podcast, podcastId)
	if result.Error != nil {
		ctx.AbortWithError(http.StatusBadRequest, result.Error)
		return
	}

	if podcast.UserId != userId {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

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

func (c PodcastsController) DeletePodcastById(ctx *gin.Context) {
	jwt, _ := auth.GetJwtClaims(ctx)
	userId := jwt.Id

	podcastId := ctx.Param("id")
	var podcast models.Podcast
	result := db.GetDb().Find(&podcast, podcastId)
	if result.Error != nil {
		ctx.AbortWithError(http.StatusBadRequest, result.Error)
		return
	}

	if podcast.UserId != userId {
		ctx.AbortWithStatus(http.StatusUnauthorized)
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
	jwt, _ := auth.GetJwtClaims(ctx)
	userId := jwt.Id

	podcastId := ctx.Param("id")
	var podcast models.Podcast
	result := db.GetDb().Find(&podcast, podcastId)
	if result.Error != nil {
		ctx.AbortWithError(http.StatusBadRequest, result.Error)
		return
	}

	if podcast.UserId != userId {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	podcastItemsResult := db.GetDb().
		Where("podcast_id = ? AND is_played = ?", podcastId, false).
		Updates(models.PodcastItem{IsPlayed: true})
	if podcastItemsResult.Error != nil {
		ctx.AbortWithError(http.StatusBadRequest, podcastItemsResult.Error)
		return
	}
	ctx.Status(http.StatusOK)
}

func (c PodcastsController) ImportPodcast(ctx *gin.Context) {
	jwt, _ := auth.GetJwtClaims(ctx)
	userId := jwt.Id

	body := dto.PodcastImportDto{}

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
		UserId:  userId,
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
			UserID:          userId,
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

func getPodcastsStats() map[uint]dto.PodcastStatsDetail {
	var podcastStats []dto.PodcastStats
	db.GetDb().Model(&models.PodcastItem{}).
		Select("podcast_id, MAX(publication_date) as latest_episode").
		Group("podcast_id").
		Find(&podcastStats)

	var playedStatus []dto.PodcastPlayedStatus
	db.GetDb().Model(&models.PodcastItem{}).
		Select("podcast_id, is_played, count(is_played) as count").
		Group("podcast_id, is_played").
		Find(&playedStatus)

	statusMap := map[dto.PodcastStatsKey]dto.PodcastPlayedStatus{}
	for _, v := range playedStatus {
		statusMap[dto.PodcastStatsKey{PodcastID: v.PodcastID, IsPlayed: v.IsPlayed}] = v
	}

	statsMap := map[uint]dto.PodcastStatsDetail{}
	for _, v := range podcastStats {
		y := dto.PodcastStatsDetail{PodcastStats: v, Counters: map[dto.PodcastStatsKey]dto.PodcastPlayedStatus{}}
		statsMap[v.PodcastID] = y
		if x, found := statusMap[dto.PodcastStatsKey{PodcastID: v.PodcastID, IsPlayed: true}]; found {
			y.Counters[dto.PodcastStatsKey{PodcastID: v.PodcastID, IsPlayed: true}] = x
		}
		if x, found := statusMap[dto.PodcastStatsKey{PodcastID: v.PodcastID, IsPlayed: false}]; found {
			y.Counters[dto.PodcastStatsKey{PodcastID: v.PodcastID, IsPlayed: false}] = x
		}
	}
	return statsMap
}
