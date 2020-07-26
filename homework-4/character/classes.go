package character

type Warrior struct {
	Character
	SwordSkill float64
}

func (w Warrior) CalcAttackPower() float64 {
	return (w.Strength + w.Agility) / 2 * (w.SwordSkill / 2)
}

type Robber struct {
	Character
	ArcherySkill float64
}

func (r Robber) CalcAttackPower() float64 {
	return (r.Strength + r.Agility) / 2 * (r.ArcherySkill / 2)
}

type Magician struct {
	Character
	SpellSkill float64
}

func (m Magician) CalcAttackPower() float64 {
	return m.Magic * (m.SpellSkill / 2)
}

type Player interface {
	CalcAttackPower() float64
	GetDefence() float64
	GetName() string
}

func Attack(p1 Player, p2 Player) float64 {
	damage := p1.CalcAttackPower() - p2.GetDefence()
	if damage < 0 {
		return 0
	}
	return damage
}
