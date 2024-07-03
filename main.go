// Using structs and slices for simplicity 
// Not creating routers, controllers etc. for simplicity
package main

import(
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
	Title string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct{
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

var movies []Movie 

func main(){
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn:"438227", Title:"Movie One", Director : &Director{Firstname:"John", Lastname:"Doe"}})
	movies = append(movies, Movie{ID: "2", Isbn:"45455". Title:"Movie Two", Director : &Director{Firstname:"Cody", Lastname: "McKeon"}})
	
	r.HandleFunc("/movies", getMovies).Method("GET")
	r.HandleFunc("/movies/{id}", getMovie).Method("GET")
	r.HandleFunc("/movies", createMovie).Method("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Method("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Method("DELETE")

	fmt.Print("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}
