package rabbitmq

import (
	"log"

	"fmt"
	"github.com/streadway/amqp"
)

func ReceiveFromMQ(exchange, key, queue string, args amqp.Table) (msgs <-chan amqp.Delivery, err error) {
	if ch == nil {
		err = fmt.Errorf("Failed to connect to RabbitMQ")
		return
	}

	err = ch.ExchangeDeclare(exchange, "direct", true, false, false, false, nil)
	if err != nil {
		log.Printf("%s: %s\n", "Failed to declare a exchange", err)
		return
	}

	_, err = ch.QueueDeclare(queue, true, false, false, false, args)
	if err != nil {
		log.Printf("%s: %s\n", "Failed to declare a queue", err)
		return
	}

	err = ch.QueueBind(queue, key, exchange, false, nil)
	if err != nil {
		log.Printf("%s: %s\n", "Failed to bind queue to exchange", err)
		return
	}

	msgs, err = ch.Consume(
		queue, // queue
		queue, // consumer
		false, // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args
	)
	if err != nil {
		log.Printf("%s: %s\n", "Failed to register a consumer", err)
		return
	}

	return
}
