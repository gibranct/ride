package queue

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Queue interface {
	connect() error
	Disconnect() error
	Publish(exchange string, message []byte) error
	Consume(queue string, callback func(msg []byte) error) error
}

type RabbitMQAdapter struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

func must[T any](data T, err error) T {
	if err != nil {
		log.Fatal(err)
	}
	return data
}

func (ra *RabbitMQAdapter) connect() error {
	amqpURI := "amqp://guest:guest@localhost:5672/"
	conn, err := amqp.Dial(amqpURI)
	if err != nil {
		return err
	}
	ra.conn = conn
	return nil
}

func (ra *RabbitMQAdapter) Disconnect() error {
	return ra.conn.Close()
}

func (ra *RabbitMQAdapter) Publish(exchange string, message []byte) error {
	return ra.channel.Publish(exchange, "", false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        message,
	})
}

func (ra *RabbitMQAdapter) Consume(queue string, callback func(msg []byte) error) error {
	msgs, err := ra.channel.Consume(queue, "", false, false, false, false, nil)
	if err != nil {
		return err
	}
	go func() {
		for msg := range msgs {
			log.Printf("Received a message: %s", msg.Body)
			err := callback(msg.Body)
			if err != nil {
				log.Printf("Error processing message: %v", err)
				msg.Nack(false, false)
			} else {
				msg.Ack(false)
			}
		}
	}()
	return nil
}

func NewRabbitMQAdapter() Queue {
	rabbit := &RabbitMQAdapter{}
	err := rabbit.connect()
	if err != nil {
		panic(err)
	}
	ch, err := rabbit.conn.Channel()
	if err != nil {
		panic(err)
	}
	rabbit.channel = ch
	must(ch.ExchangeDeclare("rideCompleted", "direct", true, false, false, false, nil), nil)
	must(ch.QueueDeclare("rideCompleted.processPayment", true, false, false, false, nil))
	must(ch.QueueDeclare("rideCompleted.generateInvoice", true, false, false, false, nil))
	must(ch.QueueDeclare("rideCompleted.sendReceipt", true, false, false, false, nil))
	must(ch.QueueBind("rideCompleted.processPayment", "", "rideCompleted", false, nil), nil)
	must(ch.QueueBind("rideCompleted.generateInvoice", "", "rideCompleted", false, nil), nil)
	must(ch.QueueBind("rideCompleted.sendReceipt", "", "rideCompleted", false, nil), nil)
	if err != nil {
		panic(err)
	}
	return rabbit
}
