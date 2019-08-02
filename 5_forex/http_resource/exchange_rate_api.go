package http_resource

import (
	"encoding/json"
	"fmt"
	"forex/errors"
	"forex/models"
	"io/ioutil"
	"net/http/httptest"
)

type ExchangeApi struct {
	httpClient HttpClient
}

func NewExchangeApi(httpClient HttpClient) ExchangeApi {
	return ExchangeApi{httpClient: httpClient}
}

func (e ExchangeApi) GetRates(conversionRequest models.ConversionRequest) (float64, error) {
	httptest.NewRecorder()
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
		rates := response["rates"].(map[string]interface{})
		return rates[conversionRequest.TargetCurrency].(float64), nil
	}
	if resp.StatusCode == 400 {
		return 0.0, &errors.BadRequest{ErrMsg: "did not understand currency"}
	}
	return 0.0, errors.InternalError(fmt.Sprintf("Add specific handler for %d status code", resp.StatusCode))
}
