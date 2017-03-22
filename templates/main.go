package main

import (
	"text/template"
	"os"
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

	tmpl := template.Must(template.New("header").Parse("the header {{ . }}"))
	tmpl.New("footer").Parse("the footer")
	tmpl.New("body").Parse(`First comes {{ template "header" . }}, then follows the body,
	and then comes {{ template "footer" }}  `)

	tmpl.ExecuteTemplate(os.Stdout, "body", people)
}
