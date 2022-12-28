package controllers

import (
	"example/hello/db"
	"example/hello/db/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UsersController struct{}

func (c UsersController) GetAllUsers(ctx *gin.Context) {
	var users []models.User
	result := db.GetDb().Find(&users)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"pippo": "culo"})
	}
	ctx.JSON(http.StatusOK, gin.H{"data": users})
}
