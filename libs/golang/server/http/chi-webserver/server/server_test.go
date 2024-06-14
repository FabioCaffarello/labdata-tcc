package webserver

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type HTTPServerTestSuite struct {
	suite.Suite
	server *Server
}

func (suite *HTTPServerTestSuite) SetupTest() {
	suite.server = NewWebServer(":40")
	suite.server.ConfigureDefaults()
}

func TestHTTPServerTestSuite(t *testing.T) {
	suite.Run(t, new(HTTPServerTestSuite))
}

func (suite *HTTPServerTestSuite) TestServerInitialization() {
	assert.NotNil(suite.T(), suite.server.router)
}

func (suite *HTTPServerTestSuite) TestMiddlewareRegistration() {
	middlewareCount := len(suite.server.router.Middlewares())
	suite.server.RegisterMiddlewares(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(w, r)
		})
	})
	assert.Equal(suite.T(), middlewareCount+1, len(suite.server.router.Middlewares()))
}

func (suite *HTTPServerTestSuite) TestRouteRegistration() {
	suite.server.RegisterRoute("GET", "/test", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/test", nil)
	suite.server.router.ServeHTTP(recorder, request)

	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
}

func (suite *HTTPServerTestSuite) TestRouteGroupRegistration() {
	suite.server.RegisterRouteGroup("/api", func(r chi.Router) {
		r.Get("/test", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		})
	})
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/api/test", nil)
	suite.server.router.ServeHTTP(recorder, request)

	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
}

func (suite *HTTPServerTestSuite) TestServerStart() {
	go func() {
		err := suite.server.Start()
		assert.Nil(suite.T(), err)
	}()
}

func (suite *HTTPServerTestSuite) TestStartHandler() {
	suite.server.RegisterRoute("GET", "/test", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/test", nil)
	suite.server.router.ServeHTTP(recorder, request)

	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
}

func (suite *HTTPServerTestSuite) TestStartHandlerGroup() {
	suite.server.RegisterRouteGroup("/api", func(r chi.Router) {
		r.Get("/test", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		})
	})
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/api/test", nil)
	suite.server.router.ServeHTTP(recorder, request)

	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
}
