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

func GenererDonneesCapteur() DonneesCapteur {
	var donneesGenerees DonneesCapteur
	donneesGenerees.Id = rand.Int()
	donneesGenerees.IdCapteur = rand.Int()
	donneesGenerees.IdAeroport = RandomAeroport()
	donneesGenerees.TypeMesure = RandomCapteur()
	donneesGenerees.ValeurMesure = rand.Float32()
	donneesGenerees.Date = time.Date(2020, 07, 07, 00,00,00, 00, time.UTC)

	return donneesGenerees
}

func RandomCapteur() string {
	var s [4]string
	s[0] = "WIND"
	s[1] = "TEMP"
	s[2] = "RAIN"
	return s[rand.Int()%3]
}

func RandomAeroport() string {
	var s [4]string
	s[0] = "CDG"
	s[1] = "AVR"
	s[2] = "AVW"
	s[3] = "MAK"
	return s[rand.Int()%4]
}