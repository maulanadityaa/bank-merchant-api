package services

import (
	"github.com/gin-gonic/gin"
	"github.com/maulanadityaa/bank-merchant-api/models/dto/request"
	"github.com/maulanadityaa/bank-merchant-api/models/dto/response"
)

type PaymentService interface {
	Pay(req request.PaymentRequest, c *gin.Context) (response.PaymentResponse, error)
}
