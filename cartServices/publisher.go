// package main

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"time"

// 	amqp "github.com/rabbitmq/amqp091-go"
// )

// func FailOnError(err error, msg string) {
// 	if err != nil {
// 		log.Panicf("%s: %s", msg, err)
// 	}
// }

// func Publisher(data, exchangeName, exchangeType, routingKey string) {
// 	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
// 	FailOnError(err, "Failed to connect to RabbitMQ")
// 	defer conn.Close()

// 	ch, err := conn.Channel()
// 	fmt.Println("Channel created:", ch)
// 	FailOnError(err, "Failed to open a channel")
// 	defer ch.Close()

// 	ch.ExchangeDeclare(exchangeName, exchangeType, false, false, false, false, nil)

// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()

// 	err = ch.PublishWithContext(ctx,
// 		exchangeName,
// 		routingKey,
// 		false,
// 		false,
// 		amqp.Publishing{
// 			ContentType: "text/plain",
// 			Body:        []byte(data),
// 		})
// 	FailOnError(err, "Failed to publish a message")
// 	log.Printf(" [x] Sent %s\n", data)
// }


