package handlers

import (
	"app/internal"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
)

// this is the handler for the products
type ProductsHandler struct {
	data   map[int]internal.Product
	lastID int
}

// this is an struct to represent the body of the request
type BodyRequestProductJSON struct {
	Name         string  `json:"name"`
	Quantity     int     `json:"quantity"`
	Code_value   string  `json:"code_value"`
	Is_published bool    `json:"is_published"`
	Expiration   string  `json:"expiration"`
	Price        float64 `json:"price"`
}

// this is an struct to represent the response of the request
type ResponseProductJSON struct {
	Id           int
	Name         string
	Quantity     int
	Code_value   string
	Is_published bool
	Expiration   string
	Price        float64
}

// this function is used to migrate all the products from the file to the map
func NewHandler() *ProductsHandler {
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

	return &ProductsHandler{
		data:   productMap,
		lastID: lastID,
	}
}

// this is used when you get a GET request to /ping
func (ph *ProductsHandler) Pong() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong"))
	}
}

// this is used when you get a GET request to /products
func (ph *ProductsHandler) GetProducts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]any{
			"message": "products found",
			"data":    ph.data,
		})
	}
}

// this is used when you get a GET request to /products/{id}
func (ph *ProductsHandler) GetProductById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idString := chi.URLParam(r, "id")

		id, err := strconv.Atoi(idString)
		if err != nil {
			w.Header().Set("Content-Type", "text/plain")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("error with the id"))
			return
		}

		product, ok := ph.data[id]
		if !ok {
			w.Header().Set("Content-Type", "text/plain")
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("product not found"))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]any{
			"message": "product found",
			"data":    product,
		})
	}
}

func (ph *ProductsHandler) CreateProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// decode the body of the request
		var body BodyRequestProductJSON
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			w.Header().Set("Content-Type", "text/plain")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("invalid body"))
			return
		}

		// create the product
		// autoincrement the lastID
		(*ph).lastID++
		// serialize the product
		product := internal.Product{
			Id:           (*ph).lastID,
			Name:         body.Name,
			Quantity:     body.Quantity,
			Code_value:   body.Code_value,
			Is_published: body.Is_published,
			Expiration:   body.Expiration,
			Price:        body.Price,
		}

		// validate the rules
		if err := (*ph).ValidateRules(&product); err != nil {
			fmt.Println(err.Error())

			w.Header().Set("Content-Type", "text/plain")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("invalid product"))
			return
		}

		// store the product in the map
		(*ph).data[product.Id] = product

		// response
		data := ResponseProductJSON{
			Id:           product.Id,
			Name:         product.Name,
			Quantity:     product.Quantity,
			Code_value:   product.Code_value,
			Is_published: product.Is_published,
			Expiration:   product.Expiration,
			Price:        product.Price,
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]any{
			"message": "product created",
			"data":    data,
		})
	}
}

func (ph *ProductsHandler) ValidateRules(product *internal.Product) error {
	// fields cannot be empty
	if product.Name == "" || product.Quantity == 0 || product.Code_value == "" || product.Expiration == "" || product.Price == 0.0 {
		return errors.New("fields cannot be empty")
	}

	// code value must be unique
	for _, p := range (*ph).data {
		if p.Code_value == product.Code_value {
			return errors.New("code value must be unique")
		}
	}

	// expiration must be a valid date
	_, err := time.Parse("02/01/2006", product.Expiration)
	if err != nil {
		return errors.New("expiration must be a valid date")
	}

	return nil
}
