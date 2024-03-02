package main

import (
	"net/http"
)

type Movie struct {
	ID int
	Title string
	Year int
}




func main () {

	db := openDb()

	createTableIfNotExists(db)
	mux := http.NewServeMux()
	mux.Handle("/movies/", handleMovies(db))
	http.ListenAndServe(":8080", mux)
}