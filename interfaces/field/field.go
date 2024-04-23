package field

import "fmt"

type Field struct {
	playerA Moveable
	playerB Moveable
}

func NewField(pA, pB Moveable) *Field {
	return &Field{
		playerA: pA,
		playerB: pB,
	}
}

func (fld *Field) MovePlayers() string {
	mvA := fld.playerA.Move()
	mvB := fld.playerB.Move()
	return fmt.Sprintf("=====\nPlayer A moves with speed %d \nand\nPlayer B moves with speed %d \n=====", mvA, mvB)
}
