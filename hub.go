package main

import (
	"log"
)

type Hub struct {
	// Registered connections.
	connections map[*connection]bool

	// Inbound messages from the connections.
	broadcast chan UserMessage

	// Outgoing messages to the connections.
	outgoing chan UserMessage

	// Register requests from the connections.
	register chan *connection

	// Unregister requests from connections.
	unregister chan *connection
}

type Message struct {
	Id   int32
	Data string
}

type UserMessage struct {
	Id      string
	Message []byte
}

var h = Hub{
	broadcast:   make(chan UserMessage),
	outgoing:   make(chan UserMessage),
	register:    make(chan *connection),
	unregister:  make(chan *connection),
	connections: make(map[*connection]bool),
}

var g = game.Manager{}

func (h *Hub) run() {
	go broadcastMessages()

	for {
		select {
		case c := <-h.register:
			h.connections[c] = true
			processMessage(UserMessage {})
		case c := <-h.unregister:
			if _, ok := h.connections[c]; ok {
				delete(h.connections, c)
				close(c.send)
			}
		case m := <-h.broadcast:
			log.Print(m)
			resp := g.ProcessMessage(m)			
		}
	}
}

func () broadcastMessages {
	for c := range h.outgoing {
		select {
		case c.send <- resp:
		default:
			delete(h.outgoing, c)
			close(c.send)
		}
	}
}
