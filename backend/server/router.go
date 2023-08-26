package server

import (
	"example/hello/config"
	"example/hello/server/routes"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	env := config.GetConfig().GetString("app_env")
	router := gin.New()
	router.SetTrustedProxies([]string{"127.0.0.1"})
	if env == "prod" {
		router.Use(static.Serve("/", static.LocalFile("./dist", false)))
	} else {
		router.Use(CORSMiddleware())
	}
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	v1 := router.Group("/api/v1")
	routes.AddLoginRoutes(v1)
	routes.AddUsersRoutes(v1)
	routes.AddPodcastsRoutes(v1)
	routes.AddPodcastItemsRoutes(v1)
	routes.AddNotificationsRoutes(v1)

	return router
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
