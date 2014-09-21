package server

import (
	"github.com/lologarithm/gopong/game"
)

type Hub struct {
	Incoming        chan game.Message      // Inbound messages from the connections.
	Outgoing        chan game.Message      // Outgoing messages to the connections.
	NewConnection   chan *Connection       // Register requests from the connections.
	CloseConnection chan *Connection       // Unregister requests from connections.
	connections     map[string]*Connection // Maps connection address to connection object.
	gameManager     *game.Manager          // Game manager to send messages to
}

//var g = game.Manager{}

func (h *Hub) Run() {
	go h.gameManager.Run()
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
	h := &Hub{
		Incoming:        make(chan game.Message),
		Outgoing:        make(chan game.Message),
		NewConnection:   make(chan *Connection),
		CloseConnection: make(chan *Connection),
		connections:     make(map[string]*Connection),
		gameManager:     game.NewManager(),
	}
	// Now link together the channels so the game gets all the messages
	h.gameManager.Outgoing = h.Outgoing
	h.gameManager.Incoming = h.Incoming
	return h
}
