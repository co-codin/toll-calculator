package main

import (
	"log"

	"github.com/co-codin/toll-calculator/aggregator/client"
)

const (
	kafkaTopic         = "obudata"
	aggregatorEndpoint = "http://127.0.0.1:3000/aggregate"
)

func main() {
	var (
		err error
		svc CalculatorServicer
	)
	svc = NewCalculatorService()
	svc = NewLogMiddleware(svc)
	aggClient := client.NewHTTPClient(aggregatorEndpoint)

	kafkaConsumer, err := NewKafkaConsumer(kafkaTopic, svc, aggClient)
	if err != nil {
		log.Fatal(err)
	}
	kafkaConsumer.Start()
}
