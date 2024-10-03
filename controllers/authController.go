package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/maulanadityaa/bank-merchant-api/models/dto/request"
	"github.com/maulanadityaa/bank-merchant-api/models/dto/response"
	"github.com/maulanadityaa/bank-merchant-api/services"
	"github.com/maulanadityaa/bank-merchant-api/services/impl"
)

type AuthController struct{}

var authService services.AuthService = impl.NewAuthService()

func NewAuthController(g *gin.RouterGroup) {
	controller := new(AuthController)

	authGroup := g.Group("/auth")
	{
		authGroup.POST("/register", controller.Register)
		authGroup.POST("/login", controller.Login)
		authGroup.POST("/logout", controller.Logout)
	}
}

func (AuthController) Register(c *gin.Context) {
	var request request.RegisterRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		response.NewResponseBadRequest(c, err.Error())
		return
	}

	result, err := authService.Register(request)
	if err != nil {
		response.NewResponseError(c, err.Error())
		return
	}

	response.NewResponseCreated(c, result)
}

func (AuthController) Login(c *gin.Context) {
	var request request.LoginRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		response.NewResponseBadRequest(c, err.Error())
		return
	}

	result, err := authService.Login(request)
	if err != nil {
		response.NewResponseError(c, err.Error())
		return
	}

	response.NewResponseOK(c, result)
}

func (AuthController) Logout(c *gin.Context) {
	result, err := authService.Logout(c)
	if err != nil {
		response.NewResponseError(c, err.Error())
		return
	}

	response.NewResponseOK(c, result)
}
