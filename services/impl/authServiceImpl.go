package impl

import (
	"errors"

	"github.com/google/uuid"
	"github.com/maulanadityaa/bank-merchant-api/models/dto/request"
	"github.com/maulanadityaa/bank-merchant-api/models/dto/response"
	"github.com/maulanadityaa/bank-merchant-api/models/entity"
	"github.com/maulanadityaa/bank-merchant-api/repositories"
	"github.com/maulanadityaa/bank-merchant-api/repositories/impl"
	"github.com/maulanadityaa/bank-merchant-api/services"
	"github.com/maulanadityaa/bank-merchant-api/utils"
)

type AuthService struct{}

var accountRepository repositories.AccountRepository = impl.NewAccountRepository()
var roleRepository repositories.RoleRepository = impl.NewRoleRepository()
var customerService services.CustomerService = NewCustomerService()
var merchantService services.MerchantService = NewMerchantService()

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (AuthService) Register(req request.RegisterRequest) (response.RegisterResponse, error) {
	role, _ := roleRepository.GetRoleByID(req.RoleID)
	hashedPassword, _ := utils.HashPassword(req.Password)

	newAccount := entity.Account{
		ID:       uuid.NewString(),
		Email:    req.Email,
		Password: hashedPassword,
		RoleID:   role.ID,
		IsLogged: false,
	}

	newUserRequest := request.UserRequest{
		Name:      req.Name,
		Balance:   req.Balance,
		AccountID: newAccount.ID,
	}

	account, err := accountRepository.AddAccount(newAccount)
	if err != nil {
		return response.RegisterResponse{}, err
	}

	var userResponse response.UserResponse

	if role.Name == "ROLE_CUSTOMER" {
		userResponse, _ = customerService.AddCustomer(newUserRequest)
	} else if role.Name == "ROLE_MERCHANT" {
		userResponse, _ = merchantService.AddMerchant(newUserRequest)
	} else {
		return response.RegisterResponse{}, errors.New("role not found")
	}

	return response.RegisterResponse{
		Email: account.Email,
		Role: response.RoleResponse{
			ID:   role.ID,
			Name: role.Name,
		},
		UserResponse: userResponse,
	}, nil
}

func (AuthService) Login(req request.LoginRequest) (response.LoginResponse, error) {
	account, err := accountRepository.GetAccountByEmail(req.Email)
	if err != nil {
		return response.LoginResponse{}, errors.New("invalid email or password")
	}

	err = utils.ComparePassword(account.Password, req.Password)
	if err != nil {
		return response.LoginResponse{}, errors.New("invalid email or password")
	}

	role, _ := roleRepository.GetRoleByID(account.RoleID)

	account.IsLogged = true
	account, err = accountRepository.UpdateAccount(account)
	if err != nil {
		return response.LoginResponse{}, err
	}

	token, err := utils.GenerateJWT(account.ID, role.Name, account.Email)
	if err != nil {
		return response.LoginResponse{}, err
	}

	return response.LoginResponse{
		Token: token,
	}, nil
}
