package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/ramalloc/go-bookstore/pkg/models"
	"github.com/ramalloc/go-bookstore/pkg/utils"
)

var NewBook models.Book

func GetBooks(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// json.NewEncoder(w).Encode(res);
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookId := params["bookId"]
	ID, err := strconv.Atoi(strings.TrimSpace(bookId))
	if err != nil {
		fmt.Println("Error while converting id string into integer")
	}
	bookDetails, _ := models.GetBookById(int64(ID))
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// json.NewEncoder(w).Encode(res);
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	createdBook := &models.Book{}
	utils.ParseBody(r, createdBook)
	b := createdBook.CreateBook()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// json.NewEncoder(w).Encode(res);
	w.Write(res)
}

func DeleteBookById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookId := params["bookId"]
	ID, err := strconv.Atoi(strings.TrimSpace(bookId))
	if err != nil {
		fmt.Println("Error while converting id string into integer")
	}
	deletedBook := models.DeleteBookById(int64(ID))
	res, _ := json.Marshal(deletedBook)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// json.NewEncoder(w).Encode(res);
	w.Write(res)
}

func UpdateBookById(w http.ResponseWriter, r *http.Request){
	var updateBook = &models.Book{}
	utils.ParseBody(r, updateBook)
	params := mux.Vars(r)
	bookId := params["bookId"]
	ID, err := strconv.Atoi(strings.TrimSpace(bookId))
	if err != nil {
		fmt.Println("Error while converting id string into integer")
	}
	fetchedBook, db := models.GetBookById(int64(ID))
	if updateBook.Name != ""{
		fetchedBook.Name = updateBook.Name
	}
	if updateBook.Author != ""{
		fetchedBook.Author = updateBook.Author
	}
	if updateBook.Publiction != ""{
		fetchedBook.Publiction = updateBook.Publiction
	}
	db.Save(&fetchedBook)
	res, _ := json.Marshal(fetchedBook)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// json.NewEncoder(w).Encode(res);
	w.Write(res)

}
