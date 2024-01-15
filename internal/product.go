package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Product struct {
	Id           int
	Name         string
	Quantity     int
	Code_value   string
	Is_published bool
	Expiration   string
	Price        float64
}

func LoadProducts() (products []Product) { // i need to change this for the storage, but for now i will use this function
	// open the file
	file, err := os.Open("/Users/martindoming/Desktop/Bootcamp/go-web/products.json")
	if err != nil {
		fmt.Println("Error abriendo el archivo:", err)
		return
	}
	defer file.Close()

	// read the file
	bytes, _ := io.ReadAll(file)

	// decode the json into the struct
	err = json.Unmarshal(bytes, &products)
	if err != nil {
		fmt.Println("Error decodificando el archivo JSON:", err)
		return nil
	}

	return
}
