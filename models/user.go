package models

import "time"

type User struct {
	ID         string    `gorm:"size:36" json:"id"`
	UserName   string    `gorm:"column:username;unique;not null;size:10" json:"username"`
	Password   string    `gorm:"not null;size:255" json:"-"`
	Email      string    `gorm:"default:null;size:50" json:"email"`
	IsVerified bool      `json:"is_verified" default:"false"`
	CreatedAt  time.Time `json:"created_at" default:"now"`
	UpdatedAt  time.Time `json:"updated_at" default:"now"`
}
