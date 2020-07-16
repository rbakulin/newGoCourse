package main

import (
	"fmt"
	"strconv"
)

var exchangeRate = 70

func convertToDollars(rubles int) int{
	return rubles * exchangeRate
}

func main()  {
	var rublesStr string
	fmt.Println("Введите сумму в рублях")
	fmt.Scanln(&rublesStr)
	rublesInt, err := strconv.Atoi(rublesStr)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	fmt.Printf("Вы получите %d $", convertToDollars(rublesInt))
}
