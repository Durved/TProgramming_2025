package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Hello, World!")
}

func F(a, b, x float64) float64 {
	numerator := math.Pow(b, 3) + math.Pow(math.Sin(a*x), 2)
	denominator := math.Acos(x*b*x) + math.Exp(-x/2)
	return numerator / denominator
}
