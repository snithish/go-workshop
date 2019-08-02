package controllers

import (
	"4_gin_web_service/models"
	"4_gin_web_service/services"
	"github.com/gin-gonic/gin/binding"
	"github.com/sirupsen/logrus"
)

type PetController interface {
	CreatePet(ctx Context)
	UpdatePet(ctx Context)
}

type petController struct {
	petService services.PetService
}

func NewPetController(petService services.PetService) PetController {
	return &petController{
		petService: petService,
	}
}

func (ctrl petController) CreatePet(ctx Context) {
	var createPetRequest models.Pet
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
	serviceError := ctrl.petService.CreatePet(createPetRequest)
	if serviceError != nil {
		logrus.Error("Creating pet failed because " + serviceError.Error())
		SendConflict(ctx)
		return
	}
	SendRequestCreated(ctx)
}

func (ctrl petController) UpdatePet(ctx Context) {
	var createPetRequest models.Pet
	bindingError := ctx.ShouldBindBodyWith(&createPetRequest, binding.JSON)
	if bindingError != nil {
		logrus.Error("Update pet request object serialization failed because " + bindingError.Error())
		SendBadRequest(ctx)
		return
	}
	validationError := createPetRequest.Validate()
	if validationError != nil {
		logrus.Error("Request validation failed because " + validationError.Error())
		SendInvalidInput(ctx)
		return
	}
	serviceError := ctrl.petService.UpdatePet(createPetRequest)
	if serviceError != nil {
		logrus.Error("Updating pet failed because " + serviceError.Error())
		SendNotFound(ctx)
		return
	}
	SendRequestOK(ctx)
}
