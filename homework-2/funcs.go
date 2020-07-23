package main

import (
	"fmt"
	"strconv"
)

func isEven(number int) bool{
	return number % 2 == 0
}

func noLeftByThree(number int) bool{
	return number % 3 == 0
}

func fibonacciNumbers()  {
	// use float to prevent int overflow
	prev := 0.0
	current := 1.0

	for i := 1; i <= 100; i++ {
		fmt.Printf("%f\n", current)
		oldCurrent := current
		current = prev + current
		prev = oldCurrent
	}

}

func main()  {
	var numberStr string
	fmt.Println("Введите число")
	fmt.Scanln(&numberStr)
	numberInt, err := strconv.Atoi(numberStr)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	if isEven(numberInt) {
		fmt.Printf("Число %s является четным\n", numberStr)
	} else {
		fmt.Printf("Число %s не является четным\n", numberStr)
	}
	if noLeftByThree(numberInt) {
		fmt.Printf("Число %s делится на 3 без отсатка\n", numberStr)
	} else {
		fmt.Printf("Число %s не делится на 3 без отсатка\n", numberStr)
	}

	fibonacciNumbers()
}
