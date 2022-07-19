package models

import (
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type User struct {
	ID         string    `gorm:"primaryKey;index;size:21" json:"id"`
	UserName   string    `gorm:"column:username;unique;not null;size:10" form:"username" json:"username" binding:"required"`
	Password   string    `gorm:"not null;size:60" form:"password" json:"-" binding:"required"`
	Email      string    `gorm:"default:null;size:50" form:"email" json:"email" default:"-"`
	IsVerified bool      `form:"is_verified" json:"is_verified" default:"false"`
	CreatedAt  time.Time `form:"created_at" json:"created_at" default:"now"`
	UpdatedAt  time.Time `form:"updated_at" json:"updated_at" default:"now"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New(21)
	u.ID = id
	return
}

func (u *User) AfterUpdate(tx *gorm.DB) (err error) {
	u.UpdatedAt = time.Now().UTC()
	return
}
