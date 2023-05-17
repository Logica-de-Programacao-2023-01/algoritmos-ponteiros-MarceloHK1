package main

import "fmt"

func soma(x *int, b int) int {
	soma := 0
	if b > 0 {
		for i := 0; i <= b; i++ {
			soma += i
		}
		*x = soma
		return *x
	} else {
		fmt.Print("numero invalido")
		return 0
	}
}
func main() {
	var x, y int
	fmt.Print("digite um número inteiro: ")
	fmt.Scan(&x)

	soma(&y, x)
	fmt.Printf("a soma de todos os numeros até %d é igual a %d", x, y)
}
