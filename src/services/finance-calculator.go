package services

import (
	models "github.com/victormanduca/personal-finances/src/models"
)

func Caculator(balance models.FinancialRequestModel) float32 {
	var inc float32 = 0
	var out float32 = 0

	for i := 0; i < len(balance.Income); i++ {
		inc += balance.Income[i]
	}

	for i := 0; i < len(balance.Outcome); i++ {
		out += balance.Outcome[i]
	}

	return float32(inc - out)
}
