package model

import (
	"encoding/json"
	"github.com/garyburd/redigo/redis"
	"github.com/my/repo/cmd/api/config"
	"log"
	"math"
	"strconv"
	"strings"
	"time"
)
type AverageValues struct {
	Rain float32	`json:"rain"`
	Wind float32	`json:"wind"`
	Temperature float32	`json:"temperature"`
}
type DonneesCapteur struct {
	Id			 int		`json:"id"`
	IdCapteur    int 		`json:"idCapteur"`
	IdAeroport   string		`json:"idAeroport"`
	TypeMesure   string		`json:"typeMesure"`
	ValeurMesure float32	`json:"valeurMesure"`
	Date         time.Time	`json:"date"`
}

type DonneesCapteurArray []DonneesCapteur

func NewDonneesCapteur(d *DonneesCapteur) {
	if d == nil {
		log.Fatal(d)
	}
	data, _ := json.Marshal(d)
	_, err := config.Db().Do("SET",  "donnee:" + strconv.Itoa(d.Id), data)
	if err != nil {
		log.Fatal(err)
	}
}

func GetSensorDataByDate(idAeroport string,date string,sensor string) []float32 {
	keys, _ := redis.Strings(config.Db().Do("Keys","*"))
	var valuesToSend []float32
	for i:=0; i<len(keys);i++{
		key := keys[i]
		var data DonneesCapteur
		jsonVal, _ := redis.String(config.Db().Do("GET","donnee:"+strings.Split(key,":")[1]))
		json.Unmarshal([]byte(jsonVal),&data)
		if data.IdAeroport == idAeroport && data.Date.Format("2006-01-02")==date && sensor == data.TypeMesure{
			valuesToSend = append(valuesToSend,data.ValeurMesure)
		}
	}
	return valuesToSend
}

func AverageSensorByAirport(idAirport string, date string, sensor string) float32 {
	values := GetSensorDataByDate(idAirport, date, sensor)
	var total float32 = 0
	for _, value := range values {
		total += value
	}
	average := total / float32(len(values))
	if math.IsNaN(float64(average)){
		average = 0
	}
	return average
}


