package controllers

import (
	"encoding/json"
	"net/http"
	"shoppingcartcrud/database"
	"shoppingcartcrud/models"
)

func CreateOneProduct(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-type","application/json")
	var prod models.Product
	//prod = models.Product{"xdxdftc", "dcrfc", "fcdtcf", 90.0, "gvvt", 7, "cffct"}
	json.NewDecoder(r.Body).Decode(&prod)
	database.DB.Create(&prod)
	json.NewEncoder(w).Encode(prod)
}

func GetOneProduct(w http.ResponseWriter, r *http.Request) {

}

func GetAllProducts(w http.ResponseWriter, r *http.Request) {

}

func DeleteOneProduct(w http.ResponseWriter, r *http.Request) {

}

func UpdateCount(w http.ResponseWriter, r *http.Request) {

}

func UpdateOneProduct(w http.ResponseWriter, r *http.Request) {

}
