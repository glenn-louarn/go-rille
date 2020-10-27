package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/my/repo/cmd/api/model"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func DonneesCapteurIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(model.AllDonneesCapteur())
}

func DonneesCapteurCreate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err)
	}

	var donnees model.DonneesCapteur

	err = json.Unmarshal(body, &donnees)

	if err != nil {
		log.Fatal(err)
	}

	println("TESTTESTTEST")
	a, err := json.Marshal(donnees)
	println(string(a))

	model.NewDonneesCapteur(&donnees)

	json.NewEncoder(w).Encode(donnees)
}

func DonneesCapteurShow(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		log.Fatal(err)
	}

	donnees := model.FindDonneesCapteurById(id)

	json.NewEncoder(w).Encode(donnees)
}

func DonneesCapteurDelete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)

	// strconv.Atoi is shorthand for ParseInt
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		log.Fatal(err)
	}

	err = model.DeleteDonneesCapteurById(id)
}
