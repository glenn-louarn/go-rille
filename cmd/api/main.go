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
	log.Fatal(http.ListenAndServe(":8081", router))
}

func InitializeRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.Methods("GET").Path("/donnees/{idAirport}/{date1}/{date2}/{sensor}").Name("SensorDataDates").HandlerFunc(controller.DataBetweenDates)
	router.Methods("GET").Path("/donnees/{idAirport}/{date}").Name("AverageSensorData").HandlerFunc(controller.DataCapteurAverageSensor)
	router.Methods("POST").Path("/donnees").Name("Create").HandlerFunc(controller.DonneesCapteurCreate)

	return router
}