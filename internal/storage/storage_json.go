package storage

import (
	"app/internal"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// i don´t know how to connect this with the repository

type StorageProductJSON struct {
	// filepath is the path of the file where the products are stored
	filePath string
}

// NewStorageProductJSON creates a new instance of a storage product json
func NewStorageProductJSON(filePath string) *StorageProductJSON {
	return &StorageProductJSON{
		filePath: filePath}
}

// ReadAll is a method that reads all products from the storage
func (s *StorageProductJSON) ReadAll() (products []internal.Product, err error) {
	// open the file
	file, err := os.Open("/Users/martindoming/Desktop/Bootcamp/go-web/products.json")
	if err != nil {
		fmt.Println("cannot open the file: ", err)
		return
	}
	defer file.Close()

	// read the file
	bytes, _ := io.ReadAll(file)

	// decode the json into the struct
	err = json.Unmarshal(bytes, &products)
	if err != nil {
		fmt.Println("cannot decode the json file: ", err)
		return
	}

	return
}

// WriteAll is a method that writes all products in the storage
func (s *StorageProductJSON) WriteAll(products []internal.Product) (err error) { // i think this function must be associated with a struct but i don't know how to do it
	// open the file
	file, err := os.Open("/Users/martindoming/Desktop/Bootcamp/go-web/products.json")
	if err != nil {
		fmt.Println("cannot open the file: ", err)
		return
	}
	defer file.Close()

	// decode the struct into a json
	bytes, err := json.Marshal(products)
	if err != nil {
		fmt.Println("cannot encode the json file: ", err)
		return
	}

	// write the file
	_, err = file.Write(bytes)
	if err != nil {
		fmt.Println("cannot write the file: ", err)
		return
	}

	return nil
}
