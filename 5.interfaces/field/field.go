package field

import "fmt"

type Movable interface {
	Move() int
}

type Field struct {
	movableA Movable
	movableB Movable
}

// Constructor function
func NewField(pA, pB Movable) *Field {
	return &Field{
		movableA: pA,
		movableB: pB,
	}
}

func (fld *Field) MovePlayers() string {
	mvA := fld.movableA.Move()
	mvB := fld.movableB.Move()
	return fmt.Sprintf("=====\nPlayer A moves with speed %d \nand\nPlayer B moves with speed %d \n=====", mvA, mvB)
}
