package server

import (
	"example/hello/server/routes"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	v1 := router.Group("/api/v1")
	routes.AddUsersRoutes(v1)
	routes.AddPodcastsRoutes(v1)

	return router
}
