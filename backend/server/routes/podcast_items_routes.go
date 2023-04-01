package routes

import (
	"example/hello/controllers"
	"example/hello/middlewares"

	"github.com/gin-gonic/gin"
)

func AddPodcastItemsRoutes(rg *gin.RouterGroup) {
	itemsController := new(controllers.PodcastItemsController)
	usersRouter := rg.Group("/podcast-items").Use(middlewares.Auth())
	{
		usersRouter.GET("/latest", itemsController.GetLatestPodcastItems)
		usersRouter.GET("/favourites", itemsController.GetFavouritesPodcastItems)
		usersRouter.GET("/:itemId", itemsController.GetPodcastItemById)
		usersRouter.POST("/:itemId/update-played-time", itemsController.UpdatePodcastItemPlayedTime)
		usersRouter.POST("/:itemId/set-completed", itemsController.SetPodcastItemCompleted)
		usersRouter.POST("/:itemId/switch-played-status", itemsController.SwitchPodcastItemPlayedStatus)
		usersRouter.POST("/:itemId/switch-fav-status", itemsController.SwitchPodcastItemFavouriteStatus)
	}
}
