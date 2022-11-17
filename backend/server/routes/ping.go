package routes

import (
	"example/hello/controllers"

	"github.com/gin-gonic/gin"
)

func AddPing(rg *gin.RouterGroup) {
	controller := new(controllers.PingController)
	pingRouter := rg.Group("/ping")
	{
		pingRouter.GET("/", controller.GetPong)
	}
}
