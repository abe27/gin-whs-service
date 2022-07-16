package models

import (
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Factory struct {
	ID          string    `gorm:"size:21" json:"id"`
	Name        string    `gorm:"unique;not null;size:50" json:"name" binding:"required"`
	Prefix      string    `gorm:"not null;size:5" json:"prefix" binding:"required"`
	Description string    `gorm:"null" json:"description" default:"-"`
	IsActive    bool      `gorm:"null" json:"is_active" default:"false"`
	CreatedAt   time.Time `json:"created_at" default:"now"`
	UpdatedAt   time.Time `json:"updated_at" default:"now"`
}

func (u *Factory) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New(21)
	u.ID = id
	return
}
