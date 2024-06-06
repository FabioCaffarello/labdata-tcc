package gorabbitmq

import (
	"context"
	"testing"

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
	ctx := context.Background()
	message := []byte("test message")
	err := suite.client.publish(ctx, "text/plain", message, "test_key")
	assert.NoError(suite.T(), err)
}

func (suite *GoRabbitMQSuite) TestDeclareQueue() {
	err := suite.client.declareQueue("test_queue", nil)
	assert.NoError(suite.T(), err)
}

func (suite *GoRabbitMQSuite) TestBindQueue() {
	err := suite.client.declareQueue("test_queue", nil)
	assert.NoError(suite.T(), err)

	err = suite.client.bindQueue("test_queue", "test_key")
	assert.NoError(suite.T(), err)
}

func (suite *GoRabbitMQSuite) TestConsume() {
	// Declare and bind the queue first
	err := suite.client.declareQueue("test_queue", nil)
	assert.NoError(suite.T(), err)

	// Bind the queue to the exchange
	err = suite.client.bindQueue("test_queue", "test_key")
	assert.NoError(suite.T(), err)

	// Test the consume method
	msgCh, err := suite.client.consume("test_consumer", "test_queue", true)
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), msgCh)
}
