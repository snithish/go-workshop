package controllers

import (
	"forex/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type ConversionController interface {
	Convert(ctx *gin.Context)
	ConvertClosureBody(request interface{}, ctx *gin.Context)
}

type conversionController struct {
}

func NewConversionController() ConversionController {
	return &conversionController{}
}

func (c conversionController) Convert(ctx *gin.Context) {
	var request models.ConversionRequest
	if err := ctx.ShouldBindBodyWith(&request, binding.JSON); err != nil {
		ctx.AbortWithStatus(400)
	}
	panic("implement me")
}

func (c conversionController) ConvertClosureBody(request interface{}, ctx *gin.Context) {
	panic("implement me")
}
