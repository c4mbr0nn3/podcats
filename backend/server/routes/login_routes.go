package routes

import (
	"example/hello/controllers"

	"github.com/gin-gonic/gin"
)

func AddLoginRoutes(rg *gin.RouterGroup) {
	controller := new(controllers.TokensController)
	loginRouter := rg.Group("/auth")
	{
		loginRouter.POST("/login", controller.CreateToken)
	}
}
