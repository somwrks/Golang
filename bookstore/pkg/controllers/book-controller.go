package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/somwrks/Golang/bookstore/pkg/models"
	"github.com/somwrks/golang/bookstore/pkg/models"
	"github.com/somwrks/golang/bookstore/pkg/utils"
)

var NewBook models.Book

func GetBook (w http.Request, r *http.Response){
	newBooks:=models.GetAllBooks()
	res , _ :=json.Marshal(newBooks)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriterHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0,0)
	if err !=nil {
		fmt.Println("error while parsing")
	}
	bookDetails, _ :=models.GetBookById(ID)
}