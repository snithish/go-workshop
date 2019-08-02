package repositories

import (
	"4_gin_web_service/models"
	"errors"
)

var pets = make(map[int]models.Pet)

var petExistsErr = errors.New("pet with given id already exists")

type PetRepository interface {
	Save(pet models.Pet) error
}

type petRepository struct {
}

func NewPetRepository() PetRepository {
	return petRepository{}
}

func (petRepository petRepository) Save(pet models.Pet) error {
	if petRepository.exists(pet.Id) {
		return petExistsErr
	}
	pets[pet.Id] = pet
	return nil
}

func (petRepository petRepository) exists(id int) bool {
	_, exists := pets[id]
	return exists
}
