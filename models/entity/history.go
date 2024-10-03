package entity

import "time"

type History struct {
	ID         string    `gorm:"primary_key" json:"id"`
	CustomerID *string   `json:"customer_id"`
	MerchantID *string   `json:"merchant_id"`
	Amount     uint      `json:"amount"`
	Action     string    `gorm:"not null" json:"action,omitempty"`
	CreatedAt  time.Time `json:"created_at"`
}

func (h *History) TableName() string {
	return "trx_history"
}
