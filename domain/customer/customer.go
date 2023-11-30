package customer

import (
	"time"
)

type Customer struct {
	CustomerID    int       `gorm:"primaryKey;autoIncrement" json:"customer_id"`
	CustomerName  string    `json:"customer_name"`
	DetailAddress string    `json:"detail_address"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
