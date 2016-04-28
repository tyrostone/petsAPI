package helpers

import (
	"fmt"
	"github.com/tyrostone/petsAPI/models"
)

var currentId int

var Pets models.Pets

func init() {
	RepoCreatePet(models.Pet{Name: "Tyler"})
	RepoCreatePet(models.Pet{Name: "Fluffy"})
}

func RepoCreatePet(pet models.Pet) models.Pet {
	currentId += 1
	pet.Id = currentId
	Pets = append(Pets, pet)
	return pet
}

func RepoGetPet(id int) models.Pet {
	for _, pet := range Pets {
		if pet.Id == id {
			return pet
		}
	}
	return models.Pet{}
}

func RepoDestroyPet(id int) error {
	for i, pet := range Pets {
		if pet.Id == id {
			Pets = append(Pets[:i], Pets[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Pet with Id %d to delete", id)
}
