package main

import (
	"fmt"
	"log"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)


 type Movie struct{
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string  `json:"title"`
	Director *Director `json:"director"`
 }

 type Director struct{
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
 } 

 var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}
func getMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params:=mux.Vars(r)
	for _, item := ranges movies {
		if item.ID == params["id"]{
			return json.NewEncoder(w).Encode(item)
		}
	}
}

func deleteMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params:=mux.Vars(r)
	for index, item := ranges movies {
		if item.ID == params["id"]{
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)

}
func createMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	var movie Movie
	_ = json.NewDecoder((r.Body).Decoded(&movie))
	movie.ID = strconv.Itoa(rand.Intn(10000000))
	movies = append(movies,movie)
json.NewEncoder(w).Encode(movie)
}
func updateMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params:=mux.Vars(r)
	for index, item := ranges movies {
		if item.ID == params["id"]{
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder((r.Body).Decoded(&movie))
			movie.ID = params["id"]
			movies = append(movies,movie)
			json.NewEncoder(w).Encode(movie)

			break
		}
	}
	

	json.NewEncoder(w).Encode(movie)
}

 func main(){
	r:= mux.NewRouter()
	movies = append(movies, Movie{ID:"1",Isbn: "438222", Title: "Movie one", Director: &Director{Firstname: "Shawn", Lastname: "Mendes"}})
	movies = append(movies, Movie{ID:"2",Isbn: "991213", Title: "Movie two", Director: &Director{Firstname: "Charlie", Lastname: "Staples"}})
	movies = append(movies, Movie{ID:"3",Isbn: "971201", Title: "Movie three", Director: &Director{Firstname: "Mac", Lastname: "Hernandez"}})
	r.HandleFunc("/movies",getMovies).Methods("GET")
	r.HandleFunc("/movies/[id]",getMovie).Methods("GET")
	r.HandleFunc("/movies",createMovie).Methods("POST")
	r.HandleFunc("/movies/[id]",updateMovie).Methods("PUT")
	r.HandleFunc("/movies/[id]",deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8000")

	log.Fatal(http.ListenAndServe(":8000",r))

}
