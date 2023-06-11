package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name               string `json:"name" gorm:"type:text"`
	Surname            string `json:"surname" gorm:"type:text"`
	Username           string `json:"username"`
	Email              string `json:"email"`
	IsAdmin            bool   `json:"is_admin" gorm:"default:false"`
	Password           string `json:"password" gorm:"type:text"`
	NeedPasswordChange bool   `json:"need_password_change" gorm:"default:true"`
}

func (user *User) HashPassword() error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

func (user *User) CheckPassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return err
	}
	return nil
}
