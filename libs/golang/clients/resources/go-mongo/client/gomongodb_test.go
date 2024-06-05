package gomongodb

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type GoMongoDBSuite struct {
	suite.Suite
	client *Client
	ctx    context.Context
}

func TestGoMongoDBSuite(t *testing.T) {
	suite.Run(t, new(GoMongoDBSuite))
}

func (suite *GoMongoDBSuite) SetupSuite() {
	suite.ctx = context.Background()
	config := Config{
		User:     "testuser",
		Password: "testpassword",
		Host:     "localhost",
		Port:     "27017",
		DBName:   "testdb",
	}

	var err error
	suite.client, err = NewClient(config)
	assert.NoError(suite.T(), err)
}

func (suite *GoMongoDBSuite) TearDownSuite() {
	err := suite.client.Disconnect(suite.ctx)
	assert.NoError(suite.T(), err)
}

func (suite *GoMongoDBSuite) TestMongoDBClient() {
	err := suite.client.Ping(suite.ctx, nil)
	assert.NoError(suite.T(), err)
}
