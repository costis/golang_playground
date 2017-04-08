package main

import "database/sql"
import (
	"encoding/json"
	_ "github.com/lib/pq"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Rubygem struct {
	Id   int `json:"id"`
	Name string `json:"name"`
}

var gems = make([]Rubygem, 10)

func main() {
	db, err := sql.Open("postgres", "user=postgres dbname=rubygems sslmode=disable")
	check(err)

	defer db.Close()

	stmt, err := db.Prepare("SELECT id, name FROM rubygems WHERE id > $1 LIMIT 100")
	check(err)
	defer stmt.Close()

	rows, err := stmt.Query(100)
	check(err)
	defer rows.Close()

	for rows.Next() {
		g := Rubygem{}

		err := rows.Scan(&g.Id, &g.Name)
		check(err)

		gems = append(gems, g)
	}

	err = rows.Err()
	check(err)

	f, err := os.Create("gems.json")
	check(err)

	jsonBytes, err := json.Marshal(gems)
	check(err)

	f.WriteString(string(jsonBytes))
	f.Close()
}
