package main

import (
	"github.com/costis/golang_playground/gpbp2/chat/trace"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"os"
)

const (
	socketBufferSize  = 1024
	messageBufferSize = 256
)

type room struct {
	forward chan []byte

	join  chan *client
	leave chan *client

	clients map[*client]bool
	tracer  trace.Tracer
}

func (r *room) run() {
	for {
		select {
		case client := <-r.join:
			r.clients[client] = true
			r.tracer.Trace("Client joined.")
		case client := <-r.leave:
			delete(r.clients, client)
			close(client.send)
			r.tracer.Trace("Client left.")
		case msg := <-r.forward:
			for client := range r.clients {
				client.send <- msg
			}
		}
	}
}

var upgrader = websocket.Upgrader{ReadBufferSize: socketBufferSize, WriteBufferSize: socketBufferSize}

func (r *room) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	socket, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Println("HTTPServe error:", err)
		return
	}

	client := &client{
		socket: socket,
		room:   r,
		send:   make(chan []byte, messageBufferSize),
	}

	r.join <- client
	defer func() { r.leave <- client }()

	go client.write()

	client.read()
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

