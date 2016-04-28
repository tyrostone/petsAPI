package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/tyrostone/petsAPI/helpers"
	"github.com/tyrostone/petsAPI/models"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
)

func Login(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var creds models.Credentials
	err := decoder.Decode(&creds)
	if err != nil {
		panic(err)
	}
	if creds.Username == "" || creds.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(models.GenerateMessage("Please provide a username and password")); err != nil {
			panic(err)
		}
	} else {
		if err := json.NewEncoder(w).Encode(creds); err != nil {
			panic(err)
		}
	}
}

func RefreshToken(w http.ResponseWriter, r *http.Request) {}

func Logout(w http.ResponseWriter, r *http.Request) {}

func PetList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(helpers.Pets); err != nil {
		panic(err)
	}
}

func PetDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	petId, err := strconv.Atoi(vars["petId"])
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	pet := helpers.RepoGetPet(petId)
	if (pet == models.Pet{}) {
		petNotFound := models.GenerateMessage("Pet ID: " + vars["petId"] + " Not Found")
		w.WriteHeader(404)
		if err := json.NewEncoder(w).Encode(petNotFound); err != nil {
			panic(err)
		}
	} else {
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(pet); err != nil {
			panic(err)
		}
	}

}

func AddPet(w http.ResponseWriter, r *http.Request) {
	var pet models.Pet
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &pet); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
	p := helpers.RepoCreatePet(pet)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(p); err != nil {
		panic(err)
	}
}

func DeletePet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	petId, err := strconv.Atoi(vars["petId"])
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err := helpers.RepoDestroyPet(petId); err != nil {
		w.WriteHeader(http.StatusNotFound)
		petNotFound := models.GenerateMessage("Pet ID: " + vars["petId"] + " Not Found")
		if err := json.NewEncoder(w).Encode(petNotFound); err != nil {
			panic(err)
		}
	} else {
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(models.GenerateMessage("Deletion successful")); err != nil {
			panic(err)
		}
	}
}
