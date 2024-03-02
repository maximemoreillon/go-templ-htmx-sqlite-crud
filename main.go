package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Movie struct {
	ID int
	Title string
	Year int
}


func handleMoviesPost(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	r.ParseForm()
	title := r.Form.Get("title")
	year, _ := strconv.Atoi(r.Form.Get("year"))
	newMovie := createMovie(db, title, year)
	MovieComponent(newMovie).Render(r.Context(),w)
}

func handleMovieUpdate(w http.ResponseWriter, r *http.Request, db *sql.DB, id int) {
	r.ParseForm()
	title := r.Form.Get("title")
	year, _ := strconv.Atoi(r.Form.Get("year"))
	updateMovie(db, id, title, year)
}

func handleMovies(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request){
		idString := strings.TrimPrefix(r.URL.Path, "/movies/")
		// Not sure if this is the right way
		if idString == "" {
			switch r.Method {
			case "POST" :
				handleMoviesPost(w,r,db)
			case "GET" :
				movies := readMovies(db)
				MoviesPage(movies).Render(r.Context(),w)
			default:
				fmt.Fprintf(w,"Cannot %s /movies", r.Method)
			}
		} else {
			id, err := strconv.Atoi(idString)
			if err != nil {
				panic(err)
			}
			switch r.Method {
			case "GET" :
				movie := readMovie(db, id)
				MoviePage(movie).Render(r.Context(),w)
			case "PUT" :
				handleMovieUpdate(w, r, db, id)
			case "DELETE" :
				deleteMovie(db, id)
			default:
				fmt.Fprintf(w,"Cannot %s /movies", r.Method)
			}
		}
	})
}


func main () {

	db := openDb()

	createTableIfNotExists(db)
	mux := http.NewServeMux()
	mux.Handle("/movies/", handleMovies(db))
	http.ListenAndServe(":8080", mux)
}