package products

import (
	"net/http"

	"github.com/anikmahidul9/ecomapi/internal/json"
)

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) ListProducts(w http.ResponseWriter, r *http.Request) {
	products := []Product{
		{ID: 1, Name: "Product 1", Price: 10.99},
		{ID: 2, Name: "Product 2", Price: 19.99},
		{ID: 3, Name: "Product 3", Price: 5.99},
	}
	json.Write(w, http.StatusOK,products)
}
