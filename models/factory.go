package models

import (
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Factory struct {
	ID          string    `gorm:"size:21" form:"id"`
	Name        string    `gorm:"unique;not null;size:50" form:"name" binding:"required"`
	Prefix      string    `gorm:"not null;size:5" form:"prefix" binding:"required"`
	Description string    `gorm:"null" form:"description" default:"-"`
	IsActive    bool      `gorm:"null" form:"is_active" default:"false"`
	CreatedAt   time.Time `form:"created_at" default:"now"`
	UpdatedAt   time.Time `form:"updated_at" default:"now"`
}

func (u *Factory) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New(21)
	u.ID = id
	return
}
