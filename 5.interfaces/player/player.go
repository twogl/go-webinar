package player

type Player struct {
	rawDamage    int
	weaponDamage int
	moveSpeed    int
}

func (plr *Player) Damage() int {
	return plr.rawDamage + plr.weaponDamage
}

func (plr *Player) Move() int {
	return plr.moveSpeed
}
