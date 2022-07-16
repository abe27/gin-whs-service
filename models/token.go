package models

import (
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type JwtToken struct {
	ID        string    `gorm:"size:21" json:"id"`
	Key       string    `gorm:"size:65" json:"key" binding:"required"`
	UserID    string    `gorm:"unique;not null;size:21" json:"user_id" binding:"required"`
	Token     string    `gorm:"size:255" json:"token" binding:"required"`
	IsActive  bool      `json:"is_active" default:"true"`
	CreatedAt time.Time `json:"created_at" default:"now"`
	UpdatedAt time.Time `json:"updated_at" default:"now"`
}

func (u *JwtToken) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New(21)
	TokenID, _ := g.New(56)
	u.ID = id
	u.Key = TokenID
	u.IsActive = true
	return
}
