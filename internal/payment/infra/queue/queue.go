package queue

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Queue interface {
	Connect() error
	Disconnect() error
	Publish(exchange string, message []byte) error
	Consume(queue string, callback func(msg []byte) error) error
}

type RabbitMQAdapter struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

func (ra *RabbitMQAdapter) Connect() error {
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
	err := rabbit.Connect()
	if err != nil {
		log.Fatalln(err)
	}
	_, err = rabbit.conn.Channel()
	if err != nil {
		log.Fatalln(err)
	}
	return rabbit
}
