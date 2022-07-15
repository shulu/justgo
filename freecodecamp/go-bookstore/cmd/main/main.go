package main

import (
	"go-bookstore/pkg/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {

	route := mux.NewRouter()
	routes.RegisterBookStoreRoutes(route)
	http.Handle("/", route)
	log.Fatal(http.ListenAndServe(":9090", route))
}
