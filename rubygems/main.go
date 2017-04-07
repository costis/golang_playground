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

type rubygem struct {
	Name string
}

func main() {
	db, err := sql.Open("postgres", "user=postgres dbname=rubygems sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	stmt, err := db.Prepare("SELECT name from rubygems WHERE id > $1 LIMIT 10")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(100)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	// do we need this? or just db.Close() is enough?

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
