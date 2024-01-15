package storage

import "app/internal"

type ProductStorageJSON interface {
	// WriteAll is a method that writes all products in the storage
	WriteAll(products []internal.Product) (err error)
	// ReadAll is a method that reads all products from the storage
	ReadAll() (products []internal.Product, err error)
}
