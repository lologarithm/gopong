package game

type Manager struct {
	Players map[uint32]Player
	Balls   []Ball
}

func NewManager() *Manager {
	return &Manager{}
}

type Player struct {
	Id       uint32
	Position Vector2
}

type Ball struct {
	Id       uint32
	Position Vector2
}
