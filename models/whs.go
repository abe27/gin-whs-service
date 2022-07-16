package models

import (
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Whs struct {
	ID          string    `gorm:"size:21" form:"id"`
	Name        string    `gorm:"unique;not null;size:25" form:"name" binding:"required"`
	Slug        string    `gorm:"not null;size:5" form:"slug" binding:"required"`
	Description string    `gorm:"null;size:50" form:"description"`
	IsActive    bool      `form:"is_active" default:"false"`
	CreatedAt   time.Time `form:"created_at" default:"now"`
	UpdatedAt   time.Time `form:"updated_at" default:"now"`
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
