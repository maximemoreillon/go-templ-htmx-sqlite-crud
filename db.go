package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func openDb () *sql.DB {
	const file string = "movies.db"
	db, err := sql.Open("sqlite3", file)
	
	if err != nil {
		panic(err)
	}
	return db
}

func createTableIfNotExists (db *sql.DB) {
	const create string = `
		CREATE TABLE IF NOT EXISTS "movies" (
		"id"	INTEGER,
		"title"	TEXT NOT NULL,
		"year"	INTEGER,
		PRIMARY KEY("id" AUTOINCREMENT)
	);`
	_, err := db.Exec(create);
	if err != nil {
		panic(err)
	}
}