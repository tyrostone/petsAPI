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
		GenerateRoute("/pets"),
		petList,
	},
	Route{
		"CreatePet",
		"POST",
		GenerateRoute("/pets"),
		addPet,
	},
	Route{
		"DestroyPet",
		"POST",
		GenerateRoute("/pets/{petId}"),
		deletePet,
	},
	Route{
		"PetsDetail",
		"GET",
		GenerateRoute("/v1/pets/{petId}"),
		petDetail,
	},
}

func GenerateRoute(route string) string {
	return "/v1" + route
}
