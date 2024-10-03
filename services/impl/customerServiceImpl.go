package impl

import (
	"errors"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/maulanadityaa/bank-merchant-api/models/dto/request"
	"github.com/maulanadityaa/bank-merchant-api/models/dto/response"
	"github.com/maulanadityaa/bank-merchant-api/models/entity"
	"github.com/maulanadityaa/bank-merchant-api/repositories"
	"github.com/maulanadityaa/bank-merchant-api/repositories/impl"
	"github.com/maulanadityaa/bank-merchant-api/utils"
	"gorm.io/gorm"
)

type CustomerService struct{}

var customerRepository repositories.CustomerRepository = impl.NewCustomerRepository()

func NewCustomerService() *CustomerService {
	return &CustomerService{}
}

func (CustomerService) AddCustomer(request request.UserRequest) (response.UserResponse, error) {
	newCustomer := entity.Customer{
		ID:        uuid.NewString(),
		Name:      request.Name,
		Balance:   request.Balance,
		AccountID: request.AccountID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	customer, err := customerRepository.AddCustomer(newCustomer)
	if err != nil {
		return response.UserResponse{}, err
	}

	return response.UserResponse{
		ID:        customer.ID,
		Name:      customer.Name,
		Balance:   customer.Balance,
		CreatedAt: customer.CreatedAt.String(),
		UpdatedAt: customer.UpdatedAt.String(),
	}, nil
}

func (CustomerService) UpdateCustomer(request request.UserUpdateRequest) (response.UserResponse, error) {
	customer, err := customerRepository.GetCustomerByID(request.ID)
	if err != nil {
		return response.UserResponse{}, err
	}

	customer.Name = request.Name
	customer.Balance = request.Balance
	customer.UpdatedAt = time.Now()

	updatedCustomer, err := customerRepository.UpdateCustomer(customer)
	if err != nil {
		return response.UserResponse{}, err
	}

	return response.UserResponse{
		ID:        updatedCustomer.ID,
		Name:      updatedCustomer.Name,
		Balance:   updatedCustomer.Balance,
		CreatedAt: updatedCustomer.CreatedAt.String(),
		UpdatedAt: updatedCustomer.UpdatedAt.String(),
	}, nil
}

func (CustomerService) GetCustomerByID(customerID string) (response.UserResponse, error) {
	customer, err := customerRepository.GetCustomerByID(customerID)
	if err != nil {
		return response.UserResponse{}, err
	}

	return response.UserResponse{
		ID:        customer.ID,
		Name:      customer.Name,
		Balance:   customer.Balance,
		CreatedAt: customer.CreatedAt.String(),
		UpdatedAt: customer.UpdatedAt.String(),
	}, nil
}

func (CustomerService) GetCustomerByAccountID(accountID string) (response.UserResponse, error) {
	customer, err := customerRepository.GetCustomerByAccountID(accountID)
	if err != nil {
		return response.UserResponse{}, err
	}

	return response.UserResponse{
		ID:        customer.ID,
		Name:      customer.Name,
		Balance:   customer.Balance,
		CreatedAt: customer.CreatedAt.String(),
		UpdatedAt: customer.UpdatedAt.String(),
	}, nil
}

func (CustomerService) GetAllCustomer(paging, rowsPerPage, name string) ([]response.UserResponse, string, string, error) {
	pagingInt, err := strconv.Atoi(paging)
	if err != nil {
		return nil, "0", "0", errors.New("invalid query parameter")
	}

	rowsPerPageInt, err := strconv.Atoi(rowsPerPage)
	if err != nil {
		return nil, "0", "0", errors.New("invalid query parameter")
	}

	var spec []func(db *gorm.DB) *gorm.DB
	spec = append(spec, utils.Paginate(pagingInt, rowsPerPageInt))

	customers, totalRows, err := customerRepository.GetAllCustomer(spec, name)
	if err != nil {
		return nil, "0", "0", err
	}

	customerResponses := make([]response.UserResponse, 0)
	for _, customer := range customers {
		customerResponses = append(customerResponses, response.UserResponse{
			ID:        customer.ID,
			Name:      customer.Name,
			Balance:   customer.Balance,
			CreatedAt: customer.CreatedAt.String(),
			UpdatedAt: customer.UpdatedAt.String(),
		})
	}

	totalPage := utils.GetTotalPage(totalRows, rowsPerPageInt)

	return customerResponses, totalRows, strconv.Itoa(totalPage), nil
}
