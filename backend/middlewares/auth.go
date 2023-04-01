package middlewares

import (
	"example/hello/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authToken := ctx.GetHeader("Authorization")
		if authToken == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized", "message": "No token provided"})
			return
		}
		err := auth.ValidateJwtToken(authToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized", "message": err.Error()})
			return
		}
		ctx.Next()
	}
}
