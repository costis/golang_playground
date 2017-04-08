package main

import "database/sql"
import (
	"bytes"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"html/template"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Rubygem struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func fetchGems() []Rubygem {
	var gems = make([]Rubygem, 10)

	tpl := template.Must(template.ParseFiles("fetch_gems.sql"))
	var sql_str bytes.Buffer
	tpl.Execute(&sql_str, nil)

	db, err := sql.Open("postgres", "user=postgres dbname=rubygems sslmode=disable")
	check(err)
	defer db.Close()

	stmt, err := db.Prepare(string(sql_str.Bytes()))
	check(err)
	defer stmt.Close()

	rows, err := stmt.Query()
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

func saveJSON(b []byte) (cnt int, er error) {
	f, err := os.Create("out.json")
	check(err)
	defer f.Close()

	cnt, e := f.Write(b)
	if e != nil {
		return cnt, e
	}

	return cnt, nil
}
