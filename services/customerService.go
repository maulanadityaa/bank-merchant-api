package services

import (
	"github.com/maulanadityaa/bank-merchant-api/models/dto/request"
	"github.com/maulanadityaa/bank-merchant-api/models/dto/response"
)

type CustomerService interface {
	AddCustomer(request request.UserRequest) (response.UserResponse, error)
	UpdateCustomer(request request.UserUpdateRequest) (response.UserResponse, error)
	GetCustomerByID(customerID string) (response.UserResponse, error)
	GetCustomerByAccountID(accountID string) (response.UserResponse, error)
	GetAllCustomer(paging, rowsPerPage, name string) ([]response.UserResponse, string, string, error)
}
