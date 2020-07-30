package main

import (
	"github.com/rbakulin/newGoCourse/homework-4/character"
	"github.com/rbakulin/newGoCourse/homework-4/contact"
)
import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	// Phone Book
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

	// Characters
	knight := character.NewWarrior("Arthur", 9, 3, 0, 10, 9)
	assassin := character.NewRobber("Altair", 3, 10, 0, 4, 6)
	witcher := character.NewMagician("Geralt", 7, 7, 3, 5, 4)
	chars := character.Characters{
		knight,
		assassin,
		witcher,
	}

	input := ""
	for {
		fmt.Println("There are characters:")
		for _, char := range chars {
			fmt.Println(char.GetName())
		}
		fmt.Println("Type character names to make them fight!")
		fmt.Println("Input character1 name:")
		fmt.Print("> ")
		if _, err := fmt.Scanln(&input); err != nil {
			fmt.Println(err)
			continue
		}
		if input == "exit" {
			break
		}
		if input == "help" {
			for _, char := range chars {
				fmt.Printf("Character %s:\nclass: %s\nattack: %f\ndefence: %f\n--------------\n",
					char.GetName(), char.GetClassName(), char.CalcAttackPower(), char.GetDefence())
			}
		}
		var char1, char2 character.Player
		for _, char := range chars {
			if strings.ToLower(char.GetName()) == strings.ToLower(input) {
				char1 = char
			}
		}
		if char1 == nil {
			fmt.Printf("No such character: %s\n", input)
			continue
		}

		input = ""
		fmt.Println("Input character2 name:")
		fmt.Print("> ")
		if _, err := fmt.Scanln(&input); err != nil {
			fmt.Println(err)
			continue
		}
		for _, char := range chars {
			if strings.ToLower(char.GetName()) == strings.ToLower(input) {
				char2 = char
			}
		}
		if char2 == nil {
			fmt.Printf("No such character: %s\n", input)
			continue
		}
		damage := character.Attack(char1, char2)
		fmt.Printf("%s attacks %s: damage is %f\n", char1.GetName(), char2.GetName(), damage)
		break
	}
}
