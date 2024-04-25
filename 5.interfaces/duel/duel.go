package duel

type Attacker interface {
	Damage() int
}

type Duel struct {
	attackerA Attacker
	attackerB Attacker
}
