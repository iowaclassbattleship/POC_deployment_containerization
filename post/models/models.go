package models

import (
	"time"
)

type Post struct {
	ID        uint `gorm:"primaryKey"`
	Author    string
	Title     string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type RequestCreatePost struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}
