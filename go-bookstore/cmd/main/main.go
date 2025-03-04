package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "gorm.io/driver/mysql"
	"github.com/ramalloc/go-bookstore/pkg/routers"
)

func main() {
	fmt.Println("Main...")
	r := mux.NewRouter()
	routers.RegisterBookstoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":9010", r))
}
