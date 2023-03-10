package controllers

import (
	"encoding/json"
	"net/http"
	"shoppingcartcrud/database"
	"shoppingcartcrud/models"

	"github.com/gorilla/mux"
)

func CreateOneProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var prod models.Product
	//prod = models.Product{"xdxdftc", "dcrfc", "fcdtcf", 90.0, "gvvt", 7, "cffct"}

	json.NewDecoder(r.Body).Decode(&prod)
	//prod.ProductID = base64.StdEncoding.EncodeToString([]byte("Apple iPhone 14"))
	//t := time.Now()
	//prod.CreatedAt = t.String()

	database.Db.Create(&prod)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(prod)
}

func checkIfProductExists(prodId string) bool {
	var prod models.Product
	database.Db.First(&prod, prodId)
	return prod.ProductID == 0
}

func GetOneProduct(w http.ResponseWriter, r *http.Request) {
	prodId := mux.Vars(r)["id"]

	if checkIfProductExists(prodId) == false {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Product Not Found")
		return
	}

	var prod models.Product
	database.Db.First(&prod, prodId)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(prod)
}

func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	var products []models.Product
	database.Db.Find(&products)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}

func DeleteOneProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	prodId := mux.Vars(r)["id"]

	if checkIfProductExists(prodId) == false {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Product Not Found")
		return
	}

	var prod models.Product
	database.Db.Delete(&prod, prodId)
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode("Product Deleted successfully")

}

/*func UpdateCount(w http.ResponseWriter, r *http.Request) {
    prodId:= mux.Vars(r)["id"]

	var prod models.Product
	database.Db.First(&prod,prodId)
	if prod.ProductID==""{
		json.NewEncoder(w).Encode("Product Not Found")
	}

	var reqprod models.Product
	json.NewDecoder(r.Body).Decode(&reqprod)
}*/

func UpdateOneProduct(w http.ResponseWriter, r *http.Request) {
	prodId := mux.Vars(r)["id"]

	if checkIfProductExists(prodId) == false {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Product Not Found")
		return
	}

	var prod models.Product
	database.Db.First(&prod, prodId)
	json.NewDecoder(r.Body).Decode(&prod)
	database.Db.Save(&prod)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(prod)
}
