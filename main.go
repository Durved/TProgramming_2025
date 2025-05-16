package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Задача A")
	fmt.Println(ListValue(1.2, 0.48, 0.7, 2.2, 0.3))

	fmt.Println("Задача B")
	fmt.Println(ListValueByList(1.2, 0.48, []float64{0.25, 0.36, 0.56, 0.94, 1.28}))
}

func F(a, b, x float64) float64 {
	numerator := math.Pow(b, 3) + math.Pow(math.Sin(a*x), 2)
	denominator := math.Acos(x*b*x) + math.Exp(-x/2)
	return numerator / denominator
}

func ListValue(a, b, x0, xn, dx float64) []float64 {
	var result []float64
	for x := x0; x <= xn; x += dx {
		result = append(result, F(a, b, x))
	}
	return result
}

func ListValueByList(a, b float64, xx []float64) []float64 {
	var result []float64
	for _, x := range xx {
		result = append(result, F(a, b, x))
	}
	return result
}
