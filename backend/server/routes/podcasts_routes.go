package routes

import (
	"example/hello/controllers"
	"example/hello/middlewares"

	"github.com/gin-gonic/gin"
)

func AddPodcastsRoutes(rg *gin.RouterGroup) {
	podcastsController := new(controllers.PodcastsController)
	routes := rg.Group("/podcasts").Use(middlewares.Auth())
	{
		routes.GET("/", podcastsController.GetAllPodcasts)
		routes.GET("/:id", podcastsController.GetPodcastById)
		routes.GET("/:id/items", podcastsController.GetPodcastItemsByPodcastId)
		routes.DELETE("/:id/remove", podcastsController.DeletePodcastById)
		routes.POST("/:id/set-all-played", podcastsController.SetAllPlayed)
		routes.POST("/import", podcastsController.ImportPodcast)
	}
}
