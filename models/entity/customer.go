package entity

import "time"

type Customer struct {
	ID        string    `gorm:"primary_key" json:"id"`
	Name      string    `json:"name"`
	Balance   uint64    `gorm:"default:0" json:"balance"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// Relationship
	// One to One with Account
	AccountID string `json:"account_id"`
	// One to Many with History
	History []History `gorm:"foreignKey:CustomerID" json:"history"`
}

func (c *Customer) TableName() string {
	return "mst_customer"
}
