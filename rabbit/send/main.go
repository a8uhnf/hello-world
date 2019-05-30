package main

import (
	"github.com/a8uhnf/hello-world/rabbit/rabbitmq"
	"github.com/apex/log"
	"github.com/jinzhu/configor"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	cfg := rabbitmq.Config{}
	configor.Load(&cfg, "./config/config.json")

	rmq, err := rabbitmq.InitRabbitMQ(cfg.AMQP)
	if err != nil {
		failOnError(err, "hello")
	}
	// defer rmq.Shutdown()
	err = rmq.PublishWithDelay("user.event.publish", []byte("hello"), 100)
	if err != nil {
		failOnError(err, "hello")
	}
}
