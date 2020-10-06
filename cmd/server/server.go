package main

import (
	"github.com/swaggo/swag/example/basic/web"
	"go-rille/internal/config"
	"go-rille/internal/json"
	"log"
	"net/http"
)

type person struct {
	Age    int
	Name   string
	Prenom string
}

func tempHandler(w http.ResponseWriter, r *http.Request) {
	p := person{Age: 12, Prenom: "Glenn", Name: "Louarn"}
	json.WriteJSON(w, p)
}

func main() {
	http.HandleFunc("/temp", tempHandler)
	port := ":" + config.GetEnv("PORT")
	log.Fatal(http.ListenAndServe(port, nil))
}
