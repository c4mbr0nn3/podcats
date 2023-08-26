package controllers

import (
	"example/hello/auth"
	"example/hello/db"
	"example/hello/db/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type NotificationsController struct{}

func (c NotificationsController) GetAllNotifications(ctx *gin.Context) {
	jwt, _ := auth.GetJwtClaims(ctx)
	userId := jwt.Id

	var notifications []models.Notification
	result := db.GetDb().Where("user_id = ? AND is_read = ?", userId, false).Order("publication_date DESC").Find(&notifications)
	if result.Error != nil {
		ctx.AbortWithError(http.StatusBadRequest, result.Error)
		return
	}

	ctx.JSON(http.StatusOK, notifications)
}

func (c NotificationsController) MarkAllAsRead(ctx *gin.Context) {
	jwt, _ := auth.GetJwtClaims(ctx)
	userId := jwt.Id

	var notifications []models.Notification
	result := db.GetDb().Where("user_id = ? AND is_read = ?", userId, false).Find(&notifications)
	if result.Error != nil {
		ctx.AbortWithError(http.StatusBadRequest, result.Error)
		return
	}

	for _, notification := range notifications {
		notification.IsRead = true
		db.GetDb().Save(&notification)
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "All notifications marked as read"})
}

func (c NotificationsController) MarkAsRead(ctx *gin.Context) {
	jwt, _ := auth.GetJwtClaims(ctx)
	userId := jwt.Id

	notificationId := ctx.Param("id")
	var notification models.Notification
	result := db.GetDb().Find(&notification, notificationId)
	if result.Error != nil {
		ctx.AbortWithError(http.StatusBadRequest, result.Error)
		return
	}

	if notification.UserId != userId {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	notification.IsRead = true
	db.GetDb().Save(&notification)

	ctx.JSON(http.StatusOK, gin.H{"message": "Notification marked as read"})
}
