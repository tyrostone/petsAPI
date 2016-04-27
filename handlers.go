package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func petList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(pets); err != nil {
		panic(err)
	}
}

func petDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	petId, err := strconv.Atoi(vars["petId"])
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	pet := RepoGetPet(petId)
	if (pet == Pet{}) {
		petNotFound := GenerateMessage("Pet ID: " + vars["petId"] + " Not Found")
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

func addPet(w http.ResponseWriter, r *http.Request) {
	var pet Pet
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
	p := RepoCreatePet(pet)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(p); err != nil {
		panic(err)
	}
}

func deletePet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	petId, err := strconv.Atoi(vars["petId"])
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err := RepoDestroyPet(petId); err != nil {
		w.WriteHeader(http.StatusNotFound)
		petNotFound := GenerateMessage("Pet ID: " + vars["petId"] + " Not Found")
		w.WriteHeader(404)
		if err := json.NewEncoder(w).Encode(petNotFound); err != nil {
			panic(err)
		}
	} else {
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(GenerateMessage("Deletion successful")); err != nil {
			panic(err)
		}
	}
}
