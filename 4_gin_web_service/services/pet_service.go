package services

import (
	"4_gin_web_service/models"
	"4_gin_web_service/repositories"
)

type PetService interface {
	CreatePet(request models.Pet) error
}

type petService struct {
	petRepository repositories.PetRepository
}

func NewPetService(petRepository repositories.PetRepository) PetService {
	return petService{
		petRepository: petRepository,
	}
}

func (petService petService) CreatePet(request models.Pet) error {
	persistError := petService.petRepository.Save(request)
	if persistError != nil {
		return persistError
	}
	return nil
}
