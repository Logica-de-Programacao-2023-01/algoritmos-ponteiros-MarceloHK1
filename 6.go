package main

import "fmt"

type Livro struct {
	titulo string
	autor  string
}

func seis(x *Livro) string {
	a := *x
	if a.autor == "Anônimo" {
		a.titulo = "Desconhecido"
	}
	return a.titulo
}
func main() {
	var x, y string

	fmt.Print("quem é o autor? ")
	fmt.Scan(&y)
	fmt.Print("qual é o titulo? ")
	fmt.Scan(&x)

	livro := Livro{
		titulo: x,
		autor:  y,
	}
	novo := seis(&livro)
	fmt.Print("novo titulo: ", novo)
}
