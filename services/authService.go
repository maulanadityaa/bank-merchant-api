package services

import (
	"github.com/gin-gonic/gin"
	"github.com/maulanadityaa/bank-merchant-api/models/dto/request"
	"github.com/maulanadityaa/bank-merchant-api/models/dto/response"
)

type AuthService interface {
	Register(req request.RegisterRequest) (response.RegisterResponse, error)
	Login(req request.LoginRequest) (response.LoginResponse, error)
	Logout(c *gin.Context) (response.LogoutResponse, error)
}
