package models

import (
	"time"
)

type Tag struct {
	ID        uint      `json:"id" gorm:"primaryKey;auto_increment"`
	Name      string    `json:"name" gorm:"type:text;not null"`
	CreatedAt time.Time `json:"created_at" gorm:"type:time;autoCreateTime;not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:time;autoUpdateTime;not null"`
	Posts     []Post    `json:"-" gorm:"many2many:post_tags;constraint:OnDelete:CASCADE"`
}

type CreateTag struct {
	Name string `json:"name" validate:"required,max=20"`
}

type UpdateTag struct {
	Name string `json:"name" validate:"required,max=20"`
}

// TableName is Database Table Name of this model
func (tag *Tag) TableName() string {
	return "tags"
}
