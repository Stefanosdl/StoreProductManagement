package model

import (
	"encoding/json"
	"net/http"

	"gorm.io/gorm"
)

var db *gorm.DB

// CreateProduct creates a new product
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	db.Create(&product)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}
