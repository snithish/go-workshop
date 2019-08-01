package models

type ConversionRequest struct {
	SourceCurrency  string `json:"sourceCurrency" binding:"required"`
	TargetCurrency  string `json:"targetCurrency" binding:"required"`
	AmountToConvert int    `json:"amountToConvert" binding:"required"`
}
