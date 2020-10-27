package model

import (
	"encoding/json"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
	"time"
)

func connect(clientId string, config Configuration) mqtt.Client {
	opts := createClientOptions(clientId, config)
	client := mqtt.NewClient(opts)
	token := client.Connect()
	for !token.WaitTimeout(3 * time.Second) {
	}
	if err := token.Error(); err != nil {
		log.Fatal(err)
	}
	return client
}

func createClientOptions(clientId string, config Configuration) *mqtt.ClientOptions {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("%s:%s", config.ADRESSE_BROKER_MQTT, config.PORT_BROKER_MQTT))
	/*
	opts.SetUsername(config.ID_CLIENT)
	password := config.MOT_DE_PASSE
	opts.SetPassword(password)
	opts.SetClientID(clientId)
	*/
	return opts
}

func Send(config Configuration, donnees DonneesCapteur) {
	jsonAEnvoyer, _ := json.Marshal(donnees)
	println("test : ")
	println(string(jsonAEnvoyer))

	client := connect("pub", config)
	client.Publish("capteur", 0, false, string(jsonAEnvoyer))
}
