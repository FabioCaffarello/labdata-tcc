package consumer

import (
	"context"
	"fmt"
	queue "libs/golang/clients/resources/go-rabbitmq/client"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

// AmqpConsumer handles consuming messages from a RabbitMQ queue.
type AmqpConsumer struct {
	rabbitMQ         *queue.Client
	rabbitMQConsumer *queue.RabbitMQConsumer
	queueName        string
	routingKey       string
	msgCh            chan []byte
	quitCh           chan struct{}
}

// NewAmqpConsumer creates a new instance of AmqpConsumer.
//
// Parameters:
//   - rmqClient: RabbitMQ client.
//   - queueName: Name of the queue to consume messages from.
//   - consumerName: Name of the consumer.
//   - routingKey: Routing key to bind the queue to.
//
// Returns:
//   - A new instance of AmqpConsumer.
func NewAmqpConsumer(rmqClient *queue.Client, queueName, consumerName, routingKey string) *AmqpConsumer {
	consumerConfig := queue.ConsumerConfig{
		ConsumerName: consumerName,
		AutoAck:      false,
		Args:         nil,
	}

	consumer := queue.NewRabbitMQConsumer(
		rmqClient,
		consumerConfig,
	)

	return &AmqpConsumer{
		rabbitMQ:         rmqClient,
		rabbitMQConsumer: consumer,
		queueName:        queueName,
		routingKey:       routingKey,
		msgCh:            make(chan []byte),
		quitCh:           make(chan struct{}),
	}
}

// GetListenerTag returns a listener tag that uniquely identifies the consumer.
//
// Returns:
//   - A string representing the listener tag.
func (al *AmqpConsumer) GetListenerTag() string {
	return fmt.Sprintf("%s:%s:%s", al.rabbitMQConsumer.ConsumerName, al.queueName, al.routingKey)
}

// Consume starts consuming messages from the queue and processes them.
//
// It listens for messages and sends them to the msgCh channel. If the quitCh
// channel receives a signal, the consumption stops.
func (al *AmqpConsumer) Consume() {
	msgCh := make(chan amqp.Delivery)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	log.Println("Starting to consume messages...")
	go al.rabbitMQConsumer.Consume(ctx, msgCh, al.queueName, al.routingKey)

mainloop:
	for {
		select {
		case msg := <-msgCh:
			if msg.Body == nil {
				log.Println("Received nil message, continuing...")
				continue
			}
			log.Printf("Received message: %s from queue: %s", string(msg.Body), al.queueName)
			al.msgCh <- msg.Body
		case <-al.quitCh:
			log.Println("Received quit signal, stopping consumer...")
			break mainloop
		}
	}
	log.Println("Consumer main loop exited")
}

// GetMsgCh returns the channel where messages are sent.
//
// Returns:
//   - A read-only channel of byte slices containing messages.
func (al *AmqpConsumer) GetMsgCh() <-chan []byte {
	return al.msgCh
}

// Stop stops the consumer by closing the quitCh channel.
func (al *AmqpConsumer) Stop() {
	close(al.quitCh)
}
