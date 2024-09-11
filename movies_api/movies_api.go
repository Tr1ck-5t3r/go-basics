package movies_api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

type Movie struct {
	id     string  `json:"id"`
	isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Year   int     `json:"year"`
	Rating float64 `json:"rating"`
}

type Director struct {
	Name   string   `json:"name"`
	Movies []string `json:"movies"`
}

var movies = []Movie{
	{Title: "The Shawshank Redemption", Year: 1994, Rating: 9.3},
	{Title: "Pulp Fiction", Year: 1994, Rating: 8.9},
	{Title: "The Godfather", Year: 1972, Rating: 9.2},
	{Title: "The Dark Knight", Year: 2008, Rating: 9.0},
}

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range movies {
		if item.Title == params["title"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Movie{})
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.Rating = rand.Float64() * 10
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}
