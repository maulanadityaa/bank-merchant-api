package impl

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/maulanadityaa/bank-merchant-api/models/dto/request"
	"github.com/maulanadityaa/bank-merchant-api/models/dto/response"
	"github.com/maulanadityaa/bank-merchant-api/repositories"
	"github.com/maulanadityaa/bank-merchant-api/repositories/impl"
	"github.com/maulanadityaa/bank-merchant-api/utils"
)

type PaymentService struct{}

var paymentRepository repositories.PaymentRepository = impl.NewPaymentRepository()

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

	newHistoryRequest := request.HistoryRequest{
		CustomerID: utils.StringToPointer(customer.ID),
		MerchantID: utils.StringToPointer(merchant.ID),
		Amount:     req.Amount,
		Action:     "PAYMENT",
	}

	history, err := historyService.AddHistory(newHistoryRequest)
	if err != nil && !history {
		return response.PaymentResponse{}, err
	}

	_, err = paymentRepository.TransferBalance(accountId, req.To, req.Amount)
	if err != nil {
		return response.PaymentResponse{}, err
	}

	customerResponse, err := customerService.GetCustomerByAccountID(accountId)
	if err != nil {
		return response.PaymentResponse{}, err
	}

	merchantResponse, err := merchantService.GetMerchantByID(req.To)
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
