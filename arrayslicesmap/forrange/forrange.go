package main

import "fmt"

func main() {
	numeros := [...]int{1, 2, 3, 4, 5} // Compilador conta o tamanho do array
	for i, numero := range numeros {
		fmt.Printf("%d) %d\n", i+1, numero)
	}

	for _, numero := range numeros {
		fmt.Println(numero)
	}
}
