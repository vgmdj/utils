package rabbitmq

import (
	"log"

	"fmt"
	"github.com/streadway/amqp"
)

func SendToMQ(exchange, key string, body []byte) (err error) {
	if ch == nil {
		err = fmt.Errorf("Failed to connect to RabbitMQ")
		return
	}

	err = ch.ExchangeDeclare(exchange, "direct", true, false, false, false, nil)
	if err != nil {
		log.Printf("%s: %s\n", "Failed to declare a exchange", err)
		return
	}

	err = ch.Publish(
		exchange, // exchange
		key,      // routing key
		false,    // mandatory
		false,    // immediate
		amqp.Publishing{
			ContentType:  "text/plain",
			Body:         body,
			DeliveryMode: amqp.Persistent,
		})
	log.Printf(" [x] Sent %s", body)
	if err != nil {
		log.Printf("%s: %s\n", "Failed to publish a message", err)
		return
	}

	return
}

func SendToDLQue(exchange, key, queue string, body []byte, args amqp.Table) (err error) {
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
		log.Println("%s: %s\n", "Failed to bind a exchange", err)
		return
	}

	err = ch.Publish(
		exchange, // exchange
		key,      // routing key
		false,    // mandatory
		false,    // immediate
		amqp.Publishing{
			ContentType:  "text/plain",
			Body:         body,
			DeliveryMode: amqp.Persistent,
		})
	log.Printf(" [x] Sent %s", body)
	if err != nil {
		log.Printf("%s: %s\n", "Failed to publish a message", err)
		return
	}

	return
}
