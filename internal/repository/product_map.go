package repository

import (
	"app/internal"
)

// in this file, i only handle what is related directly to the products

// this is the handler for the products
type ProductsMap struct {
	data   map[int]internal.Product
	lastID int
}

// NewProductsMap returns a new ProductsMap instance, filled with the products from the json file
func NewProductsMap() *ProductsMap {
	productsInfo := internal.LoadProducts()

	// convert the slice into a map
	productMap := make(map[int]internal.Product)
	for _, p := range productsInfo {
		productMap[p.Id] = p
	}

	lastID := 0
	for _, p := range productsInfo {
		if p.Id > lastID {
			lastID = p.Id
		}
	}

	return &ProductsMap{
		data:   productMap,
		lastID: lastID,
	}
}

func (ph *ProductsMap) GetAll() (products []internal.Product) {
	for _, p := range ph.data {
		products = append(products, p)
	}
	return
}

func (ph *ProductsMap) GetById(id int) (product internal.Product, err error) {
	product, ok := ph.data[id]
	if !ok {
		err = internal.ErrProductNotFound
	}
	return
}

func (ph *ProductsMap) Create(product *internal.Product) (err error) {
	// code value must be unique
	for _, p := range (*ph).data {
		if p.Code_value == product.Code_value {
			return internal.ErrRepeatedCode
		}
	}

	ph.lastID++
	product.Id = ph.lastID
	ph.data[product.Id] = *product
	return
}

func (ph *ProductsMap) Update(product *internal.Product) (err error) {
	_, ok := ph.data[product.Id]
	if !ok {
		return internal.ErrProductNotFound
	}

	ph.data[product.Id] = *product
	return
}

func (ph *ProductsMap) Delete(id int) (err error) {
	_, ok := ph.data[id]
	if !ok {
		return internal.ErrProductNotFound
	}

	delete(ph.data, id)
	return
}
