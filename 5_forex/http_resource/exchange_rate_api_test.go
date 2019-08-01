package http_resource

import (
	"bytes"
	"forex/generated_mocks"
	"forex/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestExchangeApi_ReturnsConversionValueWhenHTTPCallSucceeds(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish() // erase all mock setups

	mockHttpClient := generated_mocks.NewMockHttpClient(ctrl)

	exchangeApi := NewExchangeApi(mockHttpClient)
	conversionRequest := models.ConversionRequest{
		SourceCurrency:  "USD",
		TargetCurrency:  "GBP",
		AmountToConvert: 100,
	}

	responseBody := `{
    "rates": {
        "USD": 1.0,
        "GBP": 0.825858476
    },
    "base": "USD",
    "date": "2019-08-01"
}`
	responseBytes := []byte(responseBody)
	response := http.Response{StatusCode: 200, Body: ioutil.NopCloser(bytes.NewBuffer(responseBytes))}

	mockHttpClient.EXPECT().Get("https://api.exchangeratesapi.io/latest?base=USD&&symbols=USD,GBP").Return(&response, nil)

	rate, err := exchangeApi.GetRates(conversionRequest)

	assert.Equal(t, 0.825858476, rate)
	assert.NoError(t, err)
}

func TestExchangeApi_BadRequestWhenAPICallReturns400(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish() // erase all mock setups

	mockHttpClient := generated_mocks.NewMockHttpClient(ctrl)

	exchangeApi := NewExchangeApi(mockHttpClient)
	conversionRequest := models.ConversionRequest{
		SourceCurrency:  "USD",
		TargetCurrency:  "GBP",
		AmountToConvert: 100,
	}

	response := http.Response{StatusCode: 400}

	mockHttpClient.EXPECT().Get("https://api.exchangeratesapi.io/latest?base=USD&&symbols=USD,GBP").Return(&response, nil)

	_, err := exchangeApi.GetRates(conversionRequest)

	assert.NoError(t, err)
}
