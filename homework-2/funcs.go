package main

import (
	"fmt"
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
	var num int
	fmt.Println("Введите число")
	_, err := fmt.Scanln(&num)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	if isEven(num) {
		fmt.Printf("Число %d является четным\n", num)
	} else {
		fmt.Printf("Число %d не является четным\n", num)
	}
	if noLeftByThree(num) {
		fmt.Printf("Число %d делится на 3 без отсатка\n", num)
	} else {
		fmt.Printf("Число %d не делится на 3 без отсатка\n", num)
	}

	fibonacciNumbers()
}
