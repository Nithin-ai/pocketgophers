package main

import (
	"fmt"
	"log"
	"net/http"
	"shoppingcartcrud/controllers"
	"shoppingcartcrud/database"

	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/products/add", controllers.CreateOneProduct).Methods("POST")
	router.HandleFunc("/products/{id}", controllers.GetOneProduct).Methods("GET")
	router.HandleFunc("/products", controllers.GetAllProducts).Methods("GET")
	router.HandleFunc("/products/delete/{id}", controllers.DeleteOneProduct).Methods("DELETE")
	//router.HandleFunc("products/updatecount/{id}", controllers.UpdateCount).Methods("PUT")
	router.HandleFunc("/products/update/{id}", controllers.UpdateOneProduct).Methods("PUT")

}

func main() {
	LoadConfig()
	database.DbConnection(AppConfig.ConnectionString)

	router := mux.NewRouter().StrictSlash(true)
	RegisterRoutes(router)

	log.Println(fmt.Sprintf("Server starting at port %s", AppConfig.Port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", AppConfig.Port), router))

}
