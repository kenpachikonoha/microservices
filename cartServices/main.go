package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func FailOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func Publisher(data, exchangeName, exchangeType, routingKey string) {
	conn, err := amqp.Dial("amqp://localhost:5672")
	FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	ch.ExchangeDeclare(exchangeName, exchangeType, false, false, false, false, nil)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = ch.PublishWithContext(ctx,
		exchangeName,
		routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(data),
		})
	FailOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s\n", data)
}

func main() {
	const exchangeType string = "direct"
	const exchangeName string = "cartExchange"
	const routingKey string = "fromCart"
	const data string = `{		"cartId": "123",
		"cartItems": [
			{
				"itemId": "1",
				"itemName": "Condones",
				"itemPrice": 400
				"itemQuantity": 2
			},
			{
				"itemId": "3",
				"itemName": "Azulita",
				"itemPrice": 125,
				"itemQuantity": 1
			}]`

	Publisher(data, exchangeName, exchangeType, routingKey)
	cartHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, data)
	}

	http.HandleFunc("/cart", cartHandler)
	log.Fatal(http.ListenAndServe(":5566", nil))
}
