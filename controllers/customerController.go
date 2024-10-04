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

type CustomerController struct{}

var customerService services.CustomerService = impl.NewCustomerService()

func NewCustomerController(g *gin.RouterGroup) {
	controller := new(CustomerController)

	customerGroup := g.Group("/customers", middlewares.ValidateJWT())
	{
		customerGroup.PUT("", middlewares.AuthWithRole([]string{"ROLE_CUSTOMER"}), controller.UpdateCustomer)
		customerGroup.GET("", controller.GetAllCustomer)
		customerGroup.GET("/:id", middlewares.AuthWithRole([]string{"ROLE_CUSTOMER"}), controller.GetCustomerByID)
		customerGroup.GET("/account/:accountID", middlewares.AuthWithRole([]string{"ROLE_CUSTOMER"}), controller.GetCustomerByAccountID)
	}
}

func (CustomerController) UpdateCustomer(c *gin.Context) {
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

	result, err := customerService.UpdateCustomer(request)
	if err != nil {
		response.NewResponseError(c, err.Error())
		return
	}

	response.NewResponseOK(c, result)
}

func (CustomerController) GetAllCustomer(c *gin.Context) {
	paging := c.DefaultQuery("page", "1")
	rowsPerPage := c.DefaultQuery("rowsPerPage", "10")
	name := c.DefaultQuery("name", "")

	result, totalRows, totalPage, err := customerService.GetAllCustomer(paging, rowsPerPage, name)
	if err != nil {
		response.NewResponseError(c, err.Error())
		return
	}

	response.NewResponseOKWithPaging(c, result, paging, rowsPerPage, totalRows, totalPage)
}

func (CustomerController) GetCustomerByID(c *gin.Context) {
	id := c.Param("id")

	result, err := customerService.GetCustomerByID(id)
	if err != nil {
		response.NewResponseError(c, err.Error())
		return
	}

	response.NewResponseOK(c, result)
}

func (CustomerController) GetCustomerByAccountID(c *gin.Context) {
	accountID := c.Param("accountID")

	result, err := customerService.GetCustomerByAccountID(accountID)
	if err != nil {
		response.NewResponseError(c, err.Error())
		return
	}

	response.NewResponseOK(c, result)
}
