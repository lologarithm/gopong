package server

import (
	"github.com/lologarithm/gopong/game"
	"log"
)

type Hub struct {
	connections     map[string]*Connection // Maps connection address to connection object.
	Incoming        chan NetMessage        // Inbound messages from the connections.
	Outgoing        chan NetMessage        // Outgoing messages to the connections.
	NewConnection   chan *Connection       // Register requests from the connections.
	CloseConnection chan *Connection       // Unregister requests from connections.
	gameManager     *game.Manager
}

type NetMessage struct {
	Id   string
	Data []byte
}

//var g = game.Manager{}

func (h *Hub) Run() {
	go h.broadcastMessages()

	for {
		select {
		case c := <-h.NewConnection:
			c.Addr = c.ws.RemoteAddr().String()
			h.connections[c.Addr] = c
		case c := <-h.CloseConnection:
			if _, ok := h.connections[c.Addr]; ok {
				delete(h.connections, c.Addr)
				close(c.send)
			}
		case m := <-h.Incoming:
			log.Print(m)
		}
	}
}

func (h *Hub) broadcastMessages() {
	for {
		select {
		case msg := <-h.Outgoing:
			for _, c := range h.connections {
				c.send <- msg
			}
		}
	}
}

func NewHub() *Hub {
	return &Hub{
		Incoming:        make(chan NetMessage),
		Outgoing:        make(chan NetMessage),
		NewConnection:   make(chan *Connection),
		CloseConnection: make(chan *Connection),
		connections:     make(map[string]*Connection),
		gameManager:     game.NewManager(),
	}

}
