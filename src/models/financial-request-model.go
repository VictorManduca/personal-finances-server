package models

type FinancialRequestModel struct {
	Income  []float32 `json:"income"`
	Outcome []float32 `json:"outcome"`
}
