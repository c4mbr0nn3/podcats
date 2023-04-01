package routes

import (
	"example/hello/controllers"
	"example/hello/middlewares"

	"github.com/gin-gonic/gin"
)

func AddPodcastsRoutes(rg *gin.RouterGroup) {
	podcastsController := new(controllers.PodcastsController)
	usersRouter := rg.Group("/podcasts").Use(middlewares.Auth())
	{
		usersRouter.GET("/", podcastsController.GetAllPodcasts)
		usersRouter.GET("/:id", podcastsController.GetPodcastById)
		usersRouter.GET("/:id/items", podcastsController.GetPodcastItemsByPodcastId)
		usersRouter.DELETE("/:id/remove", podcastsController.DeletePodcastById)
		usersRouter.POST("/:id/set-all-played", podcastsController.SetAllPlayed)
		usersRouter.POST("/import", podcastsController.ImportPodcast)
	}
}
