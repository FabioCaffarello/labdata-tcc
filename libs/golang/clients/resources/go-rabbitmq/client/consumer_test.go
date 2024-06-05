package gorabbitmq

import (
	"context"
	"testing"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type GoRabbitMQConsumerSuite struct {
	suite.Suite
	client         *Client
	config         Config
	consumerConfig ConsumerConfig
}

func TestGoRabbitMQConsumerSuite(t *testing.T) {
	suite.Run(t, new(GoRabbitMQConsumerSuite))
}

func (suite *GoRabbitMQConsumerSuite) SetupSuite() {
	suite.config = Config{
		User:         "guest",
		Password:     "guest",
		Host:         "localhost",
		Port:         "5672",
		Protocol:     "amqp",
		ExchangeName: "test_exchange",
		ExchangeType: "direct",
	}

	suite.consumerConfig = ConsumerConfig{
		ConsumerName: "test_consumer",
		AutoAck:      false,
		Args:         nil,
	}
}

func (suite *GoRabbitMQConsumerSuite) SetupTest() {
	client, err := NewClient(suite.config)
	assert.NoError(suite.T(), err)
	suite.client = client
}

func (suite *GoRabbitMQConsumerSuite) TearDownTest() {
	if suite.client != nil {
		err := suite.client.Close()
		assert.NoError(suite.T(), err)
	}
}

func (suite *GoRabbitMQConsumerSuite) TestNewRabbitMQConsumer() {
	consumer := NewRabbitMQConsumer(suite.client, suite.consumerConfig)
	assert.NotNil(suite.T(), consumer)
	assert.Equal(suite.T(), suite.client, consumer.rmqClient)
	assert.Equal(suite.T(), suite.consumerConfig.AutoAck, consumer.autoAck)
	assert.Equal(suite.T(), suite.consumerConfig.Args, consumer.args)
}

func (suite *GoRabbitMQConsumerSuite) TestConsume() {
	queueName := "test_queue"
	routingKey := "test_key"
	msgCh := make(chan amqp.Delivery)

	consumer := NewRabbitMQConsumer(suite.client, suite.consumerConfig)
	go consumer.Consume(msgCh, queueName, routingKey)

	// Publish a test message
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	message := []byte("test message")
	err := suite.client.publish(ctx, "text/plain", message, routingKey)
	assert.NoError(suite.T(), err)

	select {
	case msg := <-msgCh:
		assert.Equal(suite.T(), message, msg.Body)
	case <-ctx.Done():
		suite.T().Error("Did not receive message in time")
	}
}
