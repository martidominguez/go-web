package application

import (
	"app/internal/handler"
	"app/internal/repository"
	"app/internal/service"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// NewDefaultHTTP creates a new instance of a default http server
func NewDefaultHTTP(addr string) *DefaultHTTP {
	// default config / values
	// ...

	return &DefaultHTTP{
		addr: addr,
	}
}

// DefaultHTTP is a struct that represents the default implementation of a http server
type DefaultHTTP struct {
	// addr is the address of the http server
	addr string
}

// Run runs the http server
func (h *DefaultHTTP) Run() (err error) {
	// initialize dependencies
	// - repository
	rp := repository.NewProductsMap()
	// - service
	sv := service.NewProductDefault(rp)
	// - handler
	hd := handler.NewDefaultProducts(sv)
	// - router
	rt := chi.NewRouter()

	//   endpoints
	rt.Get("/products", hd.GetAll())
	rt.Get("/products/{id}", hd.GetById())
	rt.Post("/products", hd.Create())
	rt.Put("/products/{id}", hd.Update())
	rt.Patch("/products/{id}", hd.UpdatePartial())
	rt.Delete("/products/{id}", hd.Delete())

	// run http server
	err = http.ListenAndServe(h.addr, rt)
	return
}
