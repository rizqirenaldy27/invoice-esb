package invoice

import (
	"time"
)

type Invoice struct {
	InvoiceID  uint       `gorm:"primaryKey;autoIncrement" json:"invoice_id"`
	IssueDate  *time.Time `gorm:"type:date" json:"issue_date"`
	CustomerID uint       `gorm:"index" json:"customer_id"`
	Subject    string     `gorm:"type:varchar(255)" json:"subject"`
	DueDate    *time.Time `gorm:"type:date" json:"due_date"`
	Status     string     `gorm:"type:varchar(255)" json:"status"`
	TotalItems int        `json:"total_items"`
	SubTotal   float64    `gorm:"type:decimal(10,2)" json:"sub_total"`
	Tax        float64    `gorm:"type:decimal(10,2)" json:"tax"`
	GrandTotal float64    `gorm:"type:decimal(10,2)" json:"grand_total"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
}

type InvoiceItem struct {
	InvoiceItemID uint      `gorm:"primaryKey;autoIncrement" json:"invoice_item_id"`
	InvoiceID     uint      `gorm:"index" json:"invoice_id"`
	ItemID        uint      `gorm:"index" json:"item_id"`
	Quantity      int       `json:"quantity"`
	Amount        float64   `gorm:"type:decimal(10,2)" json:"amount"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type InvoiceInput struct {
	IssueDate  string             `json:"issue_date"`
	DueDate    string             `json:"due_date"`
	CustomerID uint               `json:"customer_id"`
	Subject    string             `json:"subject"`
	Status     string             `json:"status"`
	Items      []InvoiceItemInput `json:"items"`
	TotalItems int                `json:"total_items"`
	SubTotal   float64            `json:"sub_total"`
	Tax        float64            `json:"tax"`
	GrandTotal float64            `json:"grand_total"`
}

type InvoiceItemInput struct {
	InvoiceItemID uint    `json:"invoice_item_id"`
	ItemID        uint    `json:"item_id"`
	Quantity      int     `json:"quantity"`
	UnitPrice     float64 `json:"unit_price"`
	Amount        float64 `json:"amount"`
}

type InvoiceOutput struct {
	InvoiceID  string              `json:"invoice_id"`
	IssueDate  string              `json:"issue_date"`
	DueDate    string              `json:"due_date"`
	CustomerID uint                `json:"customer_id"`
	Subject    string              `json:"subject"`
	Status     string              `json:"status"`
	Items      []InvoiceItemOutput `json:"items"`
	TotalItems int                 `json:"total_items"`
	SubTotal   float64             `json:"sub_total"`
	Tax        float64             `json:"tax"`
	GrandTotal float64             `json:"grand_total"`
}

type InvoiceItemOutput struct {
	InvoiceItemID uint    `json:"invoice_item_id"`
	ItemID        uint    `json:"item_id"`
	Quantity      int     `json:"quantity"`
	UnitPrice     float64 `json:"unit_price"`
	Amount        float64 `json:"amount"`
}

type InvoiceOutputRead struct {
	InvoiceID  string  `json:"invoice_id"`
	IssueDate  string  `json:"issue_date"`
	DueDate    string  `json:"due_date"`
	CustomerID uint    `json:"customer_id"`
	Subject    string  `json:"subject"`
	Status     string  `json:"status"`
	TotalItems int     `json:"total_items"`
	SubTotal   float64 `json:"sub_total"`
	Tax        float64 `json:"tax"`
	GrandTotal float64 `json:"grand_total"`
}

type InvoiceFilter struct {
	InvoiceID  string
	Subject    string
	CustomerID string
	TotalItems string
	IssueDate  string
	DueDate    string
	Status     string
}
