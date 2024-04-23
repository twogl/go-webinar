package duel

import (
	"fmt"
)

type Duel struct {
	playerA Attacker
	playerB Attacker
}

func NewDuel(pA, pB Attacker) *Duel {
	return &Duel{
		playerA: pA,
		playerB: pB,
	}
}

func (duel *Duel) Present() string {
	dmgA := duel.playerA.Damage()
	dmgB := duel.playerB.Damage()
	return fmt.Sprintf("=====\nPlayer A with %d DMG\nvs.\nPlayer B with %d DMG\n=====", dmgA, dmgB)
}
