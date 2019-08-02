package models

import (
	"4_gin_web_service/helpers"
	"errors"
)

var allowedStatuses = []string{"available", "pending", "sold"}

type Pet struct {
	Id        int      `json:"id" binding:"required"`
	Name      string   `json:"name" binding:"required"`
	PhotoUrls []string `json:"photoUrls" binding:"gt=0"`
	Status    string   `json:"status" binding:"required"`
}

//TODO: Move this to a tag based validation https://github.com/gin-gonic/gin#custom-validators
func (c Pet) Validate() error {
	if helpers.ContainsString(allowedStatuses, c.Status) {
		return nil
	}
	return errors.New("status is not valid")
}
