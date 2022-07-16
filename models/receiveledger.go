package models

import (
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type ReceiveLedger struct {
	gorm.Model
	ID         string    `gorm:"index;size:21" form:"id"`
	WhsID      string    `gorm:"size:21" form:"whs_id"`
	FactoryID  string    `gorm:"size:21" form:"factory_id"`
	RssGroupID string    `gorm:"size:21" form:"rss_group_id"`
	IsActive   bool      `gorm:"null" form:"is_active" default:"false"`
	CreatedAt  time.Time `form:"created_at" default:"now"`
	UpdatedAt  time.Time `form:"updated_at" default:"now"`
	Whs        Whs       `gorm:"foreignKey:WhsID;references:ID"`
	Factory    Factory   `gorm:"foreignKey:FactoryID;references:ID"`
	RssGroup   RssGroup  `gorm:"foreignKey:RssGroupID;references:ID"`
}

func (u *ReceiveLedger) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New(21)
	u.ID = id
	return
}
