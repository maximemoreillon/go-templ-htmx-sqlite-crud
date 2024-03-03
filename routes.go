package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func handleMoviesPost(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	r.ParseForm()
	title := r.Form.Get("title")
	year, _ := strconv.Atoi(r.Form.Get("year"))
	newMovie := createMovie(db, title, year)
	MovieComponent(newMovie).Render(r.Context(),w)
}

func handleMoviesGet(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	movies := readMovies(db)
	MoviesPage(movies).Render(r.Context(),w)
}

func handleMovieGet(w http.ResponseWriter, r *http.Request, db *sql.DB, id int) {
	movie := readMovie(db, id)
	MoviePage(movie).Render(r.Context(),w)
}

func handleMoviePut(w http.ResponseWriter, r *http.Request, db *sql.DB, id int) {
	r.ParseForm()
	title := r.Form.Get("title")
	year, _ := strconv.Atoi(r.Form.Get("year"))
	updateMovie(db, id, title, year)
}

func handleMovieDelete(w http.ResponseWriter, r *http.Request, db *sql.DB, id int){
	deleteMovie(db, id)
}

func handleMovies(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request){
		// Not sure if this is the right way
		// Will be fixed with Go 1.22 using /movies/{id}
		
		idString := strings.TrimPrefix(r.URL.Path, "/movies/")
		if idString == "" {
			switch r.Method {
			case http.MethodPost :
				handleMoviesPost(w,r,db)
			case http.MethodGet :
				handleMoviesGet(w,r,db)
			default:
				fmt.Fprintf(w,"Cannot %s /movies", r.Method)
			}
		} else {
			id, err := strconv.Atoi(idString)
			if err != nil {
				panic(err)
			}
			switch r.Method {
			case http.MethodGet :
				handleMovieGet(w,r,db,id)
			case http.MethodPut :
				handleMoviePut(w, r, db, id)
			case http.MethodDelete :
				handleMovieDelete(w,r,db,id)
				
			default:
				fmt.Fprintf(w,"Cannot %s /movies", r.Method)
			}
		}
	})
}
