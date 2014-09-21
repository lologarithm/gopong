package game

import (
	"encoding/json"
	"log"
	"math"
	"time"
)

const (
	SimUpdatesPerSecond = 50.0
	SimUpdateSeconds    = 1000.0 / SimUpdatesPerSecond
	SimUpdateInverse    = 1 / SimUpdatesPerSecond
)

type Manager struct {
	Incoming chan Message
	Outgoing chan Message
	Players  map[int]*Thing // For now its just left and right players
	Balls    []*Thing
}

func (m *Manager) Run() {
	lastUpdate := time.Now()
	timeout := lastUpdate.Add(time.Millisecond * SimUpdateSeconds).Sub(time.Now())
	doRun := true
	for doRun {
		select {
		case <-time.After(timeout):
			lastUpdate = time.Now()
			m.Tick()
			timeout = lastUpdate.Add(time.Millisecond * SimUpdateSeconds).Sub(time.Now())
		case msg := <-m.Incoming:
			m.handleMessage(msg)
		}
	}
}

func (m *Manager) handleMessage(msg Message) {
	log.Printf("Got message: %s", msg.Type)
	switch msg.Type {
	case "newPlayer":
		log.Printf("Adding player.")
		id := len(m.Players)
		pos := Vector2{X: -100, Y: 0}
		if id == 1 {
			pos = Vector2{X: 100, Y: 0}
			// Make the balls start moving!
			m.Balls[0].Velocity.X = 1.0
			log.Printf("STARTING GAME")
		}
		m.Players[id] = &Thing{Id: id, Position: pos}
		m.Outgoing <- msg
	case "playerUpdate":
		log.Printf("Player updated direction!")
		var pu PlayerUpdate
		err := json.Unmarshal(msg.Data, &pu)
		if err != nil {
			log.Printf("Failed to parse player update!")
		}
		m.Players[pu.Id].Velocity.Y = 1.0 * pu.Direction
		m.Outgoing <- msg
	}
}

func (m *Manager) Tick() {
	var dist float64
	for _, b := range m.Balls {
		b.UpdatePosition()
		for _, other := range m.Players {
			dist = math.Pow(float64(other.Position.X-b.Position.X), 2) + math.Pow(float64(other.Position.Y-b.Position.Y), 2)
			if dist <= math.Pow(float64(other.Size+b.Size), 2) {
				b.Velocity.X *= -1
				log.Printf("COLLISION!")
				break
			}
		}
		if b.Position.X > 100 {
			log.Printf("Player 1 has won!")
			b.Position.X = 0
		} else if b.Position.X < -100 {
			log.Printf("Player 2 has won!")
			b.Position.X = 0
		}
	}
	for _, p := range m.Players {
		p.UpdatePosition()
		//for _, other := range m.Balls {
		//	dist = math.Pow(float64(other.Position.X-p.Position.X), 2) + math.Pow(float64(other.Position.Y-p.Position.Y), 2)
		//	if dist <= math.Pow(float64(other.Size+p.Size), 2) {
		//		other.Velocity.X *= -1
		//		log.Printf("COLLISION!")
		//		break
		//	}
		//}
	}
}

func NewManager() *Manager {
	return &Manager{
		Players:  make(map[int]*Thing, 0),
		Balls:    []*Thing{&Thing{Id: 0, Size: 10}},
		Incoming: make(chan Message, 10),
	}
}

type Message struct {
	Id   string
	Type string
	Data []byte
}

type PlayerUpdate struct {
	Id        int
	Direction float64
}
