package main

import "database/sql"
import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
	_ "github.com/lib/pq"
	"io"
	"encoding/json"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

const dbConnStr = `user=postgres dbname=rubygems sslmode=disable`
const sqlFetchGemsBatch = `fetch_gems.sql`
const sqlFetchGemDetail = `fetch_gem_detail.sql`

func loadTemplate() *template.Template {
	currentPath, err := os.Getwd()
	check(err)

	return template.Must(template.ParseFiles(filepath.Join(currentPath, sqlFetchGemsBatch)))
}

type Gem struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Gems []Gem

func (r *Gems) Load() {
	gems := fetchGems()
	*r = append(*r, gems...)
}

func (r *Gems) ToJSON(writer io.Writer) {
	jsonBytes, err := json.MarshalIndent(r, "", "  ")
	check(err)

	_, err = writer.Write(jsonBytes)
	check(err)
}

func fetchGemsBatch(startId int) ([]Gem, error) {
	db, err := sql.Open("postgres", dbConnStr)
	check(err)
	defer db.Close()

	tpl := loadTemplate()
	var sqlStr bytes.Buffer

	tpl.Execute(&sqlStr, startId)
	rows, err := db.Query(string(sqlStr.Bytes()))
	check(err)
	defer rows.Close()

	fmt.Println(string(sqlStr.Bytes()))

	var gems []Gem
	for rows.Next() {
		g := Gem{}

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

func fetchGems() []Gem {
	var gems []Gem

	pos := 0
	for {
		rows, err := fetchGemsBatch(pos)
		if err != nil {
			break
		}
		gems = append(gems, rows...)

		if len(rows) == 0 {
			break
		}

		pos = rows[len(rows)-1].Id
	}

	return gems
}

func main() {
	gems := make(Gems, 0)
	gems.Load()

	// dump to file
	f, err := os.Create("out.json")
	check(err)
	defer f.Close()
	gems.ToJSON(f)
}
