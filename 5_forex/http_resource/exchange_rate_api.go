package http_resource

import (
	"encoding/json"
	"fmt"
	"forex/errors"
	"forex/models"
	"io/ioutil"
	"net/http"
)

type ExchangeApi interface {
	GetRates(conversionRequest models.ConversionRequest) (float64, error)
}

type exchangeApi struct {
	httpClient *http.Client
}

func NewExchangeApi(httpClient *http.Client) ExchangeApi {
	return &exchangeApi{httpClient: httpClient}
}

func (e exchangeApi) GetRates(conversionRequest models.ConversionRequest) (float64, error) {
	url := fmt.Sprintf("https://api.exchangeratesapi.io/latest?base=%s&&symbols=%s,%s",
		conversionRequest.SourceCurrency,
		conversionRequest.SourceCurrency,
		conversionRequest.TargetCurrency)
	resp, err := e.httpClient.Get(url)
	if err != nil {
		return 0, err
	}
	if resp.StatusCode == 200 {
		responseBodyBytes, _ := ioutil.ReadAll(resp.Body)
		var response map[string]interface{}
		json.Unmarshal(responseBodyBytes, &response)
		rates := response["rates"].(map[string]float64)
		return rates[conversionRequest.TargetCurrency], nil
	}
	if resp.StatusCode == 200 {
		responseBodyBytes, _ := ioutil.ReadAll(resp.Body)
		var response map[string]interface{}
		json.Unmarshal(responseBodyBytes, &response)
		rates := response["rates"].(map[string]float64)
		return rates[conversionRequest.TargetCurrency], nil
	}
	if resp.StatusCode == 400 {
		return 0.0, errors.BadRequest{ErrMsg: "did not understand currency"}
	}
	return 0.0, errors.InternalError(fmt.Sprintf("Add specific handler for %d status code", resp.StatusCode))
}
