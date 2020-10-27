package main

import (
	"github.com/gorilla/mux"
	"github.com/my/repo/cmd/api/config"
	"github.com/my/repo/cmd/api/controller"
	"log"
	"net/http"
)

func main() {
	config.DatabaseInit()
	router := InitializeRouter()

	// Populate database
	//model.NewCar(&model.Car{Manufacturer: "citroen", Design: "ds3", Style: "sport", Doors: 4})

	log.Fatal(http.ListenAndServe(":8080", router))
}

func InitializeRouter() *mux.Router {
	// StrictSlash is true => redirect /donneees/ to /donneees
	router := mux.NewRouter().StrictSlash(true)

	router.Methods("GET").Path("/donnees").Name("Index").HandlerFunc(controller.DonneesCapteurIndex)
	router.Methods("POST").Path("/donnees").Name("Create").HandlerFunc(controller.DonneesCapteurCreate)
	router.Methods("GET").Path("/donnees/{id}").Name("Show").HandlerFunc(controller.DonneesCapteurShow)
	router.Methods("DELETE").Path("/donnees/{id}").Name("DELETE").HandlerFunc(controller.DonneesCapteurDelete)
	return router
}