package character

type Warrior struct {
	Character
	SwordSkill float64
	className string
}

const (
	warriorClassName = "Warrior"
	robberClassName = "Robber"
	magicianClassName = "Magician"
)

// Use constructor to set default class name
func NewWarrior(name string, strength, agility, magic, defence, swordSkill float64) *Warrior {
	w := new(Warrior)
	w.Name = name
	w.Strength = strength
	w.Agility = agility
	w.Magic = magic
	w.Defence = defence
	w.SwordSkill = swordSkill
	w.className = warriorClassName
	return w
}

func (w Warrior) CalcAttackPower() float64 {
	return (w.Strength + w.Agility) / 2 * (w.SwordSkill / 2)
}

func (w Warrior) GetClassName() string {
	return w.className
}

type Robber struct {
	Character
	ArcherySkill float64
	className string
}

func NewRobber(name string, strength, agility, magic, defence, archerySkill float64) *Robber {
	r := new(Robber)
	r.Name = name
	r.Strength = strength
	r.Agility = agility
	r.Magic = magic
	r.Defence = defence
	r.ArcherySkill = archerySkill
	r.className = robberClassName
	return r
}

func (r Robber) CalcAttackPower() float64 {
	return (r.Strength + r.Agility) / 2 * (r.ArcherySkill / 2)
}

func (r Robber) GetClassName() string {
	return r.className
}

type Magician struct {
	Character
	SpellSkill float64
	className string
}

func NewMagician(name string, strength, agility, magic, defence, spellSkill float64) *Magician {
	m := new(Magician)
	m.Name = name
	m.Strength = strength
	m.Agility = agility
	m.Magic = magic
	m.Defence = defence
	m.SpellSkill = spellSkill
	m.className = magicianClassName
	return m
}

func (m Magician) CalcAttackPower() float64 {
	return m.Magic * (m.SpellSkill / 2)
}

func (m Magician) GetClassName() string {
	return m.className
}

type Player interface {
	CalcAttackPower() float64
	GetDefence() float64
	GetName() string
	GetClassName() string
}

type Characters []Player

func Attack(p1 Player, p2 Player) float64 {
	damage := p1.CalcAttackPower() - p2.GetDefence()
	if damage < 0 {
		return 0
	}
	return damage
}
