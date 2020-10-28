package model

import (
	"math/rand"
	"time"
)

type DonneesCapteur struct {
	Id			 int		`json:"id"`
	IdCapteur    int 		`json:"idCapteur"`
	IdAeroport   string		`json:"idAeroport"`
	TypeMesure   string		`json:"typeMesure"`
	ValeurMesure float32	`json:"valeurMesure"`
	Date         time.Time	`json:"date"`
}

func GenererDonneesCapteur(typeMesure string, dateMesure time.Time, idAeroport string) DonneesCapteur {
	var donneesGenerees DonneesCapteur
	donneesGenerees.Id = rand.Int()
	donneesGenerees.IdCapteur = rand.Int()
	donneesGenerees.IdAeroport = idAeroport
	donneesGenerees.TypeMesure = typeMesure
	donneesGenerees.ValeurMesure = rand.Float32()
	donneesGenerees.Date = dateMesure

	return donneesGenerees
}

func GetTypeMesureTableau() [3]string {
	var typesMesures [3]string
	typesMesures[0] = "WIND"
	typesMesures[1] = "TEMP"
	typesMesures[2] = "RAIN"
	return typesMesures
}

func RandomAeroport() string {
	var aeroports [4]string
	aeroports[0] = "CDG"
	aeroports[1] = "AVR"
	aeroports[2] = "AVW"
	aeroports[3] = "MAK"
	return aeroports[rand.Int()%4]
}

func RandomDate() time.Time {
	min := time.Date(2020, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Date(2021, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	delta := max - min

	sec := rand.Int63n(delta) + min
	return time.Unix(sec, 0)
}