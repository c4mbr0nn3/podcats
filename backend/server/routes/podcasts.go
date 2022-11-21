package routes

import (
	"example/hello/controllers"

	"github.com/gin-gonic/gin"
)

func AddPodcastsRoutes(rg *gin.RouterGroup) {
	controller := new(controllers.PodcastsController)
	usersRouter := rg.Group("/podcasts")
	{
		usersRouter.POST("/import", controller.ImportPodcast)
	}
}
