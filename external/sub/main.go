package main

import (
	"github.com/my/repo/external/sub/model"
	"time"
)
const frequenceEmissionDuration time.Duration = 2
const frequenceEmissionInt int = 2
const tempsEmission int = 10

func main()  {
	var configuration model.Configuration = model.GetConfig()

	client := model.Connect("sub", configuration)

	var nbTours = (tempsEmission * 60) / frequenceEmissionInt
	for i := nbTours ; i > 0 ; i-- {
		model.Get(client)
		time.Sleep(frequenceEmissionDuration * time.Second)
	}
}

