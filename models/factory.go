package models

import (
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Factory struct {
	ID          string    `gorm:"primaryKey;index;size:21" form:"id" json:"id"`
	Name        string    `gorm:"unique;not null;size:50" form:"name" json:"name" binding:"required"`
	Prefix      string    `gorm:"not null;size:5" form:"prefix" json:"prefix" binding:"required"`
	Description string    `gorm:"null" form:"description" json:"description" default:"-"`
	IsActive    bool      `gorm:"null" form:"is_active" json:"is_active" default:"false"`
	CreatedAt   time.Time `form:"created_at" json:"created_at" default:"now"`
	UpdatedAt   time.Time `form:"updated_at" json:"updated_at" default:"now"`
}

func (u *Factory) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New(21)
	u.ID = id
	return
}
