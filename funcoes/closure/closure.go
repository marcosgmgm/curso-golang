package main

import "fmt"

func closure() func() {
	x := 10 // Contexto de x só existe aqui
	var funcao = func() {
		fmt.Println(x)
	}
	return funcao
}

func main() {
	x := 20
	fmt.Println(x)
	imprimeX := closure()
	imprimeX()
}
