package main

import (
	"github.com/my/repo/external/sub/model"
	"time"
)
const frequenceEmissionDuration time.Duration = 1
const frequenceEmissionInt int = 2
const tempsEmission int = 1

func main()  {
	var configuration model.Configuration = model.GetConfig()

	client := model.Connect("sub", configuration)

	var nbTours = (tempsEmission * 60) / frequenceEmissionInt
	configuration.ID_CLIENT = "Sub"
	for i := nbTours ; i > 0 ; i-- {
		model.Get(configuration,client)
		time.Sleep(frequenceEmissionDuration * time.Second)
	}
}

