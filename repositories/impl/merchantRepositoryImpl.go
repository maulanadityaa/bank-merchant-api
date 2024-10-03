package impl

import (
	"github.com/maulanadityaa/bank-merchant-api/config"
	"github.com/maulanadityaa/bank-merchant-api/models/entity"
	"github.com/maulanadityaa/bank-merchant-api/utils"
	"gorm.io/gorm"
)

type MerchantRepository struct{}

func NewMerchantRepository() *MerchantRepository {
	return &MerchantRepository{}
}

func (MerchantRepository) AddMerchant(merchant entity.Merchant) (entity.Merchant, error) {
	if result := config.DB.Create(&merchant); result.Error != nil {
		return entity.Merchant{}, result.Error
	}

	return merchant, nil
}

func (MerchantRepository) UpdateMerchant(merchant entity.Merchant) (entity.Merchant, error) {
	if result := config.DB.Save(&merchant); result.Error != nil {
		return entity.Merchant{}, result.Error
	}

	return merchant, nil
}

func (MerchantRepository) GetMerchantByID(merchantID string) (entity.Merchant, error) {
	var merchant entity.Merchant
	if result := config.DB.Where("id = ?", merchantID).First(&merchant); result.Error != nil {
		return entity.Merchant{}, result.Error
	}

	return merchant, nil
}

func (MerchantRepository) GetMerchantByAccountID(accountID string) (entity.Merchant, error) {
	var merchant entity.Merchant
	if result := config.DB.Where("account_id = ?", accountID).First(&merchant); result.Error != nil {
		return entity.Merchant{}, result.Error
	}

	return merchant, nil
}

func (MerchantRepository) GetAllMerchant(spec []func(db *gorm.DB) *gorm.DB, name string) ([]entity.Merchant, string, error) {
	var merchants []entity.Merchant

	if name != "" {
		spec = append(spec, utils.SelectByName(name))
	}

	db := config.DB.Model(&entity.Merchant{}).Scopes(spec[1:]...)
	totalRows := utils.GetTotalRows(db)
	err := db.Scopes(spec[0]).Find(&merchants).Error

	return merchants, totalRows, err
}
