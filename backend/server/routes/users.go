package routes

import (
	"example/hello/controllers"

	"github.com/gin-gonic/gin"
)

func AddUsersRoutes(rg *gin.RouterGroup) {
	controller := new(controllers.UsersController)
	usersRouter := rg.Group("/users")
	{
		usersRouter.GET("/", controller.GetAllUsers)
	}
}
