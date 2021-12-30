package model

import "time"

type User struct {
	ID        int       `json:"id" gorm:"primary_key"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type AddUser struct {
	Email    string `json:"email" validate:"required,email,min=10,max=50" example:"example@example.com"`
	Password string `json:"password" validate:"required,min=6,max=20" example:"example"`
	Name     string `json:"name" validate:"required,min=2,max=20" example:"example"`
}
