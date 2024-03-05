package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
)

func handleMoviesPost(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request){
		r.ParseForm()
		title := r.Form.Get("title")
		year, _ := strconv.Atoi(r.Form.Get("year"))
		newMovie := createMovie(db, title, year)
		MovieComponent(newMovie).Render(r.Context(),w)
	})
}

func handleMoviesGet(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request){
		movies := readMovies(db)
		MoviesPage(movies).Render(r.Context(),w)
	})
}

func handleMovieGet(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request){

		idString := r.PathValue("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}

		movie := readMovie(db, id)
		MovieComponent(movie).Render(r.Context(),w)
	})
}




func handleMoviePut(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request){
		idString := r.PathValue("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}
		r.ParseForm()
		title := r.Form.Get("title")
		year, _ := strconv.Atoi(r.Form.Get("year"))
		updatedMovie := updateMovie(db, id, title, year)
		MovieComponent(updatedMovie).Render(r.Context(),w)
	})
}

func handleMovieDelete(db *sql.DB) http.HandlerFunc{
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request){
		idString := r.PathValue("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}
		deleteMovie(db, id)
		fmt.Fprint(w,idString)
	})
}




func handleMoviesEditGet(db *sql.DB) http.HandlerFunc{
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request){

		idString := r.PathValue("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}

		movie := readMovie(db, id)
		MovieEditComponent(movie).Render(r.Context(),w)
	})

}