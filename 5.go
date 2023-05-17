package main

import (
	"fmt"
	"math"
)

func cinco(x *float64) float64 {
	y := math.Pi
	media := (*x + y) / 2
	return media
}
func main() {
	var x float64
	fmt.Print("digite um numero: ")
	fmt.Scan(&x)
	final := cinco(&x)
	fmt.Print("valor novo: ", final)
}
