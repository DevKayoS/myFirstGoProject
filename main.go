package main

import "database/sql"

func main() {
}

func connectToDb() (*sql.DB, error) {
	db, err := sql.Open("", "")
	if err != nil {
	}

	defer db.Close()
	return db, nil
}
