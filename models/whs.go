package models

import "time"

type Whs struct {
	ID          string    `gorm:"size:21" json:"id"`
	Name        string    `gorm:"unique;not null;size:25" json:"name"`
	Slug        string    `gorm:"not null;size:5" json:"slug"`
	Description string    `gorm:"null;size:50" json:"description"`
	IsActive    bool      `json:"is_active" default:"false"`
	CreatedAt   time.Time `json:"created_at" default:"now"`
	UpdatedAt   time.Time `json:"updated_at" default:"now"`
}
