package router

import (
	"github.com/gin-gonic/gin"
	"github.com/maulanadityaa/bank-merchant-api/controllers"
)

func InitRouter(route *gin.RouterGroup) {
	controllers.NewAuthController(route)
	controllers.NewPaymentController(route)
	controllers.NewCustomerController(route)
	controllers.NewMerchantController(route)
}
