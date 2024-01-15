package handler_test

import (
	"app/internal"
	"app/internal/handler"
	"app/internal/repository"
	"app/internal/service"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/require"
)

func NewRequest(method, url string, body io.Reader, urlParams map[string]string, urlQuery map[string]string) (req *http.Request) {
	// old request
	req = httptest.NewRequest(method, url, body)

	// new request
	// - url params
	if urlParams != nil {
		chiKey := chi.NewRouteContext()
		for key, value := range urlParams {
			chiKey.URLParams.Add(key, value)
		}
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiKey))
	}
	// - url query
	if urlQuery != nil {
		query := req.URL.Query()
		for key, value := range urlQuery {
			query.Add(key, value)
		}
		req.URL.RawQuery = query.Encode()
	}

	// - content-type for json requests
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	return
}

func ConvertToJSON(t *testing.T, v interface{}) (b []byte) {
	b, err := json.Marshal(v)
	require.NoError(t, err)
	return
}

func TestProductsDefault_GetById(t *testing.T) {
	t.Run("success 01 - should return a product", func(t *testing.T) {
		// arrange
		// - repository
		db := []internal.Product{
			{
				Id:           1,
				Name:         "product 1",
				Quantity:     10,
				Code_value:   "123",
				Is_published: true,
				Expiration:   "14/05/2024",
				Price:        100,
			},
		}
		rp := repository.NewProductsMap(db)
		// - service
		sv := service.NewProductDefault(rp)
		// - handler
		hd := handler.NewDefaultProducts(sv)
		// - handler function
		getById := hd.GetById()

		// act
		req := NewRequest("GET", "/products/1", nil, map[string]string{"id": "1"}, nil)
		res := httptest.NewRecorder()
		getById(res, req)

		// assert
		productExpected := internal.Product{
			Id:           1,
			Name:         "product 1",
			Quantity:     10,
			Code_value:   "123",
			Is_published: true,
			Expiration:   "14/05/2024",
			Price:        100,
		}
		expectedCode := http.StatusOK
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		expectedBody := fmt.Sprintf(`{"message":"product found","data":%s}`, ConvertToJSON(t, productExpected))
		require.Equal(t, expectedCode, res.Code)
		require.Equal(t, expectedHeader, res.Header())
		require.JSONEq(t, expectedBody, res.Body.String())
	})

	t.Run("failure 01 - product not found", func(t *testing.T) {
		// arrange
		// - repository
		rp := repository.NewProductsMap(nil)
		// - service
		sv := service.NewProductDefault(rp)
		// - handler
		hd := handler.NewDefaultProducts(sv)
		// - handler function
		getById := hd.GetById()

		// act
		req := NewRequest("GET", "/products/1", nil, map[string]string{"id": "1"}, nil)
		res := httptest.NewRecorder()
		getById(res, req)

		// assert
		expectedCode := http.StatusNotFound
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		expectedBody := fmt.Sprintf(`{"status":"Not Found","message":"product not found"}`)
		require.Equal(t, expectedCode, res.Code)
		require.Equal(t, expectedHeader, res.Header())
		require.JSONEq(t, expectedBody, res.Body.String())
	})
}

func TestProductsDefault_GetAll(t *testing.T) {
	t.Run("success 01 - should return a list of products", func(t *testing.T) {
		// arrange
		// - repository
		db := []internal.Product{
			{
				Id:           1,
				Name:         "product 1",
				Quantity:     10,
				Code_value:   "123",
				Is_published: true,
				Expiration:   "14/05/2024",
				Price:        100,
			},
			{
				Id:           2,
				Name:         "product 2",
				Quantity:     20,
				Code_value:   "456",
				Is_published: true,
				Expiration:   "14/05/2024",
				Price:        200,
			},
		}
		rp := repository.NewProductsMap(db)
		// - service
		sv := service.NewProductDefault(rp)
		// - handler
		hd := handler.NewDefaultProducts(sv)
		// - handler function
		getAll := hd.GetAll()

		// act
		req := NewRequest("GET", "/products", nil, nil, nil)
		res := httptest.NewRecorder()
		getAll(res, req)

		// assert
		productsExpected := []internal.Product{
			{
				Id:           1,
				Name:         "product 1",
				Quantity:     10,
				Code_value:   "123",
				Is_published: true,
				Expiration:   "14/05/2024",
				Price:        100,
			},
			{
				Id:           2,
				Name:         "product 2",
				Quantity:     20,
				Code_value:   "456",
				Is_published: true,
				Expiration:   "14/05/2024",
				Price:        200,
			},
		}
		expectedCode := http.StatusOK
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		expectedBody := fmt.Sprintf(`{"message":"products found","data":%s}`, ConvertToJSON(t, productsExpected))
		require.Equal(t, expectedCode, res.Code)
		require.Equal(t, expectedHeader, res.Header())
		require.JSONEq(t, expectedBody, res.Body.String())
	})
}

func TestProductsDefault_Post(t *testing.T) {
	t.Run("success 01 - should create a product", func(t *testing.T) {
		// arrange
		// - repository
		rp := repository.NewProductsMap([]internal.Product{})
		// - service
		sv := service.NewProductDefault(rp)
		// - handler
		hd := handler.NewDefaultProducts(sv)
		// - handler function
		create := hd.Create()

		// act
		body := `{"name": "product 1", "quantity": 1, "code_value": "code1", "is_published": true, "expiration": "14/05/2024", "price": 1.1}`
		bodyReader := strings.NewReader(body)
		req := NewRequest("POST", "/products", bodyReader, nil, nil)
		res := httptest.NewRecorder()
		create(res, req)

		// assert
		expectedCode := http.StatusCreated
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		expectedBody := fmt.Sprintf(`{"data":{"Id":1,"Name":"product 1","Quantity":1,"Code_value":"code1","Is_published":true,"Expiration":"14/05/2024","Price":1.1},"message":"product created"}`)
		require.Equal(t, expectedCode, res.Code)
		require.Equal(t, expectedHeader, res.Header())
		require.JSONEq(t, expectedBody, res.Body.String())
	})
}

func TestProductsDefault_Delete(t *testing.T) {
	t.Run("success 01 - should delete a product", func(t *testing.T) {
		// arrange
		// - db
		db := []internal.Product{
			{
				Id:           1,
				Name:         "product 1",
				Quantity:     10,
				Code_value:   "123",
				Is_published: true,
				Expiration:   "14/05/2024",
				Price:        100,
			},
		}
		// - repository
		rp := repository.NewProductsMap(db)
		// - service
		sv := service.NewProductDefault(rp)
		// - handler
		hd := handler.NewDefaultProducts(sv)
		// - handler function
		delete := hd.Delete()

		// act
		req := NewRequest("GET", "/products/1", nil, map[string]string{"id": "1"}, nil)
		res := httptest.NewRecorder()
		delete(res, req)

		// assert
		expectedCode := http.StatusNoContent
		expectedBody := ""
		expectedHeaders := http.Header{}
		require.Equal(t, expectedCode, res.Code)
		require.Equal(t, expectedBody, res.Body.String())
		require.Equal(t, expectedHeaders, res.Header())
	})
}
