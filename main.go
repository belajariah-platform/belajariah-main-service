package main

import (
	"belajariah-main-service/config"
	"belajariah-main-service/utils"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("BELAJARIAH MAIN SERVICE INITIALIZATION")

	//create gin log
	crt, _ := os.Create("gin-port.log")
	gin.DefaultWriter = io.MultiWriter(crt)

	//get config
	configModel := config.GetConfig()

	//get db config
	db := config.ConnectDB(configModel)
	fmt.Println(db)

	//initiate router
	router := gin.New()
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	router.Use()

	router.Use(gin.Recovery())
	utils.PushLogf("BELAJARIAH MAIN SERVICE STARTED")
	fmt.Println(fmt.Sprintf("BELAJARIAH MAIN SERVICE STARTED ON PORT %d", configModel.Server.Port))

	router.Run(fmt.Sprintf(":%v", configModel.Server.Port))
}
