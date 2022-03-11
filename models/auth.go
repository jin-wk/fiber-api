package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID              int       `json:"id"`
	Email           string    `json:"email" validate:"required,email"`
	Password        string    `json:"password" validate:"required_with=password_confirm,eqfield=PasswordConfirm"`
	PasswordConfirm string    `json:"password_confirm"`
	Name            string    `json:"name" validate:"required"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

func (u *User) BeforeSave(tx *gorm.DB) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hash)
	return nil
}

type ResponseUser struct {
	ID        int
	Email     string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *ResponseUser) TableName() string {
	return "users"
}
