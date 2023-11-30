package invoice

import "time"

type Invoice struct {
	InvoiceID  uint       `gorm:"primaryKey;autoIncrement" json:"invoice_id"`
	IssueDate  *time.Time `gorm:"type:date" json:"issue_date"`
	CustomerID uint       `gorm:"index" json:"customer_id"`
	Subject    string     `gorm:"type:varchar(255)" json:"subject"`
	DueDate    *time.Time `gorm:"type:date" json:"due_date"`
	TotalItems int        `json:"total_items"`
	SubTotal   float64    `gorm:"type:decimal(10,2)" json:"sub_total"`
	Tax        float64    `gorm:"type:decimal(10,2)" json:"tax"`
	GrandTotal float64    `gorm:"type:decimal(10,2)" json:"grand_total"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
}
