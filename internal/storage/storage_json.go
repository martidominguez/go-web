package storage

import "app/internal"

type ProductStorageJSON interface { // TODO --> THE IMPLEMENTATION OF THIS FUNCTIONS
	// Write writes all products in the storage
	Write(products []internal.Product) (err error)
	// Read reads all products from the storage
	Read() (products []internal.Product, err error)
}
