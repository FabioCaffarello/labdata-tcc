package gorabbitmq

import (
	"context"
	"testing"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type GoRabbitMQSuite struct {
	suite.Suite
	client *Client
	config Config
}

func TestGoRabbitMQSuite(t *testing.T) {
	suite.Run(t, new(GoRabbitMQSuite))
}

func (suite *GoRabbitMQSuite) SetupSuite() {
	suite.config = Config{
		User:         "guest",
		Password:     "guest",
		Host:         "localhost",
		Port:         "5672",
		Protocol:     "amqp",
		ExchangeName: "test_exchange",
		ExchangeType: "direct",
	}
}

func (suite *GoRabbitMQSuite) SetupTest() {
	client, err := NewClient(suite.config)
	assert.NoError(suite.T(), err)
	suite.client = client
}

func (suite *GoRabbitMQSuite) TearDownTest() {
	if suite.client != nil {
		err := suite.client.Close()
		assert.NoError(suite.T(), err)
	}
}

func (suite *GoRabbitMQSuite) TestNewClient() {
	client, err := NewClient(suite.config)
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), client)
	assert.Equal(suite.T(), suite.config.ExchangeName, client.ExchangeName)
	assert.Equal(suite.T(), suite.config.ExchangeType, client.ExchangeType)
}

func (suite *GoRabbitMQSuite) TestConnect() {
	err := suite.client.connect()
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), suite.client.Conn)
}

func (suite *GoRabbitMQSuite) TestChannel() {
	err := suite.client.channel()
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), suite.client.Channel)
}

func (suite *GoRabbitMQSuite) TestDeclareExchange() {
	err := suite.client.declareExchange()
	assert.NoError(suite.T(), err)
}

func (suite *GoRabbitMQSuite) TestClose() {
	err := suite.client.Close()
	assert.NoError(suite.T(), err)
}

func (suite *GoRabbitMQSuite) TestPublish() {
	routingKey := "test_key_publish"
	ctx := context.Background()
	message := []byte("test message")
	err := suite.client.publish(ctx, "text/plain", message, routingKey)
	assert.NoError(suite.T(), err)
}

func (suite *GoRabbitMQSuite) TestDeclareQueue() {
	queueName := "test_queue_declare"
	q, err := suite.client.declareQueue(queueName, nil)
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), q.Name, queueName)
}

func (suite *GoRabbitMQSuite) TestBindQueue() {
	queueName := "test_queue_bind"
	routingKey := "test_key_bind"
	q, err := suite.client.declareQueue(queueName, nil)
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), q.Name, queueName)

	err = suite.client.bindQueue(q.Name, routingKey)
	assert.NoError(suite.T(), err)
}

func (suite *GoRabbitMQSuite) TestBindQueueAndPublish() {
	queueName := "test_queue_publish"
	routingKey := "test_queue_publish"
	q, err := suite.client.declareQueue(queueName, nil)
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), q.Name, queueName)

	err = suite.client.bindQueue(q.Name, routingKey)
	assert.NoError(suite.T(), err)

	ctx := context.Background()
	message := []byte("test message")
	err = suite.client.publish(ctx, "text/plain", message, routingKey)
	assert.NoError(suite.T(), err)
}

func (suite *GoRabbitMQSuite) TestConsume() {
	// Declare and bind the queue first
	queueName := "test_queue_consume"
	routingKey := "test_key_consume"
	q, err := suite.client.declareQueue(queueName, nil)
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), q.Name, queueName)

	// Bind the queue to the exchange
	err = suite.client.bindQueue(q.Name, routingKey)
	assert.NoError(suite.T(), err)

	// Test the consume method
	msgCh := make(chan amqp.Delivery, 1)
	go suite.client.consume(msgCh, "consumer-name", queueName, false)

	// Publish a test message
	message := []byte("test message")
	err = suite.client.publish(context.Background(), "text/plain", message, routingKey)
	assert.NoError(suite.T(), err)

	select {
	case msg := <-msgCh:
		assert.Equal(suite.T(), message, msg.Body)
	case <-time.After(5 * time.Second):
		assert.Fail(suite.T(), "Timed out waiting for message")
	}
}
