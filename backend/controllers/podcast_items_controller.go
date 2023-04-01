package controllers

import (
	"example/hello/auth"
	"example/hello/db"
	"example/hello/db/models"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type PodcastItemsController struct{}

func (c PodcastItemsController) GetLatestPodcastItems(ctx *gin.Context) {
	jwt, _ := auth.GetJwtClaims(ctx)
	userId := jwt.Id

	pageString := ctx.DefaultQuery("page", "1")
	page, err := strconv.Atoi(pageString)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var podcastItemsCount int64
	countQuery := db.GetDb().Table("podcast_items").
		Where("user_id = ?", userId).
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
	result := db.GetDb().
		Limit(podcastItemsPerPage).
		Offset(offset).
		Order("publication_date desc").
		Where("user_id = ?", userId).
		Find(&podcastItems)
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

func (c PodcastItemsController) GetFavouritesPodcastItems(ctx *gin.Context) {
	jwt, _ := auth.GetJwtClaims(ctx)
	userId := jwt.Id

	pageString := ctx.DefaultQuery("page", "1")
	page, err := strconv.Atoi(pageString)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var podcastItemsCount int64
	countQuery := db.GetDb().Table("podcast_items").
		Where("user_id = ?", userId).
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
		Where("user_id = ?", userId).
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

func (c PodcastItemsController) GetPodcastItemById(ctx *gin.Context) {
	jwt, _ := auth.GetJwtClaims(ctx)
	userId := jwt.Id

	itemId := ctx.Param("itemId")
	var podcastItem models.PodcastItem
	result := db.GetDb().Find(&podcastItem, itemId)
	if result.Error != nil {
		ctx.AbortWithError(http.StatusBadRequest, result.Error)
		return
	}

	if podcastItem.UserID != userId {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	ctx.JSON(http.StatusOK, podcastItem)
}

func (c PodcastItemsController) UpdatePodcastItemPlayedTime(ctx *gin.Context) {
	jwt, _ := auth.GetJwtClaims(ctx)
	userId := jwt.Id

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

	if podcastItem.UserID != userId {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	podcastItem.LastPlayPosition = timePlayed
	db.GetDb().Save(&podcastItem)

	ctx.Status(http.StatusOK)
}

func (c PodcastItemsController) SetPodcastItemCompleted(ctx *gin.Context) {
	jwt, _ := auth.GetJwtClaims(ctx)
	userId := jwt.Id

	itemId := ctx.Param("itemId")
	var podcastItem models.PodcastItem
	result := db.GetDb().Find(&podcastItem, itemId)
	if result.Error != nil {
		ctx.AbortWithError(http.StatusBadRequest, result.Error)
		return
	}

	if podcastItem.UserID != userId {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	podcastItem.IsPlayed = true
	db.GetDb().Save(&podcastItem)
	ctx.Status(http.StatusOK)
}

func (c PodcastItemsController) SwitchPodcastItemPlayedStatus(ctx *gin.Context) {
	jwt, _ := auth.GetJwtClaims(ctx)
	userId := jwt.Id

	itemId := ctx.Param("itemId")
	var podcastItem models.PodcastItem
	result := db.GetDb().Find(&podcastItem, itemId)
	if result.Error != nil {
		ctx.AbortWithError(http.StatusBadRequest, result.Error)
		return
	}

	if podcastItem.UserID != userId {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	podcastItem.IsPlayed = !podcastItem.IsPlayed
	db.GetDb().Save(&podcastItem)
	ctx.Status(http.StatusOK)
}

func (c PodcastItemsController) SwitchPodcastItemFavouriteStatus(ctx *gin.Context) {
	jwt, _ := auth.GetJwtClaims(ctx)
	userId := jwt.Id

	itemId := ctx.Param("itemId")
	var podcastItem models.PodcastItem
	result := db.GetDb().Find(&podcastItem, itemId)
	if result.Error != nil {
		ctx.AbortWithError(http.StatusBadRequest, result.Error)
		return
	}

	if podcastItem.UserID != userId {
		ctx.AbortWithStatus(http.StatusUnauthorized)
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
