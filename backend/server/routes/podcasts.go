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
		usersRouter.POST("/import", controller.ImportPodcast)
		usersRouter.POST("/test", controller.GetPodcastFake)
	}
}
