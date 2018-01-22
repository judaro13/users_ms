package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

var (
	ErrEnv        = errors.New("not found RABBIT_PATH  environment variable")
	ErrEnvChannel = errors.New("not found RABBIT_CHANNEL  environment variable")
	ErrConnect    = errors.New("failed to connect to RabbitMQ")
	ErrChannel    = errors.New("failed to open a channel")
	ErrQueue      = errors.New("failed to declare a queue")
	ErrRegister   = errors.New("failed to register a consumer")
)

func main() {
	rabbit_url := os.Getenv("RABBIT_PATH")
	rabbit_ch := os.Getenv("RABBIT_CHANNEL")
	if len(rabbit_url) == 0 {
		failOnError(ErrEnv, ErrEnv.Error())
	}
	if len(rabbit_ch) == 0 {
		failOnError(ErrEnvChannel, ErrEnvChannel.Error())
	}

	conn, err := amqp.Dial(rabbit_url)
	failOnError(err, ErrConnect.Error())
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, ErrChannel.Error())
	defer ch.Close()

	q, err := ch.QueueDeclare(
		rabbit_ch, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	failOnError(err, ErrQueue.Error())

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, ErrRegister.Error())

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever

}
