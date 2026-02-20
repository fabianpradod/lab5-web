package main

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

func main() {
	db, _ := sql.Open("sqlite", "file:tv-shows.db")
	defer db.Close()

	db.Exec("CREATE TABLE IF NOT EXISTS shows (id INTEGER PRIMARY KEY AUTOINCREMENT, title TEXT UNIQUE, current_episode TEXT)")

	shows := [][2]string{
		{"Breaking Bad", "S01E01"},
		{"The Office", "S02E03"},
		{"Stranger Things", "S01E02"},
	}

	for _, s := range shows {
		db.Exec("INSERT INTO shows (title, current_episode) VALUES (?, ?)", s[0], s[1])
	}

	var id int
	var title string
	var current_episode string

	query, _ := db.Query("SELECT id, title, current_episode FROM shows ORDER BY id")
	defer query.Close()

	for query.Next() {
		query.Scan(&id, &title, &current_episode)
		fmt.Printf("%d: %s (%s)\n", id, title, current_episode)
	}
}
