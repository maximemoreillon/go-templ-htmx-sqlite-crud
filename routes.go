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
	MovieComponent(movie).Render(r.Context(),w)
}

func handleMoviePut(w http.ResponseWriter, r *http.Request, db *sql.DB, id int) {
	r.ParseForm()
	title := r.Form.Get("title")
	year, _ := strconv.Atoi(r.Form.Get("year"))
	updatedMovie := updateMovie(db, id, title, year)
	MovieComponent(updatedMovie).Render(r.Context(),w)
}

func handleMovieDelete(w http.ResponseWriter, r *http.Request, db *sql.DB, id int){
	deleteMovie(db, id)
}

func handleMovieEditGet(w http.ResponseWriter, r *http.Request, db *sql.DB, id int){
	movie := readMovie(db, id)
	MovieEditComponent(movie).Render(r.Context(),w)
}

func handleMovies(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request){
		// Would be improved with Go 1.22 using /movies/{id}

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
		} else if strings.Contains(idString, "/edit") {
			idString = strings.TrimSuffix(idString, "/edit")
			id, err := strconv.Atoi(idString)
			if err != nil {
				panic(err)
			}
			
			switch r.Method {
			case http.MethodGet :
				handleMovieEditGet(w,r,db,id)
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
