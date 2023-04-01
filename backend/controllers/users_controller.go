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

func (c UsersController) RegisterUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if checkIfUserAlreadyPresent(user) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User already present"})
		return
	}

	if err := user.HashPassword(); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	result := db.GetDb().Create(&user)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"userId": user.ID, "email": user.Email, "username": user.Username, "message": "User created"})
}

func checkIfUserAlreadyPresent(user models.User) bool {
	var existingUser models.User
	result := db.GetDb().First(&existingUser, "username = ? OR email = ? ", user.Username, user.Email)
	return result.Error == nil
}
