package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/my/repo/cmd/api/model"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"time"
)

func DataBetweenDates(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")

	layout := "2006-01-02"
	vars := mux.Vars(r)
	idAirport := vars["idAirport"]
	sensor := vars["sensor"]
	date1, err := time.Parse(layout, vars["date1"])
	date2, err := time.Parse(layout, vars["date2"])
	if err != nil {
		fmt.Println(err)
	}
	resultToSend := make(map[string][]float32)
	days := date2.Sub(date1).Hours()/24
	if days == 0 {
		dateFormatted := time.Time.Format(date1, layout)
		resultToSend[dateFormatted] = model.GetSensorDataByDate(idAirport, dateFormatted,sensor)
	} else {
		for i := 0; i < int(math.Abs(days))+1; i++ {
			dateFormatted := time.Time.Format(date1, layout)
			resultToSend[dateFormatted] = model.GetSensorDataByDate(idAirport, dateFormatted,sensor)
			date1 = date1.AddDate(0, 0, 1)
		}
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resultToSend)
}
func DataCapteurAverageSensor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	vars := mux.Vars(r)
	idAirport, _ := vars["idAirport"]
	date, _ := vars["date"]
	averageRain := model.AverageSensorByAirport(idAirport,date,"RAIN")
	averageWind := model.AverageSensorByAirport(idAirport,date,"WIND")
	averageTemp := model.AverageSensorByAirport(idAirport,date,"TEMP")
	averageValues := model.AverageValues{Temperature: averageTemp,Rain: averageRain,Wind: averageWind}
	fmt.Printf("%+v",averageValues)
	err := json.NewEncoder(w).Encode(averageValues)
	if err !=nil{
		log.Fatal(err)
	}
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

	model.NewDonneesCapteur(&donnees)

	json.NewEncoder(w).Encode(donnees)
}