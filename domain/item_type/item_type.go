package item_type

import "time"

type ItemType struct {
	ItemTypeID   int       `gorm:"primaryKey" json:"item_type_id"`
	ItemTypeName string    `gorm:"type:varchar(255);uniqueIndex;not null" json:"item_type_name"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
