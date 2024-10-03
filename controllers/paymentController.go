package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/maulanadityaa/bank-merchant-api/middlewares"
	"github.com/maulanadityaa/bank-merchant-api/models/dto/request"
	"github.com/maulanadityaa/bank-merchant-api/models/dto/response"
	"github.com/maulanadityaa/bank-merchant-api/services"
	"github.com/maulanadityaa/bank-merchant-api/services/impl"
)

type PaymentController struct{}

var paymentService services.PaymentService = impl.NewPaymentService()

func NewPaymentController(g *gin.RouterGroup) {
	controller := new(PaymentController)

	paymentGroup := g.Group("/payment", middlewares.ValidateJWT())
	{
		paymentGroup.POST("/pay", middlewares.AuthWithRole([]string{"ROLE_CUSTOMER"}), controller.Pay)
	}
}

func (PaymentController) Pay(c *gin.Context) {
	var request request.PaymentRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		response.NewResponseBadRequest(c, err.Error())
		return
	}

	result, err := paymentService.Pay(request, c)
	if err != nil {
		response.NewResponseError(c, err.Error())
		return
	}

	response.NewResponseOK(c, result)
}
