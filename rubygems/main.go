package main

import "database/sql"
import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"

	_ "github.com/lib/pq"
	//"time"
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

const sqlTemplateFile = "fetch_gems.sql"

func loadTemplate() *template.Template {
	currentPath, err := os.Getwd()
	check(err)

	return template.Must(template.ParseFiles(filepath.Join(currentPath, sqlTemplateFile)))
}

func fetchGemsBatch(startId int) ([]Rubygem, error) {
	db, err := sql.Open("postgres", "user=postgres dbname=rubygems sslmode=disable")
	check(err)
	defer db.Close()

	tpl := loadTemplate()
	var sqlStr bytes.Buffer

	tpl.Execute(&sqlStr, startId)
	rows, err := db.Query(string(sqlStr.Bytes()))
	check(err)
	defer rows.Close()

	fmt.Println(string(sqlStr.Bytes()))

	var gems []Rubygem
	for rows.Next() {
		g := Rubygem{}

		err := rows.Scan(&g.Id, &g.Name)
		check(err)

		gems = append(gems, g)
	}
	if rows.Err() != nil {
		if rows.Err() == sql.ErrNoRows {
			log.Print(rows.Err())
			return nil, sql.ErrNoRows
		} else {
			panic(rows.Err())
		}
	}

	return gems, nil
}

func fetchGems() []Rubygem {
	var gems []Rubygem

	pos := 0
	for {
		rows, err := fetchGemsBatch(pos)
		if err != nil {
			break
		}
		gems = append(gems, rows...)

		// TODO: break if err != nil, not by checking row size.
		if len(rows) == 0 {
			break
		}

		pos = rows[len(rows)-1].Id
		fmt.Printf("The is id %d\n", pos)

		//if rows[len(rows)-1].Id > 20000 {
		//	break
		//}
	}

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
