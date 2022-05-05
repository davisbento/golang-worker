package infra

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	elastic "github.com/elastic/go-elasticsearch/v6"

	"github.com/davisbento/golang-worker/src/interfaces"
)

func MakeElasticConnection() *elastic.Client {
	cfg := elastic.Config{}
	esClient, _ := elastic.NewClient(cfg)
	_, err := esClient.Info()
	if err != nil {
		log.Fatal("client.Info() ERROR:", err)
	}
	return esClient
}

func InsertNewLog(c *elastic.Client, payload interfaces.Message) {
	dataJSON, err := json.Marshal(payload)
	if err != nil {
		log.Fatal(err)
	}
	res, err := c.Index("logs", strings.NewReader(string(dataJSON)))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("[InsertLog] Insertion Successful", res)
}
