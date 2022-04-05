package models

import "time"

type Board struct {
	Id        int       `json:"id"`
	UserId    int       `json:"user_id"`
	Title     string    `json:"string"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateBoard struct {
	Id      int    `json:"id"`
	UserId  int    `json:"user_id"`
	Title   string `json:"title" validate:"required"`
	Content string `json:"content" validate:"required"`
}

type UpdateBoard struct {
	Id      int    `json:"id" validate:"required"`
	Title   string `json:"title" validate:"required"`
	Content string `json:"content" validate:"required"`
}
