package main

import (
	"fmt"
	"math"
	"strconv"
)

func calcHypotenuse(cathetus1 float64, cathetus2 float64) float64 {
	hypotenuse := math.Sqrt(math.Pow(cathetus1, 2) + math.Pow(cathetus2, 2))
	return hypotenuse
}

func calcSquare(cathetus1 float64, cathetus2 float64) float64 {
	square := cathetus1 * cathetus2 / 2
	return square
}

func calcPerimeter(cathetus1 float64, cathetus2 float64, hypotenuse float64) float64 {
	perimeter := cathetus1 + cathetus2 + hypotenuse
	return perimeter
}

func checkErr(err error){
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
}

func main()  {
	var cathetus1Str, cathetus2Str string

	fmt.Println("Введите катет 1")
	_, scanErr := fmt.Scanln(&cathetus1Str)
	checkErr(scanErr)
	fmt.Println("Введите катет 2")
	_, scanErr = fmt.Scanln(&cathetus2Str)
	checkErr(scanErr)

	cathetus1Int, parseErr := strconv.ParseFloat(cathetus1Str, 64)
	checkErr(parseErr)
	cathetus2Int, parseErr := strconv.ParseFloat(cathetus2Str, 64)
	checkErr(parseErr)

	hypotenuse := calcHypotenuse(cathetus1Int, cathetus2Int)
	square := calcSquare(cathetus1Int, cathetus2Int)
	perimeter := calcPerimeter(cathetus1Int, cathetus2Int, hypotenuse)
	fmt.Printf("Гипотенуза: %g, Площадь: %g, Периметр: %g", hypotenuse, perimeter, square)
}
