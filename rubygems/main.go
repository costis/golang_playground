package main

import "database/sql"
import (
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

var (
	name string
)

func main() {
	db, err := sql.Open("postgres", "user=postgres dbname=rubygems sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT name from rubygems LIMIT 20")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	defer rows.Close()
	// or just rows.Close()?

	for rows.Next() {
		err := rows.Scan(&name)
		if err != nil {
			log.Fatal(nil)
		}

		fmt.Println(name)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

}
