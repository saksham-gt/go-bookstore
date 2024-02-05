package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/saksham-gt/go-bookstore/pkg/models"
	"github.com/saksham-gt/go-bookstore/pkg/utils"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("content-type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookDetails, _ := models.GetBookById(ID)

	w.Header().Set("content-type", "pkglication/json")
	res, _ := json.Marshal(bookDetails)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	b := CreateBook.CreateBook()

	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]

	ID, err := strconv.ParseInt(bookId, 0, 0)

	if err != nil {
		panic("error parsing bookId")
	}

	book := models.DeleteBook(ID)
	w.WriteHeader(http.StatusOK)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-type", "pkglication/json")
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	var updateBook = &models.Book{}
	utils.ParseBody(r, updateBook)
	bookId := vars["bookId"]

	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		panic("error parsing bookId")
	}

	bookDetails, db := models.GetBookById(ID)
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}

	db.Save(bookDetails)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("content-type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
