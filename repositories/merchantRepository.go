package repositories

import (
	"github.com/maulanadityaa/bank-merchant-api/models/entity"
	"gorm.io/gorm"
)

type MerchantRepository interface {
	AddMerchant(merchant entity.Merchant) (entity.Merchant, error)
	UpdateMerchant(merchant entity.Merchant) (entity.Merchant, error)
	GetMerchantByID(merchantID string) (entity.Merchant, error)
	GetMerchantByAccountID(accountID string) (entity.Merchant, error)
	GetAllMerchant(spec []func(db *gorm.DB) *gorm.DB, name string) ([]entity.Merchant, string, error)
}
