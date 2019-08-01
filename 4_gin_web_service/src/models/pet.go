package models

import (
	"4_gin_web_service/src/helpers"
	"errors"
)

var allowedStatuses = []string{"available","pending","sold"}

type CreatePetRequest struct {
	Id        int   `json:"id" binding:"required"`
	Name      string   `json:"name" binding:"required"`
	PhotoUrls []string `json:"photoUrls" binding:"dive,required"`
	Status    string   `json:"status" binding:"required"`
}

func (c CreatePetRequest) Validate() error {
	if helpers.ContainsString(allowedStatuses,c.Status) {
		return nil
	}
	return errors.New("status is not valid")
}