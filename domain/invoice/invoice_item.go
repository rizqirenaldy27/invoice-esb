package invoice

import "time"

type InvoiceItem struct {
	InvoiceItemID uint      `gorm:"primaryKey;autoIncrement" json:"invoice_item_id"`
	InvoiceID     uint      `gorm:"index" json:"invoice_id"`
	ItemID        uint      `gorm:"index" json:"item_id"`
	Quantity      int       `json:"quantity"`
	Amount        float64   `gorm:"type:decimal(10,2)" json:"amount"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
