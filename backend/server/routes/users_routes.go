package routes

import (
	"example/hello/controllers"
	"example/hello/middlewares"

	"github.com/gin-gonic/gin"
)

func AddUsersRoutes(rg *gin.RouterGroup) {
	controller := new(controllers.UsersController)
	routes := rg.Group("/users").Use(middlewares.Auth())
	{
		routes.GET("/", controller.GetAllUsers)
		routes.POST("/create", controller.CreateUser)
		routes.POST("/:id/update", controller.UpdateUser)
		routes.GET("/:id/reset-password", controller.ResetUserPassword)
		routes.POST("/:id/set-password", controller.SetUserPassword)
		routes.DELETE("/:id/delete", controller.DeleteUser)
	}
}
