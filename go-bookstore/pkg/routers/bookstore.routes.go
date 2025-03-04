package routers

import "github.com/gorilla/mux"

import "github.com/ramalloc/go-bookstore/pkg/controllers"

var RegisterBookstoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBookById).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBookById).Methods("DELETE")
	// router.HandleFunc("/book/delete-all", controllers.DeleteBookById).Methods("DELETE")
}
