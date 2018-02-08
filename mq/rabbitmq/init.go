package rabbitmq

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

var (
	rabbitmqHost     = ""
	rabbitmqPort     = ""
	rabbitmqVhost    = ""
	rabbitmqUserName = ""
	rabbitmqPassword = ""

	conn *amqp.Connection
	ch   *amqp.Channel
)

//InitMQ
func InitMQ(host, port, vhost, userName, password string) {

	var err error

	rabbitmqHost = host
	rabbitmqPort = port
	rabbitmqVhost = vhost
	rabbitmqUserName = userName
	rabbitmqPassword = password

	if vhost == "" {
		rabbitmqVhost = "/"
	} else if vhost[0] != '/' {
		rabbitmqVhost = "/" + vhost
	}

	dialUrl := fmt.Sprintf("amqp://%s:%s@%s:%s%s", rabbitmqUserName,
		rabbitmqPassword, rabbitmqHost, rabbitmqPort, rabbitmqVhost)
	conn, err = amqp.Dial(dialUrl)
	if err != nil {
		log.Fatalf("%s: %s\n", "Failed to connect to RabbitMQ", err)
		return
	}

	ch, err = conn.Channel()
	if err != nil {
		log.Fatalf("%s: %s\n", "Failed to open a channel", err)
		return
	}

}

//CloseQueue
func CloseQueue(queue string) {
	num, err := ch.QueueDelete(queue, false, false, false)
	if err != nil {
		log.Println(err.Error())
		return
	}

	log.Println("delete queue ", queue, num)

}
