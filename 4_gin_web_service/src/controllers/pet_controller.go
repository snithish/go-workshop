package controllers

import (
	"net/http"
)

type PetController interface {
	CreatePet(ctx Context)
}

type petController struct {
}

func NewPetController() PetController {
	return &petController{
	}
}

func (ctrl *petController) CreatePet(ctx Context) {
	ctx.JSON(http.StatusOK, &struct {
		Message string `json:"message"`
	}{
		Message: "hello",
	})
}
