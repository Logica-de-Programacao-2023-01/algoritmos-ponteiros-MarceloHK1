package main

import "fmt"

type Conta struct {
	saldo   float64
	titular string
}

func sete(x *Conta, y float64) float64 {
	a := *x
	soma := a.saldo + y
	return soma
}
func main() {
	var x, y float64
	var z string

	fmt.Print("digite o saldo atual: ")
	fmt.Scan(&x)
	fmt.Print("digite o novo valor a ser adicionado: ")
	fmt.Scan(&y)
	fmt.Print("digite o nome do titular: ")
	fmt.Scan(&z)

	conta := Conta{
		saldo:   x,
		titular: z,
	}
	final := sete(&conta, y)
	fmt.Print("o saldo final é: ", final)
}