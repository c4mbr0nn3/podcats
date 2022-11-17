package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PingController struct{}

func (c PingController) GetPong(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "pong!"})
}
