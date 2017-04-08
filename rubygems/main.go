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

func fetchGemsBatch(startId int) ([]Rubygem) {
	db, err := sql.Open("postgres", "user=postgres dbname=rubygems sslmode=disable")
	check(err)
	defer db.Close()

	tpl := template.Must(template.ParseFiles("fetch_gems.sql"))
	var sql_str bytes.Buffer

	var gems = make([]Rubygem, 10)
	tpl.Execute(&sql_str, startId)
	rows, err := db.Query(string(sql_str.Bytes()))
	check(err)
	defer rows.Close()

	for rows.Next() {
		g := Rubygem{}

		err := rows.Scan(&g.Id, &g.Name)
		check(err)

		gems = append(gems, g)
	}
	check(rows.Err())

	return gems
}

func fetchGems() []Rubygem {
	var gems = make([]Rubygem, 10)
	gems = append(gems, fetchGemsBatch(0)...)

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
