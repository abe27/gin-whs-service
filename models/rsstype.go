package models

import (
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type RssGroup struct {
	ID          string    `gorm:"primaryKey;index;size:21" form:"id"`
	Name        string    `gorm:"unique;not null;size:50" form:"name" json:"name" binding:"required"`
	Value       string    `gorm:"not null;size:10" form:"value" json:"value" binding:"required"`
	Description string    `gorm:"null" form:"description" json:"description" default:"-"`
	IsActive    bool      `gorm:"null" form:"is_active" json:"is_active" default:"false"`
	CreatedAt   time.Time `form:"created_at" default:"now"`
	UpdatedAt   time.Time `form:"updated_at" default:"now"`
}

func (u *RssGroup) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New(21)
	u.ID = id
	return
}
