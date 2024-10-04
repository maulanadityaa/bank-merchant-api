package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/maulanadityaa/bank-merchant-api/middlewares"
	"github.com/maulanadityaa/bank-merchant-api/models/dto/request"
	"github.com/maulanadityaa/bank-merchant-api/models/dto/response"
	"github.com/maulanadityaa/bank-merchant-api/services"
	"github.com/maulanadityaa/bank-merchant-api/services/impl"
	"github.com/maulanadityaa/bank-merchant-api/validators"
)

type MerchantController struct{}

var merchantService services.MerchantService = impl.NewMerchantService()

func NewMerchantController(g *gin.RouterGroup) {
	controller := new(MerchantController)

	merchantGroup := g.Group("/merchants", middlewares.ValidateJWT())
	{
		merchantGroup.PUT("", middlewares.AuthWithRole([]string{"ROLE_MERCHANT"}), controller.UpdateMerchant)
		merchantGroup.GET("", controller.GetAllMerchant)
		merchantGroup.GET("/:id", middlewares.AuthWithRole([]string{"ROLE_MERCHANT"}), controller.GetMerchantByID)
		merchantGroup.GET("/account/:accountID", middlewares.AuthWithRole([]string{"ROLE_MERCHANT"}), controller.GetMerchantByAccountID)
	}
}

func (MerchantController) UpdateMerchant(c *gin.Context) {
	var request request.UserUpdateRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		response.NewResponseBadRequest(c, err.Error())
		return
	}

	errors := validators.ValidateStruct(request)
	if errors != nil {
		response.NewResponseValidationError(c, errors)
		return
	}

	result, err := merchantService.UpdateMerchant(request)
	if err != nil {
		response.NewResponseError(c, err.Error())
		return
	}

	response.NewResponseOK(c, result)
}

func (MerchantController) GetAllMerchant(c *gin.Context) {
	paging := c.DefaultQuery("page", "1")
	rowsPerPage := c.DefaultQuery("rowsPerPage", "10")
	name := c.DefaultQuery("name", "")

	result, totalRows, totalPage, err := merchantService.GetAllMerchant(paging, rowsPerPage, name)
	if err != nil {
		response.NewResponseError(c, err.Error())
		return
	}

	response.NewResponseOKWithPaging(c, result, paging, rowsPerPage, totalRows, totalPage)
}

func (MerchantController) GetMerchantByID(c *gin.Context) {
	id := c.Param("id")

	result, err := merchantService.GetMerchantByID(id)
	if err != nil {
		response.NewResponseError(c, err.Error())
		return
	}

	response.NewResponseOK(c, result)
}

func (MerchantController) GetMerchantByAccountID(c *gin.Context) {
	accountID := c.Param("accountID")

	result, err := merchantService.GetMerchantByAccountID(accountID)
	if err != nil {
		response.NewResponseError(c, err.Error())
		return
	}

	response.NewResponseOK(c, result)
}
