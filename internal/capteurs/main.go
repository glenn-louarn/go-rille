package main

import (
	"encoding/json"
	"github.com/my/repo/internal/capteurs/model"
	"time"
)

const frequenceEmissionDuration time.Duration = 2
const frequenceEmissionInt int = 2
const tempsEmission int = 10

func main() {
	var configuration model.Configuration = model.GetConfig()

	nbTours := (tempsEmission * 60) / frequenceEmissionInt

	for i := nbTours ; i > 0 ; i-- {
		dateMesure := model.RandomDate()
		aeroportMesure := model.RandomAeroport()

		for j := 0 ; j < len(model.GetTypeMesureTableau()) ; j++ {
			donneesAleatoire := model.GenererDonneesCapteur(model.GetTypeMesureTableau()[j], dateMesure, aeroportMesure)
			val, _ := json.Marshal(donneesAleatoire)
			println(string(val))
			model.Send(configuration, donneesAleatoire)
		}
		time.Sleep(frequenceEmissionDuration * time.Second)
	}
}