package main

import (
	"github.com/a8uhnf/hello-world/rabbit/rabbitmq"
	"github.com/apex/log"
	"github.com/jinzhu/configor"
	"github.com/streadway/amqp"
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
	defer rmq.Shutdown()
	err = declareConsumer(rmq)
	if err != nil {
		failOnError(err, "hello")
	}
	for {

	}
}

// declareConsumer declares all queues and bindings for the consumer
func declareConsumer(rmq *rabbitmq.RabbitMQ) error {
	var err error

	// rmq.Queue, err = rmq.Chann.QueueDeclare("user-created-queue", true, false, false, false, nil)
	// if err != nil {
	// 	return err
	// }
	// err = rmq.Chann.QueueBind(rmq.Queue.Name, "user.event.create", rmq.Exchange, false, nil)
	// if err != nil {
	// 	return err
	// }

	delayedQueue, err := rmq.Chann.QueueDeclare("user-published-queue", true, false, false, false, nil)
	if err != nil {
		return err
	}
	err = rmq.Chann.QueueBind(delayedQueue.Name, "user.event.publish", "delayed", false, nil)
	if err != nil {
		return err
	}

	// Set our quality of service.  Since we're sharing 3 consumers on the same
	// channel, we want at least 2 messages in flight.
	err = rmq.Chann.Qos(2, 0, false)
	if err != nil {
		return err
	}

	published, err := rmq.Chann.Consume(
		"user-published-queue",
		"user-published-consumer",
		false,
		false,
		false,
		false,
		nil)
	if err != nil {
		return err
	}

	go consume(published)

	return nil
}

func consume(ds <-chan amqp.Delivery) {

	for {
		log.Debug("start listenning to consume")

		select {
		case d, ok := <-ds:
			if !ok {
				return
			}
			log.Infof("consume %s", string(d.Body))
			d.Ack(false)
		}
	}
}
