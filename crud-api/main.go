package main

import (
	"fmt"
	"log"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gorilla/mux@latest"
)


 type Movie struct{
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string  `json:"title"`
	Director *Director `json:"director"`
 }