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
	donneesGenerees.IdAeroport = "AAA"
	donneesGenerees.TypeMesure = "WIND"
	donneesGenerees.ValeurMesure = rand.Float32()
	donneesGenerees.Date = time.Date(2020, 07, 07, 00,00,00, 00, time.UTC)

	return donneesGenerees
}
