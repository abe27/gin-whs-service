package models

import (
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Whs struct {
	ID          string    `gorm:"primaryKey;index;size:21" form:"id" json:"id"`
	Name        string    `gorm:"unique;not null;size:25" form:"name" json:"name" binding:"required"`
	Slug        string    `gorm:"not null;size:5" form:"slug" json:"slug" binding:"required"`
	Description string    `gorm:"null;size:50" form:"description" json:"description"`
	IsActive    bool      `form:"is_active" json:"is_active" default:"false"`
	CreatedAt   time.Time `form:"created_at" json:"created_at" default:"now"`
	UpdatedAt   time.Time `form:"updated_at" json:"updated_at" default:"now"`
}

func (u *Whs) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New(21)
	u.ID = id
	return
}

func (u *Whs) AfterUpdate(tx *gorm.DB) (err error) {
	u.UpdatedAt = time.Now().UTC()
	return
}
