package server

import (
	"net/http"
	"resto-backend/config"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
)

type Server struct {
	config config.Config
	router *echo.Echo
}

// Validator wraps go-playground/validator and implements echo.Validator interface
type Validator struct {
	validator *validator.Validate
}

func (v *Validator) Validate(i interface{}) error {
	return v.validator.Struct(i)
}

func NewServer(config config.Config) (*Server, error) {

	server := &Server{}
	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := echo.New()

	// Use validator
	router.Validator = &Validator{validator: validator.New()}

	// Middleware
	router.GET("/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Start(address)
}
