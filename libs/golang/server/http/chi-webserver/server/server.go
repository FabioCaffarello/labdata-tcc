package webserver

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// Server represents an HTTP server with a router and address.
type Server struct {
	router *chi.Mux
	addr   string
}

// NewWebServer creates and returns a new Server instance with the specified address.
//
// Parameters:
//
//	addr: The address to run the server on.
//
// Returns:
//
//	A new Server instance.
func NewWebServer(addr string) *Server {
	return &Server{
		router: chi.NewRouter(),
		addr:   addr,
	}
}

// ConfigureDefaults sets up the default middleware for the server, including request ID, real IP, logger, recoverer, and a timeout of 60 seconds.
func (s *Server) ConfigureDefaults() {
	middlewares := []func(http.Handler) http.Handler{
		middleware.RequestID,
		middleware.RealIP,
		middleware.Logger,
		middleware.Recoverer,
		middleware.Timeout(60 * time.Second),
	}
	s.RegisterMiddlewares(middlewares...)
}

// RegisterMiddlewares adds multiple middlewares to the server.
//
// Parameters:
//
//	middlewares: A list of middleware functions to add to the server.
//
// Returns:
//
//	None.
//
// Example:
//
//	server.RegisterMiddlewares(middleware1, middleware2, middleware3)
//
// This will add middleware1, middleware2, and middleware3 to the server.
func (s *Server) RegisterMiddlewares(middlewares ...func(http.Handler) http.Handler) {
	for _, m := range middlewares {
		s.router.Use(m)
	}
}

// RegisterRoute adds a new route with an HTTP method, pattern, and handler function. If a group is specified, the route is added to that group.
//
// Parameters:
//
//	method: The HTTP method for the route.
//	pattern: The URL pattern for the route.
//	handler: The handler function for the route.
//	group: An optional group name to add the route to.
//
// Returns:
//
//	None.
//
// Example:
//
//	server.RegisterRoute("GET", "/hello", helloHandler)
//
// This will add a new route to the server that listens for GET requests on the /hello URL pattern and calls the helloHandler function.
func (s *Server) RegisterRoute(method, pattern string, handler http.HandlerFunc, group ...string) {
	if len(group) > 0 && group[0] != "" {
		r := s.router.Route(group[0], func(r chi.Router) {})
		r.MethodFunc(method, pattern, handler)
	} else {
		s.router.MethodFunc(method, pattern, handler)
	}
}

// RegisterRouteGroup registers a group of routes under a common prefix.
//
// Parameters:
//
//	prefix: The common prefix for the group of routes.
//	routes: A function that defines the routes for the group.
//
// Returns:
//
//	None.
//
// Example:
//
//	server.RegisterRouteGroup("/api", func(r chi.Router) {
//		r.Get("/hello", helloHandler)
//		r.Post("/world", worldHandler)
//	})
//
// This will add a group of routes under the /api prefix with two routes: /hello and /world.
func (s *Server) RegisterRouteGroup(prefix string, routes func(r chi.Router)) {
	s.router.Route(prefix, routes)
}

// Start runs the web server on the specified address.
//
// Parameters:
//
//	None.
//
// Returns:
//
//	An error if the server fails to start.
func (s *Server) Start() error {
	return http.ListenAndServe(s.addr, s.router)
}
