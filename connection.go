package main

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"net/http"
)

type connection struct {
	// The websocket connection.
	ws *websocket.Conn

	// Buffered channel of outbound messages.
	send chan Message
}

func (c *connection) reader() {
	for {
		_, message, err := c.ws.ReadMessage()
		if err != nil {
			break
		}
		h.broadcast <- UserMessage{Message: message, Id: c.ws.RemoteAddr().String()}
	}
	c.ws.Close()
}

func (c *connection) writer() {
	for message := range c.send {
		jsonMsg, _ := json.Marshal(message)

		err := c.ws.WriteMessage(websocket.TextMessage, jsonMsg)
		if err != nil {
			break
		}
	}
	c.ws.Close()
}

var upgrader = &websocket.Upgrader{ReadBufferSize: 1024, WriteBufferSize: 1024}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	c := &connection{send: make(chan Message, 256), ws: ws}
	h.register <- c
	defer func() { h.unregister <- c }()
	go c.writer()
	c.reader()
}
