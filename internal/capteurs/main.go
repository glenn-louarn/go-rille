package main

import (
	"github.com/my/repo/internal/capteurs/model"
	"time"
)

const frequenceEmissionDuration time.Duration = 2
const frequenceEmissionInt int = 2
const tempsEmission int = 1

func main() {
	var configuration model.Configuration = model.GetConfig()
	configuration.ID_CLIENT = ""

	nbTours := (tempsEmission * 60) / frequenceEmissionInt

	for i := nbTours ; i > 0 ; i-- {
		var donneesAleatoire model.DonneesCapteur = model.GenererDonneesCapteur()
		println(donneesAleatoire.IdCapteur)
		println(donneesAleatoire.IdAeroport)
		println(donneesAleatoire.TypeMesure)
		println(donneesAleatoire.ValeurMesure)
		println(donneesAleatoire.Date.String())
		model.Send(configuration, donneesAleatoire)
		time.Sleep(frequenceEmissionDuration * time.Second)
	}
}