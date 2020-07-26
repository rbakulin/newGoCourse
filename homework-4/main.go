package main

import (
	"./character"
	"fmt"
)

func main() {
	knight := character.Warrior {
		Character: character.Character {
			Name: "Arthur",
			Strength: 9,
			Agility: 3,
			Magic: 0,
			Defence: 10,
		},
		SwordSkill: 9,
	}
	assassin := character.Robber {
		Character: character.Character {
			Name: "Altair",
			Strength: 3,
			Agility: 10,
			Magic: 0,
			Defence: 4,
		},
		ArcherySkill: 6,
	}
	witcher := character.Magician {
		Character: character.Character {
			Name: "Geralt",
			Strength: 3,
			Agility: 10,
			Magic: 0,
			Defence: 7,
		},
		SpellSkill: 4,
	}
	witcherAttacksKnight := character.Attack(witcher, knight)
	fmt.Printf("%s attacks %s: damage is %f\n", witcher.GetName(), knight.GetName(), witcherAttacksKnight)
	assassinAttacksWithcer := character.Attack(assassin, witcher)
	fmt.Printf("%s attacks %s: damage is %f\n", assassin.GetName(), witcher.GetName(), assassinAttacksWithcer)
}
