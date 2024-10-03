package impl

import (
	"github.com/maulanadityaa/bank-merchant-api/config"
	"github.com/maulanadityaa/bank-merchant-api/models/entity"
	"github.com/maulanadityaa/bank-merchant-api/utils"
	"gorm.io/gorm"
)

type CustomerRepository struct{}

func NewCustomerRepository() *CustomerRepository {
	return &CustomerRepository{}
}

func (CustomerRepository) AddCustomer(customer entity.Customer) (entity.Customer, error) {
	if result := config.DB.Create(&customer); result.Error != nil {
		return entity.Customer{}, result.Error
	}

	return customer, nil
}

func (CustomerRepository) UpdateCustomer(customer entity.Customer) (entity.Customer, error) {
	if result := config.DB.Save(&customer); result.Error != nil {
		return entity.Customer{}, result.Error
	}

	return customer, nil
}

func (CustomerRepository) GetCustomerByID(customerID string) (entity.Customer, error) {
	var customer entity.Customer
	if result := config.DB.Where("id = ?", customerID).First(&customer); result.Error != nil {
		return entity.Customer{}, result.Error
	}

	return customer, nil
}

func (CustomerRepository) GetCustomerByAccountID(accountID string) (entity.Customer, error) {
	var customer entity.Customer
	if result := config.DB.Where("account_id = ?", accountID).First(&customer); result.Error != nil {
		return entity.Customer{}, result.Error
	}

	return customer, nil
}

func (CustomerRepository) GetAllCustomer(spec []func(db *gorm.DB) *gorm.DB, name string) ([]entity.Customer, string, error) {
	var customers []entity.Customer

	if name != "" {
		spec = append(spec, utils.SelectByName(name))
	}

	db := config.DB.Model(&entity.Customer{}).Scopes(spec[1:]...)
	totalRows := utils.GetTotalRows(db)
	err := db.Scopes(spec[0]).Find(&customers).Error

	return customers, totalRows, err
}
