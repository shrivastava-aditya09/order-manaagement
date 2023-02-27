package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/aditya-0902/order-management/pkg/models"
	"github.com/aditya-0902/order-management/pkg/utils"
	"github.com/gorilla/mux"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["id"]
	fmt.Println("the bookid is")
	fmt.Println(vars)
	fmt.Println(bookId)
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error occured while parsing")
	}
	bookDetails, _ := models.GetBooksById(ID)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
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
	bookId := vars["id"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing bookid")
	}
	book := models.DeleteBook(ID)
	res, _ := json.Marshal(book)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "pkglication/json")
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var UpdateBook = &models.Book{}
	utils.ParseBody(r, UpdateBook)
	vars := mux.Vars(r)
	bookId := vars["id"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing bookid")
	}
	bookDetails, db := models.GetBooksById(ID)
	if UpdateBook.Name != bookDetails.Name {
		UpdateBook.Name = bookDetails.Name
	}
	if UpdateBook.Author != bookDetails.Author {
		UpdateBook.Author = bookDetails.Author
	}
	if UpdateBook.Publication != bookDetails.Publication {
		UpdateBook.Publication = bookDetails.Publication
	}
	db.Save(&bookDetails)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
