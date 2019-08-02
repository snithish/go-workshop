package repositories

import (
	"4_gin_web_service/models"
	"errors"
)

var (
	ErrPetExists       = errors.New("pet with given id already exists")
	ErrPetDoesNotExist = errors.New("pet with given id already exists")
)

type PetRepository interface {
	Save(pet models.Pet) error
	Update(pet models.Pet) error
	Delete(petID int) error
	Get(petID int) (models.Pet, error)
}

type petRepository struct {
	pets map[int]models.Pet
}

func NewPetRepository() PetRepository {
	return petRepository{
		pets: make(map[int]models.Pet),
	}
}

func (petRepository petRepository) Save(pet models.Pet) error {
	if petRepository.exists(pet.Id) {
		return ErrPetExists
	}
	petRepository.pets[pet.Id] = pet
	return nil
}

func (petRepository petRepository) Update(pet models.Pet) error {
	if !petRepository.exists(pet.Id) {
		return ErrPetDoesNotExist
	}
	petRepository.pets[pet.Id] = pet
	return nil
}

func (petRepository petRepository) Delete(petID int) error {
	if !petRepository.exists(petID) {
		return ErrPetDoesNotExist
	}
	delete(petRepository.pets, petID)
	return nil
}

func (petRepository petRepository) Get(petID int) (models.Pet, error) {
	pet, exists := petRepository.pets[petID]
	if !exists {
		return models.Pet{}, ErrPetDoesNotExist
	}
	return pet, nil
}

func (petRepository petRepository) exists(id int) bool {
	_, exists := petRepository.pets[id]
	return exists
}
