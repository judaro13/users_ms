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
	ErrPublish    = errors.New("failed to publish a message")
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

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, ErrConnect.Error())
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, ErrChannel.Error())
	defer ch.Close()

}
