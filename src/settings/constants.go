package settings

import (
	"os"
)

func GetAMQPHost() string {

	aqmpHost := os.Getenv("AMQP_HOST")

	return aqmpHost
}
