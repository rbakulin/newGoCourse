package character

type Character struct {
	Name     string
	Strength float64
	Agility  float64
	Magic    float64
	Defence  float64
}

func (c Character) GetDefence() float64  {
	return c.Defence
}

func (c Character) GetName() string  {
	return c.Name
}
