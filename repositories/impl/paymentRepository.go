package impl

import (
	"errors"

	"github.com/maulanadityaa/bank-merchant-api/config"
	"github.com/maulanadityaa/bank-merchant-api/models/entity"
	"gorm.io/gorm"
)

type PaymentRepository struct{}

func NewPaymentRepository() *PaymentRepository {
	return &PaymentRepository{}
}

func (PaymentRepository) TransferBalance(from, to string, amount uint) (bool, error) {
	err := config.DB.Transaction(func(tx *gorm.DB) error {
		var customer entity.Customer
		if err := tx.Where("account_id = ?", from).First(&customer).Error; err != nil {
			return err
		}

		if customer.Balance < uint64(amount) {
			return errors.New("insufficient balance")
		}

		var merchant entity.Merchant
		if err := tx.Where("id = ?", to).First(&merchant).Error; err != nil {
			return err
		}

		customer.Balance -= uint64(amount)
		if err := tx.Save(&customer).Error; err != nil {
			return err
		}

		merchant.Balance += uint64(amount)
		if err := tx.Save(&merchant).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return false, err
	}

	return true, nil
}
