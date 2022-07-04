package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	models "github.com/victormanduca/personal-finances/src/models"
	services "github.com/victormanduca/personal-finances/src/services"
)

func Calc(c *gin.Context) {
	var body models.FinancialRequestModel

	if err := c.BindJSON(&body); err != nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	var sum float32 = services.Caculator(body)
	if sum > 0 {
		result := make(map[string]string)
		result["message"] = fmt.Sprintf("Your profit: %.2f", sum)

		c.IndentedJSON(http.StatusOK, result)
		return
	} else if sum < 0 {
		result := make(map[string]string)
		result["message"] = fmt.Sprintf("Your deficit: %.2f", sum)

		c.IndentedJSON(http.StatusOK, result)
		return
	} else {
		result := make(map[string]string)
		result["message"] = fmt.Sprintf("Your balance: %.2f", sum)

		c.IndentedJSON(http.StatusOK, result)
		return
	}
}
