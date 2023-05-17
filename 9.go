package main

import "fmt"

type livro struct {
	titulo string
	autor  string
	preco  float64
}

func nove(x *livro) float64 {
	a := *x
	novopreco := a.preco * 0.9
	return novopreco
}
func main() {
	var x float64
	fmt.Print("digite o valor do livro: ")
	fmt.Scan(&x)

	libro := livro{
		preco: x,
	}
	final := nove(&libro)
	fmt.Print("novo preço: ", final)
}