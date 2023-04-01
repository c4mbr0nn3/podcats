package routes

import (
	"example/hello/controllers"
	"example/hello/middlewares"

	"github.com/gin-gonic/gin"
)

func AddUsersRoutes(rg *gin.RouterGroup) {
	controller := new(controllers.UsersController)
	usersRouter := rg.Group("/users").Use(middlewares.Auth())
	{
		usersRouter.GET("/", controller.GetAllUsers)
		usersRouter.POST("/register", controller.RegisterUser)
	}
}
