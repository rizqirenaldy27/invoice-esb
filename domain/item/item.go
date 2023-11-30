package item

import (
	"time"
)

type Item struct {
	ItemID     int       `gorm:"primaryKey" json:"item_id"`
	ItemName   string    `gorm:"type:varchar(255);uniqueIndex;not null" json:"item_name"`
	UnitPrice  float64   `json:"unit_price"`
	ItemTypeID int       `gorm:"not null" json:"item_type_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type ItemWithTypeName struct {
	ItemID       int       `json:"item_id"`
	ItemName     string    `json:"item_name"`
	UnitPrice    float64   `json:"unit_price"`
	ItemTypeID   int       `json:"item_type_id"`
	ItemTypeName string    `json:"item_type_name"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
