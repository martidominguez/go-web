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

// this is used when you get a GET request to /products
/* func (ph *ProductsMap) GetProducts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]any{
			"message": "products found",
			"data":    ph.data,
		})
	}
} */

// this is used when you get a GET request to /products/{id}
/* func (ph *ProductsMap) GetProductById() http.HandlerFunc {
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
} */

/*
func (ph *ProductsMap) CreateProduct() http.HandlerFunc {
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
} */
