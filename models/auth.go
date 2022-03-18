package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type RegisterUser struct {
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required_with=password_confirm,eqfield=PasswordConfirm"`
	PasswordConfirm string `json:"password_confirm"`
	Name            string `json:"name" validate:"required"`
}

func (RegisterUser) TableName() string {
	return "users"
}

func (u *RegisterUser) BeforeSave(tx *gorm.DB) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hash)
	return nil
}

type LoginUser struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (LoginUser) TableName() string {
	return "users"
}

type ResponseUser struct {
	ID        int
	Email     string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (ResponseUser) TableName() string {
	return "users"
}
