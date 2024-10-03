package utils

import (
	"strings"

	"gorm.io/gorm"
)

func SelectByName(name string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("LOWER(name) LIKE ?", "%"+strings.ToLower(name)+"%")
	}
}
