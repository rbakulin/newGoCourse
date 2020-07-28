package main

import (
	"./character"
	"./contact"
	"fmt"
	"sort"
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

	phoneBook := contact.PhoneBook {
		&contact.Contact {
			Name: "Zack",
			Phone: "39999999999",
			Email: "zack@gmail.com",
		},
		&contact.Contact {
			Name: "Karen",
			Phone: "27777777777",
			Email: "karen@gmail.com",
		},
		&contact.Contact {
			Name: "Aaron",
			Phone: "18888888888",
			Email: "aaron@gmail.com",
		},
	}
	sort.Sort(phoneBook)
	for _, cont := range phoneBook {
		fmt.Println(cont.Name)
	}
}
