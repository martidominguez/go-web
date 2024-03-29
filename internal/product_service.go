package internal

import "errors"

// errors
var (
	ErrInvalidExpiration = errors.New("expiration must be a valid date")
	ErrFieldsEmpty       = errors.New("fields cannot be empty")
)

// MovieService is an interface that represents a movie service
// - business logic
// - validation
// - external services (e.g. apis, databases, etc.)
type ProductService interface {
	// GetAll gets all movies
	GetAll() (products []Product)
	// GetById gets a movie by id
	GetById(id int) (product Product, err error)
	// Create creates a product
	Create(product *Product) (err error)
	// Update updates a product
	Update(product *Product) (err error)
	// Delete deletes a product
	Delete(id int) (err error)
}
