package routes

import (
	"example/hello/controllers"
	"example/hello/middlewares"

	"github.com/gin-gonic/gin"
)

func AddNotificationsRoutes(rg *gin.RouterGroup) {
	notificationsController := new(controllers.NotificationsController)
	routes := rg.Group("/notifications").Use(middlewares.Auth())
	{
		routes.GET("/", notificationsController.GetAllNotifications)
		routes.POST("/set-all-read", notificationsController.MarkAllAsRead)
		routes.POST("/:id/set-read", notificationsController.MarkAsRead)
	}
}
