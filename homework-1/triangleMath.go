package main

import (
	"fmt"
	"math"
	"strconv"
)

func calcTriangleParams(cathetus1 float64, cathetus2 float64) (float64, float64, float64){
	hypotenuse := math.Sqrt(math.Pow(cathetus1, 2) + math.Pow(cathetus2, 2))
	perimeter := cathetus1 + cathetus2 + hypotenuse
	square := cathetus1 * cathetus2 / 2
	return hypotenuse, perimeter, square
}

func main()  {
	var cathetus1Str string
	var cathetus2Str string
	fmt.Println("Введите катет 1")
	fmt.Scanln(&cathetus1Str)
	fmt.Println("Введите катет 2")
	fmt.Scanln(&cathetus2Str)
	cathetus1Int, _ := strconv.ParseFloat(cathetus1Str, 64)
	cathetus2Int, _ := strconv.ParseFloat(cathetus2Str, 64)
	hypotenuse, perimeter, square := calcTriangleParams(cathetus1Int, cathetus2Int)
	fmt.Printf("Гипотенуза: %g, Периметр: %g, Площадь: %g", hypotenuse, perimeter, square)
}
