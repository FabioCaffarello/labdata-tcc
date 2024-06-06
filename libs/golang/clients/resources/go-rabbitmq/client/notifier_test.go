package gorabbitmq

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type GoRabbitMQNotifierSuite struct {
	suite.Suite
	client *Client
	config Config
}

func TestGoRabbitMQNotifierSuite(t *testing.T) {
	suite.Run(t, new(GoRabbitMQNotifierSuite))
}

func (suite *GoRabbitMQNotifierSuite) SetupSuite() {
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

func (suite *GoRabbitMQNotifierSuite) SetupTest() {
	client, err := NewClient(suite.config)
	assert.NoError(suite.T(), err)
	suite.client = client
}

func (suite *GoRabbitMQNotifierSuite) TearDownTest() {
	if suite.client != nil {
		err := suite.client.Close()
		assert.NoError(suite.T(), err)
	}
}

func (suite *GoRabbitMQNotifierSuite) TestNewRabbitMQNotifier() {
	notifier := NewRabbitMQNotifier(suite.client)
	assert.NotNil(suite.T(), notifier)
	assert.Equal(suite.T(), suite.client, notifier.rmqClient)
}

func (suite *GoRabbitMQNotifierSuite) TestNotify() {
	notifier := NewRabbitMQNotifier(suite.client)
	err := notifier.Notify([]byte("test message"), "test_routing_key")
	assert.NoError(suite.T(), err)
}
