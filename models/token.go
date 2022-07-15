package models

import "time"

type JwtToken struct {
	ID        string    `gorm:"size:21" json:"id"`
	Key       string    `gorm:"size:65" json:"key"`
	UserID    string    `gorm:"unique;not null;size:21" json:"user_id"`
	Token     string    `gorm:"size:255" json:"token"`
	CreatedAt time.Time `json:"created_at" default:"now"`
	UpdatedAt time.Time `json:"updated_at" default:"now"`
}
