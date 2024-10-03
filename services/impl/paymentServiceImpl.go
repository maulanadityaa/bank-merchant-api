package impl

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/maulanadityaa/bank-merchant-api/models/dto/request"
	"github.com/maulanadityaa/bank-merchant-api/models/dto/response"
	"github.com/maulanadityaa/bank-merchant-api/utils"
)

type PaymentService struct{}

func NewPaymentService() *PaymentService {
	return &PaymentService{}
}

func (service *PaymentService) Pay(req request.PaymentRequest, c *gin.Context) (response.PaymentResponse, error) {
	claims := utils.GetJWTClaims(c)
	accountId := claims["accountId"].(string)

	customer, err := customerService.GetCustomerByAccountID(accountId)
	if err != nil {
		return response.PaymentResponse{}, err
	}

	merchant, err := merchantService.GetMerchantByID(req.To)
	if err != nil {
		return response.PaymentResponse{}, err
	}

	if customer.Balance < uint64(req.Amount) {
		return response.PaymentResponse{}, errors.New("insufficient balance")
	}

	newCustomerRequest := request.UserUpdateRequest{
		ID:      customer.ID,
		Name:    customer.Name,
		Balance: customer.Balance - uint64(req.Amount),
	}

	customerResponse, err := customerService.UpdateCustomer(newCustomerRequest)
	if err != nil {
		return response.PaymentResponse{}, err
	}

	newMerchantRequest := request.UserUpdateRequest{
		ID:      merchant.ID,
		Name:    merchant.Name,
		Balance: merchant.Balance + uint64(req.Amount),
	}

	merchantResponse, err := merchantService.UpdateMerchant(newMerchantRequest)
	if err != nil {
		return response.PaymentResponse{}, err
	}

	return response.PaymentResponse{
		From:      customerResponse,
		To:        merchantResponse,
		Amount:    req.Amount,
		CreatedAt: time.Now().String(),
	}, nil
}
