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

	mux.Handle("GET /movies/", handleMoviesGet(db))
	mux.Handle("POST /movies/", handleMoviesPost(db))
	mux.Handle("GET /movies/{id}", handleMovieGet(db))
	mux.Handle("PUT /movies/{id}", handleMoviePut(db))
	mux.Handle("DELETE /movies/{id}", handleMovieDelete(db))
	mux.Handle("GET /movies/{id}/edit", handleMoviesEditGet(db))
	
	http.ListenAndServe(":8081", mux)
}