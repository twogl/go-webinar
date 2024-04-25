package field

type Movable interface {
	Move() int
}

type Field struct {
	movableA Movable
	movableB Movable
}
