package main

import (
	"flag"
	"github.com/costis/golang_playground/gpbp2/chat/trace"
	"log"
	"net/http"
	"os"
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

func newRoom(withTrace bool) *room {
	var t trace.Tracer

	if withTrace {
		t = trace.New(os.Stdout)

	} else {
		t = trace.Off()
	}

	return &room{
		forward: make(chan []byte),
		join:    make(chan *client),
		leave:   make(chan *client),
		clients: make(map[*client]bool),
		tracer:  t,
	}
}

func main() {
	var addr = flag.String("addr", "localhost:8080", "The address of the app")
	var withTrace = flag.Bool("tracing", false, "Enable tracing")
	flag.Parse()

	r := newRoom(*withTrace)

	tHandler := &TemplateHandler{filename: "templates/chat.html"}
	tLogin := &TemplateHandler{filename: "templates/login.html"}

	http.Handle("/", MustAuth(tHandler))
	http.Handle("/login", tLogin)
	http.Handle("/room", r)
	http.HandleFunc("/auth/", loginHandler)

	go r.run()

	log.Println("Starting web server at address:", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("Fuckedup", err)
	}
}
