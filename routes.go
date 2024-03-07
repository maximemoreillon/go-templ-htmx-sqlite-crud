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
		newMovie, err := createMovie(db, title, year)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		MovieComponent(newMovie).Render(r.Context(),w)
	})
}

func handleMoviesGet(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request){
		movies, err := readMovies(db)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		MoviesPage(movies).Render(r.Context(),w)
	})
}

func handleMovieGet(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request){

		idString := r.PathValue("id")
		id, err := strconv.Atoi(idString)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
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
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		r.ParseForm()
		title := r.Form.Get("title")
		year, _ := strconv.Atoi(r.Form.Get("year"))
		updatedMovie, err := updateMovie(db, id, title, year)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		MovieComponent(updatedMovie).Render(r.Context(),w)
	})
}

func handleMovieDelete(db *sql.DB) http.HandlerFunc{
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request){
		idString := r.PathValue("id")
		id, err := strconv.Atoi(idString)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = deleteMovie(db, id)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		fmt.Fprint(w,idString)
	})
}




func handleMoviesEditGet(db *sql.DB) http.HandlerFunc{
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request){

		idString := r.PathValue("id")
		id, err := strconv.Atoi(idString)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		movie := readMovie(db, id)
		MovieEditComponent(movie).Render(r.Context(),w)
	})

}