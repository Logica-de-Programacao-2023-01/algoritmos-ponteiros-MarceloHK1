package main

import "fmt"

func oito(x *int) int {
	a := *x
	if a != 1 {
		a = 1
	} else {
		a = 2
	}
	return a
}
func main() {
	var x int
	fmt.Print("digite um numero: ")
	fmt.Scan(&x)
	a := oito(&x)
	fmt.Print("novo valor: ", a)
}
