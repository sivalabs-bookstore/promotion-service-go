package utils

import "database/sql"

func getDB() (db *sql.DB) {
	// Realize the connection with mysql driver
	db, err := sql.Open("mysql", "test:test@tcp(localhost:3306)/test?charset=utf8&parseTime=True&loc=Local")

	// If error stop the application
	if err != nil {
		panic(err.Error())
	}
	// Return db object to be used by other functions
	return db
}
