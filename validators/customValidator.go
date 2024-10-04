package validators

import (
	"github.com/go-playground/validator/v10"
	"github.com/maulanadityaa/bank-merchant-api/config"
	"github.com/maulanadityaa/bank-merchant-api/models/entity"
)

func UniqueEmail(fl validator.FieldLevel) bool {
	email := fl.Field().String()
	var existingEmail entity.Account

	result := config.DB.Where("email = ?", email).First(&existingEmail)

	return result.Error != nil
}

func PositiveAmount(fl validator.FieldLevel) bool {
	amount := fl.Field().Uint()
	return amount > 0
}
