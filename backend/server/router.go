package server

import (
	"example/hello/server/routes"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.SetTrustedProxies([]string{"127.0.0.1"})
	router.Use(static.Serve("/", static.LocalFile("./dist", false)))
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	v1 := router.Group("/api/v1")
	routes.AddLoginRoutes(v1)
	routes.AddUsersRoutes(v1)
	routes.AddPodcastsRoutes(v1)
	routes.AddPodcastItemsRoutes(v1)

	return router
}
