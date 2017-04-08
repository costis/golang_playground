package main

import "database/sql"
import (
	"encoding/json"
	_ "github.com/lib/pq"
	"os"
	"fmt"
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


func fetchGems()([]Rubygem) {
	var gems = make([]Rubygem, 10)

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

	return gems
}


func main() {
	gems := fetchGems()

	jsonBytes, err := json.Marshal(gems)
	cnt, err := saveJSON(jsonBytes)
	check(err)

	fmt.Printf("Written %d bytes\n", cnt)
}

func saveJSON(b []byte)(cnt int, er error) {
	f, err := os.Create("out.json")
	check(err)
	defer f.Close()

	cnt, e := f.Write(b)
	if e != nil {
		return cnt, e
	}

	return cnt, nil
}
