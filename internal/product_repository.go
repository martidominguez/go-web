package internal

import "errors"

// errors
var (
	ErrRepeatedCode    = errors.New("code value must be unique")
	ErrProductNotFound = errors.New("product not found")
)

// product repository is an interface that defines the methods that the repository must implement
type ProductRepository interface { // maybe i should change this functions later
	// GetAll returns the list of products from the repository
	GetAll() (products []Product)
	// GetById returns a product by id from the repository
	GetById(id int) (product Product, err error)
	// Create creates a product in the repository
	Create(product *Product) (err error)
	// Update updates a product in the repository
	Update(product *Product) (err error)
	// Delete delete a product from the repository
	Delete(id int) (err error)
}
