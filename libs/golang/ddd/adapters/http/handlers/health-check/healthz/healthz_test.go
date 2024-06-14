package healthz

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type WebHealthzHandlerSuite struct {
	suite.Suite
	handler          *WebHealthzHandler
	mockTimeProvider *MockTimeProvider
}

func TestWebHealthzHandlerSuite(t *testing.T) {
	suite.Run(t, new(WebHealthzHandlerSuite))
}

func (suite *WebHealthzHandlerSuite) SetupTest() {
	suite.mockTimeProvider = &MockTimeProvider{
		currentTime: time.Now(),
	}
	suite.handler = NewWebHealthzHandler(suite.mockTimeProvider)
}

func (suite *WebHealthzHandlerSuite) TestHealthzHandlerBeforeFiveSeconds() {
	req, err := http.NewRequest("GET", "/healthz", nil)
	assert.NoError(suite.T(), err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(suite.handler.Healthz)

	handler.ServeHTTP(rr, req)

	assert.Equal(suite.T(), http.StatusInternalServerError, rr.Code)
	assert.Contains(suite.T(), rr.Body.String(), "Healthz check failed after")
}

func (suite *WebHealthzHandlerSuite) TestHealthzHandlerAfterFiveSeconds() {
	// Advance time by 5 seconds
	suite.mockTimeProvider.Advance(5 * time.Second)

	req, err := http.NewRequest("GET", "/healthz", nil)
	assert.NoError(suite.T(), err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(suite.handler.Healthz)

	handler.ServeHTTP(rr, req)

	assert.Equal(suite.T(), http.StatusOK, rr.Code)
	assert.Equal(suite.T(), "Healthz check passed", rr.Body.String())
}
