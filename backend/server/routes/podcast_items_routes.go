package routes

import (
	"example/hello/controllers"
	"example/hello/middlewares"

	"github.com/gin-gonic/gin"
)

func AddPodcastItemsRoutes(rg *gin.RouterGroup) {
	itemsController := new(controllers.PodcastItemsController)
	routes := rg.Group("/podcast-items").Use(middlewares.Auth())
	{
		routes.GET("/latest", itemsController.GetLatestPodcastItems)
		routes.GET("/favourites", itemsController.GetFavouritesPodcastItems)
		routes.GET("/:itemId", itemsController.GetPodcastItemById)
		routes.POST("/:itemId/update-played-time", itemsController.UpdatePodcastItemPlayedTime)
		routes.POST("/:itemId/set-completed", itemsController.SetPodcastItemCompleted)
		routes.POST("/:itemId/switch-played-status", itemsController.SwitchPodcastItemPlayedStatus)
		routes.POST("/:itemId/switch-fav-status", itemsController.SwitchPodcastItemFavouriteStatus)
	}
}
