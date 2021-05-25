package models

import (
	"time"
)

type Like struct {
	ID        uint `gorm:"primaryKey;autoIncrement"`
	UserID    uint `gorm:"column:user_id"`
	User      User
	PostID    uint `gorm:"column:post_id"`
	Post      Post
	CreatedAt time.Time
	UpdatedAt time.Time
}
