package main

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/andrewunifei/full-cycle-imersao/go/internal/infra/akafka"
	"github.com/andrewunifei/full-cycle-imersao/go/internal/market/dto"
	"github.com/andrewunifei/full-cycle-imersao/go/internal/market/entity"
	"github.com/andrewunifei/full-cycle-imersao/go/internal/market/transformer"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() { 
	// Entidade na aplicação principal - furo de camada
	ordersIn := make(chan *entity.Order)
	ordersOut := make(chan *entity.Order)

	wg := &sync.WaitGroup{}
	defer wg.Wait()

	kafkaMsgChan := make(chan *kafka.Message)
	configMap := &kafka.ConfigMap{
		"bootstrap.servers": "host.docker.internal:9094",
		"group.id": "group-1",
		"auto.offset.reset": "earliest",
	}

	producer := akafka.NewKafkaProducer(configMap)
	consumer := akafka.NewKafkaConsumer(configMap, []string{"input-orders"})

	go consumer.Consume(kafkaMsgChan) // Criação de Thread 2

	// Recebe do kafka
	// Realiza computações
	// Envia para o output e publica no kafka 
	book := entity.NewBook(ordersIn, ordersOut, wg)
	go book.Trade() // Criação de Thread 3

	go func() {
		for msg := range kafkaMsgChan {
			wg.Add(1)
			tradeInput := dto.TradeInput{}
			err := json.Unmarshal(msg.Value, &tradeInput)

			if err != nil {
				panic(err)
			}

			order := transformer.TransformInput(tradeInput)
			ordersIn <- order
		}	
	}()

	for res := range ordersOut {
		output := transformer.TransformOutput(res)
		outputJson, err := json.MarshalIndent(output, "", "   ")

		if err != nil {
			fmt.Println(err)
		}

		producer.Publish(outputJson, []byte("orders"), "output-orders")
	}
}