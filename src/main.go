package main

import (
	"encoding/json"
	"fmt"
	"log"

	dotenv "github.com/joho/godotenv"
	"github.com/wagslane/go-rabbitmq"

	esClient "github.com/davisbento/golang-worker/src/infra/elasticSearch"
	"github.com/davisbento/golang-worker/src/interfaces"
	"github.com/davisbento/golang-worker/src/settings"
	"github.com/davisbento/golang-worker/src/utils"
)

func main() {
	err := dotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	c := esClient.MakeElasticConnection()

	consumer, err := rabbitmq.NewConsumer(settings.GetAMQPHost())

	if err != nil {
		log.Fatal(err)
	}

	err = consumer.StartConsuming(
		func(d rabbitmq.Delivery) bool {
			message := string(d.Body)
			payload := interfaces.Message{}
			payload.Timestamp = utils.GetDateIso()
			json.Unmarshal([]byte(message), &payload)

			log.Printf("data: %v", payload)

			esClient.InsertNewLog(c, payload)

			// true to ACK, false to NACK
			return true
		},
		"logs",
		[]string{"logs_key"},
		rabbitmq.WithConsumeOptionsBindingExchange("logs"))

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("up and running")

	// block main thread so consumers run forever
	forever := make(chan struct{})
	<-forever
}
