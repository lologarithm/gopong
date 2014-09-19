package game

import (
	"encoding/json"
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
	Players  map[uint32]*Player
	Balls    []*Ball
}

func (m *Manager) Run() {
	lastUpdate := time.Now()
	timeout := lastUpdate.Add(time.Millisecond * SimUpdateSeconds).Sub(time.Now())
	doRun := true
	for doRun {
		select {
		case <-time.After(timeout):
			lastUpdate = time.Now()
			//m.Tick(true)
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
		m.Players[uint32(id)] = &Player{Id: uint32(id), Position: Vector2{X: 0, Y: 0}}
	case "playerUpdate":
		id, err := strconv.ParseInt(msg["id"], 10, 32)
		if err != nil {
			log.Printf("Failed to parse player id for update!")
		}

		err = json.Unmarshal([]byte(msg["pos"]), &m.Players[uint32(id)].Position)
		if err != nil {
			log.Printf("Failed to parse player position for update!")
		}
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
		Players: make(map[uint32]*Player, 0),
		Balls:   []*Ball{&Ball{Id: 0, Size: 10}},
	}
}
