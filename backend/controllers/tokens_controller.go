package controllers

import (
	"example/hello/auth"
	"example/hello/db"
	"example/hello/db/models"
	"example/hello/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TokenRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type TokensController struct{}

func (c TokensController) CreateToken(ctx *gin.Context) {
	var tokenRequest TokenRequest
	if err := ctx.ShouldBindJSON(&tokenRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var user models.User
	result := db.GetDb().Where("username = ?", tokenRequest.Username).First(&user)
	if result.Error != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}
	if err := user.CheckPassword(tokenRequest.Password); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}
	token, err := auth.GenerateJwtToken(int(user.ID), user.Username, user.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	userDto := dto.NewUserDto(user)

	ctx.JSON(http.StatusOK, gin.H{"token": token, "user": userDto})
}
