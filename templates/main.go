package main

import (
	"text/template"
	"os"
	"strings"
	"fmt"
)

type person struct {
	Name string
	Age  int
}

func main() {
	people := []person{
		{"john", 30},
		{"gianni", 10},
	}

	helpers := template.FuncMap{
		"up": func(v string) string { return strings.ToUpper(v) },
	}

	tmpl := template.Must(template.New("header").Parse("the header {{ . }}"))
	tmpl.Funcs(helpers)

	tmpl.New("footer").Parse("the footer")
	tmpl, err := tmpl.New("body").Parse(`
	First comes {{ template "header" . }},

	then follows the body with the names:
	{{ range . }}
		{{ up "hi" }}
	{{ end }}

	and then comes {{ template "footer" }} `)

	if err != nil {
		fmt.Println(err)
	}

	tmpl.New("person").Parse(`The person's name is {{ .Name }}'`)

	tmpl.ExecuteTemplate(os.Stdout, "body", people)
}
