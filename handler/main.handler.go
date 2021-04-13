package handler

import (
	"belajariah-main-service/model"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type mainHandler struct {
	config *model.Config
}

type MainHandler interface {
	NoRoute(ctx *gin.Context)
	Log(ctx *gin.Context)
}

func InitMainHandler(config *model.Config) MainHandler {
	return &mainHandler{
		config,
	}
}

func (mainHandler *mainHandler) NoRoute(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "BELAJARIAH MAIN SERVICE 0.0.3")
}

func (mainHandler *mainHandler) Log(ctx *gin.Context) {
	data, err := ioutil.ReadFile(mainHandler.config.Log.Filename)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"result": "ERROR READ LOG",
			"error":  err.Error(),
		})
	}

	loging := model.Logs{}
	err = json.Unmarshal(data, &loging)

	ctx.JSON(http.StatusOK, gin.H{
		"result": loging,
		"error":  "",
	})
}
