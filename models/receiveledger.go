package models

import (
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type ReceiveLedger struct {
	ID         string    `gorm:"primaryKey;index;size:21" form:"id" json:"id"`
	WhsID      string    `gorm:"size:21" form:"whs_id" json:"whs_id"`
	FactoryID  string    `gorm:"size:21" form:"factory_id" json:"factory_id"`
	RssGroupID string    `gorm:"size:21" form:"rss_group_id" json:"rss_group_id"`
	IsActive   bool      `gorm:"null" form:"is_active" json:"is_active" default:"false"`
	CreatedAt  time.Time `form:"created_at" json:"created_at" default:"now"`
	UpdatedAt  time.Time `form:"updated_at" json:"updated_at" default:"now"`
	Whs        Whs       `gorm:"foreignKey:WhsID;references:ID"`
	Factory    Factory   `gorm:"foreignKey:FactoryID;references:ID"`
	RssGroup   RssGroup  `gorm:"foreignKey:RssGroupID;references:ID"`
}

func (u *ReceiveLedger) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New(21)
	u.ID = id
	return
}
