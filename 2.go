package main

import "fmt"

func par(x *int) int {
	if *x%2 == 0 {
		*x = 0
		return *x
	} else {
		*x = 1
		return *x
	}
}
func main() {
	var x int
	fmt.Print("digite um número inteiro: ")
	fmt.Scan(&x)
	par(&x)
	fmt.Print("novo valor: ", x)
}
