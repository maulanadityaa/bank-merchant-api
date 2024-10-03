package cmd

import (
	"fmt"
	"time"

	"github.com/maulanadityaa/bank-merchant-api/config"
)

func InitApp() {
	location, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		fmt.Println(err.Error())
	}
	time.Local = location

	config.LoadConfig()
	config.ConnectDB()
}
