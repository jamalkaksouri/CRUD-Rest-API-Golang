package controllers

import (
	"encoding/json"
	"golang-crud-rest-api/database"
	"golang-crud-rest-api/models"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	product := models.Product{}
	json.NewDecoder(r.Body).Decode(&product)
	database.Instance.Create(&product)
	json.NewEncoder(w).Encode(product)
}

func checkIfProductExists(productId string) bool {
	product := models.Product{}
	database.Instance.First(&product, productId)

	return product.Id != 0
}

func GetProductById(w http.ResponseWriter, r *http.Request) {
	productId := mux.Vars(r)["id"]
	if !checkIfProductExists(productId) {
		json.NewEncoder(w).Encode("Product Not Found!")
		return
	}

	var product models.Product
	database.Instance.First(&product, productId)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	products := []models.Product{}
	database.Instance.Find(&products)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	productId := mux.Vars(r)["id"]
	if !checkIfProductExists(productId) {
		json.NewEncoder(w).Encode("Product Not Found!")
		return
	}

	product := models.Product{}
	database.Instance.First(&product, productId)
	json.NewDecoder(r.Body).Decode(&product)
	database.Instance.Save(&product)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	productId := mux.Vars(r)["id"]
	w.Header().Set("Content-Type", "application/json")
	if !checkIfProductExists(productId) {
		json.NewEncoder(w).Encode("Product Not Found!")
		return
	}

	product := models.Product{}
	database.Instance.Delete(&product, productId)
	json.NewEncoder(w).Encode("Product Deleted Successfuly!")
}
