package handler

import (
	"app/internal"
	"app/platform/web/request"
	"app/platform/web/response"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// in this file i only handle what is related to the request and response

// NewDefaultProducts returns a new DefaultMovies instance
func NewDefaultProducts(sv internal.ProductService) *DefaultProducts {
	return &DefaultProducts{
		sv: sv,
	}
}

// DefaultProducts is an implementation with handlers for the Products storage
type DefaultProducts struct {
	// sv is a movie service
	sv internal.ProductService
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

func (d *DefaultProducts) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		products := (*d).sv.GetAll()

		// response
		response.JSON(w, http.StatusOK, map[string]any{"message": "products found", "data": products})
	}
}

func (d *DefaultProducts) GetById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		idString := chi.URLParam(r, "id")

		id, err := strconv.Atoi(idString)
		if err != nil {
			response.Error(w, http.StatusInternalServerError, "error with the id")
			return
		}

		product, err := (*d).sv.GetById(id)
		if err != nil {
			response.Error(w, http.StatusNotFound, "product not found")
			return
		}

		response.JSON(w, http.StatusOK, map[string]any{"message": "product found", "data": product})
	}
}

func (d *DefaultProducts) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body BodyRequestProductJSON

		if err := request.JSON(r, &body); err != nil {
			response.Error(w, http.StatusBadRequest, "invalid body")
			return
		}

		product := internal.Product{
			Name:         body.Name,
			Quantity:     body.Quantity,
			Code_value:   body.Code_value,
			Is_published: body.Is_published,
			Expiration:   body.Expiration,
			Price:        body.Price,
		}

		if err := d.sv.Create(&product); err != nil {
			switch {
			case errors.Is(err, internal.ErrFieldsEmpty):
				response.Error(w, http.StatusBadRequest, "fields cannot be empty")
			case errors.Is(err, internal.ErrInvalidExpiration):
				response.Error(w, http.StatusBadRequest, "expiration must be a valid date")
			case errors.Is(err, internal.ErrRepeatedCode):
				response.Error(w, http.StatusBadRequest, "code value must be unique")
			default:
				response.Error(w, http.StatusInternalServerError, "internal server error")
			}
			return
		}

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

		response.JSON(w, http.StatusCreated, map[string]any{"message": "product created", "data": data})
	}
}

// validate the rules

/* func (ph *ProductsHandler) ValidateRules(product *internal.Product) error {
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
} */
