package entity

import "time"

type Blacklist struct {
	ID        string    `gorm:"primary_key" json:"id"`
	Token     string    `gorm:"unique, not null" json:"token"`
	CreatedAt time.Time `json:"created_at"`
}

func (b *Blacklist) TableName() string {
	return "mst_blacklist"
}
