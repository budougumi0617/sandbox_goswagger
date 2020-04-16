package controller

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/budougumi0617/sandbox_goswagger/gen/restapi"
	"github.com/budougumi0617/sandbox_goswagger/gen/restapi/operations"
	"github.com/go-openapi/loads"
)

// Server is graceful start server.
type Server struct {
	Listener net.Listener
	Server   *http.Server
}

// Start to accept listener.
func (s *Server) Start() error {
	return s.Server.Serve(s.Listener)
}

// Shutdown terminates server by graceful shutdown.
func (s *Server) Shutdown(ctx context.Context) error {
	return s.Server.Shutdown(ctx)
}

// NewServer returns Server.
func NewServer(l net.Listener) *Server {
	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	h := ConfigureAPI(operations.NewSampleAPI(swaggerSpec))
	// TODO: Functional Option Patternを使いたい。
	return &Server{
		Listener: l,
		Server: &http.Server{
			Handler:      h,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  120 * time.Second,
		},
	}
}

// port from net/http/httptest.
func newLocalListener() net.Listener {
	l, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		if l, err = net.Listen("tcp6", "[::1]:0"); err != nil {
			panic(fmt.Sprintf("httptest: failed to listen on a port: %v", err))
		}
	}
	return l
}
