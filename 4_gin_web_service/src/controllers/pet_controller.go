package controllers

import (
	"4_gin_web_service/src/models"
	"github.com/gin-gonic/gin/binding"
	"github.com/sirupsen/logrus"
)

type PetController interface {
	CreatePet(ctx Context)
}

type petController struct {
}

func NewPetController() PetController {
	return &petController{}
}

func (ctrl *petController) CreatePet(ctx Context) {
	var createPetRequest models.CreatePetRequest
	bindingError := ctx.ShouldBindBodyWith(&createPetRequest, binding.JSON)
	if bindingError != nil {
		logrus.Error("Create pet request object serialization failed because " + bindingError.Error())
		SendBadRequest(ctx)
		return
	}
	validationError := createPetRequest.Validate()
	if validationError != nil {
		logrus.Error("Request validation failed because " + validationError.Error())
		SendInvalidInput(ctx)
		return
	}
	SendRequestOK(ctx)
}
