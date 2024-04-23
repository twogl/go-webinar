package player

type Player struct {
	rawDamage    int
	weaponDamage int
	moveSpeed    int
}

func NewPlayer(rawDmg, weaponDmg, speed int) *Player {
	return &Player{
		rawDamage:    rawDmg,
		weaponDamage: weaponDmg,
		moveSpeed:    speed,
	}
}

func (plr *Player) Damage() int {
	return plr.rawDamage + plr.weaponDamage
}

func (plr *Player) Move() int {
	return plr.moveSpeed
}
