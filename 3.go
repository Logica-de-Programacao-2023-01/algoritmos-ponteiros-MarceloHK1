package main

import "fmt"

func reverso(x *string) (y string) {
	for _, i := range *x {
		y = string(i) + y
	}
	return
}
func main() {
	var x string
	fmt.Print("digite uma string: ")
	fmt.Scan(&x)
	reverse := reverso(&x)
	fmt.Print("o reverso é ", reverse)
}
