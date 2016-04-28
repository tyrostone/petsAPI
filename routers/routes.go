package routers

import (
	"github.com/tyrostone/petsAPI/controllers"
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
		controllers.PetList,
	},
	Route{
		"CreatePet",
		"POST",
		GenerateRoute("/pets"),
		controllers.AddPet,
	},
	Route{
		"DestroyPet",
		"DELETE",
		GenerateRoute("/pets/{petId}"),
		controllers.DeletePet,
	},
	Route{
		"PetsDetail",
		"GET",
		GenerateRoute("/pets/{petId}"),
		controllers.PetDetail,
	},
	Route{
		"Login",
		"POST",
		GenerateRoute("/login"),
		controllers.Login,
	},
	Route{
		"Login",
		"POST",
		GenerateRoute("/login"),
		controllers.Logout,
	},
	Route{
		"RefreshToken",
		"GET",
		GenerateRoute("/refresh-token"),
		controllers.RefreshToken,
	},
}

func GenerateRoute(route string) string {
	return "/v1" + route
}
