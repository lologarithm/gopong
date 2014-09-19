package game

type Player struct {
	Id       uint32
	Position Vector2
	Size     float64
}

type Ball struct {
	Id       uint32
	Position Vector2
	Velocity Vector2
	Size     float64
}

func (b *Ball) UpdatePosition() {
	b.Position.X += b.Velocity.X
	b.Position.Y += b.Velocity.Y
}
