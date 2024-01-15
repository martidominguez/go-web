package application

import (
	"app/internal"
	"app/internal/handler"
	"app/internal/middleware"
	"app/internal/repository"
	"app/internal/service"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// NewDefaultHTTP creates a new instance of a default http server
func NewDefaultHTTP(addr string, token string) *DefaultHTTP {
	// default config / values
	// ...

	return &DefaultHTTP{
		addr:  addr,
		token: token,
	}
}

// DefaultHTTP is a struct that represents the default implementation of a http server
type DefaultHTTP struct {
	// addr is the address of the http server
	addr string
	// token is the token of the http server
	token string
}

// Run runs the http server
func (h *DefaultHTTP) Run() (err error) {
	// initialize dependencies
	// - db
	db := internal.LoadProducts()
	// - repository
	rp := repository.NewProductsMap(db)
	// - service
	sv := service.NewProductDefault(rp)
	// - handler
	hd := handler.NewDefaultProducts(sv)
	// - router
	rt := chi.NewRouter()
	// - middleware
	auth := middleware.NewAuthenticator(h.token)
	logger := middleware.NewLogger()

	// endpoints
	rt.Use(logger.Log)

	// get routes without authentication
	rt.Get("/products", hd.GetAll())
	rt.Get("/products/{id}", hd.GetById())

	// rutes with authentication
	rt.Group(func(r chi.Router) {
		r.Use(auth.ValidateToken)
		r.Post("/products", hd.Create())
		r.Put("/products/{id}", hd.Update())
		r.Patch("/products/{id}", hd.UpdatePartial())
		r.Delete("/products/{id}", hd.Delete())
	})

	// run http server
	err = http.ListenAndServe(h.addr, rt)
	return
}
