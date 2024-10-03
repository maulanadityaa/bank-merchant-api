package services

import (
	"github.com/maulanadityaa/bank-merchant-api/models/dto/request"
	"github.com/maulanadityaa/bank-merchant-api/models/dto/response"
)

type MerchantService interface {
	AddMerchant(request request.UserRequest) (response.UserResponse, error)
	UpdateMerchant(request request.UserUpdateRequest) (response.UserResponse, error)
	GetMerchantByID(merchantID string) (response.UserResponse, error)
	GetMerchantByAccountID(accountID string) (response.UserResponse, error)
	GetAllMerchant(paging, rowsPerPage, name string) ([]response.UserResponse, string, string, error)
}
