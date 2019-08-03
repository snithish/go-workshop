package services

import (
	"4_gin_web_service/models"
	"4_gin_web_service/repositories"
)

type PetService interface {
	CreatePet(request models.Pet) error
	UpdatePet(request models.Pet) error
	DeletePet(petID int) error
	GetPet(petID int) (models.Pet, error)
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

func (petService petService) UpdatePet(request models.Pet) error {
	updateError := petService.petRepository.Update(request)
	if updateError != nil {
		return updateError
	}
	return nil
}

func (petService petService) DeletePet(petID int) error {
	if err := petService.petRepository.Delete(petID); err != nil {
		return petService.petRepository.Delete(petID)
	} else {
		return nil
	}
}

func (petService petService) GetPet(petID int) (models.Pet, error) {
	pet, deleteError := petService.petRepository.Get(petID)
	if deleteError != nil {
		return models.Pet{}, deleteError
	}
	return pet, nil
}
