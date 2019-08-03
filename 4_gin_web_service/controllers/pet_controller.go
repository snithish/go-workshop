package controllers

import (
	"4_gin_web_service/models"
	"4_gin_web_service/services"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type PetController interface {
	CreatePet(ctx *gin.Context)
	UpdatePet(ctx *gin.Context)
	DeletePet(ctx *gin.Context)
	GetPet(ctx *gin.Context)
}

type petController struct {
	petService services.PetService
}

func NewPetController(petService services.PetService) PetController {
	return &petController{
		petService: petService,
	}
}

func (ctrl petController) CreatePet(ctx *gin.Context) {
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
	ctrl.petService.CreatePet(createPetRequest)
	//if serviceError != nil {
	//	logrus.Error("Creating pet failed because " + serviceError.Error())
	//	SendConflict(ctx)
	//	return
	//}
	SendRequestCreated(ctx)
}

func (ctrl petController) UpdatePet(ctx *gin.Context) {
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

func (ctrl petController) DeletePet(ctx *gin.Context) {
	petIDString := ctx.Param("petID")
	if petIDString == "" {
		logrus.Error("Pet identifier not provided")
		SendBadRequest(ctx)
		return
	}
	petID, validationError := strconv.Atoi(petIDString)
	if validationError != nil {
		logrus.Error("invalid pet identifier provided")
		SendInvalidInput(ctx)
		return
	}
	serviceError := ctrl.petService.DeletePet(petID)
	if serviceError != nil {
		logrus.Error("Deleting pet failed because " + serviceError.Error())
		SendNotFound(ctx)
		return
	}
	SendRequestOK(ctx)
}

func (ctrl petController) GetPet(ctx *gin.Context) {
	petIDString := ctx.Param("petID")
	if petIDString == "" {
		logrus.Error("Pet identifier not provided")
		SendBadRequest(ctx)
		return
	}
	petID, validationError := strconv.Atoi(petIDString)
	if validationError != nil {
		logrus.Error("invalid pet identifier provided")
		SendInvalidInput(ctx)
		return
	}
	pet, serviceError := ctrl.petService.GetPet(petID)
	if serviceError != nil {
		logrus.Error("Fetching pet failed because " + serviceError.Error())
		SendNotFound(ctx)
		return
	}
	SendMessageWithStatus(ctx, http.StatusOK, pet)
}
