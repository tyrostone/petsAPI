package main

import (
	"fmt"
)

var currentId int

var pets Pets

func init() {
	RepoCreatePet(Pet{Name: "Tyler"})
	RepoCreatePet(Pet{Name: "Fluffy"})
}

func RepoCreatePet(pet Pet) Pet {
	currentId += 1
	pet.Id = currentId
	pets = append(pets, pet)
	return pet
}

func RepoGetPet(id int) Pet {
	for _, pet := range pets {
		if pet.Id == id {
			return pet
		}
	}
	return Pet{}
}

func RepoDestroyPet(id int) error {
	for i, pet := range pets {
		if pet.Id == id {
			pets = append(pets[:i], pets[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Pet with Id %d to delete", id)
}
