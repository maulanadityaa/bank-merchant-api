package cmd

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/maulanadityaa/bank-merchant-api/config"
	"github.com/maulanadityaa/bank-merchant-api/router"
	"github.com/maulanadityaa/bank-merchant-api/validators"
)

func initDomainModule(r *gin.Engine) {
	apiGroup := r.Group("/api")
	v1Group := apiGroup.Group("/v1")

	router.InitRouter(v1Group)
}

func InitApp() {
	r := gin.Default()

	location, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		fmt.Println(err.Error())
	}
	time.Local = location

	config.LoadConfig()
	config.ConnectDB()

	validators.InitValidator()

	initDomainModule(r)

	addr := flag.String("port", ":"+os.Getenv("PORT"), "Address to listen and serve")
	if err := r.Run(*addr); err != nil {
		fmt.Println(err.Error())
	}
}
