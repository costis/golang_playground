package main

import (
	"flag"
	"log"
	"net/http"
	"sync"
	"text/template"
)

type TemplateHandler struct {
	once     sync.Once
	filename string
	tmpl     *template.Template
}

func (tHandler *TemplateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tHandler.once.Do(func() {
		tHandler.tmpl = template.Must(template.ParseFiles(tHandler.filename))
	})

	tHandler.tmpl.Execute(w, r)
}

func newRoom() *room {
	return &room{
		forward: make(chan []byte),
		join:    make(chan *client),
		leave:   make(chan *client),
		clients: make(map[*client]bool),
	}
}

func main() {
	var addr = flag.String("addr", "localhost:8080", "The address of the app")
	flag.Parse()

	r := newRoom()
	http.Handle("/", &TemplateHandler{filename: "templates/chat.html"})
	http.Handle("/room", r)
	go r.run()

	log.Println("Starting web server at address:", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("Fuckedup", err)
	}
}

//func main() {
//	http.Handle("/", &TemplateHandler{filename: "templates/chat.html"})
//
//	err2 := http.ListenAndServe("localhost:8080", nil)
//	if err2 != nil {
//		log.Fatal(err2)
//	}
//}
