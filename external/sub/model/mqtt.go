package model

import (
	"bytes"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
	"net/http"
	"time"
)

func Connect(clientId string, config Configuration) mqtt.Client {
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
	return opts
}
func Get(client mqtt.Client){
	client.Subscribe("capteur",0, func(client mqtt.Client, message mqtt.Message) {
		var msg = string(message.Payload())
		fmt.Println(msg)
		req,err := http.NewRequest("POST","http://localhost:8081/donnees",bytes.NewBuffer([]byte(msg)))
		req.Header.Set("X-Custom-Header","MyValue")
		req.Header.Set("Content-Type","application/json")
		webClient := &http.Client{}
		resp,err := webClient.Do(req)
		if err!=nil{
			panic(err)
		}
		defer resp.Body.Close()
	})
}
