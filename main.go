package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/judaro13/users_ms/store"
	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

var (
	ErrEnv        = errors.New("not found RABBIT_PATH environment variable")
	ErrEnvChannel = errors.New("not found RABBIT_CHANNEL environment variable")
	ErrDBHost     = errors.New("not found DB_HOST environment variable")
	ErrDBUser     = errors.New("not found DB_USER environment variable")
	ErrDBPassword = errors.New("not found DB_PASSWORD environment variable")
	ErrDBName     = errors.New("not found DB_NAME environment variable")
	ErrConnect    = errors.New("failed to connect to RabbitMQ")
	ErrChannel    = errors.New("failed to open a channel")
	ErrQueue      = errors.New("failed to declare a queue")
	ErrRegister   = errors.New("failed to register a consumer")
)

func ValidEnvVars() {
	if len(os.Getenv("RABBIT_PATH")) == 0 {
		panic(ErrEnv)
	}
	if len(os.Getenv("RABBIT_CHANNEL")) == 0 {
		panic(ErrEnvChannel)
	}
	if len(os.Getenv("DB_HOST")) == 0 {
		panic(ErrDBHost)
	}
	if len(os.Getenv("DB_USER")) == 0 {
		panic(ErrDBUser)
	}
	if len(os.Getenv("DB_PASSWORD")) == 0 {
		panic(ErrDBPassword)
	}
	if len(os.Getenv("DB_NAME")) == 0 {
		panic(ErrDBName)
	}
}

func main() {
	ValidEnvVars()
	rabbit_url := os.Getenv("RABBIT_PATH")
	rabbit_ch := os.Getenv("RABBIT_CHANNEL")

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
			user.Store(d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever

}
