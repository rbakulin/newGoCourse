package main

import (
	"fmt"
)

const exchangeRate = 70

func convertToDollars(rubles int) int{
	return rubles * exchangeRate
}

func main()  {
	fmt.Println("Введите сумму в рублях")
	var rubles int
	_, err := fmt.Scanln(&rubles)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	fmt.Printf("Вы получите %d$", convertToDollars(rubles))
}
