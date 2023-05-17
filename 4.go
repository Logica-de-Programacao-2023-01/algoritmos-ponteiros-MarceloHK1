package main

import "fmt"

func soma2(x *int) (int) {
	numero := *x

	ultimodig := numero % 10
	numero /= 10
	penultimo := numero % 10

	soma := ultimodig + penultimo
	return soma
}
func main() {
	var x int
	fmt.Print("digite um numero inteiro: ")
	fmt.Scan(&x)
	final := soma2(&x)
	fmt.Print("novo valor: ", final)

}
