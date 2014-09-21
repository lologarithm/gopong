package server

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/lologarithm/gopong/game"
	"log"
	"net/http"
)

type Connection struct {
	// The websocket connection.
	ws *websocket.Conn

	// Buffered channel of outbound messages.
	send chan game.Message

	Addr string

	myHub *Hub
}

func (c *Connection) ReadMessagesFromNet() {
	myMessage := &game.Message{}
	for {
		_, message, err := c.ws.ReadMessage()
		if err != nil {
			break
		}

		json.Unmarshal(message, myMessage)
		c.myHub.Incoming <- *myMessage
	}
	c.ws.Close()
}

func (c *Connection) WriteMessagesToNet() {
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

func HandleWebSocket(w http.ResponseWriter, r *http.Request, hub *Hub) {
	ws, err := upgrader.Upgrade(w, r, nil)
	log.Printf("got new connection")
	if err != nil {
		log.Printf("ERROR: During handle upgrade: %s", err.Error())
		return
	}
	c := &Connection{send: make(chan game.Message, 256), ws: ws, myHub: hub}
	hub.NewConnection <- c
	defer func() { hub.CloseConnection <- c }()
	go c.WriteMessagesToNet()
	c.ReadMessagesFromNet()
}
