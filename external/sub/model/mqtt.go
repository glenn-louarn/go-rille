package model

import (
	"bytes"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"io/ioutil"
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
	/*
	opts.SetUsername(config.ID_CLIENT)
	password := config.MOT_DE_PASSE
	opts.SetPassword(password)
	opts.SetClientID(clientId)
	*/
	return opts
}
func Get(config Configuration,client mqtt.Client){
	println("test sub : ")
	client.Subscribe("capteur",0, func(client mqtt.Client, message mqtt.Message) {
		var msg = string(message.Payload())
		fmt.Println(msg)
		req,err := http.NewRequest("POST","http://localhost:8080/donnees",bytes.NewBuffer([]byte(msg)))
		req.Header.Set("X-Custom-Header","MyValue")
		req.Header.Set("Content-Type","application/json")
		webClient := &http.Client{}
		resp,err := webClient.Do(req)
		if err!=nil{
			panic(err)
		}
		defer resp.Body.Close()
		fmt.Println("response Status:",resp.Status)
		fmt.Println("response Headers:",resp.Header)
		body,_:=ioutil.ReadAll(resp.Body)
		fmt.Println("response Body:",string(body))
	})
}
