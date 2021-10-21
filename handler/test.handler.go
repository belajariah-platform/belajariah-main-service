package handler

import (
	"belajariah-main-service/model"
	"belajariah-main-service/shape"
	"belajariah-main-service/usecase"
	"belajariah-main-service/utils"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type testHandler struct {
	testUsecase usecase.TestUsecase
}

type TestHandler interface {
	GetAllTest(ctx *gin.Context)
	CorrectionTest(ctx *gin.Context)
}

func InitTestHandler(testUsecase usecase.TestUsecase) TestHandler {
	return &testHandler{
		testUsecase,
	}
}

func (testHandler *testHandler) GetAllTest(ctx *gin.Context) {
	var query model.Query
	var count int
	err := ctx.BindQuery(&query)

	if err == nil {
		var array []map[string]interface{}
		if err := json.Unmarshal([]byte(query.Filter), &array); err != nil {
			panic(err)
		}
		for _, arr := range array {
			query.Filters = append(query.Filters, model.Filter{
				Type:  arr["type"].(string),
				Field: arr["field"].(string),
				Value: arr["value"].(string),
			})
		}

		var testResult []shape.ClassTest
		testResult, count, err = testHandler.testUsecase.GetAllTest(query)
		if err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"data":  testResult,
				"count": count,
				"error": "",
			})
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"data":  testResult,
				"count": count,
				"error": err.Error(),
			})
		}

	} else {
		utils.PushLogf("err", err)
	}
}

func (testHandler *testHandler) CorrectionTest(ctx *gin.Context) {
	var test shape.ClassTestPost
	var email string
	for _, value := range ctx.Request.Header["Email"] {
		email = value
		break
	}
	if err := ctx.ShouldBindJSON(&test); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	result, msg, err := testHandler.testUsecase.CorrectionTest(test, email)
	if err == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"result":  result,
			"message": msg,
			"error":   "",
		})
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"result":  result,
			"message": msg,
			"error":   err.Error(),
		})
	}
}
