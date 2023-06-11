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
		usersRouter.POST("/create", controller.CreateUser)
		usersRouter.POST("/:id/update", controller.UpdateUser)
		usersRouter.GET("/:id/reset-password", controller.ResetUserPassword)
		usersRouter.POST("/:id/set-password", controller.SetUserPassword)
		usersRouter.DELETE("/:id/delete", controller.DeleteUser)
	}
}
