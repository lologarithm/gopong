package game

import (
	"log"
	"math"
	"strconv"
	"time"
)

const (
	SimUpdatesPerSecond = 50.0
	SimUpdateSeconds    = 1000.0 / SimUpdatesPerSecond
	SimUpdateInverse    = 1 / SimUpdatesPerSecond
)

type Manager struct {
	Incoming chan map[string]string
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

func (m *Manager) handleMessage(msg map[string]string) {
	switch msg["type"] {
	case "newPlayer":
		//id, err := strconv.ParseInt(msg["id"], 10, 32)
		id := len(m.Players)
		pos := Vector2{X: -100, Y: 0}
		if id == 1 {
			pos = Vector2{X: 100, Y: 0}
		}
		m.Players[id] = &Thing{Id: id, Position: pos}
	case "playerUpdate":
		id, err := strconv.Atoi(msg["id"])
		if err != nil {
			log.Printf("Failed to parse player id for update!")
		}

		moveDirection, err := strconv.ParseFloat(msg["moveDir"], 64)
		if err != nil {
			log.Printf("Failed to parse player position for update!")
		}
		m.Players[id].Velocity.Y = 25 * moveDirection
	}
}

func (m *Manager) Tick() {
	var dist float64
	for _, b := range m.Balls {
		b.UpdatePosition()
		for _, other := range m.Players {
			dist = math.Pow(float64(other.Position.X-b.Position.X), 2) + math.Pow(float64(other.Position.Y-b.Position.Y), 2)
			if dist <= math.Pow(float64(other.Size+b.Size), 2) {
				log.Printf("COLLISION!")
				break
			}
		}

	}
}

func NewManager() *Manager {
	return &Manager{
		Players: make(map[int]*Thing, 0),
		Balls:   []*Thing{&Thing{Id: 0, Size: 10}},
	}
}
