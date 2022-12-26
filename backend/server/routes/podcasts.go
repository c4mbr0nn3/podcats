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
		usersRouter.GET("/latest-items", controller.GetLatestPodcastItems)
		usersRouter.GET("/:id", controller.GetPodcastById)
		usersRouter.GET("/:id/item/:itemId", controller.GetPodcastItemById)
		usersRouter.POST("/import", controller.ImportPodcast)
		usersRouter.GET("/test", controller.GetPodcastFake)
	}
}
