package entity

import "time"

type History struct {
	ID         string    `gorm:"primary_key" json:"id"`
	CustomerID string    `gorm:"not null" json:"customer_id"`
	MerchantID string    `gorm:"not null" json:"merchant_id"`
	Amount     uint      `json:"amount"`
	Action     string    `gorm:"not null" json:"action,omitempty"`
	CreatedAt  time.Time `json:"created_at"`
}

func (h *History) TableName() string {
	return "trx_history"
}
