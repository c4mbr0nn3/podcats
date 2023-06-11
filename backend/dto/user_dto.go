package dto

import "example/hello/db/models"

type UserDto struct {
	ID                 uint   //`json:"id"`
	Name               string //`json:"name"`
	Surname            string //`json:"surname"`
	Username           string //`json:"username"`
	Email              string //`json:"email"`
	IsAdmin            bool   //`json:"is_admin"`
	CreatedAt          string //`json:"created_at"`
	UpdatedAt          string //`json:"updated_at"`
	NeedPasswordChange bool   //`json:"need_password_change"`
}

// constructor taking a models.User and returning a dto.UserDto
func NewUserDto(user models.User) UserDto {
	var userDto UserDto
	userDto.ID = user.ID
	userDto.Name = user.Name
	userDto.Surname = user.Surname
	userDto.Username = user.Username
	userDto.Email = user.Email
	userDto.IsAdmin = user.IsAdmin
	userDto.CreatedAt = user.CreatedAt.String()
	userDto.UpdatedAt = user.UpdatedAt.String()
	userDto.NeedPasswordChange = user.NeedPasswordChange
	return userDto
}
