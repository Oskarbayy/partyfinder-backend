package products

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
)

type ProductHandler struct {
	svc ProductService
}

func NewProductHandler(svc ProductService) *ProductHandler {
	return &ProductHandler{svc: svc}
}

// UseCase
func (ph *ProductHandler) AddProduct(w http.ResponseWriter, r *http.Request) {
	var req addProductRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	p := Product{
		ID:       "0",
		Name:     req.Name,
		Price:    req.Price,
		Quantity: req.Quantity,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second*2))
	defer cancel()
	ph.svc.AddProduct(ctx, p)

	respP := addProductResponse{
		Name:     p.Name,
		Price:    p.Price,
		Quantity: p.Quantity,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(respP)
}
