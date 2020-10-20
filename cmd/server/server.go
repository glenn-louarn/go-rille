package main

import (
	"github.com/rs/cors"
	"go-rille/internal/json"
	"net/http"
)

type aeroport struct {
	Id              string
	AeroportName    string
	AeroportInitial string
	Kmh             int
	Temperature     int
	Pluie           int
}

func main() {
	mux := http.NewServeMux()
	aeroport1 := aeroport{
		Id:              "15",
		AeroportName:    "Charle de gaule",
		AeroportInitial: "CDG",
		Kmh:             20,
		Temperature:     20,
		Pluie:           20,
	}
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		json.WriteJSON(w, aeroport1)
	})
	handler := cors.Default().Handler(mux)
	http.ListenAndServe(":8081", handler)
}
