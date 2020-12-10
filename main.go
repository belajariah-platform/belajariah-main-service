package main

import (
	"belajariah-main-service/config"
	"belajariah-main-service/handler"
	"belajariah-main-service/middleware"
	"belajariah-main-service/repository"
	"belajariah-main-service/usecase"
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

	//initiate repository
	userRepository := repository.InitUserRepository(db)

	//initiate usecase
	userUsecase := usecase.InitUserUsecase(userRepository)

	//initiate handler
	userHandler := handler.InitUserHandler(userUsecase)

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
	router.Use(middleware.CORSMiddleware())

	// routeruser
	fmt.Println(userHandler)
	// router.GET("/vessels", userHandler)

	utils.PushLogf("BELAJARIAH MAIN SERVICE STARTED")
	fmt.Println(fmt.Sprintf("BELAJARIAH MAIN SERVICE STARTED ON PORT %d", configModel.Server.Port))

	router.Run(fmt.Sprintf(":%v", configModel.Server.Port))
}
