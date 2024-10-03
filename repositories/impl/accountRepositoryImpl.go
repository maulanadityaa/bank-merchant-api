package impl

import (
	"github.com/maulanadityaa/bank-merchant-api/config"
	"github.com/maulanadityaa/bank-merchant-api/models/entity"
)

type AccountRepository struct{}

func NewAccountRepository() *AccountRepository {
	return &AccountRepository{}
}

func (AccountRepository) AddAccount(account entity.Account) (entity.Account, error) {
	if result := config.DB.Create(&account); result.Error != nil {
		return entity.Account{}, result.Error
	}

	return account, nil
}

func (AccountRepository) UpdateAccount(account entity.Account) (entity.Account, error) {
	if result := config.DB.Save(&account); result.Error != nil {
		return entity.Account{}, result.Error
	}

	return account, nil
}

func (AccountRepository) GetAccountByID(accountID string) (entity.Account, error) {
	var account entity.Account
	if result := config.DB.Where("id = ?", accountID).First(&account); result.Error != nil {
		return entity.Account{}, result.Error
	}

	return account, nil
}

func (AccountRepository) GetAccountByEmail(email string) (entity.Account, error) {
	var account entity.Account
	if result := config.DB.Where("email = ?", email).First(&account); result.Error != nil {
		return entity.Account{}, result.Error
	}

	return account, nil
}
