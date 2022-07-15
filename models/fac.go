package models

import "time"

type Factory struct {
	ID          string    `gorm:"size:21" json:"id"`
	Name        string    `gorm:"unique;not null;size:50" json:"name"`
	InvPrefix   string    `gorm:"null;size:5" json:"inv_prefix"`
	WhsPrefix   string    `gorm:"null;size:5" json:"whs_prefix"`
	RssType     string    `gorm:"null;size:5" json:"rss_type"`
	Description string    `gorm:"null" json:"description"`
	IsActive    bool      `gorm:"null" json:"is_active" default:"false"`
	CreatedAt   time.Time `json:"created_at" default:"now"`
	UpdatedAt   time.Time `json:"updated_at" default:"now"`
}
