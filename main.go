package main

import (
	"github.com/abhibvp003/docker-kafka/consumer"
	"github.com/abhibvp003/docker-kafka/producer"
	"os"
)

func main()  {
	if os.Args[1] == "consumer" {
		consumer.StartConsumer()
	} else if os.Args[1] == "producer" {
		producer.StartProducer()
	}
}
