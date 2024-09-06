package handlers

import (
	"micro/data"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// swagger:route DELETE /{id} products deleteProduct
// DeleteProducts deletes a product from database
// responses:
// 201: noContent

// DeleteProducts deletes a product from database
func (p *Products) DeleteProducts(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, _ := strconv.Atoi(vars["id"])

	p.l.Println("Handle DELETE Product", id)

	err := data.DeleteProduct(id)

	if err != data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
	}

	if err != nil {
		http.Error(rw, "Product not found", http.StatusBadRequest)
	}
}
