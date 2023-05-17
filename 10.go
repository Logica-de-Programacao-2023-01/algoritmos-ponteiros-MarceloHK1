package main

import "fmt"

func dez(x *[]int, y int) []int {
	a := *x
	qntd := 0
	numero := 2

	for qntd < y {
		if primos(numero) {
			a = append(a, numero)
			qntd++
		}
		numero++
	}
	*x = a
	return *x
}
func primos(a int) bool {
	if a < 2 {
		return false
	}

	for i := 2; i*i <= a; i++ {
		if a%i == 0 {
			return false
		}
	}
	return true
}
func main() {
	var x int
	var slice []int
	fmt.Print("digite o tamanho: ")
	fmt.Scan(&x)

	final := dez(&slice, x)
	fmt.Print("slice final: ", final)

}