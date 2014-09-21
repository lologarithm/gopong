package game

type Thing struct {
	Id       int
	Position Vector2
	Velocity Vector2
	Size     float64
}

//type Ball struct {
//	Id       uint32
//	Position Vector2
//	Velocity Vector2
//	Size     float64
//}

func (t *Thing) UpdatePosition() {
	t.Position.X += t.Velocity.X
	t.Position.Y += t.Velocity.Y
}
