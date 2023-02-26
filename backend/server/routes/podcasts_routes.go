package routes

import (
	"example/hello/controllers"

	"github.com/gin-gonic/gin"
)

func AddPodcastsRoutes(rg *gin.RouterGroup) {
	controller := new(controllers.PodcastsController)
	usersRouter := rg.Group("/podcasts")
	{
		usersRouter.GET("/", controller.GetAllPodcasts)
		usersRouter.GET("/:id", controller.GetPodcastById)
		usersRouter.DELETE("/:id/remove", controller.DeletePodcastById)
		usersRouter.GET("/latest-items", controller.GetLatestPodcastItems)
		usersRouter.GET("/favourites-items", controller.GetFavouritesPodcastItems)
		usersRouter.GET("/:id/items", controller.GetPodcastItemsByPodcastId)
		usersRouter.POST("/:id/set-all-played", controller.SetAllPlayed)
		usersRouter.GET("/:id/items/:itemId", controller.GetPodcastItemById)
		usersRouter.POST("/:id/items/:itemId/update-played-time", controller.UpdatePodcastItemPlayedTime)
		usersRouter.POST("/:id/items/:itemId/set-completed", controller.SetPodcastItemCompleted)
		usersRouter.POST("/:id/items/:itemId/switch-status", controller.SwitchPodcastItemPlayedStatus)
		usersRouter.POST("/:id/items/:itemId/switch-fav-status", controller.SwitchPodcastItemFavouriteStatus)
		usersRouter.POST("/import", controller.ImportPodcast)
		usersRouter.GET("/test", controller.GetPodcastFake)
	}
}
