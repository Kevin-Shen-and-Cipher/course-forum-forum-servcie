package models

import (
	"time"
)

type Post struct {
	ID        uint      `json:"id" gorm:"primaryKey;auto_increment"`
	Title     string    `json:"title" gorm:"type:text;not null"`
	Content   string    `json:"content" gorm:"type:text;not null"`
	Score     uint8     `json:"score" gorm:"size:5;not null"`
	Views     uint      `json:"views" gorm:"default:0;not null"`
	State     bool      `json:"state" gorm:"default:false;not null"`
	CreateBy  string    `json:"create_by" gorm:"type:text;not null"`
	CreatedAt time.Time `json:"created_at,string" gorm:"type:time;autoCreateTime;not null"`
	UpdatedAt time.Time `json:"updated_at,string" gorm:"type:time;autoUpdateTime;not null"`
}

type CreatePost struct {
	Title    string `json:"title" binding:"required"`
	Content  string `json:"content" binding:"required"`
	Score    uint8  `json:"score" binding:"required"`
	CreateBy string `json:"create_by" binding:"required"`
}

type UpdatePost struct {
	State bool `json:"state"`
}

// TableName is Database Table Name of this model
func (post *Post) TableName() string {
	return "posts"
}
