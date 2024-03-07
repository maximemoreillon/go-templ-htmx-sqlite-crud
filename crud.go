package main

import (
	"database/sql"
)



func createMovie (db *sql.DB, title string, year int) (Movie, error) {
	newMovie := Movie{0, title, year}
	res, err := db.Exec("INSERT INTO movies VALUES(null,?,?);", newMovie.Title, newMovie.Year)
	if err != nil {
		return newMovie, err
	}
	id, err := res.LastInsertId()
	return readMovie(db, int(id)), err
}


func readMovies (db *sql.DB) ([]Movie, error) {
	movies := []Movie {}
	rows, err := db.Query("SELECT * FROM movies ORDER BY id DESC LIMIT 100")

	if err != nil {
		return []Movie{}, err
	}
	
	for rows.Next() {
		movie := Movie{}
		err = rows.Scan(&movie.ID, &movie.Title, &movie.Year)
		// TODO: break out of the loop with error
		if err != nil {
			break
		}
		movies = append(movies, movie)
	}
	return movies, err
}

func readMovie (db *sql.DB, id int) Movie {
	foundMovie := Movie{}
	row:= db.QueryRow("SELECT * FROM movies WHERE id=?", id)
	row.Scan(&foundMovie.ID, &foundMovie.Title, &foundMovie.Year)
	return foundMovie
}





func updateMovie (db *sql.DB, id int, newTitle string, newYear int) (Movie, error) {
	update := `
	UPDATE movies 
	SET title=?, year=? 
	WHERE id=?`

	_, err := db.Exec(update, newTitle, newYear, id)

	return readMovie(db,id), err
}

func deleteMovie  (db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM movies WHERE id=?;", id)
	return err
}