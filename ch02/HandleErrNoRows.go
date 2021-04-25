package main

import (
	"fmt"
	"database/sql"
	"log"
	"errors"
)
// var ErrNoRows = error.New("sql: no rows in result set.")

func main() {
	db, err := sql.Open("driver-name", "database=test1")
	if err != nil {
		log.Fatal(err)
	}
	id := 1
	err := db.QueryRow("select id from Person").Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			//handler this err
			return errors.WithMessage(err, "empty result")

		}
		return err
	}
}