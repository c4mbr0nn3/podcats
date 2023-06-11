package controllers

import (
	"example/hello/db"
	"example/hello/db/models"
	"example/hello/dto"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sethvargo/go-password/password"
)

type UsersController struct{}

func (c UsersController) GetAllUsers(ctx *gin.Context) {
	var users []models.User
	result := db.GetDb().Find(&users)
	if result.Error != nil {
		ctx.AbortWithError(http.StatusBadRequest, result.Error)
		return
	}

	// map all users to UserDto
	var usersDto []dto.UserDto
	for _, user := range users {
		dto := dto.NewUserDto(user)
		usersDto = append(usersDto, dto)
	}

	ctx.JSON(http.StatusOK, usersDto)
}

func (c UsersController) CreateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if isUserAlreadyPresent(user) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User already present"})
		return
	}

	password, err := password.Generate(12, 3, 0, false, false)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	user.Password = password

	if err := user.HashPassword(); err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	result := db.GetDb().Create(&user)
	if result.Error != nil {
		ctx.AbortWithError(http.StatusInternalServerError, result.Error)
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"password": password, "message": "User created"})
}

func (c UsersController) UpdateUser(ctx *gin.Context) {
	userId := ctx.Param("id")
	var user models.User
	result := db.GetDb().Find(&user, userId)
	if result.Error != nil {
		ctx.AbortWithError(http.StatusBadRequest, result.Error)
		return
	}

	var userDto dto.UserDto
	if err := ctx.ShouldBindJSON(&userDto); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	user.Email = userDto.Email
	user.Name = userDto.Name
	user.Surname = userDto.Surname
	user.IsAdmin = userDto.IsAdmin

	result = db.GetDb().Save(&user)
	if result.Error != nil {
		ctx.AbortWithError(http.StatusInternalServerError, result.Error)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "User updated"})
}

func (c UsersController) ResetUserPassword(ctx *gin.Context) {
	userId := ctx.Param("id")
	var user models.User
	result := db.GetDb().Find(&user, userId)
	if result.Error != nil {
		ctx.AbortWithError(http.StatusBadRequest, result.Error)
		return
	}

	password, err := password.Generate(12, 3, 0, false, false)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	user.Password = password
	user.NeedPasswordChange = true

	if err := user.HashPassword(); err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	result = db.GetDb().Save(&user)
	if result.Error != nil {
		ctx.AbortWithError(http.StatusInternalServerError, result.Error)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"password": password, "message": "User password reset"})
}

func (c UsersController) SetUserPassword(ctx *gin.Context) {
	userId := ctx.Param("id")
	var user models.User
	result := db.GetDb().Find(&user, userId)
	if result.Error != nil {
		ctx.AbortWithError(http.StatusBadRequest, result.Error)
		return
	}

	var passwordDto dto.PasswordDto
	if err := ctx.ShouldBindJSON(&passwordDto); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	user.Password = passwordDto.Password
	user.NeedPasswordChange = false

	if err := user.HashPassword(); err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	result = db.GetDb().Save(&user)
	if result.Error != nil {
		ctx.AbortWithError(http.StatusInternalServerError, result.Error)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "User password set"})
}

func (c UsersController) DeleteUser(ctx *gin.Context) {
	userId := ctx.Param("id")
	var user models.User
	result := db.GetDb().Find(&user, userId)
	if result.Error != nil {
		ctx.AbortWithError(http.StatusBadRequest, result.Error)
		return
	}

	result = db.GetDb().Delete(&user)
	if result.Error != nil {
		ctx.AbortWithError(http.StatusInternalServerError, result.Error)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}

func isUserAlreadyPresent(user models.User) bool {
	var existingUser models.User
	result := db.GetDb().Where(&existingUser, "username = ? OR email = ? ", user.Username, user.Email).Limit(1).Find(&existingUser)
	return result.RowsAffected > 0
}
