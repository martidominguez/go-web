package application

import (
	"app/internal/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// ServerChi is the server using chi
type ServerChi struct {
	// address is the address to listen on
	address string
}

// NewServerChi returns a new ServerChi instance
func NewServerChi(address string) *ServerChi {
	defaultAddress := ":8080"
	if address != "" {
		defaultAddress = address
	}

	return &ServerChi{
		address: defaultAddress,
	}
}

func (s *ServerChi) Run() error {
	// dependencies
	//handler
	productsHandler := handlers.NewHandler() // maybe i can change and improve this later. check github
	// router
	router := chi.NewRouter()

	// endpoints
	router.Get("/ping", productsHandler.Pong())

	router.Get("/products", productsHandler.GetProducts())

	router.Get("/products/{id}", productsHandler.GetProductById())

	// here must be exercise d from 09-01

	router.Post("/products", productsHandler.CreateProduct())

	// run server
	return http.ListenAndServe(s.address, router)
}
