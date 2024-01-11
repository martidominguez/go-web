package service

import (
	"app/internal"
	"time"
)

// in this file i only handle what is related with the rules

// NewProductDefault creates a new instance of a product service
func NewProductDefault(rp internal.ProductRepository) *MovieDefault {
	return &MovieDefault{
		rp: rp,
	}
}

// MovieDefault is a struct that represents the default implementation of a movie service
type MovieDefault struct {
	// rp is a movie repository
	rp internal.ProductRepository
	// external services
	// ... (weather api, etc.)
}

func (d *MovieDefault) GetAll() (products []internal.Product) {
	return d.rp.GetAll()
}

func (d *MovieDefault) GetById(id int) (product internal.Product, err error) {
	product, err = (*d).rp.GetById(id)
	return
}

func (d *MovieDefault) Create(product *internal.Product) (err error) {
	// here i must validate the product
	// fields cannot be empty
	if product.Name == "" || product.Quantity == 0 || product.Code_value == "" || product.Expiration == "" || product.Price == 0.0 {
		return internal.ErrFieldsEmpty
	}

	// expiration must be a valid date
	_, err = time.Parse("02/01/2006", product.Expiration)
	if err != nil {
		return internal.ErrInvalidExpiration
	}

	// here i must call the repository to create the product
	err = (*d).rp.Create(product)
	return
}
