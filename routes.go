package main

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"PetsList",
		"GET",
		"/pets",
		petList,
	},
	Route{
		"CreatePet",
		"POST",
		"/pets",
		addPet,
	},
	Route{
		"DestroyPet",
		"POST",
		"/pets/{petId}",
		deletePet,
	},
	Route{
		"PetsDetail",
		"GET",
		"/pets/{petId}",
		petDetail,
	},
}
