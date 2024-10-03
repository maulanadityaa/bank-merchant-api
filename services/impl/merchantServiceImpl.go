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

type MerchantService struct{}

var merchantRepository repositories.MerchantRepository = impl.NewMerchantRepository()

func NewMerchantService() *MerchantService {
	return &MerchantService{}
}

func (MerchantService) AddMerchant(request request.UserRequest) (response.UserResponse, error) {
	newMerchant := entity.Merchant{
		ID:        uuid.NewString(),
		Name:      request.Name,
		Balance:   request.Balance,
		AccountID: request.AccountID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	merchant, err := merchantRepository.AddMerchant(newMerchant)
	if err != nil {
		return response.UserResponse{}, err
	}

	return response.UserResponse{
		ID:        merchant.ID,
		Name:      merchant.Name,
		Balance:   merchant.Balance,
		CreatedAt: merchant.CreatedAt.String(),
		UpdatedAt: merchant.UpdatedAt.String(),
	}, nil
}

func (MerchantService) UpdateMerchant(request request.UserUpdateRequest) (response.UserResponse, error) {
	merchant, err := merchantRepository.GetMerchantByID(request.ID)
	if err != nil {
		return response.UserResponse{}, err
	}

	merchant.Name = request.Name
	merchant.Balance = request.Balance
	merchant.UpdatedAt = time.Now()

	updatedMerchant, err := merchantRepository.UpdateMerchant(merchant)
	if err != nil {
		return response.UserResponse{}, err
	}

	return response.UserResponse{
		ID:        updatedMerchant.ID,
		Name:      updatedMerchant.Name,
		Balance:   updatedMerchant.Balance,
		CreatedAt: updatedMerchant.CreatedAt.String(),
		UpdatedAt: updatedMerchant.UpdatedAt.String(),
	}, nil
}

func (MerchantService) GetMerchantByID(merchantID string) (response.UserResponse, error) {
	merchant, err := merchantRepository.GetMerchantByID(merchantID)
	if err != nil {
		return response.UserResponse{}, err
	}

	return response.UserResponse{
		ID:        merchant.ID,
		Name:      merchant.Name,
		Balance:   merchant.Balance,
		CreatedAt: merchant.CreatedAt.String(),
		UpdatedAt: merchant.UpdatedAt.String(),
	}, nil
}

func (MerchantService) GetMerchantByAccountID(accountID string) (response.UserResponse, error) {
	merchant, err := merchantRepository.GetMerchantByAccountID(accountID)
	if err != nil {
		return response.UserResponse{}, err
	}

	return response.UserResponse{
		ID:        merchant.ID,
		Name:      merchant.Name,
		Balance:   merchant.Balance,
		CreatedAt: merchant.CreatedAt.String(),
		UpdatedAt: merchant.UpdatedAt.String(),
	}, nil
}

func (MerchantService) GetAllMerchant(paging, rowsPerPage, name string) ([]response.UserResponse, string, string, error) {
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

	merchants, totalRows, err := merchantRepository.GetAllMerchant(spec, name)
	if err != nil {
		return nil, "0", "0", err
	}

	merchantsResponses := make([]response.UserResponse, 0)
	for _, merchant := range merchants {
		merchantsResponses = append(merchantsResponses, response.UserResponse{
			ID:        merchant.ID,
			Name:      merchant.Name,
			Balance:   merchant.Balance,
			CreatedAt: merchant.CreatedAt.String(),
			UpdatedAt: merchant.UpdatedAt.String(),
		})
	}

	totalPage := utils.GetTotalPage(totalRows, rowsPerPageInt)

	return merchantsResponses, totalRows, strconv.Itoa(totalPage), nil
}
