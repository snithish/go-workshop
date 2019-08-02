package controllers

import (
	"forex/errors"
	"forex/http_resource"
	"forex/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/sirupsen/logrus"
)

type ConversionController interface {
	Convert(ctx *gin.Context)
	ConvertClosureBody(request interface{}, ctx *gin.Context)
}

type conversionController struct {
	exchangeApi http_resource.ExchangeApi
}

func NewConversionController(exchangeApi http_resource.ExchangeApi) ConversionController {
	return &conversionController{
		exchangeApi: exchangeApi,
	}
}

func (c conversionController) Convert(ctx *gin.Context) {
	var request models.ConversionRequest
	if err := ctx.ShouldBindBodyWith(&request, binding.JSON); err != nil {
		ctx.AbortWithStatus(400)
	}
	rate, err := c.exchangeApi.GetRates(request)
	logrus.Error("Error in getting exchange rate: ", err)
	switch err.(type) {
	case errors.BadRequest:
		ctx.AbortWithStatus(400)
	default:
		ctx.AbortWithStatus(500)
	}
	ctx.JSON(200, struct {
		value float64
	}{
		value: rate * float64(request.AmountToConvert),
	})
}

func (c conversionController) ConvertClosureBody(request interface{}, ctx *gin.Context) {
	panic("implement me")
}
