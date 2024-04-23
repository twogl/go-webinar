package main

import (
	"fmt"

	"github.com/twogl/go-webinar/interfaces/duel"
	"github.com/twogl/go-webinar/interfaces/field"
	"github.com/twogl/go-webinar/interfaces/player"
)

func main() {
	pA := player.NewPlayer(15, 25, 2)
	pB := player.NewPlayer(12, 23, 3)
	duel := duel.NewDuel(pA, pB)

	fmt.Println(duel.Present())

	field := field.NewField(pA, pB)

	fmt.Println(field.MovePlayers())
}
