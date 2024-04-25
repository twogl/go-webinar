package duel

import (
	"fmt"
)

type Attacker interface {
	Damage() int
}

type Duel struct {
	attackerA Attacker
	attackerB Attacker
}

// Constructor function
func NewDuel(pA, pB Attacker) *Duel {
	return &Duel{
		attackerA: pA,
		attackerB: pB,
	}
}

func (duel *Duel) Present() string {
	dmgA := duel.attackerA.Damage()
	dmgB := duel.attackerB.Damage()
	return fmt.Sprintf("=====\nPlayer A with %d DMG\nvs.\nPlayer B with %d DMG\n=====", dmgA, dmgB)
}
