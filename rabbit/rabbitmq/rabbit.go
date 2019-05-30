package rabbitmq

import (
	"fmt"
	"time"

	"github.com/apex/log"
	"github.com/pkg/errors"
	"github.com/streadway/amqp"
)

// RabbitMQ stores rabbitmq's connection information
// it also handles disconnection (purpose of URL and QueueName storage)
type RabbitMQ struct {
	URL        string
	Exchange   string
	Conn       *amqp.Connection
	Chann      *amqp.Channel
	Queue      amqp.Queue
	closeChann chan *amqp.Error
	quitChann  chan bool
}

// AMQP is the amqp configuration
type AMQP struct {
	URL      string `default:"amqp://guest:guest@127.0.0.1:5672/guest"`
	Exchange string `default:"amq.direct"`
}

// Config is the application configuration
type Config struct {
	AppName    string `json:"app-name" default:"rabbitmq"`
	AppVersion string `json:"app-version" required:"true"`

	AMQP AMQP
}

func (rmq *RabbitMQ) Load() error {
	var err error

	rmq.Conn, err = amqp.Dial(rmq.URL)
	if err != nil {
		return err
	}

	rmq.Chann, err = rmq.Conn.Channel()
	if err != nil {
		return err
	}

	log.Info("connection to rabbitMQ established")

	rmq.closeChann = make(chan *amqp.Error)
	rmq.Conn.NotifyClose(rmq.closeChann)

	// declare exchange if not exist
	err = rmq.Chann.ExchangeDeclare(rmq.Exchange, "direct", true, false, false, false, nil)
	if err != nil {
		return errors.Wrapf(err, "declaring exchange %q", rmq.Exchange)
	}
	args := make(amqp.Table)
	args["x-delayed-type"] = "direct"

	err = rmq.Chann.ExchangeDeclare("delayed", "x-delayed-message", true, false, false, false, args)
	if err != nil {
		return errors.Wrapf(err, "declaring exchange %q", "delayed")
	}

	// err = declareConsumer(rmq)
	// if err != nil {
	// 	return err
	// }
	return nil
}

func InitRabbitMQ(config AMQP) (*RabbitMQ, error) {
	rmq := &RabbitMQ{
		URL:      config.URL,
		Exchange: config.Exchange,
	}

	err := rmq.Load()
	if err != nil {
		return nil, err
	}

	rmq.quitChann = make(chan bool)
	fmt.Println("---------------")
	go rmq.handleDisconnect()

	return rmq, err
}

// handleDisconnect handle a disconnection trying to reconnect every 5 seconds
func (rmq *RabbitMQ) handleDisconnect() {
	for {
		select {
		case errChann := <-rmq.closeChann:
			if errChann != nil {
				log.Errorf("rabbitMQ disconnection: %v", errChann)
			}
		case <-rmq.quitChann:
			rmq.Conn.Close()
			log.Info("...rabbitMQ has been shut down")
			rmq.quitChann <- true
			return
		}

		log.Info("...trying to reconnect to rabbitMQ...")

		time.Sleep(5 * time.Second)

		if err := rmq.Load(); err != nil {
			log.Errorf("rabbitMQ error: %v", err)
		}
	}
}

// Shutdown closes rabbitmq's connection
func (rmq *RabbitMQ) Shutdown() {
	rmq.quitChann <- true

	log.Info("shutting down rabbitMQ's connection...")

	<-rmq.quitChann
}

// Publish sends the given body on the routingKey to the channel
func (rmq *RabbitMQ) Publish(routingKey string, body []byte) error {
	return rmq.publish(rmq.Exchange, routingKey, body, 0)
}

// PublishWithDelay sends the given body on the routingKey to the channel with a delay
func (rmq *RabbitMQ) PublishWithDelay(routingKey string, body []byte, delay int64) error {
	return rmq.publish("delayed", routingKey, body, delay)
}

func (rmq *RabbitMQ) publish(exchange string, routingKey string, body []byte, delay int64) error {
	headers := make(amqp.Table)

	log.Debugf("publishing to %q %q", routingKey, body)

	if delay != 0 {
		headers["x-delay"] = delay
	}

	return rmq.Chann.Publish(exchange, routingKey, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		Timestamp:    time.Now(),
		ContentType:  "application/json",
		Body:         body,
		Headers:      headers,
	})
}
