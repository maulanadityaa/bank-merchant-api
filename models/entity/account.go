package entity

type Account struct {
	ID       string `gorm:"primary_key" json:"id"`
	Email    string `gorm:"unique, not null" json:"email"`
	Password string `json:"password"`
	IsLogged bool   `json:"is_logged"`
	RoleID   string `json:"role_id"`

	// Relationship
	Customer Customer `gorm:"foreignKey:AccountID" json:"customer"`
	Merchant Merchant `gorm:"foreignKey:AccountID" json:"merchant"`
}

func (a *Account) TableName() string {
	return "mst_account"
}
