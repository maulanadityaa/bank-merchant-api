package entity

type Role struct {
	ID   string `gorm:"primary_key" json:"id"`
	Name string `json:"name"`

	// Relationship
	Accounts []Account `gorm:"foreignKey:RoleID" json:"accounts"`
}

func (r *Role) TableName() string {
	return "mst_role"
}
