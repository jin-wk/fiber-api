package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type RegisterUser struct {
	ID              int       `json:"id" swaggerignore:"true"`
	Email           string    `json:"email" validate:"required,email,min=10,max=50" example:"example@example.com"`
	Password        string    `json:"password" validate:"required_with=password_confirm,eqfield=PasswordConfirm,min=6,max=20" example:"example"`
	PasswordConfirm string    `json:"password_confirm" example:"example"`
	Name            string    `json:"name" validate:"required,min=2,max=20" example:"example"`
	CreatedAt       time.Time `json:"created_at" swaggerignore:"true"`
	UpdatedAt       time.Time `json:"updated_at" swaggerignore:"true"`
}

func (u *RegisterUser) BeforeSave(tx *gorm.DB) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hash)
	return nil
}

type User struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
